package ops

import (
	"github.com/blackchip-org/zc/v6/pkg/types"
	"github.com/blackchip-org/zc/v6/pkg/zc"
)

type decimalState struct {
	prec int
}

func getDecimalState(c zc.Calc) *decimalState {
	s, ok := c.State("decimal")
	if !ok {
		s = &decimalState{prec: 16}
		c.NewState("decimal", s)
	}
	return s.(*decimalState)
}

func AddDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Add(a1)
	zc.PushDecimal(c, r0)
}

/*
oper	dec
func	Dec         p0:Decimal  -- Decimal
func	DecFloat    p0:Float    -- Decimal
title	Decimal number

desc
Pops *p0* from the stack and formats it as a `Decimal`.
end

example
1e3 dec -- 1000
end
*/
func Dec(c zc.Calc) {
	x := zc.PopDecimal(c)
	zc.PushDecimal(c, x)
}

func DecFloat(c zc.Calc) {
	x := zc.PopFloat(c)
	r := types.NewDecimalFromFloat(x)
	zc.PushDecimal(c, r)
}

func DecRational(c zc.Calc) {
	s := getDecimalState(c)
	x := zc.PopRational(c)
	r := x.FloatString(s.prec)
	r = zc.RemoveTrailingZeros(r)
	zc.PushString(c, r)
}

/*
oper	dec?
title	x
aliases decimal?
func 	DecimalIs x:Str -- Bool
*/
func DecimalIs(c zc.Calc) {
	x := zc.PopString(c)
	r := zc.Decimal.Is(x)
	zc.PushBool(c, r)
}

/*
oper	dec-prec
title	x
aliases	decimal-prec
func	DecPrecGet -- Int
*/
func DecPrecGet(c zc.Calc) {
	s := getDecimalState(c)
	zc.PushInt(c, s.prec)
}

/*
oper 	dec-prec=
title	x
aliases	decimal-prec=
func 	DecPrecSet x:Int --
*/
func DecPrecSet(c zc.Calc) {
	s := getDecimalState(c)
	x := zc.PopInt(c)
	if x < 0 {
		zc.ErrInvalidArgs(c, "cannot be negative")
		return
	}
	s.prec = x
	c.SetInfo("decimal precision set to %v", s.prec)
}
