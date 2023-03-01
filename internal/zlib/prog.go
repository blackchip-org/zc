package zlib

import (
	"math"
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func opAndBitwise(z *big.Int, a *big.Int, b *big.Int) error { z.And(a, b); return nil }
func opNotBitwise(z *big.Int, a *big.Int) error             { z.Not(a); return nil }
func opOrBitwise(z *big.Int, a *big.Int, b *big.Int) error  { z.Or(a, b); return nil }
func opXor(z *big.Int, a *big.Int, b *big.Int) error        { z.Xor(a, b); return nil }

func AndBitwise(env *zc.Env) error { return funcs.EvalBinaryBigInt(env, opAndBitwise) }
func NotBitwise(env *zc.Env) error { return funcs.EvalUnaryBigInt(env, opNotBitwise) }
func OrBitwise(env *zc.Env) error  { return funcs.EvalBinaryBigInt(env, opOrBitwise) }
func Xor(env *zc.Env) error        { return funcs.EvalBinaryBigInt(env, opXor) }

func Bin(env *zc.Env) error {
	v, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	env.Stack.PushBigIntWithAttrs(v, zc.FormatAttrs{Radix: 2})
	return nil
}

func Bit(env *zc.Env) error {
	i, err := env.Stack.PopInt()
	if err != nil {
		return err
	}
	a, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	bit := a.Bit(i)
	env.Stack.PushUint(bit)
	return nil
}

func Bits(env *zc.Env) error {
	a, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	bitLen := a.BitLen()
	env.Stack.PushInt(bitLen)
	return nil
}

func Bytes(env *zc.Env) error {
	a, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	bitLen := a.BitLen()
	z := int(math.Ceil(float64(bitLen) / 8.0))
	env.Stack.PushInt(z)
	return nil
}

func Dec(env *zc.Env) error {
	v, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	env.Stack.PushBigIntWithAttrs(v, zc.FormatAttrs{Radix: 10})
	return nil
}

func Hex(env *zc.Env) error {
	v, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	env.Stack.PushBigIntWithAttrs(v, zc.FormatAttrs{Radix: 16})
	return nil
}

func Lsh(env *zc.Env) error {
	n, err := env.Stack.PopUint()
	if err != nil {
		return err
	}
	a, attrs, err := env.Stack.PopBigIntWithAttrs()
	if err != nil {
		return err
	}
	var z big.Int
	z.Lsh(a, n)
	env.Stack.PushBigIntWithAttrs(&z, attrs)
	return nil
}

func Oct(env *zc.Env) error {
	v, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	env.Stack.PushBigIntWithAttrs(v, zc.FormatAttrs{Radix: 8})
	return nil
}

func Rsh(env *zc.Env) error {
	n, err := env.Stack.PopUint()
	if err != nil {
		return err
	}
	a, attrs, err := env.Stack.PopBigIntWithAttrs()
	if err != nil {
		return err
	}
	var z big.Int
	z.Rsh(a, n)
	env.Stack.PushBigIntWithAttrs(&z, attrs)
	return nil
}
