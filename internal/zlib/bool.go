package zlib

import (
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
	"github.com/shopspring/decimal"
)

func opEqBigInt(a *big.Int, b *big.Int) (bool, error)  { return a.Cmp(b) == 0, nil }
func opNeqBigInt(a *big.Int, b *big.Int) (bool, error) { return a.Cmp(b) != 0, nil }
func opGtBigInt(a *big.Int, b *big.Int) (bool, error)  { return a.Cmp(b) > 0, nil }
func opGteBigInt(a *big.Int, b *big.Int) (bool, error) { return a.Cmp(b) >= 0, nil }
func opLtBigInt(a *big.Int, b *big.Int) (bool, error)  { return a.Cmp(b) < 0, nil }
func opLteBigInt(a *big.Int, b *big.Int) (bool, error) { return a.Cmp(b) <= 0, nil }

func opEqFix(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) == 0, nil }
func opNeqFix(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) != 0, nil }
func opGtFix(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) > 0, nil }
func opGteFix(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) >= 0, nil }
func opLtFix(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) < 0, nil }
func opLteFix(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) <= 0, nil }

func opEqString(a string, b string) (bool, error)  { return a == b, nil }
func opNeqString(a string, b string) (bool, error) { return a != b, nil }
func opGtString(a string, b string) (bool, error)  { return a > b, nil }
func opGteString(a string, b string) (bool, error) { return a >= b, nil }
func opLtString(a string, b string) (bool, error)  { return a < b, nil }
func opLteString(a string, b string) (bool, error) { return a <= b, nil }

func opAnd(a bool, b bool) (bool, error) { return a && b, nil }
func opNot(a bool) (bool, error)         { return !a, nil }
func opOr(a bool, b bool) (bool, error)  { return a || b, nil }

func EqBigInt(calc *zc.Calc) error  { return funcs.EvalCompareBigInt(calc, opEqBigInt) }
func GtBigInt(calc *zc.Calc) error  { return funcs.EvalCompareBigInt(calc, opGtBigInt) }
func GteBigInt(calc *zc.Calc) error { return funcs.EvalCompareBigInt(calc, opGteBigInt) }
func NeqBigInt(calc *zc.Calc) error { return funcs.EvalCompareBigInt(calc, opNeqBigInt) }
func LtBigInt(calc *zc.Calc) error  { return funcs.EvalCompareBigInt(calc, opLtBigInt) }
func LteBigInt(calc *zc.Calc) error { return funcs.EvalCompareBigInt(calc, opLteBigInt) }

func EqFix(calc *zc.Calc) error  { return funcs.EvalCompareFix(calc, opEqFix) }
func GtFix(calc *zc.Calc) error  { return funcs.EvalCompareFix(calc, opEqFix) }
func GteFix(calc *zc.Calc) error { return funcs.EvalCompareFix(calc, opEqFix) }
func NeqFix(calc *zc.Calc) error { return funcs.EvalCompareFix(calc, opEqFix) }
func LtFix(calc *zc.Calc) error  { return funcs.EvalCompareFix(calc, opEqFix) }
func LteFix(calc *zc.Calc) error { return funcs.EvalCompareFix(calc, opEqFix) }

func EqStr(calc *zc.Calc) error  { return funcs.EvalCompareStr(calc, opEqString) }
func GtStr(calc *zc.Calc) error  { return funcs.EvalCompareStr(calc, opGtString) }
func GteStr(calc *zc.Calc) error { return funcs.EvalCompareStr(calc, opGteString) }
func NeqStr(calc *zc.Calc) error { return funcs.EvalCompareStr(calc, opNeqString) }
func LtStr(calc *zc.Calc) error  { return funcs.EvalCompareStr(calc, opLtString) }
func LteStr(calc *zc.Calc) error { return funcs.EvalCompareStr(calc, opLteString) }

var (
	opEq = funcs.CompareOps{
		BigInt: opEqBigInt,
		Fix:    opEqFix,
		String: opEqString,
	}
	opNeq = funcs.CompareOps{
		BigInt: opNeqBigInt,
		Fix:    opNeqFix,
		String: opNeqString,
	}
	opGt = funcs.CompareOps{
		BigInt: opGtBigInt,
		Fix:    opGtFix,
		String: opGtString,
	}
	opGte = funcs.CompareOps{
		BigInt: opGteBigInt,
		Fix:    opGteFix,
		String: opGteString,
	}
	opLt = funcs.CompareOps{
		BigInt: opLtBigInt,
		Fix:    opLtFix,
		String: opLtString,
	}
	opLte = funcs.CompareOps{
		BigInt: opLteBigInt,
		Fix:    opLteFix,
		String: opLteString,
	}
)

func Eq(calc *zc.Calc) error  { return funcs.EvalCompareVal(calc, opEq) }
func Neq(calc *zc.Calc) error { return funcs.EvalCompareVal(calc, opNeq) }
func Gt(calc *zc.Calc) error  { return funcs.EvalCompareVal(calc, opGt) }
func Gte(calc *zc.Calc) error { return funcs.EvalCompareVal(calc, opGte) }
func Lt(calc *zc.Calc) error  { return funcs.EvalCompareVal(calc, opLt) }
func Lte(calc *zc.Calc) error { return funcs.EvalCompareVal(calc, opLte) }

func And(calc *zc.Calc) error { return funcs.EvalBinaryBool(calc, opAnd) }
func Not(calc *zc.Calc) error { return funcs.EvalUnaryBool(calc, opNot) }
func Or(calc *zc.Calc) error  { return funcs.EvalBinaryBool(calc, opOr) }
