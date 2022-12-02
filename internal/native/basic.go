package native

import (
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/shopspring/decimal"
)

func opAddBigInt(z *big.Int, a *big.Int, b *big.Int) { z.Add(a, b) }
func opMulBigInt(z *big.Int, a *big.Int, b *big.Int) { z.Mul(a, b) }
func opSubBigInt(z *big.Int, a *big.Int, b *big.Int) { z.Sub(a, b) }

func opAddDec(a decimal.Decimal, b decimal.Decimal) decimal.Decimal { return a.Add(b) }
func opMulDec(a decimal.Decimal, b decimal.Decimal) decimal.Decimal { return a.Mul(b) }
func opSubDec(a decimal.Decimal, b decimal.Decimal) decimal.Decimal { return a.Sub(b) }

func AddBigInt(calc *zc.Calc) error { return zc.BigInt2(calc, opAddBigInt) }
func MulBigInt(calc *zc.Calc) error { return zc.BigInt2(calc, opMulBigInt) }
func SubBigInt(calc *zc.Calc) error { return zc.BigInt2(calc, opSubBigInt) }

func AddDec(calc *zc.Calc) error { return zc.Dec2(calc, opAddDec) }
func MulDec(calc *zc.Calc) error { return zc.Dec2(calc, opMulDec) }
func SubDec(calc *zc.Calc) error { return zc.Dec2(calc, opSubDec) }

var (
	opAdd = zc.NumOp2{
		BigInt2: opAddBigInt,
		Dec2:    opAddDec,
	}
	opMul = zc.NumOp2{
		BigInt2: opMulBigInt,
		Dec2:    opMulDec,
	}
	opSub = zc.NumOp2{
		BigInt2: opSubBigInt,
		Dec2:    opSubDec,
	}
)

func Add(calc *zc.Calc) error { return zc.Num2(calc, opAdd) }
func Mul(calc *zc.Calc) error { return zc.Num2(calc, opMul) }
func Sub(calc *zc.Calc) error { return zc.Num2(calc, opSub) }
