package calc

import (
	"github.com/blackchip-org/zc/v6/pkg/coll"
	"github.com/cockroachdb/apd/v3"
)

type Decimal struct {
	coll.Stack[*apd.Decimal]
	mem   map[string]*apd.Decimal
	pool  *coll.Pool[apd.Decimal]
	Ctx   *apd.Context
	Flags apd.Condition
	Err   error
}

func NewDecimal() *Decimal {
	context := apd.BaseContext.WithPrecision(28)
	d := &Decimal{
		Ctx:  context,
		mem:  make(map[string]*apd.Decimal),
		pool: coll.NewPool[apd.Decimal](4),
	}
	return d
}

func (c *Decimal) update(cond apd.Condition, err error) {
	c.Flags |= cond
	c.Err = err
}

func (c *Decimal) Abs() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Abs(x, x))
	c.Emit("abs")
	return c
}

func (c *Decimal) Ack() *Decimal {
	c.Flags = 0
	c.Err = nil
	return c
}

func (c *Decimal) Add() *Decimal {
	if c.Err != nil {
		return c
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Add(x, x, y))
	c.pool.Recycle(y)
	c.Emit("add")
	return c
}

func (c *Decimal) Cbrt() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Cbrt(x, x))
	c.Emit("cbrt")
	return c
}

func (c *Decimal) Ceil() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Ceil(x, x))
	c.Emit("ceil")
	return c
}

func (c *Decimal) Cmp() int {
	y := c.Pop()
	x := c.Pop()
	z := x.Cmp(y)
	c.Emitf("cmp", "%v", z)
	return z
}

func (c *Decimal) CmpTotal() int {
	y := c.Pop()
	x := c.Pop()
	c.Emit("cmp-total")
	return x.CmpTotal(y)
}

func (c *Decimal) Dup() *Decimal {
	x1 := c.Top()
	x2 := c.pool.New()
	x2.Set(x1)
	c.Emit("dup")
	c.Push(x2)
	return c
}

func (c *Decimal) Exp() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Exp(x, x))
	c.Emit("exp")
	return c
}

func (c *Decimal) Floor() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Floor(x, x))
	c.Emit("floor")
	return c
}

func (c *Decimal) Load(name string) *Decimal {
	d, ok := c.mem[name]
	if !ok {
		panic("undefined: " + name)
	}
	c.Emitf("load", "%v -> %v", name, d)
	c.Push(d)
	return c
}

func (c *Decimal) Ln() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Ln(x, x))
	c.Emit("ln")
	return c
}

func (c *Decimal) Log10() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Log10(x, x))
	c.Emit("log10")
	return c
}

func (c *Decimal) Mul() *Decimal {
	if c.Err != nil {
		return c
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Mul(x, x, y))
	c.pool.Recycle(y)
	c.Emit("mul")
	return c
}

func (c *Decimal) Neg() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Neg(x, x))
	c.Emit("neg")
	return c
}

func (c *Decimal) PopFloat64() (float64, error) {
	if c.Err != nil {
		panic(c.Err)
	}
	return c.Pop().Float64()
}

func (c *Decimal) PopInt64() (int64, error) {
	if c.Err != nil {
		panic(c.Err)
	}
	return c.Pop().Int64()
}

func (c *Decimal) PopString() string {
	if c.Err != nil {
		panic(c.Err)
	}
	return c.Pop().String()
}

func (c *Decimal) Pow() *Decimal {
	if c.Err != nil {
		return c
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Pow(x, x, y))
	c.pool.Recycle(y)
	c.Emit("pow")
	return c
}

func (c *Decimal) Push(ds ...*apd.Decimal) *Decimal {
	if c.Err != nil {
		return c
	}
	for _, d := range ds {
		d2 := c.pool.New()
		d2.Set(d)
		c.Stack.Push(d2)
	}
	return c
}

func (c *Decimal) PushFloat64(fs ...float64) *Decimal {
	if c.Err != nil {
		return c
	}
	for _, f := range fs {
		d := c.pool.New()
		if _, err := d.SetFloat64(f); err != nil {
			c.pool.Recycle(d)
			c.update(0, err)
			return c
		} else {
			c.Push(d)
		}
	}
	return c
}

func (c *Decimal) PushInt(vals ...int) *Decimal {
	if c.Err != nil {
		return c
	}
	for _, val := range vals {
		c.PushInt64(int64(val))
	}
	return c
}

func (c *Decimal) PushInt64(vals ...int64) *Decimal {
	if c.Err != nil {
		return c
	}
	for _, val := range vals {
		d := c.pool.New()
		d.SetInt64(val)
		c.Push(d)
	}
	return c
}

func (c *Decimal) PushString(ss ...string) *Decimal {
	if c.Err != nil {
		return c
	}
	for _, s := range ss {
		d := c.pool.New()
		if _, _, err := c.Ctx.SetString(d, s); err != nil {
			c.pool.Recycle(d)
			c.update(0, err)
			return c
		} else {
			c.Push(d)
		}
	}
	return c
}

func (c *Decimal) Quantize(exp int32) *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Quantize(x, x, exp))
	c.Emitf("quantize", "%v", exp)
	return c
}

func (c *Decimal) Quo() *Decimal {
	if c.Err != nil {
		return c
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Quo(x, x, y))
	c.pool.Recycle(y)
	c.Emit("quo")
	return c
}

func (c *Decimal) QuoInteger() *Decimal {
	if c.Err != nil {
		return c
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.QuoInteger(x, x, y))
	c.pool.Recycle(y)
	c.Emit("quo-integer")
	return c
}

func (c *Decimal) Reduce() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	_, cond, err := c.Ctx.Reduce(x, x)
	c.update(cond, err)
	c.Emit("reduce")
	return c
}

func (c *Decimal) Rem() *Decimal {
	if c.Err != nil {
		return c
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Rem(x, x, y))
	c.pool.Recycle(y)
	c.Emit("rem")
	return c
}

func (c *Decimal) Round() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Round(x, x))
	c.Emit("round")
	return c
}

func (c *Decimal) RoundToIntegeralExact() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.RoundToIntegralExact(x, x))
	c.Emit("round-to-integral-exact")
	return c
}

func (c *Decimal) RoundToIntegeralValue() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.RoundToIntegralValue(x, x))
	c.Emit("round-to-integral-value")
	return c
}

func (c *Decimal) Save(name string) *Decimal {
	c.mem[name] = c.Pop()
	c.Emitf("save", "%v <- %v", name, c.mem[name])
	return c
}

func (c *Decimal) Sqrt() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	c.update(c.Ctx.Sqrt(x, x))
	c.Emit("sqrt")
	return c
}

func (c *Decimal) Sub() *Decimal {
	if c.Err != nil {
		return c
	}
	y := c.Pop()
	x := c.Top()
	c.update(c.Ctx.Sub(x, x, y))
	c.pool.Recycle(y)
	c.Emit("sub")
	return c
}

func (c *Decimal) Trunc() *Decimal {
	if c.Err != nil {
		return c
	}
	x := c.Top()
	fctx := *c.Ctx
	fctx.Rounding = apd.RoundDown
	c.update(fctx.RoundToIntegralExact(x, x))
	c.Emit("trunc")
	return c
}
