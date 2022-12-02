package native

import "github.com/blackchip-org/zc"

func Eq(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	b, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	r := a == b
	calc.Stack.Push(zc.FormatBool(r))
	return nil
}

func Gt(calc *zc.Calc) error {
	b, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	bi, err := zc.ParseBigInt(b)
	if err != nil {
		return err
	}

	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	ai, err := zc.ParseBigInt(a)
	if err != nil {
		return err
	}

	r := ai.Cmp(bi) > 0
	calc.Stack.Push(zc.FormatBool(r))
	return nil
}

func Lt(calc *zc.Calc) error {
	b, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	bi, err := zc.ParseBigInt(b)
	if err != nil {
		return err
	}

	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	ai, err := zc.ParseBigInt(a)
	if err != nil {
		return err
	}

	r := ai.Cmp(bi) < 0
	calc.Stack.Push(zc.FormatBool(r))
	return nil
}

func Not(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	ab, err := zc.ParseBool(a)
	if err != nil {
		return err
	}

	rb := !ab
	calc.Stack.Push(zc.FormatBool(rb))
	return nil
}
