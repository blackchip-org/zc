package funcs

import (
	"math/cmplx"

	"github.com/blackchip-org/zc/v6"
)

func AddComplex(e *zc.OpEnv) {
	y := zc.Complex.Pop(e)
	x := zc.Complex.Pop(e)
	z := x + y
	zc.Complex.Push(e, z)
}

func DivComplex(e *zc.OpEnv) {
	y := zc.Complex.Pop(e)
	x := zc.Complex.Pop(e)
	z := x / y
	zc.Complex.Push(e, z)
}

func MulComplex(e *zc.OpEnv) {
	y := zc.Complex.Pop(e)
	x := zc.Complex.Pop(e)
	z := x * y
	zc.Complex.Push(e, z)
}

func NegComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := -x
	zc.Complex.Push(e, z)
}

func PowComplex(e *zc.OpEnv) {
	y := zc.Complex.Pop(e)
	x := zc.Complex.Pop(e)
	z := cmplx.Pow(x, y)
	zc.Complex.Push(e, z)
}

func SqrtComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Sqrt(x)
	zc.Complex.Push(e, z)
}

func SubComplex(e *zc.OpEnv) {
	y := zc.Complex.Pop(e)
	x := zc.Complex.Pop(e)
	z := x - y
	zc.Complex.Push(e, z)
}
