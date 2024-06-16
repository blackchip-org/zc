package fn

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/types"
)

func BigIntAdd(e zc.OpEnv) {
	y := types.BigInt.As(e.Drop())
	x := types.BigInt.As(e.Top())
	x.Add(x, y)
}

func BigIntMul(e zc.OpEnv) {
	y := types.BigInt.As(e.Drop())
	x := types.BigInt.As(e.Top())
	x.Mul(x, y)
}
