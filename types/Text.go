package types

import (
	"fmt"

	"github.com/blackchip-org/zc/v6/state"
)

var Text = textType{}

type textType struct{}

func (t textType) Name() string { return "Text" }

func (t textType) Is(a any) bool {
	switch a.(type) {
	case string, *string:
		return true
	}
	return false
}

func (t textType) As(a any) string {
	switch t := a.(type) {
	case string:
		return t
	case *string:
		return *t
	default:
		panic(fmt.Errorf("expected string but got: %v", GoName(a)))
	}
}

func (t textType) Dup(a any) any {
	return a
}

func (t textType) Copy(src, dest any) {
	d, ok := dest.(*string)
	if !ok {
		panic(fmt.Errorf("expected *string but got: %v", GoName(dest)))
	}
	*d = t.As(src)
}

func (t textType) To(state state.State, a any) (any, bool) {
	return nil, false
}

func (t textType) Format(a any) string {
	d, ok := a.(string)
	if !ok {
		panic(fmt.Errorf("expected string but got: %v", GoName(a)))
	}
	return d
}
