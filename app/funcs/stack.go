package funcs

import "github.com/blackchip-org/zc/v6"

func Clear(c zc.Calc) {
	c.SetStack([]zc.Item{})
}

func Dup(c zc.Calc) {
	x := c.Pop()
	x2 := x
	x2.TypeVal = x.Type.Dup(x.TypeVal)
	c.Push(x)
	c.Push(x2)
}
