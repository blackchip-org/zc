package math_

import (
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func bigIntMod(z *big.Int, a *big.Int, b *big.Int) {
	var q big.Int
	q.DivMod(a, b, z)
}

func bigIntSign(z *big.Int, a *big.Int) {
	z.SetInt64(int64(a.Sign()))
}

func opAbsBigInt(z *big.Int, a *big.Int) error             { z.Abs(a); return nil }
func opAddBigInt(z *big.Int, a *big.Int, b *big.Int) error { z.Add(a, b); return nil }
func opCeilBigInt(z *big.Int, a *big.Int) error            { z.Set(a); return nil }
func opDivBigInt(z *big.Int, a *big.Int, b *big.Int) error { z.Div(a, b); return nil }
func opFloorBigInt(z *big.Int, a *big.Int) error           { z.Set(a); return nil }
func opModBigInt(z *big.Int, a *big.Int, b *big.Int) error { bigIntMod(z, a, b); return nil }
func opMulBigInt(z *big.Int, a *big.Int, b *big.Int) error { z.Mul(a, b); return nil }
func opNegBigInt(z *big.Int, a *big.Int) error             { z.Neg(a); return nil }
func opPowBigInt(z *big.Int, a *big.Int, b *big.Int) error { z.Exp(a, b, nil); return nil }
func opRemBigInt(z *big.Int, a *big.Int, b *big.Int) error { z.Rem(a, b); return nil }
func opSignBigInt(z *big.Int, a *big.Int) error            { bigIntSign(z, a); return nil }
func opSubBigInt(z *big.Int, a *big.Int, b *big.Int) error { z.Sub(a, b); return nil }

func AbsBigInt(env *zc.Env) error     { return funcs.EvalUnaryBigInt(env, opAbsBigInt) }
func AddBigInt(env *zc.Env) error     { return funcs.EvalBinaryBigInt(env, opAddBigInt) }
func CeilBigInt(env *zc.Env) error    { return funcs.EvalUnaryBigInt(env, opCeilBigInt) }
func DivBigInt(env *zc.Env) error     { return funcs.EvalBinaryBigInt(env, opDivBigInt) }
func FloorBigInt(env *zc.Env) error   { return funcs.EvalUnaryBigInt(env, opCeilBigInt) }
func ModulusBigInt(env *zc.Env) error { return funcs.EvalBinaryBigInt(env, opModBigInt) }
func MulBigInt(env *zc.Env) error     { return funcs.EvalBinaryBigInt(env, opMulBigInt) }
func NegBigInt(env *zc.Env) error     { return funcs.EvalUnaryBigInt(env, opNegBigInt) }
func PowBigInt(env *zc.Env) error     { return funcs.EvalBinaryBigInt(env, opPowBigInt) }
func RemBigInt(env *zc.Env) error     { return funcs.EvalBinaryBigInt(env, opRemBigInt) }
func SignBigInt(env *zc.Env) error    { return funcs.EvalUnaryBigInt(env, opSignBigInt) }
func SubBigInt(env *zc.Env) error     { return funcs.EvalBinaryBigInt(env, opSubBigInt) }
