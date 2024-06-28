package zc

import (
	"fmt"
	"math"
	"math/big"
	"math/cmplx"
	"strconv"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/cockroachdb/apd/v3"
	"golang.org/x/exp/constraints"
)

var (
	Any      = anyType{}
	BigFloat = BigFloatType{}
	BigInt   = BigIntType{}
	Complex  = ComplexType{}
	Decimal  = DecimalType{}
	Float64  = Float64Type{}
	Int      = IntType{}
	Int8     = Int8Type{}
	Int16    = Int16Type{}
	Int32    = Int32Type{}
	Int64    = Int64Type{}
	Rat      = RatType{}
	String   = StringType{}
	Uint     = UintType{}
	Uint8    = Uint8Type{}
	Uint16   = Uint16Type{}
	Uint32   = Uint32Type{}
	Uint64   = Uint64Type{}
)

func TypeOf(a any) Type {
	switch a.(type) {
	case *big.Int:
		return BigInt
	case *big.Float:
		return BigFloat
	case *big.Rat:
		return Rat
	case *apd.Decimal:
		return Decimal
	case complex128:
		return Complex
	case float64:
		return Float64
	case int:
		return Int
	case int8:
		return Int8
	case int16:
		return Int16
	case int32:
		return Int32
	case int64:
		return Int64
	case string:
		return String
	case uint:
		return Uint
	case uint8:
		return Uint8
	case uint16:
		return Uint16
	case uint32:
		return Uint32
	case uint64:
		return Uint64
	default:
		return Any
	}
}

var (
	poolSize    = 8
	decimalPool = NewPool[apd.Decimal](poolSize)
	floatPool   = NewPool[big.Float](poolSize)
	intPool     = NewPool[big.Int](poolSize)
	ratPool     = NewPool[big.Rat](poolSize)
)

type Type interface {
	Name() string
	AppName() string
	GoName() string
	String() string
	Equal(any, any) bool
	From(state.State, any) (any, Type, bool)
	Format(any) string
	Recycle(any)
}

// ----------------------------------------------------------------------------

type anyType struct{}

func (t anyType) Name() string    { return "Any" }
func (t anyType) AppName() string { return "Any" }
func (t anyType) GoName() string  { return "any" }
func (t anyType) String() string  { return t.AppName() }

func (t anyType) Equal(ax, ay any) bool {
	return false
}

func (t anyType) From(_ state.State, src any) (any, Type, bool) {
	return src, Any, true
}

func (t anyType) Format(a any) string {
	return fmt.Sprintf("%v", a)
}

func (t anyType) Recycle(any) {}

// ----------------------------------------------------------------------------

type BigFloatType struct{}

func (t BigFloatType) Name() string    { return "BigFloat" }
func (t BigFloatType) AppName() string { return "Float" }
func (t BigFloatType) GoName() string  { return "*big.Float" }
func (t BigFloatType) String() string  { return t.AppName() }

func (t BigFloatType) As(a any) *big.Float {
	val, ok := a.(*big.Float)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t BigFloatType) Pop(e *OpEnv) *big.Float {
	return t.As(e.Pop().Val)
}

func (t BigFloatType) Push(e *OpEnv, bf *big.Float) {
	if bf.IsInf() {
		e.Err = ErrInfinity(e, bf.Sign())
	} else {
		e.PushVal(bf)
	}
}

func (t BigFloatType) Equal(ax, ay any) bool {
	if ax == nil && ay == nil {
		return false
	}
	x, y := t.As(ax), t.As(ay)
	return x.Cmp(y) == 0
}

func (t BigFloatType) From(s state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case *big.Float:
		return v, t, true
	case *big.Int:
		bf := t.New(s)
		bf.SetInt(v)
		return bf, BigInt, true
	case complex128:
		r, i := real(v), imag(v)
		if i != 0 {
			return nil, Complex, false
		}
		bf := t.New(s)
		bf.SetFloat64(r)
		return bf, Complex, true
	case *apd.Decimal:
		bf := t.New(s)
		f, err := v.Float64()
		if err == nil {
			bf.SetFloat64(f)
			return bf, Decimal, true
		}
		_, ok := bf.SetString(v.String())
		if !ok {
			t.Recycle(bf)
			return nil, Decimal, false
		}
		return bf, Decimal, true
	case float64:
		bf := t.New(s)
		bf.SetFloat64(v)
		return bf, Float64, true
	case int:
		return intToBigFloat(s, v), Int, true
	case int8:
		return intToBigFloat(s, v), Int8, true
	case int16:
		return intToBigFloat(s, v), Int16, true
	case int32:
		return intToBigFloat(s, v), Int32, true
	case int64:
		return intToBigFloat(s, v), Int64, true
	case *big.Rat:
		f, _ := v.Float64()
		bf := t.New(s)
		bf.SetFloat64(f)
		return bf, Rat, true
	case uint:
		return uintToBigFloat(s, v), Uint, true
	case uint8:
		return uintToBigFloat(s, v), Uint8, true
	case uint16:
		return uintToBigFloat(s, v), Uint16, true
	case uint32:
		return uintToBigFloat(s, v), Uint32, true
	case uint64:
		return uintToBigFloat(s, v), Uint64, true
	case string:
		v = PreParseNumber(v)
		bf, ok := t.Parse(s, v)
		return bf, String, ok
	}
	return nil, Any, false
}

