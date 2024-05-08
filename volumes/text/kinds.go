package text

import (
	"fmt"

	"github.com/blackchip-org/zc/v6"
)

var TextKind = textKind{}

type textKind struct{}

func (k textKind) Name() string {
	return "Text"
}

func (k textKind) Is(a any) bool {
	switch a.(type) {
	case string:
		return true
	case *string:
		return true
	}
	return false
}

func (k textKind) As(a any) string {
	v, ok := a.(string)
	if !ok {
		panic(fmt.Errorf("expected string but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k textKind) Dup(a any) any {
	return a
}

func (k textKind) Copy(src any, dest any) {
	d, ok := dest.(*string)
	if !ok {
		panic(fmt.Errorf("expected *string but got: %v", zc.TypeName(dest)))
	}
	*d = k.As(src)
}

func (k textKind) To(a any) (any, bool) {
	return fmt.Sprintf("%v", a), true
}
