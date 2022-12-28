package zlib

import "github.com/blackchip-org/zc"

func N(env *zc.Env) error {
	n := env.Stack.Len()
	env.Stack.PushInt(n)
	return nil
}
