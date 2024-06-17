package fn

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/types"
)

func AddBigInt(e zc.OpEnv) {
	y := types.BigInt.As(e.Pop())
	x := types.BigInt.As(e.Top())
	x.Add(x, y)
	types.BigInt.Recycle(y)
}

func MulBigInt(e zc.OpEnv) {
	y := types.BigInt.As(e.Pop())
	x := types.BigInt.As(e.Top())
	x.Mul(x, y)
	types.BigInt.Recycle(y)
}

func PowBigInt(e zc.OpEnv) {
	y := types.BigInt.As(e.Pop())
	x := types.BigInt.As(e.Top())
	x.Exp(x, y, nil)
	types.BigInt.Recycle(y)
}

func SqrtBigInt(e zc.OpEnv) {
	x := types.BigInt.As(e.Top())
	x.Sqrt(x)
}

func SubBigInt(e zc.OpEnv) {
	y := types.BigInt.As(e.Pop())
	x := types.BigInt.As(e.Top())
	x.Sub(x, y)
	types.BigInt.Recycle(y)
}
