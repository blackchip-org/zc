package types

import (
	"fmt"

	"github.com/blackchip-org/zc/v6/calc/state"
)

var Val = valType{}

type valType struct{}

func (t valType) Name() string { return "Val" }

func (t valType) Is(v any) bool {
	return v != nil
}

func (t valType) Dup(any) any {
	panic("type Val cannot be duplicated")
}

func (t valType) Copy(any, any) {
	panic("type Val cannot be copied")
}

func (t valType) To(state state.State, a any) (any, bool) {
	return a, false
}

func (t valType) Format(a any) string {
	return fmt.Sprintf("%v", a)
}
