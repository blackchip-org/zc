package funcs

import (
	"fmt"
	"math"
	"math/big"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/pkg/calc"
)

func AndBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.And(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}

func BcdDec(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	c := calc.NewBigInt()

	c.PushInt(1)
	m := c.Pop()

	c.PushInt(0) // acc
	bytes := x.Bytes()
	for i := len(bytes) - 1; i >= 0; i-- {
		b := bytes[i]
		lo := int(b & 0xf)
		hi := int(b >> 4)
		c.PushInt((hi*10 + lo))
		c.Push(m)
		c.Mul()

		c.Add()

		c.Push(m)
		c.PushInt(100)
		c.Mul()
		m = c.Pop()
	}
	zc.BigInt.Push(e, c.Pop())
}

func Bin(e *zc.OpEnv) {
	var zero big.Int
	var b string

	ix := e.Pop()

	x := zc.BigInt.As(ix.Val)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0b%b", x)
	} else {
		b = fmt.Sprintf("0b%b", x)
	}

	ix.Repr = b
	e.Push(ix)
}

func Bit(e *zc.OpEnv) {
	i := zc.Int.Pop(e)
	if i < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", i)
		return
	}
	x := zc.BigInt.Pop(e)
	b := x.Bit(i)
	zc.Uint.Push(e, b)
	zc.BigInt.Recycle(x)
}

func Bits(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	b := x.BitLen()
	zc.Int.Push(e, b)
	zc.BigInt.Recycle(x)
}

func Bytes(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	b := int(math.Ceil(float64(x.BitLen()) / 8.0))
	zc.Int.Push(e, b)
	zc.BigInt.Recycle(x)
}

func DecBcd(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	c := calc.NewBigInt()

	c.PushInt(0)
	a := c.Pop()
	sh := uint(0)
	for {
		c.Push(x)
		c.PushInt(10)

		c.QuoRem() // x, lo
		lo := c.Pop()

		c.PushInt(10) // x, 10
		c.QuoRem()    // x, hi
		c.Lsh(4)      // hi <<= 4
		c.Push(lo)    // x, hi, lo
		c.Add()       // x, digits

		c.Lsh(sh) // digits <<= sh
		c.Push(a) // x, digits, a
		c.Add()   // x, a
		a = c.Pop()

		c.PushInt(0)
		if c.Cmp() == 0 {
			break
		}
		sh += 8
	}
	zc.BigInt.Push(e, a)
}

func DecBigInt(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	zc.BigInt.Push(e, x)
}

func Hex(e *zc.OpEnv) {
	var zero big.Int
	var b string

	ix := e.Pop()

	x := zc.BigInt.As(ix.Val)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0x%x", x)
	} else {
		b = fmt.Sprintf("0x%x", x)
	}

	ix.Repr = b
	e.Push(ix)
}

func Lsh(e *zc.OpEnv) {
	n := zc.Uint.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Lsh(x, n)
	zc.BigInt.Push(e, x)
}

func NotBigInt(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	x.Not(x)
	zc.BigInt.Push(e, x)
}

func Oct(e *zc.OpEnv) {
	var zero big.Int
	var b string

	ix := e.Pop()

	x := zc.BigInt.As(ix.Val)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0o%o", x)
	} else {
		b = fmt.Sprintf("0o%o", x)
	}

	ix.Repr = b
	e.Push(ix)
}

func OrBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Or(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}

func Rsh(e *zc.OpEnv) {
	n := zc.Uint.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Rsh(x, n)
	zc.BigInt.Push(e, x)
}

func XorBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Xor(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}
