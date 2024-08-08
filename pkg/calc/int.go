package calc

import (
	"math/big"

	"github.com/blackchip-org/zc/v6/pkg/coll"
)

type BigInt struct {
	coll.Stack[*big.Int]
	pool *coll.Pool[big.Int]
}

func NewBigInt() *BigInt {
	return &BigInt{pool: coll.NewPool[big.Int](4)}
}

func (c *BigInt) Abs() {
	x := c.Top()
	x.Abs(x)
}

func (c *BigInt) Add() {
	y := c.Pop()
	x := c.Top()
	x.Add(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) And() {
	y := c.Pop()
	x := c.Top()
	x.And(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) AndNot() {
	y := c.Pop()
	x := c.Top()
	x.AndNot(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) Binomial(n, k int64) {
	z := c.pool.New()
	z.Binomial(n, k)
	c.Push(z)
}

func (c *BigInt) Cmp() int {
	y := c.Pop()
	x := c.Pop()
	c.pool.Recycle(x, y)
	return x.Cmp(y)
}

func (c *BigInt) CmpAbs() int {
	y := c.Pop()
	x := c.Pop()
	c.pool.Recycle(x, y)
	return x.CmpAbs(y)
}

func (c *BigInt) Div() {
	y := c.Pop()
	x := c.Top()
	x.Div(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) DivMod() {
	m := c.pool.New()
	y := c.Pop()
	x := c.Pop()
	x.DivMod(x, y, m)
	c.Push(x) // div
	c.Push(m) // mod
	c.pool.Recycle(y)
}

func (c *BigInt) Dup() {
	x := c.Top()
	y := c.pool.New()
	y.Set(x)
	c.Push(y)
}

func (c *BigInt) Exp() {
	m := c.Pop()
	y := c.Pop()
	x := c.Top()
	x.Exp(x, y, m)
	c.pool.Recycle(m, y)
}

func (c *BigInt) GCD() {
	b := c.Pop()
	a := c.Pop()
	y := c.Pop()
	x := c.Top()
	x.GCD(x, y, a, b)
	c.pool.Recycle(b, a, y)
}

func (c *BigInt) GCD2() {
	b := c.Pop()
	a := c.Top()
	a.GCD(nil, nil, a, b)
	c.pool.Recycle(b)
}

func (c *BigInt) Lsh(n uint) {
	x := c.Top()
	x.Lsh(x, n)
}

func (c *BigInt) Mod() {
	y := c.Pop()
	x := c.Top()
	x.Mod(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) ModInverse() {
	n := c.Pop()
	g := c.Top()
	g.ModInverse(g, n)
	c.pool.Recycle(n)
}

func (c *BigInt) ModSqrt() {
	p := c.Pop()
	x := c.Top()
	x.ModSqrt(x, p)
	c.pool.Recycle(p)
}

func (c *BigInt) Mul() {
	y := c.Pop()
	x := c.Top()
	x.Mul(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) MulRange(a, b int64) {
	x := c.pool.New()
	x.MulRange(a, b)
	c.Push(x)
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
	y := c.Pop()
	x := c.Top()
	x.Or(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) Quo() {
	y := c.Pop()
	x := c.Top()
	x.Quo(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) QuoRem() {
	r := c.pool.New()
	y := c.Pop()
	x := c.Pop()
	x.QuoRem(x, y, r)
	c.Push(x) // quo
	c.Push(r) // rem
	c.pool.Recycle(y)
}

func (c *BigInt) Pow() {
	y := c.Pop()
	x := c.Top()
	x.Exp(x, y, nil)
	c.pool.Recycle(y)
}

func (c *BigInt) PushInt(vals ...int) {
	for _, v := range vals {
		i := c.pool.New()
		i.SetInt64(int64(v))
		c.Push(i)
	}
}

func (c *BigInt) PushInt64(vals ...int64) {
	for _, v := range vals {
		i := c.pool.New()
		i.SetInt64(v)
		c.Push(i)
	}
}

func (c *BigInt) PushUint(vals ...uint) {
	for _, v := range vals {
		i := c.pool.New()
		i.SetUint64(uint64(v))
		c.Push(i)
	}
}

func (c *BigInt) Rem() {
	y := c.Pop()
	x := c.Top()
	x.Rem(x, y)
	c.pool.Recycle(y)
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
	y := c.Pop()
	x := c.Top()
	x.Sub(x, y)
	c.pool.Recycle(y)
}

func (c *BigInt) Xor() {
	y := c.Pop()
	x := c.Top()
	x.Xor(x, y)
	c.pool.Recycle(y)
}
