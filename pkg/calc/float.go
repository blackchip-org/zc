package calc

import (
	"math/big"

	"github.com/blackchip-org/zc/v6/pkg/coll"
)

type BigFloat struct {
	coll.Stack[*big.Float]
	pool *coll.Pool[big.Float]
	prec uint
}

func NewBigFloat() *BigFloat {
	return NewBigFloatWithPrec(53)
}

func NewBigFloatWithPrec(prec uint) *BigFloat {
	return &BigFloat{
		pool: coll.NewPool[big.Float](4),
		prec: prec,
	}
}

func (c *BigFloat) Add() {
	y := c.Pop()
	x := c.Top()
	x.Add(x, y)
	c.pool.Recycle(y)
}

func (c *BigFloat) Dup() {
	y := c.Top()
	x := c.pool.New()
	x.SetPrec(c.prec)
	x.Set(y)
	c.Push(x)
}

func (c *BigFloat) PushInt(i int) {
	x := c.pool.New()
	x.SetPrec(c.prec)
	x.SetInt64(int64(i))
	c.Push(x)
}
