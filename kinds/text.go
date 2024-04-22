package kinds

import (
	"fmt"

	"github.com/blackchip-org/zc/v6/errors"
)

var (
	Text = TextKind{}
)

type TextKind struct{}

func (k TextKind) String() string {
	return "Text"
}

func (k TextKind) Is(a any) bool {
	switch a.(type) {
	case string:
		return true
	case *string:
		return true
	}
	return false
}

func (k TextKind) As(a any) string {
	v, ok := a.(string)
	if !ok {
		panic(errors.NewUnexpectedType("string", a))
	}
	return v
}

func (k TextKind) Dup(a any) any {
	return a
}

func (k TextKind) Copy(src any, dest any) {
	d, ok := dest.(*string)
	if !ok {
		panic(errors.NewUnexpectedType("*string", dest))
	}
	*d = k.As(src)
}

func (k TextKind) To(a any) (any, bool) {
	return fmt.Sprintf("%v", a), true
}
