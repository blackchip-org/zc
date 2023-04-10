package zc_

import "github.com/blackchip-org/zc"

func Version(env *zc.Env) error {
	env.Stack.Push(zc.Version)
	return nil
}
