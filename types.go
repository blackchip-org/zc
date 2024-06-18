package zc

import (
	"math/big"
)

var (
	Any    = anyType{}
	BigInt = BigIntType{}
	Int    = IntType{}
	String = StringType{}
)

var (
	bigIntPool = NewPool[big.Int](8)
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
