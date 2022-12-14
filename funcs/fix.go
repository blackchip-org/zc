package funcs

import (
	"errors"

	"github.com/blackchip-org/zc"
	"github.com/shopspring/decimal"
)

type UnaryFix func(decimal.Decimal) (decimal.Decimal, error)
type BinaryFix func(decimal.Decimal, decimal.Decimal) (decimal.Decimal, error)
type CompareFix func(decimal.Decimal, decimal.Decimal) (bool, error)

func EvalUnaryFix(calc *zc.Calc, fn UnaryFix) error {
	a, err := calc.Stack.PopFix()
	if err != nil {
		return err
	}
	b, err := fn(a)
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Value.FormatFix(b))
	return nil
}

func EvalBinaryFix(calc *zc.Calc, fn BinaryFix) (err error) {
	defer func() {
		if p := recover(); p != nil {
			msg, ok := p.(string)
			if !ok {
				panic(p)
			}
			if msg == "decimal division by 0" {
				err = errors.New("division by zero")
			} else {
				panic(p)
			}
		}
	}()

	a, b, err := calc.Stack.PopFix2()
	if err != nil {
		return
	}

	r, err := fn(a, b)
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Value.FormatFix(r))
	return nil
}

func EvalCompareFix(calc *zc.Calc, fn CompareFix) (err error) {
	a, b, err := calc.Stack.PopFix2()
	if err != nil {
		return
	}
	c, err := fn(a, b)
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Value.FormatBool(c))
	return nil
}
