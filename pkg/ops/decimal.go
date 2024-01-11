package ops

import (
	"github.com/blackchip-org/zc/v5/pkg/types"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

/*
oper	coeff
func	Coeff p0:Decimal -- BigInt
title 	Coefficient

desc
The value of *p0* without the decimal point.
end

example
12.345 coeff -- 12345
end
*/
func Coeff(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Coefficient()
	zc.PushBigInt(c, r0)
}

/*
oper	dec
func	Dec         p0:Decimal  -- Decimal
func	DecFloat    p0:Float    -- Decimal
func	DecRational p0:Rational -- Decimal
title	Decimal number

desc
Pops *p0* from the stack and formats it as a `Decimal`.
end

example
1e3 dec -- 1000
end
*/
func Dec(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	zc.PushDecimal(c, a0)
}

func DecFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := types.NewDecimalFromFloat(a0)
	zc.PushDecimal(c, r0)
}

func DecRational(c zc.Calc) {
	a0 := zc.PopRational(c)
	r0, exact := a0.Float64()
	zc.PushFloat(c, r0)
	if !exact {
		zc.Annotate(c, "inexact")
	}
}

/*
oper 	div-rem
func	DivRemDec p0:Decimal p1:Decimal p:Int32 -- r:Decimal q:Decimal
alias	dr
title	Division with remainder

desc
Divides *p0* by *p1* with the precision *p* and returns the quotient *q* and
remainder *r*. The following shows how to divide one dollar with three people
which gives a quotient of $0.33 and a remainder of one cent.
end

example
1.00 3 2 div-rem -- 0.01 # remainder | 0.33
end
*/
func DivRemDec(c zc.Calc) {
	prec := zc.PopInt32(c)
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	q, r := a0.QuoRem(a1, prec)
	zc.PushDecimal(c, r)
	zc.Annotate(c, "remainder")
	zc.PushDecimal(c, q)
}
