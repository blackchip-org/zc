package ops

import "github.com/blackchip-org/zc"

func Clear(c zc.Calc) {
	c.SetStack([]string{})
}

func Down(c zc.Calc) {
	a0, ok := c.Pop()
	if ok {
		c.SetStack(append([]string{a0}, c.Stack()...))
	}
}

func Drop(c zc.Calc) {
	c.Pop()
}

func Dup(c zc.Calc) {
	a0 := zc.PopString(c)
	zc.PushString(c, a0)
	zc.PushString(c, a0)
}

func N(c zc.Calc) {
	r0 := len(c.Stack())
	zc.PushInt(c, r0)
}

func Reverse(c zc.Calc) {
	var rs []string
	for c.StackLen() > 0 {
		a := c.MustPop()
		rs = append(rs, a)
	}
	c.SetStack(rs)
}

func Swap(c zc.Calc) {
	a1 := zc.PopString(c)
	a0 := zc.PopString(c)
	zc.PushString(c, a1)
	zc.PushString(c, a0)
}

func Take(c zc.Calc) {
	var rs []string
	n := zc.PopInt(c)
	for i := 0; i < n; i++ {
		a, ok := c.Pop()
		if !ok {
			break
		}
		rs = append([]string{a}, rs...)
	}
	c.SetStack(rs)
}

const Top = "1 take"

func Up(c zc.Calc) {
	s := c.Stack()
	c.SetStack(append(s[1:], s[0]))
}
