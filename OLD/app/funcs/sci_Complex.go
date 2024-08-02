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

func CosComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Cos(x)
	zc.Complex.Push(e, z)
}

func CoshComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Cosh(x)
	zc.Complex.Push(e, z)
}

func CotComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Cot(x)
	zc.Complex.Push(e, z)
}

func ExpComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Exp(x)
	zc.Complex.Push(e, z)
}

func LogComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Log(x)
	zc.Complex.Push(e, z)
}

func Log10Complex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Log10(x)
	zc.Complex.Push(e, z)
}

func SinComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Sin(x)
	zc.Complex.Push(e, z)
}

func SinhComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Sinh(x)
	zc.Complex.Push(e, z)
}

func TanComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Tan(x)
	zc.Complex.Push(e, z)
}

func TanhComplex(e *zc.OpEnv) {
	x := zc.Complex.Pop(e)
	z := cmplx.Tanh(x)
	zc.Complex.Push(e, z)
}
