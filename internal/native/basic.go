package native

import (
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/shopspring/decimal"
)

func opAddBigInt(z *big.Int, a *big.Int, b *big.Int) { z.Add(a, b) }
func opDivBigInt(z *big.Int, a *big.Int, b *big.Int) { z.Div(a, b) }
func opMulBigInt(z *big.Int, a *big.Int, b *big.Int) { z.Mul(a, b) }
func opPowBigInt(z *big.Int, a *big.Int, b *big.Int) { z.Exp(a, b, nil) }
func opSubBigInt(z *big.Int, a *big.Int, b *big.Int) { z.Sub(a, b) }

func opAddDec(a decimal.Decimal, b decimal.Decimal) decimal.Decimal { return a.Add(b) }
func opDivDec(a decimal.Decimal, b decimal.Decimal) decimal.Decimal { return a.Div(b) }
func opMulDec(a decimal.Decimal, b decimal.Decimal) decimal.Decimal { return a.Mul(b) }
func opPowDec(a decimal.Decimal, b decimal.Decimal) decimal.Decimal { return a.Pow(b) }
func opSubDec(a decimal.Decimal, b decimal.Decimal) decimal.Decimal { return a.Sub(b) }

func AddBigInt(calc *zc.Calc) error { return zc.BigIntNumOp2(calc, opAddBigInt) }
func DivBigInt(calc *zc.Calc) error { return zc.BigIntNumOp2(calc, opDivBigInt) }
func MulBigInt(calc *zc.Calc) error { return zc.BigIntNumOp2(calc, opMulBigInt) }
func PowBigInt(calc *zc.Calc) error { return zc.BigIntNumOp2(calc, opPowBigInt) }
func SubBigInt(calc *zc.Calc) error { return zc.BigIntNumOp2(calc, opSubBigInt) }

func AddDec(calc *zc.Calc) error { return zc.DecimalNumOp2(calc, opAddDec) }
func DivDec(calc *zc.Calc) error { return zc.DecimalNumOp2(calc, opDivDec) }
func MulDec(calc *zc.Calc) error { return zc.DecimalNumOp2(calc, opMulDec) }
func PowDec(calc *zc.Calc) error { return zc.DecimalNumOp2(calc, opPowDec) }
func SubDec(calc *zc.Calc) error { return zc.DecimalNumOp2(calc, opSubDec) }

var (
	opAdd = zc.FuncsNumOp2{
		BigInt:  opAddBigInt,
		Decimal: opAddDec,
	}
	opMul = zc.FuncsNumOp2{
		BigInt:  opMulBigInt,
		Decimal: opMulDec,
	}
	opPow = zc.FuncsNumOp2{
		BigInt:  opPowBigInt,
		Decimal: opPowDec,
	}
	opSub = zc.FuncsNumOp2{
		BigInt:  opSubBigInt,
		Decimal: opSubDec,
	}
)

func Add(calc *zc.Calc) error { return zc.NumOp2(calc, opAdd) }
func Mul(calc *zc.Calc) error { return zc.NumOp2(calc, opMul) }
func Pow(calc *zc.Calc) error { return zc.NumOp2(calc, opPow) }
func Sub(calc *zc.Calc) error { return zc.NumOp2(calc, opSub) }
