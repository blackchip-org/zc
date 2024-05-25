package types

import "fmt"

var Str = strType{}

type strType struct{}

func (t strType) Name() string { return "Str" }

func (t strType) Is(a any) bool {
	switch a.(type) {
	case string, *string:
		return true
	}
	return false
}

func (t strType) As(a any) string {
	switch t := a.(type) {
	case string:
		return t
	case *string:
		return *t
	default:
		panic(fmt.Errorf("expected string but got: %v", GoName(a)))
	}
}

func (t strType) Dup(a any) any {
	return a
}

func (t strType) Copy(src, dest any) {
	d, ok := dest.(*string)
	if !ok {
		panic(fmt.Errorf("expected *string but got: %v", GoName(dest)))
	}
	*d = t.As(src)
}

func (t strType) To(a any) (any, bool) {
	return nil, false
}
