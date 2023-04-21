package ops

import (
	"math"
	"strconv"

	"github.com/blackchip-org/zc/pkg/zc"
)

func AbsFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Abs(a0)
	zc.PushFloat(c, r0)
}

func AcosFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Acos(a0)
	zc.PushFloat(c, r0)
}

func AcoshFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Acosh(a0)
	zc.PushFloat(c, r0)
}

func AsinFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Asin(a0)
	zc.PushFloat(c, r0)
}

func AsinhFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Asinh(a0)
	zc.PushFloat(c, r0)
}

func AtanFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Atan(a0)
	zc.PushFloat(c, r0)
}

func AtanhFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Atanh(a0)
	zc.PushFloat(c, r0)
}

func AddFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 + a1
	zc.PushFloat(c, r0)
}

func CeilFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Ceil(a0)
	zc.PushFloat(c, r0)
}

func CosFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Cos(a0)
	zc.PushFloat(c, r0)
}

func CoshFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Cosh(a0)
	zc.PushFloat(c, r0)
}

func DivFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)

	if a1 == 0 {
		zc.ErrDivisionByZero(c, a0, a1)
		return
	}

	r0 := a0 / a1
	zc.PushFloat(c, r0)
}

var E = "2.718281828459045"

func EqFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 == a1
	zc.PushBool(c, r0)
}

func ExpFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Exp(a0)
	zc.PushFloat(c, r0)
}

func FloorFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Floor(a0)
	zc.PushFloat(c, r0)
}

func GtFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 > a1
	zc.PushBool(c, r0)
}

func GteFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 >= a1
	zc.PushBool(c, r0)
}

func LogFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Log(a0)
	zc.PushFloat(c, r0)
}

func Log10Float(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Log10(a0)
	zc.PushFloat(c, r0)
}

func Log2Float(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Log2(a0)
	zc.PushFloat(c, r0)
}

func LtFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 < a1
	zc.PushBool(c, r0)
}

func LteFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 <= a1
	zc.PushBool(c, r0)
}

func ModFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)

	if a1 == 0 {
		zc.ErrModuloByZero(c, a0, a1)
		return
	}

	r0 := math.Mod(a0, a1)
	zc.PushFloat(c, r0)
}

func MulFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 * a1
	zc.PushFloat(c, r0)
}

func NegFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := -a0
	zc.PushFloat(c, r0)
}

func NeqFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 != a1
	zc.PushBool(c, r0)
}

var Pi = "3.141592653589793"

func PowFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := math.Pow(a0, a1)
	zc.PushFloat(c, r0)
}

func RemFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := math.Remainder(a0, a1)
	zc.PushFloat(c, r0)
}

func ScientificNotation(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := strconv.FormatFloat(a0, 'e', -1, 64)
	zc.PushString(c, r0)
}

func SignFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)

	r0 := 0
	if a0 > 0 {
		r0 = 1
	}
	if a0 < 0 {
		r0 = -1
	}

	zc.PushInt(c, r0)
}

func SinFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Sin(a0)
	zc.PushFloat(c, r0)
}

func SinhFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Sinh(a0)
	zc.PushFloat(c, r0)
}

func SubFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 - a1
	zc.PushFloat(c, r0)
}

func TanFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Tan(a0)
	zc.PushFloat(c, r0)
}

func TanhFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Tanh(a0)
	zc.PushFloat(c, r0)
}
