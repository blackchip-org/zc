package funcs

import (
	"slices"
	"strings"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/msg"
)

func Apply(c zc.Calc) {
	nArgs := int(zc.Uint32.Pop(c))
	opName := zc.String.Pop(c)
	if c.Len() < nArgs {
		c.Raise(msg.ErrNotEnoughArgs())
		return
	}
	var args []string
	for i := 0; i < nArgs; i++ {
		args = slices.Insert(args, 0, zc.String.Pop(c))
	}
	zc.String.Push(c, (strings.Join(args, " ")))
	c.Eval(opName)
}

func Eval(c zc.Calc) {
	fn := zc.String.Pop(c)
	c.Eval(fn)
}

func Filter(c zc.Calc) {
	var rs []zc.Item
	opName := zc.String.Pop(c)
	for _, v := range c.Stack() {
		dc := c.New()
		dc.Push(v)
		dc.Eval(opName)
		if dc.Len() == 0 {
			c.Raise(msg.ErrNoReturnValues())
			return
		}
		out := dc.Pop()
		r, ok, _ := zc.Bool.Parse(nil, out.Val())
		if !ok {
			c.Raise(msg.ErrUnexpectedType(out.Val()))
			return
		}
		if r.(bool) {
			rs = append(rs, v)
		}
	}
	c.SetStack(rs)
}

func Fold(c zc.Calc) {
	opName := zc.String.Pop(c)
	for c.Len() > 1 {
		before := c.Len()
		if err := c.Eval(opName); err != nil {
			return
		}
		if c.Len() >= before {
			c.Raise(msg.ErrDoesNotReduce())
			return
		}
	}
}

func Map(c zc.Calc) {
	var rs []zc.Item
	expr := zc.String.Pop(c)
	for _, a := range zc.DupItems(c.Stack()) {
		dc := c.New()
		dc.Push(a)
		if err := dc.Eval(expr); err != nil {
			c.Raise(err)
			return
		}
		if dc.Len() > 0 {
			rs = append(rs, dc.Pop())
		}
	}
	c.SetStack(rs)
}

func Repeat(c zc.Calc) {
	n := zc.Uint.Pop(c)
	opName := zc.String.Pop(c)
	for i := uint(0); i < n; i++ {
		c.Eval(opName)
	}
}
