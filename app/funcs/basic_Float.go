package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddBigFloat(e *zc.OpEnv) {
	y := zc.BigFloat.Pop(e)
	x := zc.BigFloat.Pop(e)
	x.Add(x, y)
	zc.BigFloat.Push(e, x)
	zc.BigFloat.Recycle(y)
}

func DivBigFloat(e *zc.OpEnv) {
	var zero big.Float
	y := zc.BigFloat.Pop(e)
	x := zc.BigFloat.Pop(e)
	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Quo(x, y)
	zc.BigFloat.Push(e, x)
	zc.BigFloat.Recycle(y)
}

func MulBigFloat(e *zc.OpEnv) {
	y := zc.BigFloat.Pop(e)
	x := zc.BigFloat.Pop(e)
	x.Mul(x, y)
	zc.BigFloat.Push(e, x)
	zc.BigFloat.Recycle(y)
}

func NegBigFloat(e *zc.OpEnv) {
	x := zc.BigFloat.Pop(e)
	x.Neg(x)
	zc.BigFloat.Push(e, x)
}

func SignBigFloat(e *zc.OpEnv) {
	x := zc.BigFloat.Pop(e)
	s := x.Sign()
	zc.Int.Push(e, s)
}

func SqrtBigFloat(e *zc.OpEnv) {
	var zero big.Float
	x := zc.BigFloat.Pop(e)
	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	x.Sqrt(x)
	zc.BigFloat.Push(e, x)
}

func SubBigFloat(e *zc.OpEnv) {
	y := zc.BigFloat.Pop(e)
	x := zc.BigFloat.Pop(e)
	x.Sub(x, y)
	zc.BigFloat.Push(e, x)
	zc.BigFloat.Recycle(y)
}
