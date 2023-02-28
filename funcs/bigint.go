package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc"
)

type UnaryBigInt func(*big.Int, *big.Int) error
type BinaryBigInt func(*big.Int, *big.Int, *big.Int) error
type CompareBigInt func(*big.Int, *big.Int) (bool, error)

func EvalUnaryBigInt(env *zc.Env, fn UnaryBigInt) error {
	a, attrs, err := env.Stack.PopBigIntWithAttrs()
	if err != nil {
		return err
	}
	var c big.Int
	if err := fn(&c, a); err != nil {
		return err
	}
	env.Stack.PushBigIntWithAttrs(&c, attrs)
	return nil
}

func EvalBinaryBigInt(env *zc.Env, fn BinaryBigInt) error {
	x, y, err := env.Stack.Pop2()
	if err != nil {
		return err
	}

	a, err := env.Calc.ParseBigInt(x)
	if err != nil {
		return err
	}

	b, err := env.Calc.ParseBigInt(y)
	if err != nil {
		return err
	}

	var c big.Int
	if err := fn(&c, a, b); err != nil {
		return err
	}

	attrs := zc.ParseFormatAttrs(x, y)
	attrs.ApplyLayout = env.Calc.AutoFormat
	env.Stack.PushBigIntWithAttrs(&c, attrs)
	return nil
}

func EvalCompareBigInt(env *zc.Env, fn CompareBigInt) error {
	a, b, err := env.Stack.PopBigInt2()
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
