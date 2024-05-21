package floats

import (
	"fmt"

	"github.com/blackchip-org/zc/v6"
)

var Float64Kind = float64Kind{}

type float64Kind struct{}

func (k float64Kind) Name() string { return "Float64" }

func (k float64Kind) Is(a any) bool {
	switch a.(type) {
	case float64, *float64:
		return true
	}
	return false
}

func (k float64Kind) As(a any) float64 {
	v, ok := a.(float64)
	if !ok {
		panic(fmt.Errorf("expected float64 but got: %v", zc.TypeName(a)))
	}
	return v
}

func (k float64Kind) Dup(a any) any {
	return a
}

func (k float64Kind) Copy(src, dest any) {
	d, ok := dest.(*float64)
	if !ok {
		panic(fmt.Errorf("expected *float64 but got: %v", zc.TypeName(dest)))
	}
	s := k.As(src)
	*d = s
}

func (k float64Kind) To(a any) (any, bool) {
	return nil, false
}
