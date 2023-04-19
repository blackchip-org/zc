package zlib

import (
	"math/cmplx"

	"github.com/blackchip-org/zc"
)

func AbsComplex(e zc.Env) {
	a0 := zc.PopComplex(e)
	r0 := cmplx.Abs(a0)
	zc.PushFloat(e, r0)
}

func AddComplex(e zc.Env) {
	a1 := zc.PopComplex(e)
	a0 := zc.PopComplex(e)
	r0 := a0 + a1
	zc.PushComplex(e, r0)
}

func DivComplex(e zc.Env) {
	a1 := zc.PopComplex(e)
	a0 := zc.PopComplex(e)

	if real(a1) == 0 && imag(a1) == 0 {
		e.Error(zc.ErrDivisionByZero)
		return
	}

	r0 := a0 / a1
	zc.PushComplex(e, r0)
}

func EqComplex(e zc.Env) {
	a1 := zc.PopComplex(e)
	a0 := zc.PopComplex(e)
	r0 := a0 == a1
	zc.PushBool(e, r0)
}

func MulComplex(e zc.Env) {
	a1 := zc.PopComplex(e)
	a0 := zc.PopComplex(e)
	r0 := a0 * a1
	zc.PushComplex(e, r0)
}

func NeqComplex(e zc.Env) {
	a1 := zc.PopComplex(e)
	a0 := zc.PopComplex(e)
	r0 := a0 != a1
	zc.PushBool(e, r0)
}

func PowComplex(e zc.Env) {
	a1 := zc.PopComplex(e)
	a0 := zc.PopComplex(e)
	r0 := cmplx.Pow(a0, a1)
	zc.PushComplex(e, r0)
}

func SubComplex(e zc.Env) {
	a1 := zc.PopComplex(e)
	a0 := zc.PopComplex(e)
	r0 := a0 - a1
	zc.PushComplex(e, r0)
}
