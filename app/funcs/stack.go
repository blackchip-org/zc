package funcs

import (
	"slices"

	"github.com/blackchip-org/zc/v6"
)

func Clear(c zc.Calc) {
	c.SetStack([]zc.Item{})
}

func Copy(c zc.Calc) {
	if c.Len() == 0 {
		c.Raise(zc.ErrStackEmpty)
		return
	}
	s := zc.DupItems(c.Stack())
	t := c.Temp()
	t = append(t, s...)
	c.SetTemp(t)
	c.Notify("copied")
}

func Down(c zc.Calc) {
	if c.Len() == 0 {
		c.Raise(zc.ErrStackEmpty)
		return
	}
	if c.Len() == 1 {
		return
	}
	x := c.Pop()
	c.SetStack(append([]zc.Item{x}, c.Stack()...))
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

func Flip(c zc.Calc) {
	s := c.Stack()
	t := c.Temp()
	c.SetStack(t)
	c.SetTemp(s)
}

func Label(c zc.Calc) {
	l := zc.String.Pop(c)
	c.SetLabel(l)
}

func Load(c zc.Calc) {
	n := zc.String.Pop(c)
	c.Load(n)
}

func Pop(c zc.Calc) {
	t := c.Temp()
	if len(t) == 0 {
		c.Raise(zc.ErrStackEmpty)
		return
	}
	l := len(t)
	top, t := t[l-1], t[:l-1]
	c.Push(top)
	c.SetTemp(t)
}

func PopAll(c zc.Calc) {
	s := c.Stack()
	t := c.Temp()
	if len(t) == 0 {
		c.Raise(zc.ErrStackEmpty)
		return
	}
	s = append(s, t...)
	c.SetStack(s)
	c.SetTemp([]zc.Item{})
}

func Push(c zc.Calc) {
	if c.Len() == 0 {
		c.Raise(zc.ErrStackEmpty)
		return
	}
	x := c.Pop()
	t := c.Temp()
	t = append(t, x)
	c.SetTemp(t)
}

func PushAll(c zc.Calc) {
	if c.Len() == 0 {
		c.Raise(zc.ErrStackEmpty)
		return
	}
	s := c.Stack()
	t := c.Temp()
	t = append(t, s...)
	c.SetTemp(t)
	c.SetStack([]zc.Item{})
}

func Reverse(c zc.Calc) {
	s := c.Stack()
	slices.Reverse(s)
	c.SetStack(s)
}

func Rotate(c zc.Calc) {
	z := c.Pop()
	y := c.Pop()
	x := c.Pop()
	c.Push(z)
	c.Push(x)
	c.Push(y)
}

func Size(c zc.Calc) {
	zc.Uint.Push(c, uint(c.Len()))
}

func Store(c zc.Calc) {
	n := zc.String.Pop(c)
	c.Store(n)
	c.Notify("stored")
}

func Swap(c zc.Calc) {
	y := c.Pop()
	x := c.Pop()
	c.Push(y)
	c.Push(x)
}

func Take(c zc.Calc) {
	n := int(zc.Uint.Pop(c))
	s := c.Stack()
	l := len(s)
	if n > l {
		n = l
	}
	c.SetStack(s[l-n:])
}

func Tuck(c zc.Calc) {
	y := c.Pop()
	x := c.Pop()
	z := y.Type.Dup(y.TypeVal)
	c.Push(zc.Item{TypeVal: z, Type: y.Type})
	c.Push(x)
	c.Push(y)
}

func Unit(c zc.Calc) {
	u := zc.String.Pop(c)
	c.SetUnit(u)
}

func Up(c zc.Calc) {
	s := c.Stack()
	c.SetStack(append(s[1:], s[0]))
}