func (t BigFloatType) Parse(s state.State, str string) (*big.Float, bool) {
	// Try to set with a float64 first. SetString("42.42") can give different
	// results than SetFloat64(42.42)
	bf := t.New(s)
	if f, err := strconv.ParseFloat(str, 64); err == nil {
		bf.SetFloat64(f)
		return bf, true
	}
	_, ok := bf.SetString(str)
	if !ok {
		t.Recycle(bf)
		return nil, false
	}
	return bf, true
}

func (t BigFloatType) MustParse(s state.State, str string) *big.Float {
	bf, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return bf
}

func (t BigFloatType) Format(v any) string {
	return t.As(v).Text('f', -1)
}

func (t BigFloatType) New(s state.State) *big.Float {
	conf := state.ForConf(s)
	f := floatPool.New()
	f.SetPrec(conf.FloatPrec)
	return f
}

func (t BigFloatType) Recycle(v any) {
	floatPool.Recycle(v.(*big.Float))
}

// ----------------------------------------------------------------------------

type BigIntType struct{}

func (t BigIntType) Name() string    { return "BigInt" }
func (t BigIntType) AppName() string { return "Int" }
func (t BigIntType) GoName() string  { return "*big.Int" }
func (t BigIntType) String() string  { return t.AppName() }

func (t BigIntType) As(a any) *big.Int {
	val, ok := a.(*big.Int)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t BigIntType) Pop(e *OpEnv) *big.Int {
	return t.As(e.Pop().Val)
}

func (t BigIntType) Push(e *OpEnv, v *big.Int) {
	e.PushVal(v)
}

func (t BigIntType) Equal(ax, ay any) bool {
	if ax == nil && ay == nil {
		return false
	}
	x, y := t.As(ax), t.As(ay)
	return x.Cmp(y) == 0
}

func (t BigIntType) From(s state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case *big.Int:
		return v, t, true
	case *big.Float:
		if !v.IsInt() {
			return nil, Any, false
		}
		bi := t.New()
		v.Int(bi)
		return bi, BigFloat, true
	case complex128:
		r, i := real(v), imag(v)
		if i != 0 {
			return nil, Complex, false
		}
		bi, ok := float64ToBigInt(r)
		return bi, Complex, ok
	case *apd.Decimal:
		i64, err := v.Int64()
		if err != nil {
			return nil, Decimal, false
		}
		bi := t.New()
		bi.SetInt64(i64)
		return bi, Decimal, true
	case float64:
		bi, ok := float64ToBigInt(v)
		return bi, Float64, ok
	case int:
		return intToBigInt(v), Int, true
	case int8:
		return intToBigInt(v), Int8, true
	case int16:
		return intToBigInt(v), Int16, true
	case int32:
		return intToBigInt(v), Int32, true
	case int64:
		return intToBigInt(v), Int64, true
	case *big.Rat:
		if !v.IsInt() {
			return nil, Rat, false
		}
		bi := t.New()
		bi.Set(v.Num())
		return bi, Rat, true
	case uint:
		return uintToBigInt(v), Uint, true
	case uint8:
		return uintToBigInt(v), Uint8, true
	case uint16:
		return uintToBigInt(v), Uint16, true
	case uint32:
		return uintToBigInt(v), Uint32, true
	case uint64:
		return uintToBigInt(v), Uint64, true
	case string:
		bi, ok := t.Parse(s, v)
		return bi, String, ok
	}
	return nil, Any, false
}

func (t BigIntType) Parse(_ state.State, str string) (*big.Int, bool) {
	bi := t.New()
	str = PreParseNumber(str)
	_, ok := bi.SetString(str, 0)
	return bi, ok
}

func (t BigIntType) MustParse(s state.State, str string) *big.Int {
	bi, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return bi
}

func (t BigIntType) Format(a any) string {
	return t.As(a).String()
}

func (t BigIntType) New() *big.Int {
	return intPool.New()
}

func (t BigIntType) Recycle(v any) {
	intPool.Recycle(v.(*big.Int))
}

// ----------------------------------------------------------------------------
type ComplexType struct{}

func (t ComplexType) Name() string    { return "Complex" }
func (t ComplexType) AppName() string { return "Complex" }
func (t ComplexType) GoName() string  { return "complex128" }
func (t ComplexType) String() string  { return t.AppName() }

func (t ComplexType) As(a any) complex128 {
	val, ok := a.(complex128)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t ComplexType) Pop(e *OpEnv) complex128 {
	return t.As(e.Pop().Val)
}

func (t ComplexType) Push(e *OpEnv, v complex128) {
	switch {
	case cmplx.IsInf(v):
		e.Err = ErrInfinity(e, 0)
	case cmplx.IsNaN(v):
		e.Err = ErrNotANumber(e)
	default:
		e.PushVal(v)
	}
}

func (t ComplexType) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t ComplexType) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case complex128:
		return src, Complex, true
	case *big.Int:
		if !v.IsInt64() {
			return nil, BigInt, false
		}
		i := v.Int64()
		return complex(float64(i), 0), BigInt, true
	case *apd.Decimal:
		f, err := v.Float64()
		return complex(f, 0), Decimal, err == nil
	case float64:
		return complex(v, 0), Float64, true
	case int:
		return complex(float64(v), 0), Int, true
	case int8:
		return complex(float64(v), 0), Int8, true
	case int16:
		return complex(float64(v), 0), Int16, true
	case int32:
		return complex(float64(v), 0), Int32, true
	case int64:
		return complex(float64(v), 0), Int64, true
	case *big.Rat:
		f64, _ := v.Float64()
		if math.IsInf(f64, 0) {
			return nil, Rat, false
		}
		// TODO: Set flag for inexact?
		return complex(f64, 0), Rat, true
	case uint:
		return complex(float64(v), 0), Uint, true
	case uint8:
		return complex(float64(v), 0), Uint8, true
	case uint16:
		return complex(float64(v), 0), Uint16, true
	case uint32:
		return complex(float64(v), 0), Uint32, true
	case uint64:
		return complex(float64(v), 0), Uint64, true
	case string:
		c, err := strconv.ParseComplex(v, 128)
		return c, String, err == nil
	}
	return nil, Any, false
}

