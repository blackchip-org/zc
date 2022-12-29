package zlib

import (
	"fmt"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

var (
	opAbs = funcs.UnaryOps{
		BigInt: opAbsBigInt,
		Fixed:  opAbsFixed,
	}
	opAdd = funcs.BinaryOps{
		BigInt: opAddBigInt,
		Fixed:  opAddFixed,
	}
	opCeil = funcs.UnaryOps{
		BigInt: opCeilBigInt,
		Fixed:  opCeilFixed,
	}
	opFloor = funcs.UnaryOps{
		BigInt: opFloorBigInt,
		Fixed:  opFloorFixed,
	}
	opMul = funcs.BinaryOps{
		BigInt: opMulBigInt,
		Fixed:  opMulFixed,
	}
	opMod = funcs.BinaryOps{
		BigInt: opModBigInt,
		Fixed:  opModFixed,
	}
	opNeg = funcs.UnaryOps{
		BigInt: opNegBigInt,
		Fixed:  opNegFixed,
	}
	opPow = funcs.BinaryOps{
		BigInt: opPowBigInt,
		Fixed:  opPowFixed,
	}
	opRem = funcs.BinaryOps{
		BigInt: opRemBigInt,
		Fixed:  opRemFixed,
	}
	opSign = funcs.UnaryOps{
		BigInt: opSignBigInt,
		Fixed:  opSignFixed,
	}
	opSub = funcs.BinaryOps{
		BigInt: opSubBigInt,
		Fixed:  opSubFixed,
	}
)

func Abs(env *zc.Env) error   { return funcs.EvalUnaryNum(env, opAbs) }
func Add(env *zc.Env) error   { return funcs.EvalBinaryNum(env, opAdd) }
func Ceil(env *zc.Env) error  { return funcs.EvalUnaryNum(env, opCeil) }
func Floor(env *zc.Env) error { return funcs.EvalUnaryNum(env, opFloor) }
func Mod(env *zc.Env) error   { return funcs.EvalBinaryNum(env, opMod) }
func Mul(env *zc.Env) error   { return funcs.EvalBinaryNum(env, opMul) }
func Neg(env *zc.Env) error   { return funcs.EvalUnaryNum(env, opNeg) }
func Pow(env *zc.Env) error   { return funcs.EvalBinaryNum(env, opPow) }
func Rem(env *zc.Env) error   { return funcs.EvalBinaryNum(env, opRem) }
func Sign(env *zc.Env) error  { return funcs.EvalUnaryNum(env, opSign) }
func Sqrt(env *zc.Env) error  { return funcs.EvalUnaryFloat(env, opSqrtFloat) }
func Sub(env *zc.Env) error   { return funcs.EvalBinaryNum(env, opSub) }

func Round(env *zc.Env) error {
	places, err := env.Stack.PopInt32()
	if err != nil {
		return err
	}
	value, err := env.Stack.PopFixed()
	if err != nil {
		return err
	}
	fn, ok := zc.RoundingFuncsFix[env.Calc.RoundingMode]
	if !ok {
		return fmt.Errorf("invalid rounding mode: %v", env.Calc.RoundingMode)
	}
	r := fn(value, places)
	env.Stack.PushFixed(r)
	return nil
}
