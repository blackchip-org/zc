package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddBigInt(e *zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Pop())
	x.Add(x, y)
	e.PushVal(x)
	zc.BigInt.Recycle(y)
}

func DivBigInt(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Pop())
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Div(x, y)
	e.PushVal(x)
}

func ModBigInt(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Pop())
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Mod(x, y)
	e.PushVal(x)
}

func MulBigInt(e *zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Pop())
	x.Mul(x, y)
	e.PushVal(x)
	zc.BigInt.Recycle(y)
}

func NegBigInt(e *zc.OpEnv) {
	x := zc.BigInt.As(e.Pop())
	x.Neg(x)
	e.PushVal(x)
}

func PowBigInt(e *zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Pop())
	x.Exp(x, y, nil)
	e.PushVal(x)
	zc.BigInt.Recycle(y)
}

func RemBigInt(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Pop())
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Rem(x, y)
	e.PushVal(x)
}

func SignBigInt(e *zc.OpEnv) {
	x := zc.BigInt.As(e.Pop())
	e.PushVal(x.Sign())
	zc.BigInt.Recycle(x)
}

func SqrtBigInt(e *zc.OpEnv) {
	var zero big.Int
	x := zc.BigInt.As(e.Pop())
	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x.String())
		return
	}
	x.Sqrt(x)
	e.PushVal(x)
}

func SubBigInt(e *zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Pop())
	x.Sub(x, y)
	e.PushVal(x)
	zc.BigInt.Recycle(y)
}