func (t ComplexType) Format(v any) string {
	f := strconv.FormatComplex(t.As(v), 'f', -1, 128)
	f = f[1 : len(f)-1]
	return f
}

func (t ComplexType) Recycle(v any) {
}

// ----------------------------------------------------------------------------

type DecimalType struct{}

func (t DecimalType) Name() string    { return "Decimal" }
func (t DecimalType) AppName() string { return "Dec" }
func (t DecimalType) GoName() string  { return "*apd.Decimal" }
func (t DecimalType) String() string  { return t.AppName() }

func (t DecimalType) As(a any) *apd.Decimal {
	val, ok := a.(*apd.Decimal)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t DecimalType) Pop(e *OpEnv) *apd.Decimal {
	return t.As(e.Pop().Val)
}

func (t DecimalType) Push(e *OpEnv, v *apd.Decimal) {
	switch v.Form {
	case apd.Infinite:
		// FIXME: Does the Sign have the direction?
		e.Err = ErrInfinity(e, 0)
	case apd.NaN:
		e.Err = ErrNotANumber(e)
	default:
		e.PushVal(v)
	}
}

func (t DecimalType) Equal(ax, ay any) bool {
	if ax == nil && ay == nil {
		return false
	}
	x, y := t.As(ax), t.As(ay)
	return x.Cmp(y) == 0
}

func (t DecimalType) From(s state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case *apd.Decimal:
		return v, t, true
	case complex128:
		r, i := real(v), imag(v)
		if i != 0 {
			return nil, Complex, false
		}
		d := t.New()
		d.SetFloat64(r)
		return d, Complex, true
	case *big.Int:
		d := t.New()
		d.SetString(v.String())
		return d, BigInt, true
	case float64:
		d := decimalPool.New()
		d.SetFloat64(v)
		return d, Float64, true
	case int:
		return intToDecimal(v), Int, true
	case int8:
		return intToDecimal(v), Int8, true
	case int16:
		return intToDecimal(v), Int16, true
	case int32:
		return intToDecimal(v), Int32, true
	case int64:
		return intToDecimal(v), Int64, true
	case *big.Rat:
		conf := state.ForConf(s)
		s := v.FloatString(int(conf.DecPrec))
		d := t.New()
		d.SetString(s)
		return d, Rat, true
	case uint:
		return uintToDecimal(v), Uint, true
	case uint8:
		return uintToDecimal(v), Uint8, true
	case uint16:
		return uintToDecimal(v), Uint16, true
	case uint32:
		return uintToDecimal(v), Uint32, true
	case uint64:
		return uintToDecimal(v), Uint64, true
	case string:
		d, ok := t.Parse(s, v)
		return d, String, ok
	}
	return nil, Any, false
}

func (t DecimalType) Parse(e state.State, str string) (*apd.Decimal, bool) {
	d := t.New()
	str = PreParseNumber(str)
	_, _, err := d.SetString(str)
	return d, err == nil
}

func (t DecimalType) MustParse(e state.State, str string) *apd.Decimal {
	d, ok := t.Parse(e, str)
	if !ok {
		panic(str)
	}
	return d
}

func (t DecimalType) Format(v any) string {
	f := t.As(v).Text('f')
	f = RemoveTrailingZeros(f)
	f = FormatExponent(f)
	return f
}

func (t DecimalType) New() *apd.Decimal {
	return decimalPool.New()
}

func (t DecimalType) Recycle(v any) {
	decimalPool.Recycle(v.(*apd.Decimal))
}

// ----------------------------------------------------------------------------

type Float64Type struct{}

func (t Float64Type) Name() string    { return "Float64" }
func (t Float64Type) AppName() string { return "Float/64" }
func (t Float64Type) GoName() string  { return "float64" }
func (t Float64Type) String() string  { return t.AppName() }

func (t Float64Type) As(a any) float64 {
	val, ok := a.(float64)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Float64Type) Pop(e *OpEnv) float64 {
	return t.As(e.Pop().Val)
}

func (t Float64Type) Push(e *OpEnv, v float64) {
	switch {
	case math.IsInf(v, 1):
		e.Err = ErrInfinity(e, 1)
	case math.IsInf(v, -1):
		e.Err = ErrInfinity(e, -1)
	case math.IsNaN(v):
		e.Err = ErrNotANumber(e)
	default:
		e.PushVal(v)
	}
}

