package zlib

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

func AbsBigInt(calc *zc.Calc) error   { return funcs.EvalUnaryBigInt(calc, opAbsBigInt) }
func AddBigInt(calc *zc.Calc) error   { return funcs.EvalBinaryBigInt(calc, opAddBigInt) }
func CeilBigInt(calc *zc.Calc) error  { return funcs.EvalUnaryBigInt(calc, opCeilBigInt) }
func DivBigInt(calc *zc.Calc) error   { return funcs.EvalBinaryBigInt(calc, opDivBigInt) }
func FloorBigInt(calc *zc.Calc) error { return funcs.EvalUnaryBigInt(calc, opCeilBigInt) }
func ModBigInt(calc *zc.Calc) error   { return funcs.EvalBinaryBigInt(calc, opModBigInt) }
func MulBigInt(calc *zc.Calc) error   { return funcs.EvalBinaryBigInt(calc, opMulBigInt) }
func NegBigInt(calc *zc.Calc) error   { return funcs.EvalUnaryBigInt(calc, opNegBigInt) }
func PowBigInt(calc *zc.Calc) error   { return funcs.EvalBinaryBigInt(calc, opPowBigInt) }
func RemBigInt(calc *zc.Calc) error   { return funcs.EvalBinaryBigInt(calc, opRemBigInt) }
func SignBigInt(calc *zc.Calc) error  { return funcs.EvalUnaryBigInt(calc, opSignBigInt) }
func SubBigInt(calc *zc.Calc) error   { return funcs.EvalBinaryBigInt(calc, opSubBigInt) }
