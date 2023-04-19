package ops

import (
	"math"

	"github.com/blackchip-org/zc"
)

func AbsFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Abs(a0)
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

func EqFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 == a1
	zc.PushBool(c, r0)
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

func SubFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 - a1
	zc.PushFloat(c, r0)
}