func (t Float64Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Float64Type) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case float64:
		return v, Float64, true
	case *big.Float:
		f, acc := v.Float64()
		return f, BigFloat, acc == big.Exact
	case complex128:
		r, i := real(v), imag(v)
		return r, Complex, i == 0
	case int:
		return float64(v), Int, true
	case int8:
		return float64(v), Int8, true
	case int16:
		return float64(v), Int16, true
	case int32:
		return float64(v), Int32, true
	case int64:
		return float64(v), Int64, true
	case *big.Rat:
		f, _ := v.Float64()
		return f, Rat, !math.IsNaN(f)
	case uint:
		return float64(v), Uint, true
	case uint8:
		return float64(v), Uint8, true
	case uint16:
		return float64(v), Uint16, true
	case uint32:
		return float64(v), Uint32, true
	case uint64:
		return float64(v), Uint64, true
	case string:
		v = PreParseNumber(v)
		f64, err := strconv.ParseFloat(v, 64)
		return f64, String, err == nil
	}
	return nil, Any, false
}

func (t Float64Type) Format(v any) string {
	return strconv.FormatFloat(t.As(v), 'f', -1, 64)
}

func (t Float64Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type IntType struct{}

func (t IntType) Name() string    { return "Int" }
func (t IntType) AppName() string { return "Int/s" }
func (t IntType) GoName() string  { return "int" }
func (t IntType) String() string  { return t.AppName() }

func (t IntType) As(a any) int {
	val, ok := a.(int)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t IntType) Pop(e *OpEnv) int {
	return t.As(e.Pop().Val)
}

func (t IntType) Push(e *OpEnv, v int) {
	e.PushVal(v)
}

func (t IntType) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t IntType) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case int:
		return src, Int, true
	case *big.Float:
		i64, acc := v.Int64()
		return int(i64), BigFloat, acc == big.Exact && isIntRange(i64)
	case *big.Int:
		i64 := v.Int64()
		return int(i64), BigInt, v.IsInt64() && isIntRange(i64)
	case complex128:
		r, i := real(v), imag(v)
		return int(r), Complex, i == 0 && isInt(r) && isIntRangeF(r)
	case *apd.Decimal:
		i64, err := v.Int64()
		return int(i64), Decimal, err == nil
	case int8:
		return int(v), Int8, true
	case int16:
		return int(v), Int16, true
	case int32:
		return int(v), Int32, true
	case int64:
		return int(v), Int64, isIntRange(v)
	case *big.Rat:
		i64, ok := ratToInt64(v)
		return int(i64), Rat, ok && isIntRange(i64)
	case uint:
		return int(v), Uint8, true
	case uint8:
		return int(v), Uint8, true
	case uint16:
		return int(v), Uint16, true
	case uint32:
		return int(v), Uint32, true
	case uint64:
		return int(v), Uint64, isIntRangeU(v)
	case string:
		v = PreParseNumber(v)
		i, err := strconv.ParseInt(v, 0, 0)
		return int(i), String, err == nil
	}
	return nil, Any, false
}

func (t IntType) Format(v any) string {
	return strconv.FormatInt(int64(t.As(v)), 10)
}

func (t IntType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Int8Type struct{}

func (t Int8Type) Name() string    { return "Int8" }
func (t Int8Type) AppName() string { return "Int/s8" }
func (t Int8Type) GoName() string  { return "int8" }
func (t Int8Type) String() string  { return t.AppName() }

func (t Int8Type) As(a any) int8 {
	val, ok := a.(int8)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Int8Type) Pop(e *OpEnv) int8 {
	return t.As(e.Pop().Val)
}

func (t Int8Type) Push(e *OpEnv, v int8) {
	e.PushVal(v)
}

func (t Int8Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Int8Type) From(s state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case int8:
		return v, t, true
	case *big.Float:
		f, ok := bigFloatToFloat64(v)
		return int8(f), BigFloat, ok && isInt8RangeF(f)
	case *big.Int:
		i64 := v.Int64()
		return int8(i64), BigInt, v.IsInt64() && isInt8Range(i64)
	case complex128:
		r, i := real(v), imag(v)
		return int8(r), Complex, i == 0 && isInt8RangeF(r)
	case *apd.Decimal:
		i64, ok := decimalToInt64(v)
		return int8(i64), Decimal, ok && isInt8Range(i64)
	case float64:
		return int8(v), Float64, isInt8RangeF(v)
	case int:
		return int8(v), Int, isInt8Range(int64(v))
	case int16:
		return int8(v), Int16, isInt8Range(int64(v))
	case int32:
		return int8(v), Int32, isInt8Range(int64(v))
	case int64:
		return int8(v), Int64, isInt8Range(int64(v))
	case *big.Rat:
		i64, ok := ratToInt64(v)
		return int8(i64), Rat, ok && isInt8Range(i64)
	case uint:
		return int8(v), Uint, isInt8Range(int64(v))
	case uint8:
		return int8(v), Uint8, isInt8Range(int64(v))
	case uint16:
		return int8(v), Uint16, isInt8Range(int64(v))
	case uint32:
		return int8(v), Uint32, isInt8Range(int64(v))
	case uint64:
		return int8(v), Uint64, isInt8Range(int64(v))
	case string:
		i8, ok := t.Parse(s, v)
		return i8, String, ok
	}
	return nil, Any, false
}

func (t Int8Type) Parse(_ state.State, str string) (int8, bool) {
	str = PreParseNumber(str)
	i64, err := strconv.ParseInt(str, 0, 8)
	return int8(i64), err == nil
}

func (t Int8Type) MustParse(s state.State, str string) int8 {
	i8, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return i8
}

func (t Int8Type) Format(v any) string {
	return strconv.FormatInt(int64(t.As(v)), 10)
}

