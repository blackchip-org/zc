package funcs

import "github.com/blackchip-org/zc"

type CompareOps struct {
	BigInt  CompareBigInt
	Decimal CompareDecimal
	String  CompareStr
}

func EvalCompareVal(env *zc.Env, ops CompareOps) error {
	a, b, err := env.Stack.Pop2()
	if err != nil {
		return err
	}

	var result bool

	switch {
	case env.Calc.IsBigInt(a) && env.Calc.IsBigInt(b):
		x, y := env.Calc.MustParseBigInt(a), env.Calc.MustParseBigInt(b)
		result, err = ops.BigInt(x, y)
	case env.Calc.IsDecimal(a) && env.Calc.IsDecimal(b):
		x, y := env.Calc.MustParseDecimal(a), env.Calc.MustParseDecimal(b)
		result, err = ops.Decimal(x, y)
	default:
		result, err = ops.String(a, b)
	}
	if err != nil {
		return err
	}
	env.Stack.PushBool(result)
	return nil
}
