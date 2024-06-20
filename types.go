package zc

import (
	"math/big"
	"strconv"

	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/cockroachdb/apd/v3"
	"github.com/shopspring/decimal"
)

var (
	Any       = anyType{}
	BigInt    = BigIntType{}
	BigFloat  = BigFloatType{}
	Decimal   = DecimalType{}
	DecimalSS = DecimalSSType{}
	Float64   = Float64Type{}
	Int       = IntType{}
	Int32     = Int32Type{}
	String    = StringType{}
	Uint      = UintType{}
)

var (
	poolSize     = 8
	bigIntPool   = NewPool[big.Int](poolSize)
	decimalPool  = NewPool[apd.Decimal](poolSize)
	bigFloatPool = NewPool[big.Float](poolSize)
)

type Type interface {
	Name() string
	From(any) (any, Type, bool)
	Recycle(any)
}

// ----------------------------------------------------------------------------

type anyType struct{}

func (t anyType) Name() string { return "Any" }

func (t anyType) From(src any) (any, Type, bool) {
	return src, Any, true
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
	}
	return nil, nil, false
}

func (t BigIntType) New() *big.Int {
	return bigIntPool.New()
}

func (t BigIntType) Recycle(v any) {
	bigIntPool.Recycle(v.(*big.Int))
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
	e.PushVal(v)
}

func (t Float64Type) From(src any) (any, Type, bool) {
	return nil, nil, false
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

func (t IntType) Recycle(v any) {}

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

func (t Int32Type) Recycle(v any) {}

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

func (t UintType) Recycle(v any) {}
