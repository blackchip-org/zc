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

func AcosComplex128(e *zc.OpEnv) {
	x := zc.Complex128.Pop(e)
	z := cmplx.Acos(x)
	zc.Complex128.Push(e, z)
}

func AcoshComplex128(e *zc.OpEnv) {
	x := zc.Complex128.Pop(e)
	z := cmplx.Acosh(x)
	zc.Complex128.Push(e, z)
}

func AsinComplex128(e *zc.OpEnv) {
	x := zc.Complex128.Pop(e)
	z := cmplx.Asin(x)
	zc.Complex128.Push(e, z)
}

func AsinhComplex128(e *zc.OpEnv) {
	x := zc.Complex128.Pop(e)
	z := cmplx.Asinh(x)
	zc.Complex128.Push(e, z)
}

func AtanComplex128(e *zc.OpEnv) {
	x := zc.Complex128.Pop(e)
	z := cmplx.Atan(x)
	zc.Complex128.Push(e, z)
}

func AtanhComplex128(e *zc.OpEnv) {
	x := zc.Complex128.Pop(e)
	z := cmplx.Atanh(x)
	zc.Complex128.Push(e, z)
}
