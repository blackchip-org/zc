package funcs

import (
	"errors"

	"github.com/blackchip-org/zc"
	"github.com/shopspring/decimal"
)

type UnaryDecimal func(decimal.Decimal) (decimal.Decimal, error)
type BinaryDecimal func(decimal.Decimal, decimal.Decimal) (decimal.Decimal, error)
type CompareDecimal func(decimal.Decimal, decimal.Decimal) (bool, error)

func EvalUnaryDecimal(env *zc.Env, fn UnaryDecimal) error {
	a, err := env.Stack.PopDecimal()
	if err != nil {
		return err
	}
	b, err := fn(a)
	if err != nil {
		return err
	}
	env.Stack.PushDecimal(b)
	return nil
}

func EvalBinaryDecimal(env *zc.Env, fn BinaryDecimal) (err error) {
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

	sa, sb, err := env.Stack.Pop2()
	if err != nil {
		return
	}
	a, err := env.Calc.ParseDecimal(sa)
	if err != nil {
		return
	}
	b, err := env.Calc.ParseDecimal(sb)
	if err != nil {
		return
	}

	r, err := fn(a, b)
	if err != nil {
		return err
	}

	attrs := zc.ParseFormatAttrs(sa, sb)
	attrs.ApplyLayout = env.Calc.AutoFormat
	env.Stack.PushDecimalWithAttrs(r, attrs)
	return nil
}

func EvalCompareDecimal(env *zc.Env, fn CompareDecimal) (err error) {
	a, b, err := env.Stack.PopDecimal2()
	if err != nil {
		return
	}
	c, err := fn(a, b)
	if err != nil {
		return err
	}
	env.Stack.PushBool(c)
	return nil
}
