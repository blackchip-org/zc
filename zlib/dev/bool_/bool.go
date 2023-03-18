package bool_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func opAnd(a bool, b bool) (bool, error) { return a && b, nil }
func opNot(a bool) (bool, error)         { return !a, nil }
func opOr(a bool, b bool) (bool, error)  { return a || b, nil }

var (
	opEq = funcs.CompareOps{
		BigInt:  opEqBigInt,
		Decimal: opEqDecimal,
		String:  opEqString,
	}
	opNeq = funcs.CompareOps{
		BigInt:  opNeqBigInt,
		Decimal: opNeqDecimal,
		String:  opNeqString,
	}
	opGt = funcs.CompareOps{
		BigInt:  opGtBigInt,
		Decimal: opGtDecimal,
		String:  opGtString,
	}
	opGte = funcs.CompareOps{
		BigInt:  opGteBigInt,
		Decimal: opGteDecimal,
		String:  opGteString,
	}
	opLt = funcs.CompareOps{
		BigInt:  opLtBigInt,
		Decimal: opLtDecimal,
		String:  opLtString,
	}
	opLte = funcs.CompareOps{
		BigInt:  opLteBigInt,
		Decimal: opLteDecimal,
		String:  opLteString,
	}
)

func Eq(env *zc.Env) error  { return funcs.EvalCompareVal(env, opEq) }
func Neq(env *zc.Env) error { return funcs.EvalCompareVal(env, opNeq) }
func Gt(env *zc.Env) error  { return funcs.EvalCompareVal(env, opGt) }
func Gte(env *zc.Env) error { return funcs.EvalCompareVal(env, opGte) }
func Lt(env *zc.Env) error  { return funcs.EvalCompareVal(env, opLt) }
func Lte(env *zc.Env) error { return funcs.EvalCompareVal(env, opLte) }

func And(env *zc.Env) error { return funcs.EvalBinaryBool(env, opAnd) }
func Not(env *zc.Env) error { return funcs.EvalUnaryBool(env, opNot) }
func Or(env *zc.Env) error  { return funcs.EvalBinaryBool(env, opOr) }
