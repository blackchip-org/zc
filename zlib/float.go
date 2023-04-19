package zlib

import (
	"math"

	"github.com/blackchip-org/zc"
)

func AbsFloat(e zc.Env) {
	a0 := zc.PopFloat(e)
	r0 := math.Abs(a0)
	zc.PushFloat(e, r0)
}

func AddFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 + a1
	zc.PushFloat(e, r0)
}

func CeilFloat(e zc.Env) {
	a0 := zc.PopFloat(e)
	r0 := math.Ceil(a0)
	zc.PushFloat(e, r0)
}

func DivFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)

	if a1 == 0 {
		e.Error(zc.ErrDivisionByZero)
		return
	}

	r0 := a0 / a1
	zc.PushFloat(e, r0)
}

func EqFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 == a1
	zc.PushBool(e, r0)
}

func FloorFloat(e zc.Env) {
	a0 := zc.PopFloat(e)
	r0 := math.Floor(a0)
	zc.PushFloat(e, r0)
}

func GtFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 > a1
	zc.PushBool(e, r0)
}

func GteFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 >= a1
	zc.PushBool(e, r0)
}

func LtFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 < a1
	zc.PushBool(e, r0)
}

func LteFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 <= a1
	zc.PushBool(e, r0)
}

func ModFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)

	if a1 == 0 {
		e.Error(zc.ErrDivisionByZero)
		return
	}

	r0 := math.Mod(a0, a1)
	zc.PushFloat(e, r0)
}

func MulFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 * a1
	zc.PushFloat(e, r0)
}

func NegFloat(e zc.Env) {
	a0 := zc.PopFloat(e)
	r0 := -a0
	zc.PushFloat(e, r0)
}

func NeqFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 != a1
	zc.PushBool(e, r0)
}

func PowFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := math.Pow(a0, a1)
	zc.PushFloat(e, r0)
}

func RemFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := math.Remainder(a0, a1)
	zc.PushFloat(e, r0)
}

func SignFloat(e zc.Env) {
	a0 := zc.PopFloat(e)

	r0 := 0
	if a0 > 0 {
		r0 = 1
	}
	if a0 < 0 {
		r0 = -1
	}

	zc.PushInt(e, r0)
}

func SubFloat(e zc.Env) {
	a1 := zc.PopFloat(e)
	a0 := zc.PopFloat(e)
	r0 := a0 - a1
	zc.PushFloat(e, r0)
}
