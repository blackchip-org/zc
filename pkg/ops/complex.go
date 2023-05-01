package ops

import (
	"fmt"
	"math/cmplx"

	"github.com/blackchip-org/zc/pkg/types"
	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper 	abs
func 	AbsComplex p0:Complex -- Complex
title	Distance from zero in complex plane

desc
The distance of *p0* from zero in the complex plane.
end

example
2+2i abs 5 round -- 2.82843
end
*/
func AbsComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Abs(a0)
	zc.PushFloat(c, r0)
}

/*
oper 	acos
func	AcosComplex p0:Complex -- Complex
title	Inverse cosine

desc
Inverse cosine of *p0*.
end

example
0.5+1i acos 5 round -- 1.22136-0.92613i
end
*/

func AcosComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Acos(a0)
	zc.PushComplex(c, r0)
}

/*
oper	acosh
func	AcoshComplex p0:Complex -- Complex
title	Inverse hyperbolic cosine

desc
Inverse cosine of *p0*
end

example
2+2i acosh 5 round -- 1.73432+0.81655i
end
*/
func AcoshComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Acosh(a0)
	zc.PushComplex(c, r0)
}

/*
oper	asin
func	AsinComplex p0:Complex -- Complex
title	Inverse sine

desc
Inverse sine of *p0*.
end

example
0.5+2i asin 5 round -- 0.22102+1.46572i
end
*/
func AsinComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Asin(a0)
	zc.PushComplex(c, r0)
}

/*
oper	asinh
func	AsinComplex p0:Complex -- Complex
title 	Inverse hyperbolic sine

desc
Inverse hyperbolic sine of *p0*.
end

example
2+2i asinh 5 round -- 0.75425+1.73432i
end
*/
func AsinhComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Asinh(a0)
	zc.PushComplex(c, r0)
}

/*
oper	atan
func	AtanComplex p0:Complex -- Complex
title	Inverse tangent

desc
Inverse tangent of *p0*.
end

example
0.5+2i atan 5 round -- 1.42155+0.50037i
end
*/
func AtanComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Atan(a0)
	zc.PushComplex(c, r0)
}

/*
oper	atanh
func	AtanhComplex p0:Complex -- Complex
title	Inverse hyperbolic tangent

desc
Inverse hyperbolic tangent of *p0*.
end

example
0.5+2i atanh 5 round --0.09642+1.12656i
end
*/
func AtanhComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Atanh(a0)
	zc.PushComplex(c, r0)
}

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
oper	conj
func	Conj p0:Complex -- Complex
title 	Complex conjugate

desc
The complex conjugate of *p0*
end

example
5+3i conj -- 5-3i
end
*/
func Conj(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Conj(a0)
	zc.PushComplex(c, r0)
}

/*
oper	cos
func	CosComplex p0:Complex -- Complex
title	Cosine

desc
Cosine of *p0*.
end

example
2+2i cos 5 round -- -1.56563-3.29789i
end
*/
func CosComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Cos(a0)
	zc.PushComplex(c, r0)
}

/*
oper 	cosh
func	CoshComplex p0:Complex -- Complex
title	Hyperbolic cosine

desc
Hyperbolic cosine of *p0*.
end

example
2+2i cosh 5 round -- -1.56563+3.29789i
end
*/
func CoshComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Cosh(a0)
	zc.PushComplex(c, r0)
}

/*
oper	cot
func	CotComplex p0:Complex -- Complex
title	Cotangent

desc
Cotangent of *p0*.
end

example
2+3i cot 5 round -- -0.00374-0.99676i
end
*/
func CotComplex(c zc.Calc) {
	p0 := zc.PopComplex(c)
	r0 := cmplx.Cot(p0)
	zc.PushComplex(c, r0)
}

