package zlib

import (
	"github.com/blackchip-org/zc"
)

func Eval(env *zc.Env) error {
	node, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	return zc.Eval(env, "<eval>", []byte(node))
}
