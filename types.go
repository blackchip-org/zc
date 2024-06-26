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
	"github.com/shopspring/decimal"
)

var (
	Any       = anyType{}
	BigInt    = BigIntType{}
	BigFloat  = BigFloatType{}
	Complex   = ComplexType{}
	Decimal   = DecimalType{}
	DecimalSS = DecimalSSType{}
	Float64   = Float64Type{}
	Int       = IntType{}
	Int8      = Int8Type{}
	Int32     = Int32Type{}
	Int64     = Int64Type{}
	Rat       = RatType{}
	String    = StringType{}
	Uint      = UintType{}
	Uint8     = Uint8Type{}
	Uint16    = Uint16Type{}
	Uint32    = Uint32Type{}
	Uint64    = Uint64Type{}
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
	case decimal.Decimal:
		return DecimalSS
	case float64:
		return Float64
	case int:
		return Int
	case int8:
		return Int8
	case int32:
		return Int32
	case string:
		return String
	case uint:
		return Uint
	case uint8:
		return Uint8
	default:
		return Any
	}
}

var (
	poolSize     = 8
	bigFloatPool = NewPool[big.Float](poolSize)
	bigIntPool   = NewPool[big.Int](poolSize)
	decimalPool  = NewPool[apd.Decimal](poolSize)
	ratPool      = NewPool[big.Rat](poolSize)
)

type Type interface {
	Name() string
	From(any) (any, Type, bool)
	Format(any) string
	Recycle(any)
}

// ----------------------------------------------------------------------------

type anyType struct{}

func (t anyType) Name() string { return "Any" }

func (t anyType) From(src any) (any, Type, bool) {
	return src, Any, true
}

func (t anyType) Format(a any) string {
	return fmt.Sprintf("%v", a)
}

func (t anyType) Recycle(any) {}

// ----------------------------------------------------------------------------

type BigIntType struct{}

func (t BigIntType) Name() string { return "Int" }

func (t BigIntType) As(a any) *big.Int {
	val, ok := a.(*big.Int)
	if !ok {
		panic(ErrWrongGoType("*big.Int", a))
	}
	return val
}

func (t BigIntType) Pop(e *OpEnv) *big.Int {
	return t.As(e.Pop().Val)
}

func (t BigIntType) Push(e *OpEnv, v *big.Int) {
	e.PushVal(v)
}

func (t BigIntType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case *big.Int:
		return v, t, true
	case int:
		bi := bigIntPool.New()
		bi.SetInt64(int64(v))
		return bi, Int, true
	case string:
		bi := bigIntPool.New()
		v = PreParseNumber(v)
		_, ok := bi.SetString(v, 0)
		return bi, String, ok
	case uint:
		bi := bigIntPool.New()
		bi.SetUint64(uint64(v))
		return bi, Uint, true
	case uint8:
		bi := bigIntPool.New()
		bi.SetUint64(uint64(v))
		return bi, Uint, true
	}
	return nil, nil, false
}

func (t BigIntType) Format(a any) string {
	return t.As(a).String()
}

func (t BigIntType) New() *big.Int {
	return bigIntPool.New()
}

func (t BigIntType) Recycle(v any) {
	bigIntPool.Recycle(v.(*big.Int))
}

// ----------------------------------------------------------------------------
type ComplexType struct{}

func (t ComplexType) Name() string { return "Complex" }

