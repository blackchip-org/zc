package funcs

import "github.com/blackchip-org/zc"

type UnaryOps struct {
	BigInt UnaryBigInt
	Fix    UnaryFix
}

type BinaryOps struct {
	BigInt BinaryBigInt
	Fix    BinaryFix
}

func EvalUnaryNum(calc *zc.Calc, ops UnaryOps) error {
	a, err := calc.Stack.Get()
	if err != nil {
		return err
	}

	switch {
	case calc.IsBigInt(a):
		return EvalUnaryBigInt(calc, ops.BigInt)
	default:
		return EvalUnaryFix(calc, ops.Fix)
	}
}

func EvalBinaryNum(calc *zc.Calc, ops BinaryOps) error {
	a, b, err := calc.Peek2()
	if err != nil {
		return err
	}

	switch {
	case calc.IsBigInt(a) && calc.IsBigInt(b):
		return EvalBinaryBigInt(calc, ops.BigInt)
	default:
		return EvalBinaryFix(calc, ops.Fix)
	}
}
