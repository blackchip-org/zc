package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc"
)

type UnaryBigInt func(*big.Int, *big.Int) error
type BinaryBigInt func(*big.Int, *big.Int, *big.Int) error
type CompareBigInt func(*big.Int, *big.Int) (bool, error)

func EvalUnaryBigInt(env *zc.Env, fn UnaryBigInt) error {
	a, r, err := env.Stack.PopBigIntWithRadix()
	if err != nil {
		return err
	}
	var c big.Int
	if err := fn(&c, a); err != nil {
		return err
	}
	env.Stack.PushBigIntWithRadix(&c, r)
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

	radix := resolveRadix(zc.ParseRadix(x), zc.ParseRadix(y))
	env.Stack.PushBigIntWithRadix(&c, radix)
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

func resolveRadix(rx int, ry int) int {
	switch {
	case rx == 16 || ry == 16:
		return 16
	case rx == 8 || ry == 8:
		return 8
	case rx == 2 || ry == 2:
		return 2
	}
	return 10
}