func (t ComplexType) As(a any) complex128 {
	val, ok := a.(complex128)
	if !ok {
		panic(ErrWrongGoType("complex128", a))
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

func (t ComplexType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case complex128:
		return src, Complex, true
	case *apd.Decimal:
		f, err := v.Float64()
		return f, Decimal, err == nil
	case *big.Int:
		if !v.IsInt64() {
			return nil, nil, false
		}
		i := v.Int64()
		return complex(float64(i), 0), Int, true
	case int:
		return complex(float64(v), 0), Int, true
	case int32:
		return complex(float64(v), 0), Int, true
	case float64:
		return complex(v, 0), Int, true
	case string:
		c, err := strconv.ParseComplex(v, 128)
		return c, String, err == nil
	}
	return nil, nil, false
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

func (t DecimalType) Name() string { return "Dec" }

func (t DecimalType) As(a any) *apd.Decimal {
	val, ok := a.(*apd.Decimal)
	if !ok {
		panic(ErrWrongGoType("*apd.Decimal", a))
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

func (t DecimalType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case *apd.Decimal:
		return v, t, true
	case *big.Int:
		// FIXME: slow
		d := decimalPool.New()
		d.SetString(v.String())
		return d, BigInt, true
	case int:
		d := decimalPool.New()
		d.SetInt64(int64(v))
		return d, Int, true
	case float64:
		d := decimalPool.New()
		d.SetFloat64(v)
		return d, Float64, true
	case string:
		d := decimalPool.New()
		v = PreParseNumber(v)
		_, _, err := d.SetString(v)
		return d, String, err == nil
	}
	return nil, nil, false
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

type DecimalSSType struct{}

func (t DecimalSSType) Name() string { return "Dec/ss" }

func (t DecimalSSType) As(a any) decimal.Decimal {
	val, ok := a.(decimal.Decimal)
	if !ok {
		panic(ErrWrongGoType("decimal.Decimal", a))
	}
	return val
}

func (t DecimalSSType) Pop(e *OpEnv) decimal.Decimal {
	return t.As(e.Pop().Val)
}

func (t DecimalSSType) Push(e *OpEnv, v decimal.Decimal) {
	e.PushVal(v)
}

func (t DecimalSSType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case decimal.Decimal:
		return v, t, true
	case *big.Int:
		d := decimal.NewFromBigInt(v, 0)
		return d, BigInt, true
	case int:
		d := decimal.NewFromInt(int64(v))
		return d, Int, true
	case float64:
		d := decimal.NewFromFloat(v)
		return d, Float64, true
	case string:
		v = PreParseNumber(v)
		d, err := decimal.NewFromString(v)
		return d, String, err == nil
	}
	return nil, nil, false
}

func (t DecimalSSType) Format(v any) string {
	return t.As(v).String()
}

func (t DecimalSSType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type BigFloatType struct{}

func (t BigFloatType) Name() string { return "Float" }

func (t BigFloatType) As(a any) *big.Float {
	val, ok := a.(*big.Float)
	if !ok {
		panic(ErrWrongGoType("*big.Float", a))
	}
	return val
}

func (t BigFloatType) Pop(e *OpEnv) *big.Float {
	conf := state.ForConf(e.State)
	bf := t.As(e.Pop().Val)
	bf.SetPrec(conf.FloatPrec)
	return bf
}

func (t BigFloatType) Push(e *OpEnv, bf *big.Float) {
	if bf.IsInf() {
		e.Err = ErrInfinity(e, bf.Sign())
	} else {
		e.PushVal(bf)
	}
}

func (t BigFloatType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case *big.Float:
		return v, t, true
	case *apd.Decimal:
		// FIXME: slow
		bf := bigFloatPool.New()
		bf.SetString(v.String())
		return bf, Decimal, true
	case int:
		bf := bigFloatPool.New()
		bf.SetInt64(int64(v))
		return bf, Int, true
	case string:
		bf := bigFloatPool.New()
		v = PreParseNumber(v)
		_, ok := bf.SetString(v)
		return bf, String, ok
	}
	return nil, nil, false
}

func (t BigFloatType) Format(v any) string {
	return t.As(v).Text('f', -1)
}

func (t BigFloatType) New() *big.Float {
	return bigFloatPool.New()
}

func (t BigFloatType) Recycle(v any) {
	bigFloatPool.Recycle(v.(*big.Float))
}

// ----------------------------------------------------------------------------

type Float64Type struct{}

func (t Float64Type) Name() string { return "Float/64" }

func (t Float64Type) As(a any) float64 {
	val, ok := a.(float64)
	if !ok {
		panic(ErrWrongGoType("float64", a))
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

func (t Float64Type) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case float64:
		return v, Float64, true
	case string:
		v = PreParseNumber(v)
		f64, err := strconv.ParseFloat(v, 64)
		return f64, String, err == nil
	}
	return nil, nil, false
}

func (t Float64Type) Format(v any) string {
	return strconv.FormatFloat(t.As(v), 'f', -1, 64)
}

func (t Float64Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type IntType struct{}

func (t IntType) Name() string { return "Int/s" }

func (t IntType) As(a any) int {
	val, ok := a.(int)
	if !ok {
		panic(ErrWrongGoType("int", a))
	}
	return val
}

func (t IntType) Pop(e *OpEnv) int {
	return t.As(e.Pop().Val)
}

func (t IntType) Push(e *OpEnv, v int) {
	e.PushVal(v)
}

func (t IntType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case int:
		return src, Int, true
	case string:
		v = PreParseNumber(v)
		i, err := strconv.ParseInt(v, 0, 0)
		return int(i), String, err == nil
	}
	return nil, nil, false
}

func (t IntType) Format(v any) string {
	return strconv.FormatInt(int64(t.As(v)), 10)
}

func (t IntType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Int8Type struct{}

func (t Int8Type) Name() string { return "Int/s8" }

func (t Int8Type) As(a any) int8 {
	val, ok := a.(int8)
	if !ok {
		panic(ErrWrongGoType("int8", a))
	}
	return val
}

func (t Int8Type) Pop(e *OpEnv) int8 {
	return t.As(e.Pop().Val)
}

func (t Int8Type) Push(e *OpEnv, v int8) {
	e.PushVal(v)
}

func (t Int8Type) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case int8:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		i8, err := strconv.ParseInt(v, 0, 8)
		return int8(i8), String, err == nil
	}
	return nil, nil, false
}

func (t Int8Type) Format(v any) string {
	return strconv.FormatInt(int64(t.As(v)), 10)
}

func (t Int8Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Int32Type struct{}

func (t Int32Type) Name() string { return "Int/s32" }

func (t Int32Type) As(a any) int32 {
	val, ok := a.(int32)
	if !ok {
		panic(ErrWrongGoType("int32", a))
	}
	return val
}

func (t Int32Type) Pop(e *OpEnv) int32 {
	return t.As(e.Pop().Val)
}

func (t Int32Type) Push(e *OpEnv, v int32) {
	e.PushVal(v)
}

func (t Int32Type) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case int32:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		i32, err := strconv.ParseInt(v, 0, 32)
		return int32(i32), String, err == nil
	}
	return nil, nil, false
}

func (t Int32Type) Format(v any) string {
	return strconv.FormatInt(int64(t.As(v)), 10)
}

func (t Int32Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Int64Type struct{}

func (t Int64Type) Name() string { return "Int/s64" }

func (t Int64Type) As(a any) int64 {
	val, ok := a.(int64)
	if !ok {
		panic(ErrWrongGoType("int64", a))
	}
	return val
}

func (t Int64Type) Pop(e *OpEnv) int64 {
	return t.As(e.Pop().Val)
}

func (t Int64Type) Push(e *OpEnv, v int64) {
	e.PushVal(v)
}

func (t Int64Type) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case int64:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		i64, err := strconv.ParseInt(v, 0, 64)
		return int64(i64), String, err == nil
	}
	return nil, nil, false
}

