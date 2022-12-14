package zlib

import "github.com/blackchip-org/zc"

func N(calc *zc.Calc) error {
	n := calc.Stack.Len()
	calc.Stack.Push(calc.Val.FormatInt(n))
	return nil
}
