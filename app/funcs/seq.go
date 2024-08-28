package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/pkg/calc"
)

func Fibonacci(c zc.Calc) {
	n := zc.Uint.Pop(c)
	switch {
	case n == 0:
		zc.BigInt.Push(c, big.NewInt(0))
	case n == 1:
		zc.BigInt.Push(c, big.NewInt(1))
	default:
		ic := calc.NewBigInt()
		ic.PushInt(0)
		ic.PushInt(1)
		for i := uint(2); i <= n; i++ {
			ic.Dup()
			ic.Rotate()
			ic.Add()
		}
		zc.BigInt.Push(c, ic.Pop())
	}
}

func FibonacciBigFloat(c zc.Calc) {
	n := zc.Uint.Pop(c)
	switch {
	case n == 0:
		zc.BigFloat.Push(c, big.NewFloat(0))
	case n == 1:
		zc.BigFloat.Push(c, big.NewFloat(1))
	default:
		fc := calc.NewBigFloatWithPrec(zc.PrecFloat128)
		fc.PushInt(0)
		fc.PushInt(1)
		for i := uint(2); i <= n; i++ {
			fc.Dup()
			fc.Rotate()
			fc.Add()
		}
		zc.BigFloat.Push(c, fc.Pop())
	}
}

func Sequence(c zc.Calc) {
	to := zc.BigInt.Pop(c)
	from := zc.BigInt.Pop(c)
	one := big.NewInt(1)

	if from.Cmp(to) <= 0 {
		i := from
		for i.Cmp(to) <= 0 {
			a := new(big.Int)
			a.Set(i)
			zc.BigInt.Push(c, a)
			i.Add(i, one)
		}
	} else {
		i := from
		for i.Cmp(to) >= 0 {
			a := new(big.Int)
			a.Set(i)
			zc.BigInt.Push(c, a)
			i.Sub(i, one)
		}
	}
}
