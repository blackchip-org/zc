package zc

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/blackchip-org/dms"
	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/app/vars"
	"github.com/blackchip-org/zc/v6/pkg/coll"
	"github.com/blackchip-org/zc/v6/types"
	"github.com/cockroachdb/apd/v3"
)

var (
	Any      = AnyType{}
	BigFloat = BigFloatType{}
	BigInt   = BigIntType{}
	Bool     = BoolType{}
	Complex  = ComplexType{}
	Decimal  = DecimalType{}
	DMS      = DMSType{}
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

const (
	PrecFloat128 = 113
	PrecDec      = vars.DefaultPrec
)

var (
	poolSize  = 8
	decPool   = coll.NewPool[apd.Decimal](poolSize)
	floatPool = coll.NewPool[big.Float](poolSize)
	intPool   = coll.NewPool[big.Int](poolSize)
	ratPool   = coll.NewPool[big.Rat](poolSize)
)

// ----------------------------------------------------------------------------
type AnyType struct{}

func (t AnyType) AppName() string { return "Any" }
func (t AnyType) GoName() string  { return "any" }

func (t AnyType) Parse(_ coll.State, str string) (any, bool, error) {
	return str, true, nil
}

func (t AnyType) Format(a any) string {
	return fmt.Sprint(a)
}

func (t AnyType) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type BigFloatType struct{}

func (t BigFloatType) AppName() string { return "Float/128" }
func (t BigFloatType) GoName() string  { return "*big.Float" }

func (t BigFloatType) New() *big.Float {
	f := floatPool.New()
	f.SetPrec(113)
	return f
}

func (t BigFloatType) Recycle(vals ...*big.Float) {
	for _, val := range vals {
		floatPool.Recycle(val)
	}
}

func (t BigFloatType) As(a any) *big.Float {
	v, ok := a.(*big.Float)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t BigFloatType) Push(c Calc, val *big.Float) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t BigFloatType) Pop(c Calc) *big.Float {
	return t.As(c.Pop().TypeVal)
}

func (t BigFloatType) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseDecimal(str)
	v := t.New()
	_, ok := v.SetString(str)
	if !ok {
		t.Recycle(v)
		return nil, false, nil
	}
	return v, true, nil
}

func (t BigFloatType) Format(a any) string {
	v := t.As(a)
	return FormatExponent(v.Text('g', -1))
}

func (t BigFloatType) Dup(a any) any {
	i := t.New()
	i.Set(t.As(a))
	return i
}

// ----------------------------------------------------------------------------
type BigIntType struct{}

func (t BigIntType) AppName() string { return "Int" }
func (t BigIntType) GoName() string  { return "*big.Int" }

func (t BigIntType) New() *big.Int {
	return intPool.New()
}

func (t BigIntType) Recycle(vals ...*big.Int) {
	for _, val := range vals {
		intPool.Recycle(val)
	}
}

func (t BigIntType) As(a any) *big.Int {
	v, ok := a.(*big.Int)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t BigIntType) Push(c Calc, val *big.Int) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t BigIntType) Pop(c Calc) *big.Int {
	return t.As(c.Pop().TypeVal)
}

func (t BigIntType) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	v := t.New()
	_, ok := v.SetString(str, 0)
	if !ok {
		t.Recycle(v)
		return nil, false, nil
	}
	return v, true, nil
}

func (t BigIntType) Format(a any) string {
	v := t.As(a)
	return v.String()
}

func (t BigIntType) Dup(a any) any {
	i := t.New()
	i.Set(t.As(a))
	return i
}

// ----------------------------------------------------------------------------
type BoolType struct{}

func (t BoolType) AppName() string { return "Bool" }
func (t BoolType) GoName() string  { return "bool" }

func (t BoolType) As(a any) bool {
	v, ok := a.(bool)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t BoolType) Push(c Calc, val bool) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t BoolType) Pop(c Calc) bool {
	return t.As(c.Pop().TypeVal)
}

func (t BoolType) Parse(_ coll.State, str string) (any, bool, error) {
	istr := strings.ToLower(str)
	switch istr {
	case "true":
		return true, true, nil
	case "false":
		return false, true, nil
	default:
		return false, false, nil
	}
}

func (t BoolType) Format(a any) string {
	b := t.As(a)
	if b {
		return "true"
	}
	return "false"
}

func (t BoolType) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type ComplexType struct{}

func (t ComplexType) AppName() string { return "Complex" }
func (t ComplexType) GoName() string  { return "complex128" }

