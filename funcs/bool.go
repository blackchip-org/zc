package funcs

import "github.com/blackchip-org/zc"

type UnaryBool func(bool) (bool, error)
type BinaryBool func(bool, bool) (bool, error)

func EvalUnaryBool(calc *zc.Calc, fn UnaryBool) error {
	a, err := calc.Stack.PopBool()
	if err != nil {
		return err
	}
	b, err := fn(a)
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Val.FormatBool(b))
	return nil
}

func EvalBinaryBool(calc *zc.Calc, fn BinaryBool) error {
	a, b, err := calc.Stack.PopBool2()
	if err != nil {
		return err
	}
	c, err := fn(a, b)
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Val.FormatBool(c))
	return nil
}
