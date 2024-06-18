package zc

import (
	"math/big"
)

var (
	Any    = anyType{}
	BigInt = BigIntType{}
	Int    = IntType{}
)

var (
	bigIntPool = NewPool[big.Int](8)
)

type anyType struct{}

func (t anyType) From(src any) (any, Type, bool) {
	return src, Any, true
}

func (t anyType) Recycle(any) {}

type BigIntType struct{}

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
	}
	return nil, nil, false
}

func (t BigIntType) New() *big.Int {
	return bigIntPool.New()
}

func (t BigIntType) Recycle(v any) {
	bigIntPool.Recycle(v.(*big.Int))
}

type IntType struct{}

func (t IntType) From(src any) (any, Type, bool) {
	return nil, nil, false
}

func (t IntType) Recycle(v any) {}
