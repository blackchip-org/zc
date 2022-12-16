package funcs

import "github.com/blackchip-org/zc"

type CompareStr func(string, string) (bool, error)

func EvalCompareStr(calc *zc.Calc, fn CompareStr) error {
	a, b, err := calc.Stack.Pop2()
	if err != nil {
		return err
	}
	c, err := fn(a, b)
	if err != nil {
		return err
	}
	calc.Stack.PushBool(c)
	return nil
}
