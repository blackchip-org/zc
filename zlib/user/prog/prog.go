package prog

import (
	"fmt"
	"math"
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
	"github.com/blackchip-org/zc/types"
)

var (
	And = zc.FuncGeneric(ops.And)
	Not = zc.FuncGeneric(ops.Not)
	Or  = zc.FuncGeneric(ops.Or)
)

func Bin(env *zc.Env) error {
	v, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	env.Stack.Push(formatWithRadix(v, 2))
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
	env.Stack.PushBigInt(v)
	return nil
}

func Hex(env *zc.Env) error {
	v, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	env.Stack.Push(formatWithRadix(v, 16))
	return nil
}

func Lsh(env *zc.Env) error {
	n, err := env.Stack.PopUint()
	if err != nil {
		return err
	}
	a, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	var z big.Int
	z.Lsh(a, n)
	env.Stack.PushBigInt(&z)
	return nil
}

func Oct(env *zc.Env) error {
	v, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	env.Stack.Push(formatWithRadix(v, 8))
	return nil
}

func Rsh(env *zc.Env) error {
	n, err := env.Stack.PopUint()
	if err != nil {
		return err
	}
	a, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	var z big.Int
	z.Rsh(a, n)
	env.Stack.PushBigInt(&z)
	return nil
}

func Xor(env *zc.Env) error {
	y, err := env.Stack.PopBigInt()
	if y == nil {
		return err
	}
	x, err := env.Stack.PopBigInt()
	if x == nil {
		return err
	}
	var z big.Int
	z.Xor(x, y)
	env.Stack.PushBigInt(&z)
	return nil
}

func formatWithRadix(v *big.Int, radix int) string {
	sign := ""
	if v.Sign() < 0 {
		sign = "-"
	}
	var absV big.Int
	absV.Abs(v)

	switch radix {
	case 16:
		return fmt.Sprintf("%v0x%x", sign, &absV)
	case 8:
		return fmt.Sprintf("%v0o%o", sign, &absV)
	case 2:
		return fmt.Sprintf("%v0b%b", sign, &absV)
	}
	return types.BigInt.Format(v)
}
