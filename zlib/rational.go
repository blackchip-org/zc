package zlib

import (
	"math/big"

	"github.com/blackchip-org/zc"
)

var zeroRat big.Rat

func AbsRational(c zc.Calc) {
	var r0 big.Rat
	a0 := zc.PopRational(c)
	r0.Abs(a0)
	zc.PushRational(c, &r0)
}

func AddRational(c zc.Calc) {
	var r0 big.Rat
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0.Add(a0, a1)
	zc.PushRational(c, &r0)
}

func DivRational(c zc.Calc) {
	var r0 big.Rat
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)

	if a1.Cmp(&zeroRat) == 0 {
		c.SetError(zc.ErrDivisionByZero)
		return
	}

	r0.Quo(a0, a1)
	zc.PushRational(c, &r0)
}

func EqRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(c, r0)
}

func GtRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(c, r0)
}

func GteRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(c, r0)
}

func LtRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(c, r0)
}

func LteRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(c, r0)
}

func MulRational(c zc.Calc) {
	var r0 big.Rat
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0.Mul(a0, a1)
	zc.PushRational(c, &r0)
}

func NegRational(c zc.Calc) {
	var r0 big.Rat
	a0 := zc.PopRational(c)
	r0.Neg(a0)
	zc.PushRational(c, &r0)
}

func NeqRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) != 0
	zc.PushBool(c, r0)
}

func SignRational(c zc.Calc) {
	a0 := zc.PopRational(c)
	r0 := a0.Sign()
	zc.PushInt(c, r0)
}

func SubRational(c zc.Calc) {
	var r0 big.Rat
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0.Sub(a0, a1)
	zc.PushRational(c, &r0)
}
