package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddRat(e *zc.OpEnv) {
	y := zc.Rat.Pop(e)
	x := zc.Rat.Pop(e)
	x.Add(x, y)
	zc.Rat.Push(e, x)
	zc.Rat.Recycle(y)
}

func DivRat(e *zc.OpEnv) {
	var zero big.Rat
	y := zc.Rat.Pop(e)
	x := zc.Rat.Pop(e)
	defer zc.Rat.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}

	x.Quo(x, y)
	zc.Rat.Push(e, x)
}

func MulRat(e *zc.OpEnv) {
	y := zc.Rat.Pop(e)
	x := zc.Rat.Pop(e)
	x.Mul(x, y)
	zc.Rat.Push(e, x)
	zc.Rat.Recycle(y)
}

func NegRat(e *zc.OpEnv) {
	x := zc.Rat.Pop(e)
	x.Neg(x)
	zc.Rat.Push(e, x)
}

func SignRat(e *zc.OpEnv) {
	x := zc.Rat.Pop(e)
	s := x.Sign()
	zc.Int.Push(e, s)
	zc.Rat.Recycle(x)
}

func SubRat(e *zc.OpEnv) {
	y := zc.Rat.Pop(e)
	x := zc.Rat.Pop(e)
	x.Sub(x, y)
	zc.Rat.Push(e, x)
	zc.Rat.Recycle(y)
}
