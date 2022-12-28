package funcs

import "github.com/blackchip-org/zc"

type UnaryBool func(bool) (bool, error)
type BinaryBool func(bool, bool) (bool, error)

func EvalUnaryBool(env *zc.Env, fn UnaryBool) error {
	a, err := env.Stack.PopBool()
	if err != nil {
		return err
	}
	b, err := fn(a)
	if err != nil {
		return err
	}
	env.Stack.PushBool(b)
	return nil
}

func EvalBinaryBool(env *zc.Env, fn BinaryBool) error {
	a, b, err := env.Stack.PopBool2()
	if err != nil {
		return err
	}
	c, err := fn(a, b)
	if err != nil {
		return err
	}
	env.Stack.PushBool(c)
	return nil
}
