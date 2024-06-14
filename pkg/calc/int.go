package calc

import (
	"math/big"
	"math/rand"
)

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
	x := c.Top()
	x.Abs(x)
}

func (c *BigInt) Add() {
	y := c.Release()
	x := c.Top()
	x.Add(x, y)
}

func (c *BigInt) And() {
	y := c.Release()
	x := c.Top()
	x.And(x, y)
}

func (c *BigInt) AndNot() {
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

func (c *BigInt) Cmp() int {
	y := c.Release()
	x := c.Release()
	return x.Cmp(y)
}

func (c *BigInt) CmpAbs() int {
	y := c.Release()
	x := c.Release()
	return x.CmpAbs(y)
}

func (c *BigInt) Div() {
	y := c.Release()
	x := c.Top()
	x.Div(x, y)
}

func (c *BigInt) DivMod() {
	y := c.Release()
	x := c.Release()
	x.DivMod(x, y, y)
	c.Push(y) // mod
	c.Push(x)
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
	m := c.Release()
	y := c.Release()
	x := c.Top()
	x.Exp(x, y, m)
}

func (c *BigInt) GCD() {
	b := c.Release()
	a := c.Release()
	y := c.Release()
	x := c.Top()
	x.GCD(x, y, a, b)
}

func (c *BigInt) GCD2() {
	b := c.Release()
	a := c.Top()
	a.GCD(nil, nil, a, b)
}

func (c *BigInt) Lsh(n uint) {
	x := c.Top()
	x.Lsh(x, n)
}

func (c *BigInt) Mod() {
	y := c.Release()
	x := c.Top()
	x.Mod(x, y)
}

func (c *BigInt) ModInverse() {
	n := c.Release()
	g := c.Top()
	g.ModInverse(g, n)
}

func (c *BigInt) ModSqrt() {
	p := c.Release()
	x := c.Top()
	x.ModSqrt(x, p)
}

func (c *BigInt) Mul() {
	y := c.Release()
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
	y := c.Release()
	x := c.Top()
	x.Or(x, y)
}

func (c *BigInt) Quo() {
	y := c.Release()
	x := c.Top()
	x.Quo(x, y)
}

func (c *BigInt) QuoRem() {
	y := c.Release()
	x := c.Release()
	x.QuoRem(x, y, y)
	c.Push(y) // rem
	c.Push(x)
}

func (c *BigInt) Pow() {
	y := c.Release()
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

func (c *BigInt) Rand(rnd *rand.Rand) {
	n := c.Top()
	n.Rand(rnd, n)
}

func (c *BigInt) Rem() {
	y := c.Release()
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

func (c *BigInt) Sign() int {
	return c.Top().Sign()
}

func (c *BigInt) Sqrt() {
	x := c.Top()
	x.Sqrt(x)
}

func (c *BigInt) Sub() {
	y := c.Release()
	x := c.Top()
	x.Sub(x, y)
}

func (c *BigInt) Xor() {
	y := c.Release()
	x := c.Top()
	x.Xor(x, y)
}
