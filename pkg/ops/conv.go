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
oper	denom
func	Denom p0:Rational -- BigInt
title	Denominator

desc
The denominator of rational number *p0*.
end

example
2/3 denom -- 3
end
*/
func Denom(c zc.Calc) {
	a0 := zc.PopRational(c)
	r0 := a0.Denom()
	zc.PushBigInt(c, r0)
}

/*
oper	imag
func	Imag p0:Complex -- Float
title	Imaginary number from complex

desc
The imaginary number part of complex number *p0*
end

example
3+4i imag -- 4
end
*/
func Imag(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := imag(a0)
	zc.PushFloat(c, r0)
}

/*
oper	num
func	Num p0:Rational -- BigInt
title	Numerator

desc
The numerator of rational number *p0*.
end

example
2/3 num -- 2
end
*/
func Num(c zc.Calc) {
	a0 := zc.PopRational(c)
	r0 := a0.Num()
	zc.PushBigInt(c, r0)
}

/*
oper	real
func	Real p0:Complex -- Float
title	Real number from complex

desc
The real number part of complex number *p0*
end

example
3+4i real -- 3
end
*/
func Real(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := real(a0)
	zc.PushFloat(c, r0)
}
