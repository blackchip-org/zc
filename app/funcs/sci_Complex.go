package funcs

import (
	"math/cmplx"

	"github.com/blackchip-org/zc/v6"
)

func AbsComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	d := cmplx.Abs(x)
	zc.Float64.Push(e, d)
}

func AcosComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Acos(x)
	zc.Complex.Push(e, z)
}

func AcoshComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Acosh(x)
	zc.Complex.Push(e, z)
}

func AsinComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Asin(x)
	zc.Complex.Push(e, z)
}

func AsinhComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Asinh(x)
	zc.Complex.Push(e, z)
}

func AtanComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Atan(x)
	zc.Complex.Push(e, z)
}

func AtanhComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Atanh(x)
	zc.Complex.Push(e, z)
}
