package funcs

import (
	"github.com/blackchip-org/zc"
)

type UnaryOps struct {
	BigInt  UnaryBigInt
	Decimal UnaryDecimal
	Float   UnaryFloat
}

type BinaryOps struct {
	BigInt  BinaryBigInt
	Decimal BinaryDecimal
}

func EvalUnaryNum(env *zc.Env, ops UnaryOps) error {
	a, err := env.Stack.Peek()
	if err != nil {
		return err
	}

	switch {
	case env.Calc.IsBigInt(a):
		return EvalUnaryBigInt(env, ops.BigInt)
	case ops.Decimal != nil:
		return EvalUnaryDecimal(env, ops.Decimal)
	case ops.Float != nil:
		return EvalUnaryFloat(env, ops.Float)
	}
	panic("unsupported operation")
}

func EvalBinaryNum(env *zc.Env, ops BinaryOps) error {
	a, b, err := env.Stack.Peek2()
	if err != nil {
		return err
	}

	switch {
	case env.Calc.IsBigInt(a) && env.Calc.IsBigInt(b):
		return EvalBinaryBigInt(env, ops.BigInt)
	default:
		return EvalBinaryDecimal(env, ops.Decimal)
	}
}
