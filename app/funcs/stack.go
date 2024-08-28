package funcs

import "github.com/blackchip-org/zc/v6"

func Clear(c zc.Calc) {
	c.SetStack([]zc.Item{})
}

func Drop(c zc.Calc) {
	c.Pop()
}

func Dup(c zc.Calc) {
	x := c.Pop()
	x2 := x
	x2.TypeVal = x.Type.Dup(x.TypeVal)
	c.Push(x)
	c.Push(x2)
}

func N(c zc.Calc) {
	zc.Uint.Push(c, uint(c.Len()))
}

func Rotate(c zc.Calc) {
	z := c.Pop()
	y := c.Pop()
	x := c.Pop()
	c.Push(z)
	c.Push(x)
	c.Push(y)
}

func Swap(c zc.Calc) {
	y := c.Pop()
	x := c.Pop()
	c.Push(y)
	c.Push(x)
}

func Tuck(c zc.Calc) {
	y := c.Pop()
	x := c.Pop()
	z := y.Type.Dup(y.TypeVal)
	c.Push(zc.Item{TypeVal: z, Type: y.Type})
	c.Push(x)
	c.Push(y)
}
