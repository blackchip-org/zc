package zlib

import "github.com/blackchip-org/zc"

func N(calc *zc.Calc) error {
	n := calc.Stack.Len()
	calc.Stack.PushInt(n)
	return nil
}
