package funcs

import (
	"github.com/blackchip-org/zc/v6"
)

func AbsBigInt(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	x.Abs(x)
	zc.BigInt.Push(e, x)
}
