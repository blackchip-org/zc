package boolean

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/zc/v6"
)

var BoolKind = boolKind{}

type boolKind struct{}

func (k boolKind) Name() string { return "Bool" }

func (k boolKind) Is(a any) bool {
	switch a.(type) {
	case bool, *bool:
		return true
	}
	return false
}

func (k boolKind) As(a any) bool {
	v, ok := a.(bool)
	if !ok {
		panic(fmt.Errorf("expected bool but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k boolKind) Dup(a any) any {
	return a
}

func (k boolKind) Copy(src, dest any) {
	d, ok := dest.(*bool)
	if !ok {
		panic(fmt.Errorf("expected *bool but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k boolKind) To(a any) (any, bool) {
	switch v := a.(type) {
	case string:
		v = strings.ToLower(v)
		switch v {
		case "true":
			return true, true
		case "false":
			return false, true
		}
	}
	return nil, false
}
