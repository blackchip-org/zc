package funcs

import (
	"github.com/blackchip-org/zc/v6"
)

func Eval(c zc.Calc) {
	fn := zc.String.Pop(c)
	c.Eval(fn)
}

func Map(c zc.Calc) {
	var rs []zc.Item
	opName := zc.String.Pop(c)
	for _, a := range c.Stack() {
		dc := c.New()
		dc.Push(a)
		if err := dc.Eval(opName); err != nil {
			c.Raise(err)
			return
		}
		if dc.Len() > 0 {
			rs = append(rs, dc.Pop())

		}
	}
	c.SetStack(rs)
}
