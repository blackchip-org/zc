package funcs

import (
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AddBigInt(e *zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(*e.PopTop())
	x.Add(x, y)
	zc.BigInt.Recycle(y)
}

func DivBigInt(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(*e.PopTop())
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Div(x, y)
}

func ModBigInt(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(*e.PopTop())
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Mod(x, y)
}

func MulBigInt(e *zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(*e.PopTop())
	x.Mul(x, y)
	zc.BigInt.Recycle(y)
}

func NegBigInt(e *zc.OpEnv) {
	x := zc.BigInt.As(*e.PopTop())
	x.Neg(x)
}

func PowBigInt(e *zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(*e.PopTop())
	x.Exp(x, y, nil)
	zc.BigInt.Recycle(y)
}

func RemBigInt(e *zc.OpEnv) {
	var zero big.Int
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(*e.PopTop())
	defer zc.BigInt.Recycle(y)

	if y.Cmp(&zero) == 0 {
		e.Err = zc.ErrDivisionByZero(e)
		return
	}
	x.Rem(x, y)
}

func SignBigInt(e *zc.OpEnv) {
	x := e.PopTop()
	x.Val = zc.BigInt.As(*x).Sign()
}

func SqrtBigInt(e *zc.OpEnv) {
	var zero big.Int
	x := zc.BigInt.As(*e.PopTop())
	if x.Cmp(&zero) < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x.String())
		return
	}
	x.Sqrt(x)
}

func SubBigInt(e *zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(*e.PopTop())
	x.Sub(x, y)
	zc.BigInt.Recycle(y)
}
