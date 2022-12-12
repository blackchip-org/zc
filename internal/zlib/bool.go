package zlib

import (
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/shopspring/decimal"
)

func opEqBigInt(a *big.Int, b *big.Int) bool  { return a.Cmp(b) == 0 }
func opNeqBigInt(a *big.Int, b *big.Int) bool { return a.Cmp(b) != 0 }
func opGtBigInt(a *big.Int, b *big.Int) bool  { return a.Cmp(b) > 0 }
func opGteBigInt(a *big.Int, b *big.Int) bool { return a.Cmp(b) >= 0 }
func opLtBigInt(a *big.Int, b *big.Int) bool  { return a.Cmp(b) < 0 }
func opLteBigInt(a *big.Int, b *big.Int) bool { return a.Cmp(b) <= 0 }

func opEqDecimal(a decimal.Decimal, b decimal.Decimal) bool  { return a.Cmp(b) == 0 }
func opNeqDecimal(a decimal.Decimal, b decimal.Decimal) bool { return a.Cmp(b) != 0 }
func opGtDecimal(a decimal.Decimal, b decimal.Decimal) bool  { return a.Cmp(b) > 0 }
func opGteDecimal(a decimal.Decimal, b decimal.Decimal) bool { return a.Cmp(b) >= 0 }
func opLtDecimal(a decimal.Decimal, b decimal.Decimal) bool  { return a.Cmp(b) < 0 }
func opLteDecimal(a decimal.Decimal, b decimal.Decimal) bool { return a.Cmp(b) <= 0 }

func opEqString(a string, b string) bool  { return a == b }
func opNeqString(a string, b string) bool { return a != b }
func opGtString(a string, b string) bool  { return a > b }
func opGteString(a string, b string) bool { return a >= b }
func opLtString(a string, b string) bool  { return a < b }
func opLteString(a string, b string) bool { return a <= b }

func opAnd(a bool, b bool) bool { return a && b }
func opOr(a bool, b bool) bool  { return a || b }

func EqBigInt(calc *zc.Calc) error  { return zc.BigIntCompOp(calc, opEqBigInt) }
func GtBigInt(calc *zc.Calc) error  { return zc.BigIntCompOp(calc, opGtBigInt) }
func GteBigInt(calc *zc.Calc) error { return zc.BigIntCompOp(calc, opGteBigInt) }
func NeqBigInt(calc *zc.Calc) error { return zc.BigIntCompOp(calc, opNeqBigInt) }
func LtBigInt(calc *zc.Calc) error  { return zc.BigIntCompOp(calc, opLtBigInt) }
func LteBigInt(calc *zc.Calc) error { return zc.BigIntCompOp(calc, opLteBigInt) }

func EqDec(calc *zc.Calc) error  { return zc.DecCompOp(calc, opEqDecimal) }
func GtDec(calc *zc.Calc) error  { return zc.DecCompOp(calc, opEqDecimal) }
func GteDec(calc *zc.Calc) error { return zc.DecCompOp(calc, opEqDecimal) }
func NeqDec(calc *zc.Calc) error { return zc.DecCompOp(calc, opEqDecimal) }
func LtDec(calc *zc.Calc) error  { return zc.DecCompOp(calc, opEqDecimal) }
func LteDec(calc *zc.Calc) error { return zc.DecCompOp(calc, opEqDecimal) }

func EqStr(calc *zc.Calc) error  { return zc.StringCompOp(calc, opEqString) }
func GtStr(calc *zc.Calc) error  { return zc.StringCompOp(calc, opGtString) }
func GteStr(calc *zc.Calc) error { return zc.StringCompOp(calc, opGteString) }
func NeqStr(calc *zc.Calc) error { return zc.StringCompOp(calc, opNeqString) }
func LtStr(calc *zc.Calc) error  { return zc.StringCompOp(calc, opLtString) }
func LteStr(calc *zc.Calc) error { return zc.StringCompOp(calc, opLteString) }

var (
	opEq = zc.FuncsCompOp{
		BigInt:  opEqBigInt,
		Decimal: opEqDecimal,
		String:  opEqString,
	}
	opNeq = zc.FuncsCompOp{
		BigInt:  opNeqBigInt,
		Decimal: opNeqDecimal,
		String:  opNeqString,
	}
	opGt = zc.FuncsCompOp{
		BigInt:  opGtBigInt,
		Decimal: opGtDecimal,
		String:  opGtString,
	}
	opGte = zc.FuncsCompOp{
		BigInt:  opGteBigInt,
		Decimal: opGteDecimal,
		String:  opGteString,
	}
	opLt = zc.FuncsCompOp{
		BigInt:  opLtBigInt,
		Decimal: opLtDecimal,
		String:  opLtString,
	}
	opLte = zc.FuncsCompOp{
		BigInt:  opLteBigInt,
		Decimal: opLteDecimal,
		String:  opLteString,
	}
)

func Eq(calc *zc.Calc) error  { return zc.CompOp(calc, opEq) }
func Neq(calc *zc.Calc) error { return zc.CompOp(calc, opNeq) }
func Gt(calc *zc.Calc) error  { return zc.CompOp(calc, opGt) }
func Gte(calc *zc.Calc) error { return zc.CompOp(calc, opGte) }
func Lt(calc *zc.Calc) error  { return zc.CompOp(calc, opLt) }
func Lte(calc *zc.Calc) error { return zc.CompOp(calc, opLte) }

func And(calc *zc.Calc) error { return zc.BoolOp2(calc, opAnd) }
func Or(calc *zc.Calc) error  { return zc.BoolOp2(calc, opOr) }

func Not(calc *zc.Calc) error {
	a, err := calc.PopBool()
	if err != nil {
		return err
	}

	r := !a
	calc.Stack.Push(zc.FormatBool(r))
	return nil
}
