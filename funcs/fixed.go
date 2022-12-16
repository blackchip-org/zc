package funcs

import (
	"errors"

	"github.com/blackchip-org/zc"
	"github.com/shopspring/decimal"
)

type UnaryFixed func(decimal.Decimal) (decimal.Decimal, error)
type BinaryFixed func(decimal.Decimal, decimal.Decimal) (decimal.Decimal, error)
type CompareFixed func(decimal.Decimal, decimal.Decimal) (bool, error)

func EvalUnaryFixed(calc *zc.Calc, fn UnaryFixed) error {
	a, err := calc.Stack.PopFixed()
	if err != nil {
		return err
	}
	b, err := fn(a)
	if err != nil {
		return err
	}
	calc.Stack.PushFixed(b)
	return nil
}

func EvalBinaryFixed(calc *zc.Calc, fn BinaryFixed) (err error) {
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

	a, b, err := calc.Stack.PopFixed2()
	if err != nil {
		return
	}

	r, err := fn(a, b)
	if err != nil {
		return err
	}
	calc.Stack.PushFixed(r)
	return nil
}

func EvalCompareFixed(calc *zc.Calc, fn CompareFixed) (err error) {
	a, b, err := calc.Stack.PopFixed2()
	if err != nil {
		return
	}
	c, err := fn(a, b)
	if err != nil {
		return err
	}
	calc.Stack.PushBool(c)
	return nil
}
