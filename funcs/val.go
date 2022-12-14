package funcs

import "github.com/blackchip-org/zc"

type CompareOps struct {
	BigInt CompareBigInt
	Fix    CompareFix
	String CompareStr
}

func EvalCompareVal(calc *zc.Calc, ops CompareOps) error {
	a, b, err := calc.Stack.Pop2()
	if err != nil {
		return err
	}

	var result bool

	switch {
	case calc.Value.IsBigInt(a) && calc.Value.IsBigInt(b):
		x, y := calc.Value.MustParseBigInt(a), calc.Value.MustParseBigInt(b)
		result, err = ops.BigInt(x, y)
	case calc.Value.IsFix(a) && calc.Value.IsFix(b):
		x, y := calc.Value.MustParseFix(a), calc.Value.MustParseFix(b)
		result, err = ops.Fix(x, y)
	default:
		result, err = ops.String(a, b)
	}
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Value.FormatBool(result))
	return nil
}
