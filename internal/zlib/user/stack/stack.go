package stack

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
	z := env.Stack.Items()[index]
	env.Stack.Clear()
	env.Stack.Push(z)
	return nil
}

func Reverse(env *zc.Env) error {
	items := env.Stack.ItemsReversed()
	env.Stack.Clear()
	for _, item := range items {
		env.Stack.Push(item)
	}
	return nil
}

func N(env *zc.Env) error {
	n := env.Stack.Len()
	env.Stack.PushInt(n)
	return nil
}
