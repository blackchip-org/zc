package calc

import (
	"github.com/blackchip-org/zc/v6/pkg/coll"
	"github.com/cockroachdb/apd/v3"
)

type Decimal struct {
	coll.Stack[*apd.Decimal]
	pool  *coll.Pool[apd.Decimal]
	Ctx   *apd.Context
	Flags apd.Condition
	Err   error
}

func NewDecimal() *Decimal {
	context := apd.BaseContext.WithPrecision(28)
	return &Decimal{
		Ctx:  context,
		pool: coll.NewPool[apd.Decimal](4),
	}
}

func (c *Decimal) update(cond apd.Condition, err error) {
	c.Flags |= cond
	c.Err = err
}

func (c *Decimal) Abs() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Abs(x, x))
}

func (c *Decimal) Add() {
	if c.Err != nil {
		return
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Add(x, x, y))
	c.pool.Recycle(y)
}

func (c *Decimal) Cbrt() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Cbrt(x, x))
}

func (c *Decimal) Ceil() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Ceil(x, x))
}

func (c *Decimal) Exp() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Exp(x, x))
}

func (c *Decimal) Floor() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Floor(x, x))
}

func (c Decimal) Ln() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Ln(x, x))
}

func (c Decimal) Log10() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Log10(x, x))
}

func (c *Decimal) Mul() {
	if c.Err != nil {
		return
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Mul(x, x, y))
	c.pool.Recycle(y)
}

func (c Decimal) Neg() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Neg(x, x))
}

func (c Decimal) PopFloat64() (float64, error) {
	return c.Pop().Float64()
}

func (c Decimal) PopInt64() (int64, error) {
	return c.Pop().Int64()
}

func (c Decimal) PopString() string {
	return c.Pop().String()
}

func (c Decimal) PushFloat64(f float64) {
	if c.Err != nil {
		return
	}
	d := c.pool.New()
	if _, err := d.SetFloat64(f); err != nil {
		c.pool.Recycle(d)
		c.update(0, err)
	} else {
		c.Push(d)
	}
}

func (c Decimal) PushInt(i int) {
	c.PushInt64(int64(i))
}

func (c Decimal) PushInt64(i int64) {
	if c.Err != nil {
		return
	}
	d := c.pool.New()
	d.SetInt64(i)
	c.Push(d)
}
