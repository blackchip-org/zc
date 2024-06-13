package calc

import "math/big"

type BigInt struct {
	Calc[*big.Int]
}

func cloneBigInt(src *big.Int) *big.Int {
	var dest big.Int
	dest.Set(src)
	return &dest
}

func newBigIntEnv() Calc[*big.Int] {
	return Calc[*big.Int]{Clone: cloneBigInt}
}

func NewBigInt() *BigInt {
	return &BigInt{Calc: newBigIntEnv()}
}

func (c *BigInt) Abs() {
	if c.err != nil {
		return
	}
	c.AssertArgs(1)
	x := c.Top()
	x.Abs(x)
}

func (c *BigInt) Add() {
	if c.err != nil {
		return
	}
	c.AssertArgs(2)

	y := c.Release()
	x := c.Top()
	x.Add(x, y)
}

func (c *BigInt) And() {
	if c.err != nil {
		return
	}
	c.AssertArgs(2)

	y := c.Release()
	x := c.Top()
	x.And(x, y)
}

func (c *BigInt) AndNot() {
	if c.err != nil {
		return
	}
	c.AssertArgs(2)

	y := c.Release()
	x := c.Top()
	x.AndNot(x, y)
}

func (c *BigInt) Binomial(n, k int64) {
	z, ok := c.Recycle()
	if ok {
		z.Binomial(n, k)
	} else {
		z := new(big.Int)
		z.Binomial(n, k)
		c.Push(z)
	}
}

func (c *BigInt) Bit(i int) uint {
	if c.err != nil {
		return 0
	}
	c.AssertArgs(1)
	return c.Top().Bit(i)
}

func (c *BigInt) BitLen() int {
	if c.err != nil {
		return 0
	}
	c.AssertArgs(0)
	return c.Top().BitLen()
}

func (c *BigInt) Dup() {
	if c.err != nil {
		return
	}
	c.AssertArgs(1)

	x := c.Top()
	y, ok := c.Recycle()
	if ok {
		y.Set(x)
	} else {
		y := new(big.Int)
		y.Set(x)
		c.Push(y)
	}
}

func (c *BigInt) PushInt(vals ...int) {
	for _, v := range vals {
		i, ok := c.Recycle()
		if ok {
			i.SetInt64(int64(v))
		} else {
			i = big.NewInt(int64(v))
			c.Push(i)
		}
	}
}

func (c *BigInt) Mul() {
	if c.err != nil {
		return
	}
	c.AssertArgs(2)

	y := c.Release()
	x := c.Top()
	x.Mul(x, y)
}

func (c *BigInt) Pow() {
	if c.err != nil {
		return
	}
	c.AssertArgs(2)

	y := c.Release()
	x := c.Top()
	x.Exp(x, y, nil)
}

func (c *BigInt) Sqrt() {
	if c.err != nil {
		return
	}
	c.AssertArgs(1)

	x := c.Top()
	x.Sqrt(x)
}

func (c *BigInt) Sub() {
	if c.err != nil {
		return
	}
	c.AssertArgs(2)

	y := c.Release()
	x := c.Top()
	x.Sub(x, y)
}
