package ops

import "github.com/blackchip-org/zc"

func AbsDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Abs()
	zc.PushDecimal(c, r0)
}

func AddDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Add(a1)
	zc.PushDecimal(c, r0)
}

func CeilDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Ceil()
	zc.PushDecimal(c, r0)
}

func DivDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)

	if a1.IsZero() {
		c.SetError(zc.ErrDivisionByZero)
		return
	}

	r0 := a0.Div(a1)
	zc.PushDecimal(c, r0)
}

func EqDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(c, r0)
}

func FloorDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Floor()
	zc.PushDecimal(c, r0)
}

func GtDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(c, r0)
}

func GteDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(c, r0)
}

func LtDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(c, r0)
}

func LteDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(c, r0)
}

func ModDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)

	if a1.IsZero() {
		c.SetError(zc.ErrDivisionByZero)
		return
	}

	r0 := a0.Mod(a1)
	zc.PushDecimal(c, r0)
}

func MulDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Mul(a1)
	zc.PushDecimal(c, r0)
}

func NegDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Neg()
	zc.PushDecimal(c, r0)
}

func NeqDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) != 0
	zc.PushBool(c, r0)
}

func SignDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Sign()
	zc.PushInt(c, r0)
}

func SubDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Sub(a1)
	zc.PushDecimal(c, r0)
}
