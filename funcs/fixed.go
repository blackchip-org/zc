package funcs

import (
	"errors"

	"github.com/blackchip-org/zc"
	"github.com/shopspring/decimal"
)

type UnaryFixed func(decimal.Decimal) (decimal.Decimal, error)
type BinaryFixed func(decimal.Decimal, decimal.Decimal) (decimal.Decimal, error)
type CompareFixed func(decimal.Decimal, decimal.Decimal) (bool, error)

func EvalUnaryFixed(env *zc.Env, fn UnaryFixed) error {
	a, err := env.Stack.PopFixed()
	if err != nil {
		return err
	}
	b, err := fn(a)
	if err != nil {
		return err
	}
	env.Stack.PushFixed(b)
	return nil
}

func EvalBinaryFixed(env *zc.Env, fn BinaryFixed) (err error) {
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
	a, err := env.Calc.ParseFixed(sa)
	if err != nil {
		return
	}
	b, err := env.Calc.ParseFixed(sb)
	if err != nil {
		return
	}

	r, err := fn(a, b)
	if err != nil {
		return err
	}

	attrs := zc.ParseFormatAttrs(sa, sb)
	env.Stack.PushFixedWithAttrs(r, attrs)
	return nil
}

func EvalCompareFixed(env *zc.Env, fn CompareFixed) (err error) {
	a, b, err := env.Stack.PopFixed2()
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