func (t Int8Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Int16Type struct{}

func (t Int16Type) Name() string    { return "Int16" }
func (t Int16Type) AppName() string { return "Int/s16" }
func (t Int16Type) GoName() string  { return "int16" }
func (t Int16Type) String() string  { return t.AppName() }

func (t Int16Type) As(a any) int16 {
	val, ok := a.(int16)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Int16Type) Pop(e *OpEnv) int16 {
	return t.As(e.Pop().Val)
}

func (t Int16Type) Push(e *OpEnv, v int16) {
	e.PushVal(v)
}

func (t Int16Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Int16Type) From(s state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case int16:
		return v, t, true
	case *big.Float:
		f, ok := bigFloatToFloat64(v)
		return int16(f), BigFloat, ok && isInt16RangeF(f)
	case *big.Int:
		i64 := v.Int64()
		return int16(i64), BigInt, v.IsInt64() && isInt16Range(i64)
	case complex128:
		r, i := real(v), imag(v)
		return int16(r), Complex, i == 0 && isInt16RangeF(r)
	case *apd.Decimal:
		i64, ok := decimalToInt64(v)
		return int16(i64), Decimal, ok && isInt16Range(i64)
	case float64:
		return int16(v), Float64, isInt16RangeF(v)
	case int:
		return int16(v), Int, isInt16Range(int64(v))
	case int8:
		return int16(v), Int8, true
	case int32:
		return int16(v), Int32, isInt16Range(int64(v))
	case int64:
		return int16(v), Int64, isInt16Range(int64(v))
	case *big.Rat:
		i64, ok := ratToInt64(v)
		return int16(i64), Rat, ok && isInt16Range(i64)
	case uint:
		return int16(v), Uint, isInt16Range(int64(v))
	case uint8:
		return int16(v), Uint8, true
	case uint16:
		return int16(v), Uint16, isInt16Range(int64(v))
	case uint32:
		return int16(v), Uint32, isInt16Range(int64(v))
	case uint64:
		return int16(v), Uint64, isInt16Range(int64(v))
	case string:
		i16, ok := t.Parse(s, v)
		return i16, String, ok
	}
	return nil, Any, false
}

func (t Int16Type) Parse(_ state.State, str string) (int16, bool) {
	str = PreParseNumber(str)
	i64, err := strconv.ParseInt(str, 0, 16)
	return int16(i64), err == nil
}

func (t Int16Type) MustParse(s state.State, str string) int16 {
	i16, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return i16
}

func (t Int16Type) Format(v any) string {
	return strconv.FormatInt(int64(t.As(v)), 10)
}

func (t Int16Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Int32Type struct{}

func (t Int32Type) Name() string    { return "Int32" }
func (t Int32Type) AppName() string { return "Int/s32" }
func (t Int32Type) GoName() string  { return "int32" }
func (t Int32Type) String() string  { return t.AppName() }

func (t Int32Type) As(a any) int32 {
	val, ok := a.(int32)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Int32Type) Pop(e *OpEnv) int32 {
	return t.As(e.Pop().Val)
}

func (t Int32Type) Push(e *OpEnv, v int32) {
	e.PushVal(v)
}

func (t Int32Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Int32Type) From(s state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case int32:
		return v, t, true
	case *big.Float:
		f, ok := bigFloatToFloat64(v)
		return int32(f), BigFloat, ok && isInt32RangeF(f)
	case *big.Int:
		i64 := v.Int64()
		return int32(i64), BigInt, v.IsInt64() && isInt32Range(i64)
	case complex128:
		r, i := real(v), imag(v)
		return int32(r), Complex, i == 0 && isInt32RangeF(r)
	case *apd.Decimal:
		i64, ok := decimalToInt64(v)
		return int32(i64), Decimal, ok && isInt32Range(i64)
	case float64:
		return int32(v), Float64, isInt32RangeF(v)
	case int:
		return int32(v), Int, isInt32Range(int64(v))
	case int8:
		return int32(v), Int8, true
	case int16:
		return int32(v), Int16, true
	case int64:
		return int32(v), Int64, isInt32Range(int64(v))
	case *big.Rat:
		i64, ok := ratToInt64(v)
		return int32(i64), Rat, ok && isInt32Range(i64)
	case uint:
		return int32(v), Uint, isInt16Range(int64(v))
	case uint8:
		return int32(v), Uint8, true
	case uint16:
		return int32(v), Uint16, true
	case uint32:
		return int32(v), Uint32, isInt32Range(int64(v))
	case uint64:
		return int32(v), Uint64, isInt32Range(int64(v))
	case string:
		i32, ok := t.Parse(s, v)
		return i32, String, ok
	}
	return nil, Any, false
}

func (t Int32Type) Parse(_ state.State, str string) (int32, bool) {
	str = PreParseNumber(str)
	i64, err := strconv.ParseInt(str, 0, 32)
	return int32(i64), err == nil
}

func (t Int32Type) MustParse(s state.State, str string) int32 {
	i32, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return i32
}

func (t Int32Type) Format(v any) string {
	return strconv.FormatInt(int64(t.As(v)), 10)
}

