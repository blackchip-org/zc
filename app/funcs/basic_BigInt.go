package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Add(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}

func MulBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Mul(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}

func NegBigInt(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	x.Neg(x)
	zc.BigInt.Push(e, x)
}

func PowBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Exp(x, y, nil)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}

func SignBigInt(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	zc.Int.Push(e, x.Sign())
	zc.BigInt.Recycle(x)
}

func SqrtBigInt(e *zc.OpEnv) {
	var zero big.Int
	x := zc.BigInt.Pop(e)
	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x.String())
		return
	}
	x.Sqrt(x)
	zc.BigInt.Push(e, x)
}

func SubBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Sub(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}

func TDiv(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Quo(x, y)
	zc.BigInt.Push(e, x)
}

func TDivRem(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	r := zc.BigInt.New()
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.QuoRem(x, y, r)
	zc.BigInt.Push(e, x)
	e.Label("quo")
	zc.BigInt.Push(e, r)
	e.Label("rem")
}

func RemBigInt(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Mod(x, y)
	zc.BigInt.Push(e, x)
}
