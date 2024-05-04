package floatb

import (
	"math/big"

	"github.com/blackchip-org/zc/v6/errors"
)

var FloatKind = floatKind{}

type floatKind struct{}

func (k floatKind) Name() string { return "Float" }

func (k floatKind) Is(a any) bool {
	switch a.(type) {
	case *big.Float:
		return true
	}
	return false
}

func (k floatKind) As(a any) *big.Float {
	v, ok := a.(*big.Float)
	if !ok {
		panic(errors.UnexpectedType("*big.Float", a))
	}
	return v
}

func (k floatKind) Dup(a any) any {
	var r big.Float
	v := k.As(a)
	r.Set(v)
	return &r
}

func (k floatKind) Copy(src, dest any) {
	d, ok := dest.(*big.Float)
	if !ok {
		panic(errors.UnexpectedType("*big.Float", dest))
	}
	s := k.As(src)
	d.Set(s)
}

func (k floatKind) To(a any) (any, bool) {
	switch v := a.(type) {
	case float64:
		return big.NewFloat(v), true
	}
	return nil, false
}