func (t Int32Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Int64Type struct{}

func (t Int64Type) Name() string    { return "Int64" }
func (t Int64Type) AppName() string { return "Int/s64" }
func (t Int64Type) GoName() string  { return "int64" }
func (t Int64Type) String() string  { return t.AppName() }

func (t Int64Type) As(a any) int64 {
	val, ok := a.(int64)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Int64Type) Pop(e *OpEnv) int64 {
	return t.As(e.Pop().Val)
}

func (t Int64Type) Push(e *OpEnv, v int64) {
	e.PushVal(v)
}

func (t Int64Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Int64Type) From(s state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case int64:
		return v, t, true
	case *big.Float:
		f, ok := bigFloatToFloat64(v)
		return int64(f), BigFloat, ok && isInt64RangeF(f)
	case *big.Int:
		i64 := v.Int64()
		return i64, BigInt, v.IsInt64()
	case complex128:
		r, i := real(v), imag(v)
		return int64(r), Complex, i == 0 && isInt64RangeF(r)
	case *apd.Decimal:
		i64, ok := decimalToInt64(v)
		return i64, Decimal, ok
	case float64:
		return int64(v), Float64, isInt64RangeF(v)
	case int:
		return int64(v), Int, true
	case int8:
		return int64(v), Int8, true
	case int16:
		return int64(v), Int16, true
	case int32:
		return int64(v), Int32, true
	case *big.Rat:
		i64, ok := ratToInt64(v)
		return i64, Rat, ok
	case uint:
		return int64(v), Uint, true
	case uint8:
		return int64(v), Uint8, true
	case uint16:
		return int64(v), Uint16, true
	case uint32:
		return int64(v), Uint32, true
	case uint64:
		return int64(v), Uint64, v <= math.MaxInt64
	case string:
		i64, ok := t.Parse(s, v)
		return i64, String, ok
	}
	return nil, Any, false
}

func (t Int64Type) Parse(_ state.State, str string) (int64, bool) {
	str = PreParseNumber(str)
	i64, err := strconv.ParseInt(str, 0, 64)
	return i64, err == nil
}

func (t Int64Type) MustParse(s state.State, str string) int64 {
	i64, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return i64
}

func (t Int64Type) Format(v any) string {
	return strconv.FormatInt(t.As(v), 10)
}

func (t Int64Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type RatType struct{}

func (t RatType) Name() string    { return "Rat" }
func (t RatType) AppName() string { return "Rat" }
func (t RatType) GoName() string  { return "*big.Rat" }
func (t RatType) String() string  { return t.AppName() }

func (t RatType) As(a any) *big.Rat {
	val, ok := a.(*big.Rat)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t RatType) Pop(e *OpEnv) *big.Rat {
	return t.As(e.Pop().Val)
}

func (t RatType) Push(e *OpEnv, r *big.Rat) {
	e.PushVal(r)
}

func (t RatType) Equal(ax, ay any) bool {
	if ax == nil && ay == nil {
		return false
	}
	x, y := t.As(ax), t.As(ay)
	return x.Cmp(y) == 0
}

func (t RatType) From(s state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case *big.Rat:
		return v, t, true
	case float64:
		r := t.New()
		r.SetFloat64(v)
		return r, Float64, true
	case int:
		r := t.New()
		r.SetInt64(int64(v))
		return r, Int64, true
	case int8:
		r := t.New()
		r.SetInt64(int64(v))
		return r, Int64, true
	case int16:
		r := t.New()
		r.SetInt64(int64(v))
		return r, Int64, true
	case int32:
		r := t.New()
		r.SetInt64(int64(v))
		return r, Int64, true
	case int64:
		r := t.New()
		r.SetInt64(int64(v))
		return r, Int64, true
	case string:
		r, ok := t.Parse(s, v)
		return r, String, ok
	}
	return nil, Any, false
}

func (t RatType) Parse(_ state.State, s string) (*big.Rat, bool) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		r := t.New()
		r.SetInt64(i)
		return r, true
	}
	f, err := strconv.ParseFloat(s, 64)
	if err == nil {
		r := t.New()
		r.SetFloat64(f)
		return r, true
	}

	sc := scan.NewScannerFromString("", s)

	var sign, whole, num, denom int64

	scan.SignedIntRule.Eval(sc)
	s1 := sc.Emit().Val
	i1, err := strconv.ParseInt(s1, 10, 64)
	if err != nil {
		return nil, false
	}
	if i1 < 0 {
		sign = -1
		i1 = i1 * -1
	} else {
		sign = 1
	}
	switch sc.This {
	case '_', '-', ' ':
		whole = i1
	case '/':
		num = i1
	default:
		return nil, false
	}
	sc.Skip()

	scan.IntRule.Eval(sc)
	s2 := sc.Emit().Val
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		return nil, false
	}
	if whole != 0 {
		num = i2
		if sc.This != '/' {
			return nil, false
		}
		sc.Skip()
		scan.IntRule.Eval(sc)
		s3 := sc.Emit().Val
		i3, err := strconv.ParseInt(s3, 10, 64)
		if err != nil {
			return nil, false
		}
		denom = i3
	} else {
		denom = i2
	}

	if whole != 0 {
		num = num + (denom * whole)
	}
	num = num * sign
	r := t.New()
	r.SetFrac64(num, denom)
	return r, true
}

func (t RatType) MustParse(s state.State, str string) *big.Rat {
	r, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return r
}

func (t RatType) Format(v any) string {
	r := t.As(v)
	n := r.Num().Int64()
	d := r.Denom().Int64()

	if n > d {
		w := n / d
		n := n % d
		if n == 0 && d == 1 {
			return fmt.Sprintf("%v", w)
		}
		return fmt.Sprintf("%v %v/%v", w, n, d)
	}
	return r.RatString()
}

func (t RatType) New() *big.Rat {
	return ratPool.New()
}

func (t RatType) Recycle(v any) {
	ratPool.Recycle(v.(*big.Rat))
}