func (t ComplexType) As(a any) complex128 {
	v, ok := a.(complex128)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t ComplexType) Push(c Calc, val complex128) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t ComplexType) Pop(c Calc) complex128 {
	return t.As(c.Pop().TypeVal)
}

func (t ComplexType) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseDecimal(str)
	c, err := strconv.ParseComplex(str, 128)
	return c, err == nil, nil
}

func (t ComplexType) Format(a any) string {
	f := strconv.FormatComplex(t.As(a), 'f', -1, 128)
	f = f[1 : len(f)-1]
	return f
}

func (t ComplexType) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type DecimalType struct{}

func (t DecimalType) AppName() string { return "Dec" }
func (t DecimalType) GoName() string  { return "*apd.Decimal" }

func (t DecimalType) New() *apd.Decimal {
	return decPool.New()
}

func (t DecimalType) Recycle(vals ...*apd.Decimal) {
	for _, val := range vals {
		decPool.Recycle(val)
	}
}

func (t DecimalType) As(a any) *apd.Decimal {
	v, ok := a.(*apd.Decimal)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t DecimalType) Push(c Calc, val *apd.Decimal) {
	switch val.Form {
	case apd.Infinite:
		// FIXME: Does the Sign have the direction?
		c.Raise(ErrInfinity(0))
	case apd.NaN:
		c.Raise(ErrNotANumber)
	default:
		c.Push(Item{TypeVal: val, Type: t})
	}
}

func (t DecimalType) Pop(c Calc) *apd.Decimal {
	return t.As(c.Pop().TypeVal)
}

func (t DecimalType) Parse(state coll.State, str string) (any, bool, error) {
	d := vars.ForConf(state).DecMath
	str = PreParseDecimal(str)
	v := t.New()
	_, cond, err := d.SetString(v, str)
	switch {
	case cond.Overflow() || (err != nil && err.Error() == "exponent out of range"):
		t.Recycle(v)
		return nil, false, ErrOverflow(str)
	case cond.Underflow():
		t.Recycle(v)
		return nil, false, ErrUnderflow(str)
	case err != nil:
		t.Recycle(v)
		return nil, false, nil
	}
	return v, true, nil
}

func (t DecimalType) Format(a any) string {
	v := t.As(a)
	v.Reduce(v)
	f := v.Text('f')
	//f = RemoveTrailingZeros(f)
	f = FormatExponent(f)
	return f
}

func (t DecimalType) Dup(a any) any {
	d := t.New()
	d.Set(t.As(a))
	return d
}

// ----------------------------------------------------------------------------
type DMSType struct{}

func (t DMSType) AppName() string { return "DMS" }
func (t DMSType) GoName() string  { return "types.DMS" }

func (t DMSType) As(a any) types.DMS {
	v, ok := a.(types.DMS)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t DMSType) Push(c Calc, val types.DMS) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t DMSType) Pop(c Calc) types.DMS {
	return t.As(c.Pop().TypeVal)
}

func (t DMSType) Parse(state coll.State, str string) (any, bool, error) {
	conf := vars.ForConf(state)
	p := dms.NewDefaultParser()
	f, err := p.ParseFields(str)
	if err != nil {
		return nil, false, nil
	}
	d, err := types.NewDMSFromFields(conf.DecMath, f)
	if err != nil {
		return nil, false, nil
	}
	return d, true, nil
}

func (t DMSType) Format(a any) string {
	d := t.As(a)
	return types.FormatDMS(d, dms.SecUnit, -1)
}

func (t DMSType) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Float64Type struct{}

func (t Float64Type) AppName() string { return "Float/64" }
func (t Float64Type) GoName() string  { return "float64" }

func (t Float64Type) As(a any) float64 {
	v, ok := a.(float64)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Float64Type) Push(c Calc, val float64) {
	switch {
	case math.IsNaN(val):
		c.Raise(ErrNotANumber)
	case math.IsInf(val, 1):
		c.Raise(ErrInfinity(1))
	case math.IsInf(val, -1):
		c.Raise(ErrInfinity(-1))
	default:
		c.Push(Item{TypeVal: val, Type: t})
	}
}

func (t Float64Type) Pop(c Calc) float64 {
	return t.As(c.Pop().TypeVal)
}

func (t Float64Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseDecimal(str)
	f64, err := strconv.ParseFloat(str, 64)
	return f64, err == nil, nil
}

func (t Float64Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatFloat(v, 'g', -1, 64)
}

