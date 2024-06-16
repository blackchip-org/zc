package app

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/types"
)

type Calc struct {
	zc.Stack[*zc.Item]
}

func (c *Calc) PushVal(vals ...any) {
	for _, val := range vals {
		c.Push(&zc.Item{Val: val})
	}
}

func (c *Calc) Do(op zc.Op) {
	top := c.Top()
	cTop, ok := types.BigInt.From(c, top)
	if !ok {
		panic("cannot convert")
	}
	top.Val = cTop

	next := c.Next()
	cNext, ok := types.BigInt.From(c, next)
	if !ok {
		panic("cannot convert")
	}
	next.Val = cNext

	op.Func(zc.OpEnv{
		Stack: &c.Stack,
	})
}
