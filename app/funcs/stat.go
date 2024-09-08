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

func VariancePop(c zc.Calc) {
	variance(c, 0)
}

func VarianceSamp(c zc.Calc) {
	variance(c, -1)
}

func variance(c zc.Calc, nadj int) {
	n := c.Len()
	if n == 0 {
		return
	}
	data := zc.DupItems(c.Stack())
	if err := c.Eval("average dec"); err != nil {
		return
	}
	mean := zc.Decimal.Pop(c)
	c.SetStack(data)
	zc.Decimal.Push(c, mean)
	c.Eval("[sub square] [map] 2 apply sum")
	zc.Int.Push(c, n+nadj)
	c.Eval("div")
}
