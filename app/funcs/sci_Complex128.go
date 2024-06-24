package funcs

import (
	"math/cmplx"

	"github.com/blackchip-org/zc/v6"
)

func AbsComplex128(e *zc.OpEnv) {
	x := zc.Complex128.Pop(e)
	d := cmplx.Abs(x)
	zc.Float64.Push(e, d)
}
