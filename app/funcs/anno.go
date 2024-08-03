package funcs

import "github.com/blackchip-org/zc/v6"

func Label(c zc.Calc) {
	label := zc.String.Pop(c)
	item := c.Pop()
	item.Label = label
	c.Push(item)
}

func Unit(c zc.Calc) {
	unit := zc.String.Pop(c)
	item := c.Pop()
	item.Unit = unit
	c.Push(item)
}
