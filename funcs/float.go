package funcs

import (
	"github.com/blackchip-org/zc"
)

type UnaryFloat func(float64) (float64, error)
type UnaryFloatSafe func(float64) float64

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

func EvalUnaryFloatSafe(env *zc.Env, fn UnaryFloatSafe) error {
	a, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	z := fn(a)
	if err != nil {
		return err
	}
	env.Stack.PushFloat(z)
	return nil
}
