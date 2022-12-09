package native

import (
	"fmt"

	"github.com/blackchip-org/zc"
)

func Bin(calc *zc.Calc) error {
	v, err := calc.PopBigInt()
	if err != nil {
		return err
	}
	calc.Stack.Push(fmt.Sprintf("0b%b", v))
	return nil
}

func Hex(calc *zc.Calc) error {
	v, err := calc.PopBigInt()
	if err != nil {
		return err
	}
	calc.Stack.Push(fmt.Sprintf("0x%x", v))
	return nil
}

func Oct(calc *zc.Calc) error {
	v, err := calc.PopBigInt()
	if err != nil {
		return err
	}
	calc.Stack.Push(fmt.Sprintf("0o%o", v))
	return nil
}
