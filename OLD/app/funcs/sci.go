package funcs

import (
	"math"

	"github.com/blackchip-org/zc/v6"
)

func AbsFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Abs(x)
	zc.Float64.Push(e, z)
}

func AcosFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Acos(x)
	zc.Float64.Push(e, z)
}

func AcoshFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Acosh(x)
	zc.Float64.Push(e, z)
}

func AsinFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Asin(x)
	zc.Float64.Push(e, z)
}

func AsinhFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Asinh(x)
	zc.Float64.Push(e, z)
}

func AtanFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Atan(x)
	zc.Float64.Push(e, z)
}

func Atan2Float64(e *zc.OpEnv) {
	y := zc.Float64.Pop(e)
	x := zc.Float64.Pop(e)
	z := math.Atan2(x, y)
	zc.Float64.Push(e, z)
}

func AtanhFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Atanh(x)
	zc.Float64.Push(e, z)
}

func CbrtFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	if x < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", x)
		return
	}
	z := math.Cbrt(x)
	zc.Float64.Push(e, z)
}

func CosFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Cos(x)
	zc.Float64.Push(e, z)
}

func CoshFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Cosh(x)
	zc.Float64.Push(e, z)
}

func ExpFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Exp(x)
	zc.Float64.Push(e, z)
}

func LogFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Log(x)
	zc.Float64.Push(e, z)
}

func Log10Float64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Log10(x)
	zc.Float64.Push(e, z)
}

func SinFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Sin(x)
	zc.Float64.Push(e, z)
}

func SinhFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Sinh(x)
	zc.Float64.Push(e, z)
}

func TanFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Tan(x)
	zc.Float64.Push(e, z)
}

func TanhFloat64(e *zc.OpEnv) {
	x := zc.Float64.Pop(e)
	z := math.Tanh(x)
	zc.Float64.Push(e, z)
}
