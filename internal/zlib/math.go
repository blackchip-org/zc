package zlib

import (
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
	"github.com/shopspring/decimal"
)

func bigIntMod(z *big.Int, a *big.Int, b *big.Int) {
	var q big.Int
	q.DivMod(a, b, z)
}

func bigIntSign(z *big.Int, a *big.Int) {
	z.SetInt64(int64(a.Sign()))
}

func fixRem(a decimal.Decimal, b decimal.Decimal) decimal.Decimal {
	_, rem := a.QuoRem(a, zc.Places)
	return rem
}

func fixSign(a decimal.Decimal) decimal.Decimal {
	return decimal.NewFromInt(int64(a.Sign()))
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

func opAbsFix(a decimal.Decimal) (decimal.Decimal, error)                    { return a.Abs(), nil }
func opAddFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Add(b), nil }
func opCeilFix(a decimal.Decimal) (decimal.Decimal, error)                   { return a.Ceil(), nil }
func opDivFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Div(b), nil }
func opFloorFix(a decimal.Decimal) (decimal.Decimal, error)                  { return a.Floor(), nil }
func opMulFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Mul(b), nil }
func opNegFix(a decimal.Decimal) (decimal.Decimal, error)                    { return a.Neg(), nil }
func opModFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Mod(b), nil }
func opPowFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Pow(b), nil }
func opRemFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Pow(b), nil }
func opSignFix(a decimal.Decimal) (decimal.Decimal, error)                   { return fixSign(a), nil }
func opSubFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Sub(b), nil }

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

func AbsFix(calc *zc.Calc) error   { return funcs.EvalUnaryFix(calc, opAbsFix) }
func AddFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opAddFix) }
func CeilFix(calc *zc.Calc) error  { return funcs.EvalUnaryFix(calc, opCeilFix) }
func DivFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opDivFix) }
func FloorFix(calc *zc.Calc) error { return funcs.EvalUnaryFix(calc, opCeilFix) }
func ModFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opModFix) }
func MulFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opMulFix) }
func NegFix(calc *zc.Calc) error   { return funcs.EvalUnaryFix(calc, opNegFix) }
func PowFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opPowFix) }
func RemFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opRemFix) }
func SignFix(calc *zc.Calc) error  { return funcs.EvalUnaryFix(calc, opSignFix) }
func SubFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opSubFix) }

var (
	opAbs = funcs.UnaryOps{
		BigInt: opAbsBigInt,
		Fix:    opAbsFix,
	}
	opAdd = funcs.BinaryOps{
		BigInt: opAddBigInt,
		Fix:    opAddFix,
	}
	opCeil = funcs.UnaryOps{
		BigInt: opCeilBigInt,
		Fix:    opCeilFix,
	}
	opFloor = funcs.UnaryOps{
		BigInt: opFloorBigInt,
		Fix:    opFloorFix,
	}
	opMul = funcs.BinaryOps{
		BigInt: opMulBigInt,
		Fix:    opMulFix,
	}
	opMod = funcs.BinaryOps{
		BigInt: opModBigInt,
		Fix:    opModFix,
	}
	opNeg = funcs.UnaryOps{
		BigInt: opNegBigInt,
		Fix:    opNegFix,
	}
	opPow = funcs.BinaryOps{
		BigInt: opPowBigInt,
		Fix:    opPowFix,
	}
	opRem = funcs.BinaryOps{
		BigInt: opRemBigInt,
		Fix:    opRemFix,
	}
	opSign = funcs.UnaryOps{
		BigInt: opSignBigInt,
		Fix:    opSignFix,
	}
	opSub = funcs.BinaryOps{
		BigInt: opSubBigInt,
		Fix:    opSubFix,
	}
)

func Abs(calc *zc.Calc) error   { return funcs.EvalUnaryNum(calc, opAbs) }
func Add(calc *zc.Calc) error   { return funcs.EvalBinaryNum(calc, opAdd) }
func Ceil(calc *zc.Calc) error  { return funcs.EvalUnaryNum(calc, opCeil) }
func Floor(calc *zc.Calc) error { return funcs.EvalUnaryNum(calc, opFloor) }
func Mod(calc *zc.Calc) error   { return funcs.EvalBinaryNum(calc, opMod) }
func Mul(calc *zc.Calc) error   { return funcs.EvalBinaryNum(calc, opMul) }
func Neg(calc *zc.Calc) error   { return funcs.EvalUnaryNum(calc, opNeg) }
func Pow(calc *zc.Calc) error   { return funcs.EvalBinaryNum(calc, opPow) }
func Rem(calc *zc.Calc) error   { return funcs.EvalBinaryNum(calc, opRem) }
func Sign(calc *zc.Calc) error  { return funcs.EvalUnaryNum(calc, opSign) }
func Sub(calc *zc.Calc) error   { return funcs.EvalBinaryNum(calc, opSub) }
