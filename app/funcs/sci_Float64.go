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
