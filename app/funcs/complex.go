package funcs

import (
	"math/cmplx"

	"github.com/blackchip-org/zc/v6"
)

func SquareRootComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	zc.Complex.Push(c, cmplx.Sqrt(x))
}
