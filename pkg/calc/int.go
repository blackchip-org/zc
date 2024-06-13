package calc

import "math/big"

type BigInt struct {
	env *Env[*big.Int]
}

func NewBigInt() *BigInt {
	return &BigInt{env: NewBigIntEnv()}
}

func (c *BigInt) Dup() {
	x := c.env.Top()
	y, ok := c.env.Recycle()
	if ok {
		y.Set(x)
	} else {
		y := new(big.Int)
		y.Set(x)
		c.env.Push(y)
	}
}

func (c *BigInt) Pop() *big.Int {
	return c.env.Pop()
}

func (c *BigInt) PushInt(v int) {
	i, ok := c.env.Recycle()
	if ok {
		i.SetInt64(int64(v))
	} else {
		i = big.NewInt(int64(v))
		c.env.Push(i)
	}
}

func (c *BigInt) Add() {
	y := c.env.Release()
	x := c.env.Release()
	x.Add(x, y)
	c.env.Push(x)
}

func (c *BigInt) Mul() {
	y := c.env.Release()
	x := c.env.Release()
	x.Mul(x, y)
	c.env.Push(x)
}
