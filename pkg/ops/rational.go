package ops

import (
	"math/big"

	"github.com/blackchip-org/zc/v5/pkg/zc"
)

/*
oper	dec
func	DecRational    p0:Rational    -- Decimal
title	Decimal number

desc
Pops *p0* from the stack and formats it as a `Decimal`.
end

example
1/2 dec -- 0.5
end
*/

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
oper	inv
func	Inv p0:Rational -- Rational
title	Inverse

desc
Returns the inverse which is 1/*p0*.
end

example
1/2 inv -- 2
end
*/
func Inv(c zc.Calc) {
	var r0 big.Rat
	a0 := zc.PopRational(c)
	r0.Inv(a0)
	zc.PushRational(c, &r0)
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
oper	rational
func	RationalBigInt n:BigInt d:BigInt -- Rational
func 	RationalFloat  p0:Float          -- Rational
alias	rat
title	Convert to a Rational

desc
A rational number using an integer numerator *n* and denominator *d*,
or using a floating point value *p0*.
end

example
c 1 2 rat -- 1/2
c 0.5 rat -- 1/2
end
*/

func RationalBigInt(c zc.Calc) {
	var r0 big.Rat
	d := zc.PopBigInt(c)
	n := zc.PopBigInt(c)
	r0.SetFrac(n, d)
	zc.PushRational(c, &r0)
}

func RationalFloat(c zc.Calc) {
	var r0 big.Rat
	a0 := zc.PopFloat(c)
	res := r0.SetFloat64(a0)
	if res == nil {
		zc.ErrInvalidArgs(c, "not finite")
		return
	}
	zc.PushRational(c, &r0)
}

/*
oper	rational?
func	RationalIs p0:Str -- Bool
title 	Checks value can be parsed as a rational

desc
Returns `true` if the value *p0* can be parsed as a Rational.
end

example
c 1/2 rational? -- true
c 1+2i rational? -- false
end
*/
func RationalIs(c zc.Calc) {
	a0 := zc.PopString(c)
	r0 := zc.Rational.Is(a0)
	zc.PushBool(c, r0)
}
