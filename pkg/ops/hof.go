package ops

import (
	"fmt"

	"github.com/blackchip-org/zc"
)

var ErrNoReduce = func(name string) error { return fmt.Errorf("%v: function does not reduce", name) }

func Fold(c zc.Calc) {
	fn := zc.PopString(c)
	for c.StackLen() > 1 {
		before := c.StackLen()
		c.Eval(fn)
		if c.Error() != nil {
			return
		}
		if c.StackLen() >= before {
			c.SetError(ErrNoReduce(fn))
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
