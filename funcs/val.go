package funcs

import "github.com/blackchip-org/zc"

type CompareOps struct {
	BigInt CompareBigInt
	Fix    CompareFix
	String CompareStr
}

func EvalCompareVal(calc *zc.Calc, ops CompareOps) error {
	a, b, err := calc.Pop2()
	if err != nil {
		return err
	}

	var result bool

	switch {
	case zc.IsBigInt(a) && zc.IsBigInt(b):
		x, y := zc.MustParseBigInt(a), zc.MustParseBigInt(b)
		result, err = ops.BigInt(x, y)
	case zc.IsDecimal(a) && zc.IsDecimal(b):
		x, y := zc.MustParseDecimal(a), zc.MustParseDecimal(b)
		result, err = ops.Fix(x, y)
	default:
		result, err = ops.String(a, b)
	}
	if err != nil {
		return err
	}
	calc.Stack.Push(zc.FormatBool(result))
	return nil
}
