package funcs

import (
	"github.com/blackchip-org/zc"
)

type UnaryFloat func(float64) (float64, error)

func EvalUnaryFloat(env *zc.Env, fn UnaryFloat) error {
	a, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	z, err := fn(a)
	if err != nil {
		return err
	}
	env.Stack.PushFloat(z)
	return nil
}
