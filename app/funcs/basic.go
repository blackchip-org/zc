package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Add(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func DivBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.Div(x, y)
	zc.BigInt.Push(c, x)
}

func DivModBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	m := zc.BigInt.New()
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.DivMod(x, y, m)
	zc.BigInt.Push(c, x)
	c.Label("quo")
	zc.BigInt.Push(c, m)
	c.Label("mod")
}

func ModBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.Mod(x, y)
	zc.BigInt.Push(c, x)
}

func MulBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Mul(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func NegBigInt(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	x.Neg(x)
	zc.BigInt.Push(c, x)
}

func PowBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Exp(x, y, nil)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func RemBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.Rem(x, y)
	zc.BigInt.Push(c, x)
}

func SignBigInt(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	zc.Int.Push(c, x.Sign())
	zc.BigInt.Recycle(x)
}

func SqrtBigInt(c zc.Calc) {
	var zero big.Int
	x := zc.BigInt.Pop(c)
	if x.Cmp(&zero) < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", x.String()))
		return
	}
	x.Sqrt(x)
	zc.BigInt.Push(c, x)
}

func SubBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	x.Sub(x, y)
	zc.BigInt.Push(c, x)
	zc.BigInt.Recycle(y)
}

func QuoBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.Quo(x, y)
	zc.BigInt.Push(c, x)
}

func QuoRemBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	r := zc.BigInt.New()
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}
	x.QuoRem(x, y, r)
	zc.BigInt.Push(c, x)
	c.Label("quo")
	zc.BigInt.Push(c, r)
	c.Label("rem")
}
