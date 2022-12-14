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
	case calc.IsBigInt(a) && calc.IsBigInt(b):
		x, y := calc.MustParseBigInt(a), calc.MustParseBigInt(b)
		result, err = ops.BigInt(x, y)
	case calc.IsDecimal(a) && calc.IsDecimal(b):
		x, y := calc.MustParseDecimal(a), calc.MustParseDecimal(b)
		result, err = ops.Fix(x, y)
	default:
		result, err = ops.String(a, b)
	}
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.FormatBool(result))
	return nil
}
