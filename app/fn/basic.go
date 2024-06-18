package fn

import (
	"github.com/blackchip-org/zc/v6"
)

func AddBigInt(e zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Top())
	x.Add(x, y)
	zc.BigInt.Recycle(y)
}

func MulBigInt(e zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Top())
	x.Mul(x, y)
	zc.BigInt.Recycle(y)
}

func PowBigInt(e zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Top())
	x.Exp(x, y, nil)
	zc.BigInt.Recycle(y)
}

func SqrtBigInt(e zc.OpEnv) {
	x := zc.BigInt.As(e.Top())
	x.Sqrt(x)
}

func SubBigInt(e zc.OpEnv) {
	y := zc.BigInt.As(e.Pop())
	x := zc.BigInt.As(e.Top())
	x.Sub(x, y)
	zc.BigInt.Recycle(y)
}