func (t Int64Type) Format(v any) string {
	return strconv.FormatInt(t.As(v), 10)
}

func (t Int64Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type RatType struct{}

func (t RatType) Name() string { return "Rat" }

func (t RatType) As(a any) *big.Rat {
	val, ok := a.(*big.Rat)
	if !ok {
		panic(ErrWrongGoType("*big.Rat", a))
	}
	return val
}

func (t RatType) Pop(e *OpEnv) *big.Rat {
	return t.As(e.Pop().Val)
}

func (t RatType) Push(e *OpEnv, r *big.Rat) {
	e.PushVal(r)
}

func (t RatType) From(src any) (any, Type, bool) {
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
		r, ok := t.Parse(v)
		return r, String, ok
	}
	return nil, nil, false
}

func (t RatType) Parse(s string) (*big.Rat, bool) {
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

func (t StringType) Name() string { return "Text" }

func (t StringType) As(a any) string {
	val, ok := a.(string)
	if !ok {
		panic(ErrWrongGoType("string", a))
	}
	return val
}

func (t StringType) Pop(e *OpEnv) string {
	return t.As(e.Pop().Val)
}

func (t StringType) Push(e *OpEnv, s string) {
	e.PushVal(s)
}

func (t StringType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case string:
		return v, t, true
	}
	return nil, nil, false
}

func (t StringType) Format(v any) string {
	return t.As(v)
}

