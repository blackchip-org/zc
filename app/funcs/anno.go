package funcs

import "github.com/blackchip-org/zc/v6"

func Label(c zc.Calc) {
	l := zc.String.Pop(c)
	c.SetLabel(l)
}

func Unit(c zc.Calc) {
	u := zc.String.Pop(c)
	c.SetUnit(u)
}
