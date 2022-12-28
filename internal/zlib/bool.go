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
		Fixed:  opEqFixed,
		String: opEqString,
	}
	opNeq = funcs.CompareOps{
		BigInt: opNeqBigInt,
		Fixed:  opNeqFixed,
		String: opNeqString,
	}
	opGt = funcs.CompareOps{
		BigInt: opGtBigInt,
		Fixed:  opGtFixed,
		String: opGtString,
	}
	opGte = funcs.CompareOps{
		BigInt: opGteBigInt,
		Fixed:  opGteFixed,
		String: opGteString,
	}
	opLt = funcs.CompareOps{
		BigInt: opLtBigInt,
		Fixed:  opLtFixed,
		String: opLtString,
	}
	opLte = funcs.CompareOps{
		BigInt: opLteBigInt,
		Fixed:  opLteFixed,
		String: opLteString,
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
