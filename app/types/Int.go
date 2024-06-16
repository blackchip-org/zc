package types

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

var (
	BigInt = BigIntType{}
)

type BigIntType struct{}

func (t BigIntType) As(item *zc.Item) *big.Int {
	val, ok := item.Val.(*big.Int)
	if !ok {
		panic(zc.ErrWrongGoType("*big.Int", item.Val))
	}
	return val
}

func (t BigIntType) Recycle(r zc.Recycler) (*big.Int, bool) {
	v, ok := r.Recycle()
	if !ok {
		return nil, false
	}
	if bi, ok := v.Val.(*big.Int); ok {
		return bi, true
	}
	return nil, false
}

func (t BigIntType) New(r zc.Recycler) *big.Int {
	bi, ok := t.Recycle(r)
	if !ok {
		bi = new(big.Int)
		return bi
	}
	return bi
}

func (t BigIntType) From(r zc.Recycler, src *zc.Item) (any, bool) {
	switch v := src.Val.(type) {
	case *big.Int:
		return v, true
	case int:
		bi := t.New(r)
		bi.SetInt64(int64(v))
		return bi, true
	}
	return nil, false
}

type IntType struct{}
