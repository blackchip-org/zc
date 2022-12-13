package zlib

import (
	"strings"

	"github.com/blackchip-org/zc"
)

func Len(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	r := len(a)
	calc.Stack.Push(zc.FormatInt(r))
	return nil
}

func StartsWith(calc *zc.Calc) error {
	prefix, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	str, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	r := strings.HasPrefix(str, prefix)
	calc.Stack.Push(zc.FormatBool(r))
	return nil
}