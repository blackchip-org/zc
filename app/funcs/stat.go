package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/pkg/calc"
)

func Fact(c zc.Calc) {
	ic := calc.NewBigInt()
	n := zc.Uint.Pop(c)

	ic.PushInt(1)
	for i := uint(1); i <= n; i++ {
		ic.PushUint(i)
		ic.Mul()
	}
	zc.BigInt.Push(c, ic.Pop())
}

func FactBigFloat(c zc.Calc) {
	n := zc.Uint.Pop(c)
	acc := big.NewFloat(1)
	bi := big.NewFloat(0)
	for i := uint(1); i <= n; i++ {
		bi.SetUint64(uint64(i))
		acc.Mul(acc, bi)
	}
	zc.BigFloat.Push(c, acc)
}
