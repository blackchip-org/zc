package zlib

import (
	"strings"

	"github.com/blackchip-org/zc"
)

func Len(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	r := len(a)
	env.Stack.Push(env.Calc.FormatInt(r))
	return nil
}

func StartsWith(env *zc.Env) error {
	prefix, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	str, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	r := strings.HasPrefix(str, prefix)
	env.Stack.PushBool(r)
	return nil
}
