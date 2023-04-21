package ops

import "github.com/blackchip-org/zc/pkg/zc"

func Eval(c zc.Calc) {
	fn := zc.PopString(c)
	c.Eval(fn)
}

func Filter(c zc.Calc) {
	var rs []string
	fn := zc.PopString(c)
	for _, v := range c.Stack() {
		dc := c.Derive()
		dc.Push(v)
		dc.Eval(fn)
		out, ok := dc.Pop()
		if !ok {
			zc.ErrInvalidFunc(c, fn, "no results")
			return
		}
		r, err := zc.Bool.Parse(out)
		if err != nil {
			c.SetError(err)
			return
		}
		if r {
			rs = append(rs, v)
		}
	}
	c.SetStack(rs)
}

func Fold(c zc.Calc) {
	fn := zc.PopString(c)
	for c.StackLen() > 1 {
		before := c.StackLen()
		c.Eval(fn)
		if c.Error() != nil {
			return
		}
		if c.StackLen() >= before {
			zc.ErrInvalidFunc(c, fn, "does not reduce")
			return
		}
	}
}

func Map(c zc.Calc) {
	var rs []string
	fn := zc.PopString(c)
	for _, a := range c.Stack() {
		dc := c.Derive()
		dc.Push(a)
		if err := dc.Eval(fn); err != nil {
			c.SetError(err)
			return
		}
		if r0, ok := dc.Pop(); ok {
			rs = append(rs, r0)
		}
	}
	c.SetStack(rs)
}

func Repeat(c zc.Calc) {
	n := zc.PopInt(c)
	fn := zc.PopString(c)
	for i := 0; i < n; i++ {
		c.Eval(fn)
	}
}
