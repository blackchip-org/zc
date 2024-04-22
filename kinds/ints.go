package kinds

import (
	"math/big"

	"github.com/blackchip-org/zc/v6/errors"
)

var (
	Int     = IntKind{}
	IntArch = IntArchKind{}
)

type IntKind struct{}

func (k IntKind) String() string { return "Int" }

func (k IntKind) Is(a any) bool {
	switch a.(type) {
	case *big.Int:
		return true
	}
	return false
}

func (k IntKind) As(a any) *big.Int {
	v, ok := a.(*big.Int)
	if !ok {
		panic(errors.NewUnexpectedType("*big.Int", a))
	}
	return v
}

func (k IntKind) Dup(a any) any {
	var r big.Int
	v := k.As(a)
	r.Set(v)
	return &r
}

func (k IntKind) Copy(src, dest any) {
	d, ok := dest.(*big.Int)
	if !ok {
		panic(errors.NewUnexpectedType("*big.Int", dest))
	}
	s := k.As(src)
	d.Set(s)
}

func (k IntKind) To(a any) (any, bool) {
	switch v := a.(type) {
	case int:
		return big.NewInt(int64(v)), true
	}
	return nil, false
}

// ----------------------------------------------------------------------------
type IntArchKind struct{}

func (k IntArchKind) String() string { return "IntArch" }

func (k IntArchKind) Is(a any) bool {
	switch a.(type) {
	case int:
		return true
	}
	return false
}

func (k IntArchKind) As(a any) int {
	v, ok := a.(int)
	if !ok {
		panic(errors.NewUnexpectedType("int", a))
	}
	return v
}

func (k IntArchKind) Dup(a any) any {
	return a
}

func (k IntArchKind) Copy(src, dest any) {
	d, ok := dest.(*int)
	if !ok {
		panic(errors.NewUnexpectedType("*int", dest))
	}
	s := k.As(src)
	*d = s
}

func (k IntArchKind) To(a any) (any, bool) {
	return nil, false
}
