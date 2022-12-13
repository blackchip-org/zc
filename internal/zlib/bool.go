package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func opAnd(a bool, b bool) (bool, error) { return a && b, nil }
func opNot(a bool) (bool, error)         { return !a, nil }
func opOr(a bool, b bool) (bool, error)  { return a || b, nil }

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
