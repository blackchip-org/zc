package zlib

import (
	"math/big"

	"github.com/blackchip-org/zc"
)

var zeroRat big.Rat

func AbsRational(e zc.Env) {
	var r0 big.Rat
	a0 := zc.PopRational(e)
	r0.Abs(a0)
	zc.PushRational(e, &r0)
}

func AddRational(e zc.Env) {
	var r0 big.Rat
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0.Add(a0, a1)
	zc.PushRational(e, &r0)
}

func DivRational(e zc.Env) {
	var r0 big.Rat
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)

	if a1.Cmp(&zeroRat) == 0 {
		e.Error(zc.ErrDivisionByZero)
		return
	}

	r0.Quo(a0, a1)
	zc.PushRational(e, &r0)
}

func EqRational(e zc.Env) {
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(e, r0)
}

func GtRational(e zc.Env) {
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(e, r0)
}

func GteRational(e zc.Env) {
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(e, r0)
}

func LtRational(e zc.Env) {
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(e, r0)
}

func LteRational(e zc.Env) {
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(e, r0)
}

func MulRational(e zc.Env) {
	var r0 big.Rat
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0.Mul(a0, a1)
	zc.PushRational(e, &r0)
}

func NegRational(e zc.Env) {
	var r0 big.Rat
	a0 := zc.PopRational(e)
	r0.Neg(a0)
	zc.PushRational(e, &r0)
}

func NeqRational(e zc.Env) {
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0 := a0.Cmp(a1) != 0
	zc.PushBool(e, r0)
}

func SignRational(e zc.Env) {
	a0 := zc.PopRational(e)
	r0 := a0.Sign()
	zc.PushInt(e, r0)
}

func SubRational(e zc.Env) {
	var r0 big.Rat
	a1 := zc.PopRational(e)
	a0 := zc.PopRational(e)
	r0.Sub(a0, a1)
	zc.PushRational(e, &r0)
}
