package native

import (
	"math/big"

	"github.com/blackchip-org/zc"
)

func Add(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	b, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	ai, err := zc.ParseBigInt(a)
	if err != nil {
		return err
	}
	bi, err := zc.ParseBigInt(b)
	if err != nil {
		return err
	}

	var zi big.Int
	zi.Add(ai, bi)
	calc.Stack.Push(zc.FormatBigInt(&zi))
	return nil
}