func (t Float64Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type IntType struct{}

func (t IntType) AppName() string { return "Int/s" }
func (t IntType) GoName() string  { return "int" }

func (t IntType) As(a any) int {
	v, ok := a.(int)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t IntType) Push(c Calc, val int) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t IntType) Pop(c Calc) int {
	return t.As(c.Pop().TypeVal)
}

func (t IntType) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseInt(str, 0, 0)
	return int(i64), err == nil, nil
}

func (t IntType) Format(a any) string {
	v := t.As(a)
	return strconv.FormatInt(int64(v), 10)
}

func (t IntType) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Int8Type struct{}

func (t Int8Type) AppName() string { return "Int/s8" }
func (t Int8Type) GoName() string  { return "int8" }

func (t Int8Type) As(a any) int8 {
	v, ok := a.(int8)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Int8Type) Push(c Calc, val int8) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t Int8Type) Pop(c Calc) int8 {
	return t.As(c.Pop().TypeVal)
}

func (t Int8Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseInt(str, 0, 8)
	return int8(i64), err == nil, nil
}

func (t Int8Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatInt(int64(v), 10)
}

func (t Int8Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Int16Type struct{}

func (t Int16Type) AppName() string { return "Int/s16" }
func (t Int16Type) GoName() string  { return "int16" }

func (t Int16Type) As(a any) int16 {
	v, ok := a.(int16)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Int16Type) Push(c Calc, val int16) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t Int16Type) Pop(c Calc) int16 {
	return t.As(c.Pop().TypeVal)
}

func (t Int16Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseInt(str, 0, 16)
	return int16(i64), err == nil, nil
}

func (t Int16Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatInt(int64(v), 10)
}

func (t Int16Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Int32Type struct{}

func (t Int32Type) AppName() string { return "Int/s32" }
func (t Int32Type) GoName() string  { return "int32" }

func (t Int32Type) As(a any) int32 {
	v, ok := a.(int32)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Int32Type) Push(c Calc, val int32) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t Int32Type) Pop(c Calc) int32 {
	return t.As(c.Pop().TypeVal)
}

func (t Int32Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseInt(str, 0, 32)
	return int32(i64), err == nil, nil
}

func (t Int32Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatInt(int64(v), 10)
}

func (t Int32Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Int64Type struct{}

func (t Int64Type) AppName() string { return "Int/s64" }
func (t Int64Type) GoName() string  { return "int64" }

func (t Int64Type) As(a any) int64 {
	v, ok := a.(int64)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Int64Type) Push(c Calc, val int64) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t Int64Type) Pop(c Calc) int64 {
	return t.As(c.Pop().TypeVal)
}

func (t Int64Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseInt(str, 0, 64)
	return int32(i64), err == nil, nil
}

func (t Int64Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatInt(int64(v), 10)
}

func (t Int64Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type RatType struct{}

func (t RatType) AppName() string { return "Rat" }
func (t RatType) GoName() string  { return "*big.Rat" }

func (t RatType) New() *big.Rat {
	return ratPool.New()
}

func (t RatType) Recycle(vals ...*big.Rat) {
	for _, val := range vals {
		ratPool.Recycle(val)
	}
}

func (t RatType) As(a any) *big.Rat {
	v, ok := a.(*big.Rat)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t RatType) Push(c Calc, val *big.Rat) {
	c.Push(Item{TypeVal: val, Type: t})

}

func (t RatType) Pop(c Calc) *big.Rat {
	return t.As(c.Pop().TypeVal)
}

func (t RatType) Parse(_ coll.State, str string) (any, bool, error) {
	return parseRat(str)
}

func (t RatType) Format(a any) string {
	r := t.As(a)
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

func (t RatType) Dup(a any) any {
	r := t.New()
	r.Set(t.As(a))
	return r
}

// ----------------------------------------------------------------------------
type StringType struct{}

func (t StringType) AppName() string { return "Text" }
func (t StringType) GoName() string  { return "string" }

func (t StringType) As(a any) string {
	v, ok := a.(string)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t StringType) Push(c Calc, val string) {
	c.Push(Item{TypeVal: val, Type: t})

}

func (t StringType) Pop(c Calc) string {
	return t.As(c.Pop().TypeVal)
}

func (t StringType) Parse(_ coll.State, str string) (any, bool, error) {
	return str, true, nil
}

func (t StringType) Format(a any) string {
	return t.As(a)
}

func (t StringType) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type UintType struct{}

func (t UintType) AppName() string { return "Int/u" }
func (t UintType) GoName() string  { return "uint" }

func (t UintType) As(a any) uint {
	v, ok := a.(uint)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t UintType) Push(c Calc, val uint) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t UintType) Pop(c Calc) uint {
	return t.As(c.Pop().TypeVal)
}

func (t UintType) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseUint(str, 0, 0)
	return uint(i64), err == nil, nil
}

