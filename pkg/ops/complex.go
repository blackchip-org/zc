package ops

import (
	"math/cmplx"

	"github.com/blackchip-org/zc"
)

func AbsComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Abs(a0)
	zc.PushFloat(c, r0)
}

func AddComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 + a1
	zc.PushComplex(c, r0)
}

func DivComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)

	if real(a1) == 0 && imag(a1) == 0 {
		zc.ErrDivisionByZero(c, a0, a1)
		return
	}

	r0 := a0 / a1
	zc.PushComplex(c, r0)
}

func EqComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 == a1
	zc.PushBool(c, r0)
}

func MulComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 * a1
	zc.PushComplex(c, r0)
}

func NeqComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 != a1
	zc.PushBool(c, r0)
}

func PowComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := cmplx.Pow(a0, a1)
	zc.PushComplex(c, r0)
}

func SubComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 - a1
	zc.PushComplex(c, r0)
}
