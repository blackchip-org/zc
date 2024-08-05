package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Add(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func DivBigFloat(c zc.Calc) {
	var zero big.Float
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)

	if y.Cmp(&zero) == 0 {
		c.Raise(zc.ErrDivisionByZero)
		return
	}

	x.Quo(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func MulBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Mul(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func NegBigFloat(c zc.Calc) {
	x := zc.BigFloat.Pop(c)
	x.Neg(x)
	zc.BigFloat.Push(c, x)
}

func SignBigFloat(c zc.Calc) {
	x := zc.BigFloat.Pop(c)
	zc.Int.Push(c, x.Sign())
	zc.BigFloat.Recycle(x)
}

func SqrtBigFloat(c zc.Calc) {
	var zero big.Float
	x := zc.BigFloat.Pop(c)
	if x.Cmp(&zero) < 0 {
		c.Raise(zc.ErrInvalidArg("%v < 0", x.String()))
		return
	}
	x.Sqrt(x)
	zc.BigFloat.Push(c, x)
}

func SubBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Sub(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}
