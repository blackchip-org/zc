package funcs

import (
	"math"
	"math/big"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vars"
	"github.com/blackchip-org/zc/v6/msg"
)

func AddBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Add(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func AddFloat32(c zc.Calc) {
	y := zc.Float32.Pop(c)
	x := zc.Float32.Pop(c)
	zc.Float32.Push(c, x+y)
}

func AddFloat64(c zc.Calc) {
	y := zc.Float64.Pop(c)
	x := zc.Float64.Pop(c)
	zc.Float64.Push(c, x+y)
}

func CbrtFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	if x < 0 {
		c.Raise(msg.ErrInvalidArg("%v < 0", x))
		return
	}
	z := math.Cbrt(x)
	zc.Float64.Push(c, z)
}

func DivBigFloat(c zc.Calc) {
	var zero big.Float
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)

	if y.Cmp(&zero) == 0 {
		c.Raise(msg.ErrDivisionByZero())
		return
	}

	x.Quo(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func DivFloat32(c zc.Calc) {
	y := zc.Float32.Pop(c)
	x := zc.Float32.Pop(c)
	if y == 0 {
		c.Raise(msg.ErrDivisionByZero())
		return
	}
	zc.Float32.Push(c, x/y)
}

func DivFloat64(c zc.Calc) {
	y := zc.Float64.Pop(c)
	x := zc.Float64.Pop(c)
	if y == 0 {
		c.Raise(msg.ErrDivisionByZero())
		return
	}
	zc.Float64.Push(c, x/y)
}

func ExpFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Exp(x)
	zc.Float64.Push(c, z)
}

func LogFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Log(x)
	zc.Float64.Push(c, z)
}

func Log10Float64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Log10(x)
	zc.Float64.Push(c, z)
}

func MulBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Mul(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func MulFloat32(c zc.Calc) {
	y := zc.Float32.Pop(c)
	x := zc.Float32.Pop(c)
	zc.Float32.Push(c, x*y)
}

func MulFloat64(c zc.Calc) {
	y := zc.Float64.Pop(c)
	x := zc.Float64.Pop(c)
	zc.Float64.Push(c, x*y)
}

func NegBigFloat(c zc.Calc) {
	x := zc.BigFloat.Pop(c)
	x.Neg(x)
	zc.BigFloat.Push(c, x)
}

func NegFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	zc.Float64.Push(c, -x)
}

func PowFloat64(c zc.Calc) {
	y := zc.Float64.Pop(c)
	x := zc.Float64.Pop(c)
	zc.Float64.Push(c, math.Pow(x, y))
}

func PrecFloat(c zc.Calc) {
	v := vars.ForFloat(c)
	prec := zc.Uint.Pop(c)
	v.Prec = prec
	c.Notify(msg.PrecisionSet(prec))
}

func PrecFloatGet(c zc.Calc) {
	v := vars.ForFloat(c)
	zc.Uint.Push(c, v.Prec)
	c.SetLabel(msg.Precision())
}

func SignBigFloat(c zc.Calc) {
	x := zc.BigFloat.Pop(c)
	zc.Int.Push(c, x.Sign())
	zc.BigFloat.Recycle(x)
}

func SignFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	if x > 0 {
		zc.Float64.Push(c, 1.0)
	} else if x < 0 {
		zc.Float64.Push(c, -1.0)
	} else {
		zc.Float64.Push(c, 0.0)
	}
}

func SqrtBigFloat(c zc.Calc) {
	var zero big.Float
	x := zc.BigFloat.Pop(c)
	if x.Cmp(&zero) < 0 {
		c.Raise(msg.ErrInvalidArg("%v < 0", x.String()))
		return
	}
	x.Sqrt(x)
	zc.BigFloat.Push(c, x)
}

func SqrtFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	if x < 0 {
		c.Raise(msg.ErrInvalidArg("%v < 0", x))
		return
	}
	zc.Float64.Push(c, math.Sqrt(x))
}

func SubBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	x.Sub(x, y)
	zc.BigFloat.Push(c, x)
	zc.BigFloat.Recycle(y)
}

func SubFloat64(c zc.Calc) {
	y := zc.Float64.Pop(c)
	x := zc.Float64.Pop(c)
	zc.Float64.Push(c, x-y)
}

func SubFloat32(c zc.Calc) {
	y := zc.Float32.Pop(c)
	x := zc.Float32.Pop(c)
	zc.Float32.Push(c, x-y)
}
