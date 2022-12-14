package funcs

import "github.com/blackchip-org/zc"

type UnaryOps struct {
	BigInt UnaryBigInt
	Fixed  UnaryFixed
}

type BinaryOps struct {
	BigInt BinaryBigInt
	Fixed  BinaryFixed
}

func EvalUnaryNum(calc *zc.Calc, ops UnaryOps) error {
	a, err := calc.Stack.Peek()
	if err != nil {
		return err
	}

	switch {
	case calc.Val.IsBigInt(a):
		return EvalUnaryBigInt(calc, ops.BigInt)
	default:
		return EvalUnaryFixed(calc, ops.Fixed)
	}
}

func EvalBinaryNum(calc *zc.Calc, ops BinaryOps) error {
	a, b, err := calc.Stack.Peek2()
	if err != nil {
		return err
	}

	switch {
	case calc.Val.IsBigInt(a) && calc.Val.IsBigInt(b):
		return EvalBinaryBigInt(calc, ops.BigInt)
	default:
		return EvalBinaryFixed(calc, ops.Fixed)
	}
}
