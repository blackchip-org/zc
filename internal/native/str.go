package native

import "github.com/blackchip-org/zc"

func Len(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	r := len(a)
	calc.Stack.Push(zc.FormatInt(r))
	return nil
}
