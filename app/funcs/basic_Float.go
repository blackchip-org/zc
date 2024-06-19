package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddBigFloat(e *zc.OpEnv) {
	y := zc.BigFloat.As(e.Pop())
	x := zc.BigFloat.As(e.Pop())
	x.Add(x, y)
	zc.BigFloat.Push(e, x)
	zc.BigFloat.Recycle(y)
}

func DivBigFloat(e *zc.OpEnv) {
	var zero big.Float
	y := zc.BigFloat.As(e.Pop())
	x := zc.BigFloat.As(e.Pop())
	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Quo(x, y)
	zc.BigFloat.Push(e, x)
	zc.BigFloat.Recycle(y)
}

func MulBigFloat(e *zc.OpEnv) {
	y := zc.BigFloat.As(e.Pop())
	x := zc.BigFloat.As(e.Pop())
	x.Mul(x, y)
	zc.BigFloat.Push(e, x)
	zc.BigFloat.Recycle(y)
}

func NegBigFloat(e *zc.OpEnv) {
	x := zc.BigFloat.As(e.Pop())
	x.Neg(x)
	zc.BigFloat.Push(e, x)
}

func SignBigFloat(e *zc.OpEnv) {
	x := zc.BigFloat.As(e.Pop())
	s := x.Sign()
	e.PushVal(s)
}

func SqrtBigFloat(e *zc.OpEnv) {
	var zero big.Float
	x := zc.BigFloat.As(e.Pop())
	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
	}
	x.Sqrt(x)
	zc.BigFloat.Push(e, x)
}

func SubBigFloat(e *zc.OpEnv) {
	y := zc.BigFloat.As(e.Pop())
	x := zc.BigFloat.As(e.Pop())
	x.Sub(x, y)
	zc.BigFloat.Push(e, x)
	zc.BigFloat.Recycle(y)
}