// ----------------------------------------------------------------------------

type StringType struct{}

func (t StringType) Name() string    { return "String" }
func (t StringType) AppName() string { return "Text" }
func (t StringType) GoName() string  { return "string" }
func (t StringType) String() string  { return t.AppName() }

func (t StringType) As(a any) string {
	val, ok := a.(string)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t StringType) Pop(e *OpEnv) string {
	return t.As(e.Pop().Val)
}

func (t StringType) Push(e *OpEnv, s string) {
	e.PushVal(s)
}

func (t StringType) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t StringType) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case string:
		return v, t, true
	}
	return nil, Any, false
}

func (t StringType) Format(v any) string {
	return t.As(v)
}

func (t StringType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type UintType struct{}

func (t UintType) Name() string    { return "Uint" }
func (t UintType) AppName() string { return "Int/u" }
func (t UintType) GoName() string  { return "uint" }
func (t UintType) String() string  { return t.AppName() }

func (t UintType) As(a any) uint {
	val, ok := a.(uint)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t UintType) Pop(e *OpEnv) uint {
	return t.As(e.Pop().Val)
}

func (t UintType) Push(e *OpEnv, ui uint) {
	e.PushVal(ui)
}

func (t UintType) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t UintType) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint:
		return v, t, true
	case *big.Int:
		if !v.IsUint64() {
			return nil, Any, false
		}
		ui := v.Uint64()
		return ui, BigInt, true
	case string:
		v = PreParseNumber(v)
		ui, err := strconv.ParseUint(v, 0, 0)
		return uint(ui), String, err == nil
	}
	return nil, Any, false
}

func (t UintType) Parse(_ state.State, str string) (uint, bool) {
	str = PreParseNumber(str)
	u64, err := strconv.ParseUint(str, 0, 0)
	return uint(u64), err == nil
}

func (t UintType) MustParse(s state.State, str string) uint {
	u, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return u
}

func (t UintType) Format(v any) string {
	return strconv.FormatUint(uint64(t.As(v)), 10)
}

func (t UintType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Uint8Type struct{}

func (t Uint8Type) Name() string    { return "Uint8" }
func (t Uint8Type) AppName() string { return "Int/u8" }
func (t Uint8Type) GoName() string  { return "uint8" }
func (t Uint8Type) String() string  { return t.AppName() }

func (t Uint8Type) As(a any) uint8 {
	val, ok := a.(uint8)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Uint8Type) Pop(e *OpEnv) uint8 {
	return t.As(e.Pop().Val)
}

func (t Uint8Type) Push(e *OpEnv, v uint8) {
	e.PushVal(v)
}

func (t Uint8Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Uint8Type) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint8:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		u8, err := strconv.ParseUint(v, 0, 8)
		return uint8(u8), String, err == nil
	}
	return nil, Any, false
}

func (t Uint8Type) Parse(_ state.State, str string) (uint8, bool) {
	str = PreParseNumber(str)
	u64, err := strconv.ParseUint(str, 0, 8)
	return uint8(u64), err == nil
}

func (t Uint8Type) MustParse(s state.State, str string) uint8 {
	u8, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return u8
}

func (t Uint8Type) Format(v any) string {
	return strconv.FormatUint(uint64(t.As(v)), 10)
}

func (t Uint8Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Uint16Type struct{}

func (t Uint16Type) Name() string    { return "Uint16" }
func (t Uint16Type) AppName() string { return "Int/u16" }
func (t Uint16Type) GoName() string  { return "uint16" }
func (t Uint16Type) String() string  { return t.AppName() }

func (t Uint16Type) As(a any) uint16 {
	val, ok := a.(uint16)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Uint16Type) Pop(e *OpEnv) uint16 {
	return t.As(e.Pop().Val)
}

func (t Uint16Type) Push(e *OpEnv, v uint16) {
	e.PushVal(v)
}

func (t Uint16Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Uint16Type) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint16:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		u16, err := strconv.ParseUint(v, 0, 16)
		return uint16(u16), String, err == nil
	}
	return nil, Any, false
}

func (t Uint16Type) Parse(_ state.State, str string) (uint16, bool) {
	str = PreParseNumber(str)
	u64, err := strconv.ParseUint(str, 0, 16)
	return uint16(u64), err == nil
}

func (t Uint16Type) MustParse(s state.State, str string) uint16 {
	u16, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return u16
}

func (t Uint16Type) Format(v any) string {
	return strconv.FormatUint(uint64(t.As(v)), 10)
}

func (t Uint16Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Uint32Type struct{}

func (t Uint32Type) Name() string    { return "Uint32" }
func (t Uint32Type) AppName() string { return "Int/u32" }
func (t Uint32Type) GoName() string  { return "uint32" }
func (t Uint32Type) String() string  { return t.AppName() }

func (t Uint32Type) As(a any) uint32 {
	val, ok := a.(uint32)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Uint32Type) Pop(e *OpEnv) uint32 {
	return t.As(e.Pop().Val)
}

func (t Uint32Type) Push(e *OpEnv, v uint32) {
	e.PushVal(v)
}

func (t Uint32Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Uint32Type) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint32:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		u32, err := strconv.ParseUint(v, 0, 32)
		return uint32(u32), String, err == nil
	}
	return nil, Any, false
}

func (t Uint32Type) Parse(_ state.State, str string) (uint32, bool) {
	str = PreParseNumber(str)
	u64, err := strconv.ParseUint(str, 0, 32)
	return uint32(u64), err == nil
}

