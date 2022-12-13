package zlib

import (
	"fmt"

	"github.com/blackchip-org/zc"
)

func Print(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	fmt.Print(a)
	return nil
}