func (t UintType) Format(a any) string {
	v := t.As(a)
	return strconv.FormatUint(uint64(v), 10)
}

func (t UintType) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Uint8Type struct{}

func (t Uint8Type) AppName() string { return "Int/u8" }
func (t Uint8Type) GoName() string  { return "uint8" }

func (t Uint8Type) As(a any) uint8 {
	v, ok := a.(uint8)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Uint8Type) Push(c Calc, val uint8) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t Uint8Type) Pop(c Calc) uint8 {
	return t.As(c.Pop().TypeVal)
}

func (t Uint8Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseUint(str, 0, 8)
	return uint8(i64), err == nil, nil
}

func (t Uint8Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatUint(uint64(v), 10)
}

func (t Uint8Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Uint16Type struct{}

func (t Uint16Type) AppName() string { return "Int/u16" }
func (t Uint16Type) GoName() string  { return "uint16" }

func (t Uint16Type) As(a any) uint16 {
	v, ok := a.(uint16)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Uint16Type) Push(c Calc, val uint16) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t Uint16Type) Pop(c Calc) uint16 {
	return t.As(c.Pop().TypeVal)
}

func (t Uint16Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseUint(str, 0, 16)
	return uint16(i64), err == nil, nil
}

func (t Uint16Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatUint(uint64(v), 10)
}

func (t Uint16Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Uint32Type struct{}

func (t Uint32Type) AppName() string { return "Int/u32" }
func (t Uint32Type) GoName() string  { return "uint32" }

func (t Uint32Type) As(a any) uint32 {
	v, ok := a.(uint32)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Uint32Type) Push(c Calc, val uint32) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t Uint32Type) Pop(c Calc) uint32 {
	return t.As(c.Pop().TypeVal)
}

func (t Uint32Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseUint(str, 0, 32)
	return uint32(i64), err == nil, nil
}

func (t Uint32Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatUint(uint64(v), 10)
}

func (t Uint32Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
type Uint64Type struct{}

func (t Uint64Type) AppName() string { return "Int/u64" }
func (t Uint64Type) GoName() string  { return "uint64" }

func (t Uint64Type) As(a any) uint64 {
	v, ok := a.(uint64)
	if !ok {
		panic(ErrWrongGoType(t, a))
	}
	return v
}

func (t Uint64Type) Push(c Calc, val uint64) {
	c.Push(Item{TypeVal: val, Type: t})
}

func (t Uint64Type) Pop(c Calc) uint64 {
	return t.As(c.Pop().TypeVal)
}

func (t Uint64Type) Parse(_ coll.State, str string) (any, bool, error) {
	str = PreParseInt(str)
	i64, err := strconv.ParseUint(str, 0, 64)
	return uint32(i64), err == nil, nil
}

func (t Uint64Type) Format(a any) string {
	v := t.As(a)
	return strconv.FormatUint(uint64(v), 10)
}

func (t Uint64Type) Dup(a any) any {
	return a
}

// ----------------------------------------------------------------------------
func parseRat(str string) (*big.Rat, bool, error) {
	r := Rat.New()
	_, ok := r.SetString(str)
	if ok {
		return r, true, nil
	}

	whole := Rat.New()
	defer Rat.Recycle(whole)

	num, denom := BigInt.New(), BigInt.New()
	defer BigInt.Recycle(num, denom)

	s := scan.NewScannerFromString("", str)
	scan.SignedIntRule.Eval(s)
	_, ok = whole.SetString(s.Emit().Val)
	if !ok {
		Rat.Recycle(r)
		return nil, false, nil
	}

	switch s.This {
	case ' ', '_', '-':
		s.Skip()
	default:
		Rat.Recycle(r)
		return nil, false, nil
	}

	scan.IntRule.Eval(s)
	_, ok = num.SetString(s.Emit().Val, 10)
	if !ok {
		Rat.Recycle(r)
		return nil, false, nil
	}

	if s.This == '/' {
		s.Skip()
	} else {
		Rat.Recycle(r)
		return nil, false, nil
	}

	scan.IntRule.Eval(s)
	_, ok = denom.SetString(s.Emit().Val, 10)
	if !ok {
		Rat.Recycle(r)
		return nil, false, nil
	}

	r.SetFrac(num, denom)
	r.Add(r, whole)
	return r, true, nil
}
