package app

import (
	"github.com/blackchip-org/zc/v6"
)

type Calc struct {
	zc.Stack[zc.Item]
}

func NewCalc() *Calc {
	return &Calc{}
}

func (c *Calc) PushVal(vals ...any) {
	for _, val := range vals {
		c.Push(zc.Item{Val: val})
	}
}

func (c *Calc) Do(op zc.Op) {
	nParams := len(op.Params)
	for i, param := range op.Params {
		arg := c.Get(nParams - i - 1)
		convVal, argType, ok := param.From(arg.Val)
		if !ok {
			panic("cannot convert")
		}
		if argType != param {
			argType.Recycle(arg.Val)
			arg.Val = convVal
			c.Set(nParams-i-1, arg)
		}
	}
	op.Func(zc.OpEnv{
		Stack: &c.Stack,
	})
}
