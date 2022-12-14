package funcs

import "github.com/blackchip-org/zc"

type CompareOps struct {
	BigInt CompareBigInt
	Fixed  CompareFixed
	String CompareStr
}

func EvalCompareVal(calc *zc.Calc, ops CompareOps) error {
	a, b, err := calc.Stack.Pop2()
	if err != nil {
		return err
	}

	var result bool

	switch {
	case calc.Val.IsBigInt(a) && calc.Val.IsBigInt(b):
		x, y := calc.Val.MustParseBigInt(a), calc.Val.MustParseBigInt(b)
		result, err = ops.BigInt(x, y)
	case calc.Val.IsFixed(a) && calc.Val.IsFixed(b):
		x, y := calc.Val.MustParseFixed(a), calc.Val.MustParseFixed(b)
		result, err = ops.Fixed(x, y)
	default:
		result, err = ops.String(a, b)
	}
	if err != nil {
		return err
	}
	calc.Stack.Push(calc.Val.FormatBool(result))
	return nil
}
