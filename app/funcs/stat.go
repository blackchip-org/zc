package funcs

import (
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
