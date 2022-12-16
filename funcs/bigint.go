package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc"
)

type UnaryBigInt func(*big.Int, *big.Int) error
type BinaryBigInt func(*big.Int, *big.Int, *big.Int) error
type CompareBigInt func(*big.Int, *big.Int) (bool, error)

func EvalUnaryBigInt(calc *zc.Calc, fn UnaryBigInt) error {
	a, r, err := calc.Stack.PopBigIntWithRadix()
	if err != nil {
		return err
	}
	var c big.Int
	if err := fn(&c, a); err != nil {
		return err
	}
	calc.Stack.Push(calc.Val.FormatBigIntBase(&c, r))
	return nil
}

func EvalBinaryBigInt(calc *zc.Calc, fn BinaryBigInt) error {
	x, y, err := calc.Stack.Pop2()
	if err != nil {
		return err
	}

	a, err := calc.Val.ParseBigInt(x)
	if err != nil {
		return err
	}

	b, err := calc.Val.ParseBigInt(y)
	if err != nil {
		return err
	}

	var c big.Int
	if err := fn(&c, a, b); err != nil {
		return err
	}

	radix := resolveRadix(zc.ParseRadix(x), zc.ParseRadix(y))
	calc.Stack.Push(calc.Val.FormatBigIntBase(&c, radix))
	return nil
}

func EvalCompareBigInt(calc *zc.Calc, fn CompareBigInt) error {
	a, b, err := calc.Stack.PopBigInt2()
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
