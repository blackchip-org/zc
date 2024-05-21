package floatb

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6"
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
		panic(fmt.Errorf("expected *big.Float but got %v", zc.TypeName(a)))
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
		panic(fmt.Errorf("expected *big.Float but got %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	d.Set(s)
}

func (k floatKind) To(a any) (any, bool) {
	switch v := a.(type) {
	case float64:
		return big.NewFloat(v), true
	case string:
		var f big.Float
		_, ok := f.SetString(v)
		return &f, ok
	}
	return nil, false
}
