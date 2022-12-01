package native

import "github.com/blackchip-org/zc"

func Eq(calc *zc.Calc) error {
	a, err := calc.Stack().Pop()
	if err != nil {
		return err
	}

	b, err := calc.Stack().Pop()
	if err != nil {
		return err
	}

	r := a == b
	calc.Stack().Push(zc.FormatBool(r))
	return nil
}
