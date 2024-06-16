package calc

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

type BigInt struct {
	zc.Stack[*big.Int]
}

func (c *BigInt) Abs() {
	x := c.Top()
	x.Abs(x)
}

func (c *BigInt) Add() {
	y := c.Drop()
	x := c.Top()
	x.Add(x, y)
}

func (c *BigInt) And() {
	y := c.Drop()
	x := c.Top()
	x.And(x, y)
}

func (c *BigInt) AndNot() {
	y := c.Drop()
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

func (c *BigInt) Cmp() int {
	y := c.Top()
	x := c.Next()
	return x.Cmp(y)
}

func (c *BigInt) CmpAbs() int {
	y := c.Top()
	x := c.Next()
	return x.CmpAbs(y)
}

func (c *BigInt) Div() {
	y := c.Drop()
	x := c.Top()
	x.Div(x, y)
}

func (c *BigInt) DivMod() {
	m, ok := c.Recycle()
	if !ok {
		m = new(big.Int)
	}

	y := c.Drop()
	x := c.Drop()
	x.DivMod(x, y, m)
	c.Push(x) // div
	c.Push(m) // mod
}

func (c *BigInt) Dup() {
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

func (c *BigInt) Exp() {
	m := c.Drop()
	y := c.Drop()
	x := c.Top()
	x.Exp(x, y, m)
}

func (c *BigInt) GCD() {
	b := c.Drop()
	a := c.Drop()
	y := c.Drop()
	x := c.Top()
	x.GCD(x, y, a, b)
}

func (c *BigInt) GCD2() {
	b := c.Drop()
	a := c.Top()
	a.GCD(nil, nil, a, b)
}

func (c *BigInt) Lsh(n uint) {
	x := c.Top()
	x.Lsh(x, n)
}

func (c *BigInt) Mod() {
	y := c.Drop()
	x := c.Top()
	x.Mod(x, y)
}

func (c *BigInt) ModInverse() {
	n := c.Drop()
	g := c.Top()
	g.ModInverse(g, n)
}

func (c *BigInt) ModSqrt() {
	p := c.Drop()
	x := c.Top()
	x.ModSqrt(x, p)
}

func (c *BigInt) Mul() {
	y := c.Drop()
	x := c.Top()
	x.Mul(x, y)
}

func (c *BigInt) MulRange(a, b int64) {
	x, ok := c.Recycle()
	if ok {
		x.MulRange(a, b)
	} else {
		x := new(big.Int)
		x.MulRange(a, b)
		c.Push(x)
	}
}

func (c *BigInt) Neg() {
	x := c.Top()
	x.Neg(x)
}

func (c *BigInt) Not() {
	x := c.Top()
	x.Not(x)
}

func (c *BigInt) Or() {
	y := c.Drop()
	x := c.Top()
	x.Or(x, y)
}

func (c *BigInt) Pop() *big.Int {
	var r big.Int
	x := c.Drop()
	r.Set(x)
	return &r
}

func (c *BigInt) Quo() {
	y := c.Drop()
	x := c.Top()
	x.Quo(x, y)
}

func (c *BigInt) QuoRem() {
	r, ok := c.Recycle()
	if !ok {
		r = new(big.Int)
	}
	y := c.Drop()
	x := c.Drop()
	x.QuoRem(x, y, r)
	c.Push(x) // quo
	c.Push(r) // rem
}

func (c *BigInt) Pow() {
	y := c.Drop()
	x := c.Top()
	x.Exp(x, y, nil)
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

func (c *BigInt) Rem() {
	y := c.Drop()
	x := c.Top()
	x.Rem(x, y)
}

func (c *BigInt) Rsh(n uint) {
	x := c.Top()
	x.Rsh(x, n)
}

func (c *BigInt) SetBit(i int, b uint) {
	x := c.Top()
	x.SetBit(x, i, b)
}

func (c *BigInt) Sqrt() {
	x := c.Top()
	x.Sqrt(x)
}

func (c *BigInt) Sub() {
	y := c.Drop()
	x := c.Top()
	x.Sub(x, y)
}

func (c *BigInt) Xor() {
	y := c.Drop()
	x := c.Top()
	x.Xor(x, y)
}