func (t StringType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type UintType struct{}

func (t UintType) Name() string { return "Int/u" }

func (t UintType) As(a any) uint {
	val, ok := a.(uint)
	if !ok {
		panic(ErrWrongGoType("uint", a))
	}
	return val
}

func (t UintType) Pop(e *OpEnv) uint {
	return t.As(e.Pop().Val)
}

func (t UintType) Push(e *OpEnv, ui uint) {
	e.PushVal(ui)
}

func (t UintType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint:
		return v, t, true
	case *big.Int:
		if !v.IsUint64() {
			return nil, nil, false
		}
		ui := v.Uint64()
		return ui, BigInt, true
	case string:
		v = PreParseNumber(v)
		ui, err := strconv.ParseUint(v, 0, 0)
		return uint(ui), String, err == nil
	}
	return nil, nil, false
}

func (t UintType) Format(v any) string {
	return strconv.FormatUint(uint64(t.As(v)), 10)
}

func (t UintType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Uint8Type struct{}

func (t Uint8Type) Name() string { return "Int/u8" }

func (t Uint8Type) As(a any) uint8 {
	val, ok := a.(uint8)
	if !ok {
		panic(ErrWrongGoType("uint8", a))
	}
	return val
}

func (t Uint8Type) Pop(e *OpEnv) uint8 {
	return t.As(e.Pop().Val)
}

func (t Uint8Type) Push(e *OpEnv, v uint8) {
	e.PushVal(v)
}

func (t Uint8Type) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint8:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		u8, err := strconv.ParseUint(v, 0, 8)
		return uint8(u8), String, err == nil
	}
	return nil, nil, false
}

func (t Uint8Type) Format(v any) string {
	return strconv.FormatUint(uint64(t.As(v)), 10)
}

func (t Uint8Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Uint16Type struct{}

func (t Uint16Type) Name() string { return "Int/u16" }

func (t Uint16Type) As(a any) uint16 {
	val, ok := a.(uint16)
	if !ok {
		panic(ErrWrongGoType("uint16", a))
	}
	return val
}

func (t Uint16Type) Pop(e *OpEnv) uint16 {
	return t.As(e.Pop().Val)
}

func (t Uint16Type) Push(e *OpEnv, v uint16) {
	e.PushVal(v)
}

func (t Uint16Type) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint16:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		u16, err := strconv.ParseUint(v, 0, 16)
		return uint16(u16), String, err == nil
	}
	return nil, nil, false
}

func (t Uint16Type) Format(v any) string {
	return strconv.FormatUint(uint64(t.As(v)), 10)
}

func (t Uint16Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Uint32Type struct{}

func (t Uint32Type) Name() string { return "Int/u32" }

func (t Uint32Type) As(a any) uint32 {
	val, ok := a.(uint32)
	if !ok {
		panic(ErrWrongGoType("uint32", a))
	}
	return val
}

func (t Uint32Type) Pop(e *OpEnv) uint32 {
	return t.As(e.Pop().Val)
}

func (t Uint32Type) Push(e *OpEnv, v uint32) {
	e.PushVal(v)
}

func (t Uint32Type) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint32:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		u32, err := strconv.ParseUint(v, 0, 32)
		return uint32(u32), String, err == nil
	}
	return nil, nil, false
}

func (t Uint32Type) Format(v any) string {
	return strconv.FormatUint(uint64(t.As(v)), 10)
}

func (t Uint32Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Uint64Type struct{}

func (t Uint64Type) Name() string { return "Int/u64" }

func (t Uint64Type) As(a any) uint64 {
	val, ok := a.(uint64)
	if !ok {
		panic(ErrWrongGoType("uint64", a))
	}
	return val
}

func (t Uint64Type) Pop(e *OpEnv) uint64 {
	return t.As(e.Pop().Val)
}

func (t Uint64Type) Push(e *OpEnv, v uint64) {
	e.PushVal(v)
}

func (t Uint64Type) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint64:
		return v, t, true
	case string:
		v = PreParseNumber(v)
		u64, err := strconv.ParseUint(v, 0, 64)
		return uint64(u64), String, err == nil
	}
	return nil, nil, false
}

func (t Uint64Type) Format(v any) string {
	return strconv.FormatUint(t.As(v), 10)
}

func (t Uint64Type) Recycle(v any) {}
