package types

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

var (
	BigInt = BigIntType{}
	Int    = IntType{}
)

var (
	bigIntPool = zc.NewPool[big.Int](8)
)

type BigIntType struct{}

func (t BigIntType) As(item *zc.Item) *big.Int {
	val, ok := item.Val.(*big.Int)
	if !ok {
		panic(zc.ErrWrongGoType("*big.Int", item.Val))
	}
	return val
}

func (t BigIntType) From(src any) (any, zc.Type, bool) {
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

func (t BigIntType) Recycle(v any) {
	bigIntPool.Recycle(v.(*big.Int))
}

type IntType struct{}

func (t IntType) From(src any) (any, zc.Type, bool) {
	return nil, nil, false
}

func (t IntType) Recycle(v any) {}
