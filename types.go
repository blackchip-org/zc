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

func (t BigIntType) As(item Item) *big.Int {
	val, ok := item.Val.(*big.Int)
	if !ok {
		panic(ErrWrongGoType("*big.Int", item.Val))
	}
	return val
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
		_, ok := bi.SetString(v, 0)
		return bi, String, ok
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

func (t DecimalType) As(item Item) *apd.Decimal {
	val, ok := item.Val.(*apd.Decimal)
	if !ok {
		panic(ErrWrongGoType("*apd.Decimal", item.Val))
	}
	return val
}

func (t DecimalType) Pop(e *OpEnv) *apd.Decimal {
	return t.As(e.Pop())
}

func (t DecimalType) Push(e *OpEnv, v *apd.Decimal) {
	e.PushVal(v)
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

func (t DecimalSSType) As(item Item) decimal.Decimal {
	val, ok := item.Val.(decimal.Decimal)
	if !ok {
		panic(ErrWrongGoType("decimal.Decimal", item.Val))
	}
	return val
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
		d, err := decimal.NewFromString(v)
		return d, String, err == nil
	}
	return nil, nil, false
}

func (t DecimalSSType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type BigFloatType struct{}

func (t BigFloatType) Name() string { return "Float" }

func (t BigFloatType) As(item Item) *big.Float {
	val, ok := item.Val.(*big.Float)
	if !ok {
		panic(ErrWrongGoType("*big.Float", item.Val))
	}
	return val
}

func (t BigFloatType) Pop(e *OpEnv) *big.Float {
	conf := state.ForConf(e.State)
	bf := t.As(e.Pop())
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

func (t Float64Type) Name() string { return "float/64" }

func (t Float64Type) From(src any) (any, Type, bool) {
	return nil, nil, false
}

func (t Float64Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type IntType struct{}

func (t IntType) Name() string { return "int" }

func (t IntType) From(src any) (any, Type, bool) {
	return nil, nil, false
}

func (t IntType) Recycle(v any) {}

// ----------------------------------------------------------------------------

type Int32Type struct{}

func (t Int32Type) Name() string { return "int/32" }

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
		i32, err := strconv.ParseInt(v, 0, 32)
		return int32(i32), String, err == nil
	}
	return nil, nil, false
}

func (t Int32Type) Recycle(v any) {}

// ----------------------------------------------------------------------------

type StringType struct{}

func (t StringType) Name() string { return "Text" }

func (t StringType) As(item Item) string {
	val, ok := item.Val.(string)
	if !ok {
		panic(ErrWrongGoType("string", item.Val))
	}
	return val
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

func (t UintType) Name() string { return "int/u" }

func (t UintType) As(item Item) uint {
	val, ok := item.Val.(uint)
	if !ok {
		panic(ErrWrongGoType("uint", item.Val))
	}
	return val
}

func (t UintType) Push(e *OpEnv, ui uint) {
	e.PushVal(ui)
}

func (t UintType) From(src any) (any, Type, bool) {
	switch v := src.(type) {
	case uint:
		return v, t, true
	case string:
		ui, err := strconv.ParseUint(v, 0, 0)
		return uint(ui), String, err == nil
	}
	return nil, nil, false
}

func (t UintType) Recycle(v any) {}
