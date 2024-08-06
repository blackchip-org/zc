package funcs

import (
	"fmt"
	"math"
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AndBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.And(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func Bin(c zc.Calc) {
	var zero big.Int
	var b string

	x := zc.BigInt.Pop(c)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0b%b", x)
	} else {
		b = fmt.Sprintf("0b%b", x)
	}
	zc.String.Push(c, b)
	zc.BigInt.Recycle(x)
}

func Bit(c zc.Calc) {
	i := zc.Int.Pop(c)
	if i < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", i))
		return
	}
	x := zc.BigInt.Pop(c)
	b := x.Bit(i)
	zc.Uint.Push(c, b)
	zc.BigInt.Recycle(x)
}

func Bits(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	b := x.BitLen()
	zc.Int.Push(c, b)
	zc.BigInt.Recycle(x)
}

func Bytes(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	b := int(math.Ceil(float64(x.BitLen()) / 8.0))
	zc.Int.Push(c, b)
	zc.BigInt.Recycle(x)
}

func DecBigInt(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	zc.BigInt.Push(c, x)
}

func Hex(c zc.Calc) {
	var zero big.Int
	var b string

	x := zc.BigInt.Pop(c)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0x%x", x)
	} else {
		b = fmt.Sprintf("0x%x", x)
	}

	zc.String.Push(c, b)
	zc.BigInt.Recycle(x)
}

func Lsh(c zc.Calc) {
	n := zc.Uint.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Lsh(x, n)
	zc.BigInt.Push(c, x)
}

func NotBigInt(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	x.Not(x)
	zc.BigInt.Push(c, x)
}

func Oct(c zc.Calc) {
	var zero big.Int
	var b string

	x := zc.BigInt.Pop(c)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0o%o", x)
	} else {
		b = fmt.Sprintf("0o%o", x)
	}

	zc.String.Push(c, b)
	zc.BigInt.Recycle(x)
}

func OrBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Or(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func Rsh(c zc.Calc) {
	n := zc.Uint.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Rsh(x, n)
	zc.BigInt.Push(c, x)
}

func XorBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Xor(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}
