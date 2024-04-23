package ints

import "github.com/blackchip-org/zc/v6/errors"

var IntArchKind = intArchKind{}

type intArchKind struct{}

func (k intArchKind) Name() string { return "IntArch" }

func (k intArchKind) Is(a any) bool {
	switch a.(type) {
	case int:
		return true
	}
	return false
}

func (k intArchKind) As(a any) int {
	v, ok := a.(int)
	if !ok {
		panic(errors.NewUnexpectedType("int", a))
	}
	return v
}

func (k intArchKind) Dup(a any) any {
	return a
}

func (k intArchKind) Copy(src, dest any) {
	d, ok := dest.(*int)
	if !ok {
		panic(errors.NewUnexpectedType("*int", dest))
	}
	s := k.As(src)
	*d = s
}

func (k intArchKind) To(a any) (any, bool) {
	return nil, false
}
