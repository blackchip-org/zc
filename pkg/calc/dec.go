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

func (c *Decimal) Clear() {
	c.Flags = 0
	c.Err = nil
}

func (c *Decimal) Cmp() int {
	y := c.Pop()
	x := c.Pop()
	c.pool.Recycle(x, y)
	return x.Cmp(y)
}

func (c *Decimal) CmpTotal() int {
	y := c.Pop()
	x := c.Pop()
	c.pool.Recycle(x, y)
	return x.CmpTotal(y)
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

func (c *Decimal) Modf() {
	y := c.Pop()
	x := c.Top()
	x.Modf(x, y)
	c.pool.Recycle(y)
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

func (c Decimal) Pow() {
	if c.Err != nil {
		return
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Pow(x, x, y))
	c.pool.Recycle(y)
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

func (c Decimal) PushInt(vals ...int) {
	if c.Err != nil {
		return
	}
	for _, val := range vals {
		c.PushInt64(int64(val))
	}
}

func (c Decimal) PushInt64(vals ...int64) {
	if c.Err != nil {
		return
	}
	for _, val := range vals {
		d := c.pool.New()
		d.SetInt64(val)
		c.Push(d)
	}
}

func (c Decimal) PushString(s string) {
	if c.Err != nil {
		return
	}
	d := c.pool.New()
	if _, _, err := c.Ctx.SetString(d, s); err != nil {
		c.pool.Recycle(d)
		c.update(0, err)
	} else {
		c.Push(d)
	}
}

func (c Decimal) Quantize(exp int32) {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Quantize(x, x, exp))
}

func (c Decimal) Quo() {
	if c.Err != nil {
		return
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Quo(x, x, y))
	c.pool.Recycle(y)
}

func (c Decimal) QuoInteger() {
	if c.Err != nil {
		return
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.QuoInteger(x, x, y))
	c.pool.Recycle(y)
}

func (c Decimal) Reduce() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	_, cond, err := c.Ctx.Reduce(x, x)
	c.update(cond, err)
}

func (c Decimal) Rem() {
	if c.Err != nil {
		return
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Rem(x, x, y))
	c.pool.Recycle(y)
}

func (c Decimal) Round() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Round(x, x))
}

func (c Decimal) RoundToIntegeralExact() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.RoundToIntegralExact(x, x))
}

func (c Decimal) RoundToIntegeralValue() {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.RoundToIntegralValue(x, x))
}

func (c Decimal) Sqrt(exp int32) {
	if c.Err != nil {
		return
	}
	x := c.Top()
	c.update(c.Ctx.Quantize(x, x, exp))
}

func (c Decimal) Sub() {
	if c.Err != nil {
		return
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Sub(x, x, y))
	c.pool.Recycle(y)
}
