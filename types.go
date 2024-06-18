package zc

import (
	"math/big"

	"github.com/cockroachdb/apd/v3"
)

var (
	Any     = anyType{}
	BigInt  = BigIntType{}
	Decimal = DecimalType{}
	Float64 = Float64Type{}
	Int     = IntType{}
	String  = StringType{}
)

var (
	poolSize    = 8
	bigIntPool  = NewPool[big.Int](poolSize)
	decimalPool = NewPool[apd.Decimal](poolSize)
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

type StringType struct{}

func (t StringType) Name() string { return "Text" }

func (t StringType) From(src any) (any, Type, bool) {
	return nil, nil, false
}

func (t StringType) Recycle(v any) {}
