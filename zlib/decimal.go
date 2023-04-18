package zlib

import "github.com/blackchip-org/zc"

func AbsDecimal(e zc.Env) {
	a0 := zc.PopDecimal(e)
	r0 := a0.Abs()
	zc.PushDecimal(e, r0)
}

func AddDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Add(a1)
	zc.PushDecimal(e, r0)
}

func CeilDecimal(e zc.Env) {
	a0 := zc.PopDecimal(e)
	r0 := a0.Ceil()
	zc.PushDecimal(e, r0)
}

func DivDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)

	if a1.IsZero() {
		e.Error(zc.ErrDivisionByZero)
		return
	}

	r0 := a0.Div(a1)
	zc.PushDecimal(e, r0)
}

func EqDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(e, r0)
}

func FloorDecimal(e zc.Env) {
	a0 := zc.PopDecimal(e)
	r0 := a0.Floor()
	zc.PushDecimal(e, r0)
}

func GtDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(e, r0)
}

func GteDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(e, r0)
}

func LtDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(e, r0)
}

func LteDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(e, r0)
}

func ModDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)

	if a1.IsZero() {
		e.Error(zc.ErrDivisionByZero)
		return
	}

	r0 := a0.Mod(a1)
	zc.PushDecimal(e, r0)
}

func MulDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Mul(a1)
	zc.PushDecimal(e, r0)
}

func NegDecimal(e zc.Env) {
	a0 := zc.PopDecimal(e)
	r0 := a0.Neg()
	zc.PushDecimal(e, r0)
}

func NeqDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Cmp(a1) != 0
	zc.PushBool(e, r0)
}

func SignDecimal(e zc.Env) {
	a0 := zc.PopDecimal(e)
	r0 := a0.Sign()
	zc.PushInt(e, r0)
}

func SubDecimal(e zc.Env) {
	a1 := zc.PopDecimal(e)
	a0 := zc.PopDecimal(e)
	r0 := a0.Sub(a1)
	zc.PushDecimal(e, r0)
}
