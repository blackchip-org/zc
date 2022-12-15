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

func Round(calc *zc.Calc) error {
	places, err := calc.Stack.PopInt32()
	if err != nil {
		return err
	}
	value, err := calc.Stack.PopFixed()
	if err != nil {
		return err
	}
	fn, ok := zc.RoundingFuncsFix[calc.Val.RoundingMode]
	if !ok {
		return fmt.Errorf("invalid rounding mode: %v", calc.Val.RoundingMode)
	}
	r := fn(value, places)
	calc.Stack.Push(calc.Val.FormatFixed(r))
	return nil
}
