package intb

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

var IntKind = intKind{}

type intKind struct{}

func (k intKind) Name() string { return "Int" }

func (k intKind) Is(a any) bool {
	switch a.(type) {
	case *big.Int:
		return true
	}
	return false
}

func (k intKind) As(a any) *big.Int {
	v, ok := a.(*big.Int)
	if !ok {
		panic(fmt.Errorf("expected *big.Int but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k intKind) Dup(a any) any {
	var r big.Int
	v := k.As(a)
	r.Set(v)
	return &r
}

func (k intKind) Copy(src, dest any) {
	d, ok := dest.(*big.Int)
	if !ok {
		panic(fmt.Errorf("expected *big.Int but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	d.Set(s)
}

func (k intKind) To(a any) (any, bool) {
	switch v := a.(type) {
	case int:
		return big.NewInt(int64(v)), true
	case string:
		var i big.Int
		_, ok := i.SetString(v, 0)
		return &i, ok
	}
	return nil, false
}
