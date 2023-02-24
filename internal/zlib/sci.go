package zlib

import (
	"strconv"

	"github.com/blackchip-org/zc"
)

func ScientificNotation(env *zc.Env) error {
	f, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	env.Stack.Push(strconv.FormatFloat(f, 'e', -1, 64))
	return nil
}
