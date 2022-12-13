package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

var (
	opAbs = funcs.UnaryOps{
		BigInt: opAbsBigInt,
		Fix:    opAbsFix,
	}
	opAdd = funcs.BinaryOps{
		BigInt: opAddBigInt,
		Fix:    opAddFix,
	}
	opCeil = funcs.UnaryOps{
		BigInt: opCeilBigInt,
		Fix:    opCeilFix,
	}
	opFloor = funcs.UnaryOps{
		BigInt: opFloorBigInt,
		Fix:    opFloorFix,
	}
	opMul = funcs.BinaryOps{
		BigInt: opMulBigInt,
		Fix:    opMulFix,
	}
	opMod = funcs.BinaryOps{
		BigInt: opModBigInt,
		Fix:    opModFix,
	}
	opNeg = funcs.UnaryOps{
		BigInt: opNegBigInt,
		Fix:    opNegFix,
	}
	opPow = funcs.BinaryOps{
		BigInt: opPowBigInt,
		Fix:    opPowFix,
	}
	opRem = funcs.BinaryOps{
		BigInt: opRemBigInt,
		Fix:    opRemFix,
	}
	opSign = funcs.UnaryOps{
		BigInt: opSignBigInt,
		Fix:    opSignFix,
	}
	opSub = funcs.BinaryOps{
		BigInt: opSubBigInt,
		Fix:    opSubFix,
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
