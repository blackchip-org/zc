package zlib

import (
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func opAndBitwise(z *big.Int, a *big.Int, b *big.Int) error { z.And(a, b); return nil }
func opNotBitwise(z *big.Int, a *big.Int) error             { z.Not(a); return nil }
func opOrBitwise(z *big.Int, a *big.Int, b *big.Int) error  { z.Or(a, b); return nil }
func opXor(z *big.Int, a *big.Int, b *big.Int) error        { z.Xor(a, b); return nil }

func AndBitwise(calc *zc.Calc) error { return funcs.EvalBinaryBigInt(calc, opAndBitwise) }
func NotBitwise(calc *zc.Calc) error { return funcs.EvalUnaryBigInt(calc, opNotBitwise) }
func OrBitwise(calc *zc.Calc) error  { return funcs.EvalBinaryBigInt(calc, opOrBitwise) }
func Xor(calc *zc.Calc) error        { return funcs.EvalBinaryBigInt(calc, opXor) }

func Bin(calc *zc.Calc) error {
	v, err := calc.Stack.PopBigInt()
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Val.FormatBigIntBase(v, 2))
	return nil
}

func Bit(calc *zc.Calc) error {
	i, err := calc.Stack.PopInt()
	if err != nil {
		return err
	}
	a, err := calc.Stack.PopBigInt()
	if err != nil {
		return err
	}
	bit := a.Bit(i)
	calc.Stack.Push(calc.Val.FormatInt(int(bit)))
	return nil
}

func Dec(calc *zc.Calc) error {
	v, err := calc.Stack.PopBigInt()
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Val.FormatBigIntBase(v, 10))
	return nil
}

func Hex(calc *zc.Calc) error {
	v, err := calc.Stack.PopBigInt()
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Val.FormatBigIntBase(v, 16))
	return nil
}

func LenBitwise(calc *zc.Calc) error {
	a, err := calc.Stack.PopBigInt()
	if err != nil {
		return err
	}
	bitLen := a.BitLen()
	calc.Stack.Push(calc.Val.FormatInt(bitLen))
	return nil
}

func Lsh(calc *zc.Calc) error {
	n, err := calc.Stack.PopUint()
	if err != nil {
		return err
	}
	a, r, err := calc.Stack.PopBigIntWithRadix()
	if err != nil {
		return err
	}
	var z big.Int
	z.Lsh(a, n)
	calc.Stack.Push(calc.Val.FormatBigIntBase(&z, r))
	return nil
}

func Oct(calc *zc.Calc) error {
	v, err := calc.Stack.PopBigInt()
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Val.FormatBigIntBase(v, 8))
	return nil
}

func Rsh(calc *zc.Calc) error {
	n, err := calc.Stack.PopUint()
	if err != nil {
		return err
	}
	a, r, err := calc.Stack.PopBigIntWithRadix()
	if err != nil {
		return err
	}
	var z big.Int
	z.Rsh(a, n)
	calc.Stack.Push(calc.Val.FormatBigIntBase(&z, r))
	return nil
}