func (t Uint32Type) MustParse(s state.State, str string) uint32 {
	u32, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return u32
}

func (t Uint32Type) Format(v any) string {
	return strconv.FormatUint(uint64(t.As(v)), 10)
}

func (t Uint32Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Uint64Type struct{}

func (t Uint64Type) Name() string    { return "Uint64" }
func (t Uint64Type) AppName() string { return "Int/u64" }
func (t Uint64Type) GoName() string  { return "uint64" }
func (t Uint64Type) String() string  { return t.AppName() }

func (t Uint64Type) As(a any) uint64 {
	val, ok := a.(uint64)
	if !ok {
		panic(ErrWrongGoType(t.GoName(), a))
	}
	return val
}

func (t Uint64Type) Pop(e *OpEnv) uint64 {
	return t.As(e.Pop().Val)
}

func (t Uint64Type) Push(e *OpEnv, v uint64) {
	e.PushVal(v)
}

func (t Uint64Type) Equal(ax, ay any) bool {
	x, y := t.As(ax), t.As(ay)
	return x == y
}

func (t Uint64Type) From(_ state.State, src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint64:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		u64, err := strconv.ParseUint(v, 0, 64)
		return uint64(u64), String, err == nil
	}
	return nil, Any, false
}

func (t Uint64Type) Parse(_ state.State, str string) (uint64, bool) {
	str = PreParseNumber(str)
	u64, err := strconv.ParseUint(str, 0, 0)
	return u64, err == nil
}

func (t Uint64Type) MustParse(s state.State, str string) uint64 {
	u64, ok := t.Parse(s, str)
	if !ok {
		panic(str)
	}
	return u64
}

func (t Uint64Type) Format(v any) string {
	return strconv.FormatUint(t.As(v), 10)
}

func (t Uint64Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

func Equal(a, b any) bool {
	at := TypeOf(a)
	bt := TypeOf(b)
	if at != bt {
		return false
	}
	return at.Equal(a, b)
}

func bigFloatToFloat64(bf *big.Float) (float64, bool) {
	f, _ := bf.Float64()
	return f, !math.IsInf(f, 0)
}

func decimalToInt64(d *apd.Decimal) (int64, bool) {
	i64, err := d.Int64()
	return i64, err == nil
}

func float64ToInt64(f float64) (int64, bool) {
	if f != math.Trunc(f) {
		return 0, false
	}
	if f >= float64(math.MaxInt64) || f <= float64(math.MinInt64) {
		return 0, false
	}
	return int64(f), true
}

func float64ToBigInt(f float64) (*big.Int, bool) {
	i64, ok := float64ToInt64(f)
	if !ok {
		return nil, false
	}
	bi := BigInt.New()
	bi.SetInt64(i64)
	return bi, true
}

func intToBigInt[T constraints.Signed](i T) *big.Int {
	bi := BigInt.New()
	bi.SetInt64(int64(i))
	return bi
}

func intToBigFloat[T constraints.Signed](s state.State, i T) *big.Float {
	bf := BigFloat.New(s)
	bf.SetInt64(int64(i))
	return bf
}

func intToDecimal[T constraints.Signed](i T) *apd.Decimal {
	d := Decimal.New()
	d.SetInt64(int64(i))
	return d
}

func ratToInt64(r *big.Rat) (int64, bool) {
	fmt.Println(r)
	if !r.Denom().IsInt64() || r.Denom().Int64() != 1 {
		return 0, false
	}
	return r.Num().Int64(), r.Num().IsInt64()
}

func uintToBigInt[T constraints.Unsigned](i T) *big.Int {
	bi := BigInt.New()
	bi.SetUint64(uint64(i))
	return bi
}

func uintToBigFloat[T constraints.Unsigned](s state.State, i T) *big.Float {
	bf := BigFloat.New(s)
	bf.SetUint64(uint64(i))
	return bf
}

func uintToDecimal[T constraints.Unsigned](i T) *apd.Decimal {
	s := fmt.Sprintf("%v", i)
	d := Decimal.New()
	d.SetString(s)
	return d
}

func isIntRange(i int64) bool {
	return i <= math.MaxInt && i >= math.MinInt
}

func isInt8Range(i int64) bool {
	return i <= math.MaxInt8 && i >= math.MinInt8
}

func isInt16Range(i int64) bool {
	return i <= math.MaxInt16 && i >= math.MinInt16
}

func isInt32Range(i int64) bool {
	return i <= math.MaxInt32 && i >= math.MinInt32
}

func isIntRangeU(i uint64) bool {
	return i <= math.MaxInt
}

func isIntRangeF(f float64) bool {
	return f == math.Trunc(f) && f < float64(math.MaxInt) && f > float64(math.MinInt)
}

func isInt8RangeF(f float64) bool {
	return f == math.Trunc(f) && f < float64(math.MaxInt8) && f > float64(math.MinInt8)
}

func isInt16RangeF(f float64) bool {
	return f == math.Trunc(f) && f < float64(math.MaxInt16) && f > float64(math.MinInt16)
}

func isInt32RangeF(f float64) bool {
	return f == math.Trunc(f) && f < float64(math.MaxInt32) && f > float64(math.MinInt32)
}

func isInt64RangeF(f float64) bool {
	return f == math.Trunc(f) && f < float64(math.MaxInt64) && f > float64(math.MinInt64)
}

func isInt(f float64) bool {
	return math.Trunc(f) == f
}