/*
oper	exp
func	ExpComplex p0:Complex -- Complex
title	Natural exponential

desc
Natural exponential of *p0*.
end

example
2+2i exp 5 round -- -3.07493+6.71885i
end
*/
func ExpComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Exp(a0)
	zc.PushComplex(c, r0)
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
oper	log
func	LogComplex p0:Complex -- Complex
title	Natural logarithm

desc
Natural logarithm of *p0*.
end

example
8+2i log 5 round -- 2.10975+0.24498i
end
*/
func LogComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Log(a0)
	zc.PushComplex(c, r0)
}

/*
oper	log10
func	Log10Complex p0:Complex -- Complex
title	Decimal logarithm

desc
Decimal logarithm of *p0*.
end

example
50+20i log10 5 round -- 1.7312+0.16525i
end
*/
func Log10Complex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Log10(a0)
	zc.PushComplex(c, r0)
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

/*
oper	round
func	RoundComplex p0:Complex n:Int -- Complex
alias	r
title	Round to a given precision

desc
Rounds the real and imaginary numbers of *p0* to *n* digits using the half
up rounding mode.
end

example
1.047197551196598+5.948293i 3 round -- 1.047+5.948i
end
*/
func RoundComplex(c zc.Calc) {
	n := zc.PopInt32(c)
	a0 := zc.PopComplex(c)
	r := types.NewDecimalFromFloat(real(a0))
	i := types.NewDecimalFromFloat(imag(a0))
	posSign := ""
	if i.IsPositive() {
		posSign = "+"
	}
	r0 := fmt.Sprintf("%v%v%vi", r.Round(n), posSign, i.Round(n))
	zc.PushString(c, r0)
}

/*
oper	phase
func	PhaseComplex p0:Complex -- Float
title	Phase (argument)

desc
The phase, or argument, of *p0* in the range of [-π, π]
end

example
1+1i phase 5 round -- 0.7854
end
*/

func PhaseComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Phase(a0)
	zc.PushFloat(c, r0)
}

/*
oper	polar
func	PolarComplex p0:Complex -- r:Float θ:Float
title	Complex to polar

desc
The absolute value *r* and phase *θ* of *p0*.
end

example
2i polar pi div 1 round -- 2 | 0.5
end
*/
func PolarComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r, θ := cmplx.Polar(a0)
	zc.PushFloat(c, r)
	zc.PushFloat(c, θ)
}

/*
oper	rect
func	RectComplex r:Float θ:Float -- Complex
title	Polar to complex

desc
The complex number with polar coordinates *r* and *θ*.
end

example
2 0.5 pi mul rect imag -- 2
end
*/
func RectComplex(c zc.Calc) {
	θ := zc.PopFloat(c)
	r := zc.PopFloat(c)
	r0 := cmplx.Rect(r, θ)
	zc.PushComplex(c, r0)
}

/*
oper	sin
func 	SinComplex p0:Complex -- Complex
title	Sine

desc
Sine of *p0*.
end

example
2+2i sin 5 round -- 3.42095-1.50931i
end
*/
func SinComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Sin(a0)
	zc.PushComplex(c, r0)
}

/*
oper	sinh
func	SinhComplex p0:Complex -- Complex
title 	Hyperbolic sine

desc
Hyperbolic sine of *p0*.
end

example
2+2i sinh 5 round -- -1.50931+3.42095i
end
*/
func SinhComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Sinh(a0)
	zc.PushComplex(c, r0)
}

/*
oper	tan
func	TanComplex p0:Complex -- Complex
title	Tangent

desc
Tangent of *p0*.
end

example
2+2i tan 5 round -- -0.02839+1.02384i
end
*/
func TanComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Tan(a0)
	zc.PushComplex(c, r0)
}

/*
oper 	tanh
func	TanhComplex p0:Complex -- Complex
title	Hyperbolic tangent

desc
Hyperbolic tangent of *p0*.
end

example
2+2i tanh 5 round -- 1.02384-0.02839i
end
*/
func TanhComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Tanh(a0)
	zc.PushComplex(c, r0)
}
