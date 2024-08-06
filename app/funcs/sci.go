package funcs

import (
	"math"
	"math/cmplx"

	"github.com/blackchip-org/zc/v6"
)

// ----------------------------------------------------------------------------
func AcosFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Acos(x)
	zc.Float64.Push(c, z)
}

func AcosComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Acos(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func AcoshFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Acosh(x)
	zc.Float64.Push(c, z)
}

func AcoshComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Acosh(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func AsinFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Asin(x)
	zc.Float64.Push(c, z)
}

func AsinComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Asin(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func AsinhFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Asinh(x)
	zc.Float64.Push(c, z)
}

func AsinhComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Asinh(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func AtanFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Atan(x)
	zc.Float64.Push(c, z)
}

func AtanComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Atan(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func Atan2Float64(c zc.Calc) {
	y := zc.Float64.Pop(c)
	x := zc.Float64.Pop(c)
	z := math.Atan2(x, y)
	zc.Float64.Push(c, z)
}

// ----------------------------------------------------------------------------
func AtanhFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Atanh(x)
	zc.Float64.Push(c, z)
}

func AtanhComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Atanh(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func CosFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Cos(x)
	zc.Float64.Push(c, z)
}

func CosComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Cos(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func CoshFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Cosh(x)
	zc.Float64.Push(c, z)
}

func CoshComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Cosh(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func CotComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Cot(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func LogFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Log(x)
	zc.Float64.Push(c, z)
}

func LogComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Log(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func SinFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Sin(x)
	zc.Float64.Push(c, z)
}

func SinComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Sin(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func SinhFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Sinh(x)
	zc.Float64.Push(c, z)
}

func SinhComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Sinh(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func TanFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Tan(x)
	zc.Float64.Push(c, z)
}

func TanComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Tan(x)
	zc.Complex.Push(c, z)
}

// ----------------------------------------------------------------------------
func TanhFloat64(c zc.Calc) {
	x := zc.Float64.Pop(c)
	z := math.Tanh(x)
	zc.Float64.Push(c, z)
}

func TanhComplex(c zc.Calc) {
	x := zc.Complex.Pop(c)
	z := cmplx.Tanh(x)
	zc.Complex.Push(c, z)
}
