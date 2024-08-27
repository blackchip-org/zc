package calc

import (
	"github.com/blackchip-org/zc/v6/pkg/coll"
	"github.com/cockroachdb/apd/v3"
)

type Decimal struct {
	stack coll.Stack[*apd.Decimal]
	mem   map[string]*apd.Decimal
	pool  *coll.Pool[apd.Decimal]
	Ctx   *apd.Context
	Flags apd.Condition
	Err   error
}

func NewDecimal() *Decimal {
	context := apd.BaseContext.WithPrecision(28)
	return &Decimal{
		Ctx:  context,
		mem:  make(map[string]*apd.Decimal),
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
	x := c.stack.Top()
	c.update(c.Ctx.Abs(x, x))
}

func (c *Decimal) Ack() {
	c.Flags = 0
	c.Err = nil
}

func (c *Decimal) Add() {
	if c.Err != nil {
		return
	}
	y := c.stack.Pop()
	x := c.stack.Top()
	c.update(c.Ctx.Add(x, x, y))
	c.pool.Recycle(y)
}

func (c *Decimal) Cbrt() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Cbrt(x, x))
}

func (c *Decimal) Ceil() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Ceil(x, x))
}

func (c *Decimal) Cmp() int {
	y := c.stack.Pop()
	x := c.stack.Pop()
	return x.Cmp(y)
}

func (c *Decimal) CmpTotal() int {
	y := c.stack.Pop()
	x := c.stack.Pop()
	return x.CmpTotal(y)
}

func (c *Decimal) Dup() {
	x1 := c.stack.Top()
	x2 := c.pool.New()
	x2.Set(x1)
	c.stack.Push(x2)
}

func (c *Decimal) Exp() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Exp(x, x))
}

func (c *Decimal) Floor() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Floor(x, x))
}

func (c *Decimal) Len() int {
	return c.stack.Len()
}

func (c *Decimal) Load(name string) {
	d, ok := c.mem[name]
	if !ok {
		panic("undefined: " + name)
	}
	c.Push(d)
}

func (c *Decimal) Ln() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Ln(x, x))
}

func (c *Decimal) Log10() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Log10(x, x))
}

func (c *Decimal) Mul() {
	if c.Err != nil {
		return
	}
	y := c.stack.Pop()
	x := c.stack.Top()
	c.update(c.Ctx.Mul(x, x, y))
	c.pool.Recycle(y)
}

func (c *Decimal) Neg() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Neg(x, x))
}

func (c *Decimal) Pop() *apd.Decimal {
	if c.Err != nil {
		panic(c.Err)
	}
	return c.stack.Pop()
}

func (c *Decimal) PopFloat64() (float64, error) {
	if c.Err != nil {
		panic(c.Err)
	}
	return c.stack.Pop().Float64()
}

func (c *Decimal) PopInt64() (int64, error) {
	if c.Err != nil {
		panic(c.Err)
	}
	return c.stack.Pop().Int64()
}

func (c *Decimal) PopString() string {
	if c.Err != nil {
		panic(c.Err)
	}
	return c.stack.Pop().String()
}

func (c *Decimal) Pow() {
	if c.Err != nil {
		return
	}
	y := c.stack.Pop()
	x := c.stack.Top()
	c.update(c.Ctx.Pow(x, x, y))
	c.pool.Recycle(y)
}

func (c *Decimal) Push(d *apd.Decimal) {
	if c.Err != nil {
		return
	}
	c.stack.Push(d)
}

func (c *Decimal) PushFloat64(f float64) {
	if c.Err != nil {
		return
	}
	d := c.pool.New()
	if _, err := d.SetFloat64(f); err != nil {
		c.pool.Recycle(d)
		c.update(0, err)
	} else {
		c.stack.Push(d)
	}
}

func (c *Decimal) PushInt(vals ...int) {
	if c.Err != nil {
		return
	}
	for _, val := range vals {
		c.PushInt64(int64(val))
	}
}

func (c *Decimal) PushInt64(vals ...int64) {
	if c.Err != nil {
		return
	}
	for _, val := range vals {
		d := c.pool.New()
		d.SetInt64(val)
		c.stack.Push(d)
	}
}

func (c *Decimal) PushString(s string) {
	if c.Err != nil {
		return
	}
	d := c.pool.New()
	if _, _, err := c.Ctx.SetString(d, s); err != nil {
		c.pool.Recycle(d)
		c.update(0, err)
	} else {
		c.stack.Push(d)
	}
}

func (c *Decimal) Quantize(exp int32) {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Quantize(x, x, exp))
}

func (c *Decimal) Quo() {
	if c.Err != nil {
		return
	}
	y := c.stack.Pop()
	x := c.stack.Top()
	c.update(c.Ctx.Quo(x, x, y))
	c.pool.Recycle(y)
}

func (c *Decimal) QuoInteger() {
	if c.Err != nil {
		return
	}
	y := c.stack.Pop()
	x := c.stack.Top()
	c.update(c.Ctx.QuoInteger(x, x, y))
	c.pool.Recycle(y)
}

func (c *Decimal) Reduce() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	_, cond, err := c.Ctx.Reduce(x, x)
	c.update(cond, err)
}

func (c *Decimal) Rem() {
	if c.Err != nil {
		return
	}
	y := c.stack.Pop()
	x := c.stack.Top()
	c.update(c.Ctx.Rem(x, x, y))
	c.pool.Recycle(y)
}

func (c *Decimal) Round() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Round(x, x))
}

func (c *Decimal) RoundToIntegeralExact() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.RoundToIntegralExact(x, x))
}

func (c *Decimal) RoundToIntegeralValue() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.RoundToIntegralValue(x, x))
}

func (c *Decimal) Save(name string) {
	c.mem[name] = c.stack.Pop()
}

func (c *Decimal) Sqrt() {
	if c.Err != nil {
		return
	}
	x := c.stack.Top()
	c.update(c.Ctx.Sqrt(x, x))
}

func (c *Decimal) Sub() {
	if c.Err != nil {
		return
	}
	y := c.stack.Pop()
	x := c.stack.Top()
	c.update(c.Ctx.Sub(x, x, y))
	c.pool.Recycle(y)
}
