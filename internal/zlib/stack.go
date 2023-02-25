package zlib

import (
	"fmt"

	"github.com/blackchip-org/zc"
)

func At(env *zc.Env) error {
	index, err := env.Stack.PopInt()
	if err != nil {
		return err
	}
	if index < 0 || index >= env.Stack.Len() {
		return fmt.Errorf("invalid stack index '%v'", index)
	}
	env.Stack.Push(env.Stack.Items()[env.Stack.Len()-index-1])
	return nil
}

func N(env *zc.Env) error {
	n := env.Stack.Len()
	env.Stack.PushInt(n)
	return nil
}
