package funcs

import "github.com/blackchip-org/zc"

type UnaryOps struct {
	BigInt UnaryBigInt
	Fixed  UnaryFixed
	Float  UnaryFloat
}

type BinaryOps struct {
	BigInt BinaryBigInt
	Fixed  BinaryFixed
}

func EvalUnaryNum(env *zc.Env, ops UnaryOps) error {
	a, err := env.Stack.Peek()
	if err != nil {
		return err
	}

	switch {
	case env.Calc.IsBigInt(a):
		return EvalUnaryBigInt(env, ops.BigInt)
	case ops.Fixed != nil:
		return EvalUnaryFixed(env, ops.Fixed)
	case ops.Float != nil :
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
		return EvalBinaryFixed(env, ops.Fixed)
	}
}
