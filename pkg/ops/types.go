package ops

import (
	"github.com/blackchip-org/zc/pkg/types"
	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper	complex
func	Complex r:Float i:Float -- Float
title	Complex from real and imaginary

desc
A complex number from a real *r* and an imaginary *i* number.
end

example
6 -- 6
12 -- 6 | 12
complex -- 6+12i
end
*/
func Complex(c zc.Calc) {
	i := zc.PopFloat(c)
	r := zc.PopFloat(c)
	r0 := complex(r, i)
	zc.PushComplex(c, r0)
}

/*
oper	dec
func	Dec p0:Decimal -- Decimal
func	DecFloat p0:Float -- Decimal
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
