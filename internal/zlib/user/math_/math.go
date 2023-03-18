package math_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

var (
	opAbs = funcs.UnaryOps{
		BigInt:  opAbsBigInt,
		Decimal: opAbsDecimal,
	}
	opAdd = funcs.BinaryOps{
		BigInt:  opAddBigInt,
		Decimal: opAddDecimal,
	}
	opCeil = funcs.UnaryOps{
		BigInt:  opCeilBigInt,
		Decimal: opCeilDecimal,
	}
	opFloor = funcs.UnaryOps{
		BigInt:  opFloorBigInt,
		Decimal: opFloorDecimal,
	}
	opMul = funcs.BinaryOps{
		BigInt:  opMulBigInt,
		Decimal: opMulDecimal,
	}
	opMod = funcs.BinaryOps{
		BigInt:  opModBigInt,
		Decimal: opModDecimal,
	}
	opNeg = funcs.UnaryOps{
		BigInt:  opNegBigInt,
		Decimal: opNegDecimal,
	}
	opPow = funcs.BinaryOps{
		BigInt:  opPowBigInt,
		Decimal: opPowDecimal,
	}
	opRem = funcs.BinaryOps{
		BigInt:  opRemBigInt,
		Decimal: opRemDecimal,
	}
	opSign = funcs.UnaryOps{
		BigInt:  opSignBigInt,
		Decimal: opSignDecimal,
	}
	opSub = funcs.BinaryOps{
		BigInt:  opSubBigInt,
		Decimal: opSubDecimal,
	}
)

func Abs(env *zc.Env) error     { return funcs.EvalUnaryNum(env, opAbs) }
func Add(env *zc.Env) error     { return funcs.EvalBinaryNum(env, opAdd) }
func Ceil(env *zc.Env) error    { return funcs.EvalUnaryNum(env, opCeil) }
func Floor(env *zc.Env) error   { return funcs.EvalUnaryNum(env, opFloor) }
func Modulus(env *zc.Env) error { return funcs.EvalBinaryNum(env, opMod) }
func Mul(env *zc.Env) error     { return funcs.EvalBinaryNum(env, opMul) }
func Neg(env *zc.Env) error     { return funcs.EvalUnaryNum(env, opNeg) }
func Pow(env *zc.Env) error     { return funcs.EvalBinaryNum(env, opPow) }
func Rem(env *zc.Env) error     { return funcs.EvalBinaryNum(env, opRem) }
func Sign(env *zc.Env) error    { return funcs.EvalUnaryNum(env, opSign) }
func Sqrt(env *zc.Env) error    { return funcs.EvalUnaryFloat(env, opSqrtFloat) }
func Sub(env *zc.Env) error     { return funcs.EvalBinaryNum(env, opSub) }

func Sum(env *zc.Env) error {
	for env.Stack.Len() > 1 {
		if err := Add(env); err != nil {
			return err
		}
	}
	return nil
}
