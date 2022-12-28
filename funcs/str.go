package funcs

import "github.com/blackchip-org/zc"

type CompareStr func(string, string) (bool, error)

func EvalCompareStr(env *zc.Env, fn CompareStr) error {
	a, b, err := env.Stack.Pop2()
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
