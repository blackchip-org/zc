package ops

import (
	"math"
	"math/big"
	"math/cmplx"

	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper	abs
func	AbsBigInt p0:BigInt -- BigInt
func	AbsDecimal p0:Decimal -- Decimal
func	AbsFloat p0:Float -- Float
func	AbsRational p0:Rational -- Rational
func 	AbsComplex p0:Complex -- Complex
title	Absolute value

desc
For non-complex numbers: If *p0* is less than zero, the negated value of *p0*,
otherwise *p0*.

For complex numbers: The distance of *p0* from zero in the complex plane.
end

example
-6 -- -6
abs -- 6
end
*/
func AbsBigInt(c zc.Calc) {
	var r0 big.Int
	a0 := zc.PopBigInt(c)
	r0.Abs(a0)
	zc.PushBigInt(c, &r0)
}

func AbsDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Abs()
	zc.PushDecimal(c, r0)
}

func AbsFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Abs(a0)
	zc.PushFloat(c, r0)
}

func AbsRational(c zc.Calc) {
	var r0 big.Rat
	a0 := zc.PopRational(c)
	r0.Abs(a0)
	zc.PushRational(c, &r0)
}

func AbsComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Abs(a0)
	zc.PushFloat(c, r0)
}

/*
oper 	acos
func	AcosFloat p0:Float -- Float
title	Inverse cosine

desc
Inverse cosine of *p0* in radians.
end

example
0.5 acos -- 1.0471975511965976
end
*/
func AcosFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Acos(a0)
	zc.PushFloat(c, r0)
}

/*
oper	acosh
func	AcoshFloat p0:Float -- Float
title	Inverse hyperbolic cosine

desc
Inverse cosine of *p0* in radians.
end

example
2 acosh -- 1.3169578969248166
end
*/
func AcoshFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Acosh(a0)
	zc.PushFloat(c, r0)
}

/*
oper	asin
func	AsinFloat p0:Float -- Float
title	Inverse sine

desc
Inverse sine of *p0* in radians.
end

example
0.5 asin -- 0.5235987755982989
end
*/
func AsinFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Asin(a0)
	zc.PushFloat(c, r0)
}

/*
oper	asinh
func	AsinhFloat p0:Float -- Float
title 	Inverse hyperbolic sine

desc
Inverse hyperbolic sine of *p0* in radians.
end

example
2 asinh -- 1.4436354751788103
end
*/
func AsinhFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Asinh(a0)
	zc.PushFloat(c, r0)
}

/*
oper	atan
func	AtanFloat p0:Float -- Float
title	Inverse tangent

desc
Inverse tangent of *p0* in radians.
end

example
0.5 atan -- 0.4636476090008061
end
*/
func AtanFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Atan(a0)
	zc.PushFloat(c, r0)
}

/*
oper	atanh
func	AtanhFloat p0:Float -- Float
title	Inverse hyperbolic tangent

desc
Inverse hyperbolic tangent of *p0* in radians.
end

example
0.5 atanh -- 0.5493061443340548
end
*/
func AtanhFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Atanh(a0)
	zc.PushFloat(c, r0)
}

/*
oper	ceil
func	NoOp p0:BigInt -- BigInt
func	CeilDecimal p0:Decimal -- Decimal
func	CeilFloat p0:Float -- Float
title	Ceiling

desc
The nearest integer value greater than or equal to *p0*.
end

example
6.12 -- 6.12
ceil -- 7
end
*/
func CeilDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Ceil()
	zc.PushDecimal(c, r0)
}

func CeilFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Ceil(a0)
	zc.PushFloat(c, r0)
}

/*
oper	cos
func	CosFloat p0:Float -- Float
title	Cosine

desc
Cosine of *p0* in radians.
end

example
2 cos -- -0.4161468365471424
end
*/
func CosFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Cos(a0)
	zc.PushFloat(c, r0)
}

/*
oper 	cosh
func	CoshFloat p0:Float -- Float
title	Hyperbolic cosine

desc
Hyperbolic cosine of *p0* in radians.
end

example
2 cosh -- 3.7621956910836314
end
*/
func CoshFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Cosh(a0)
	zc.PushFloat(c, r0)
}

/*
oper	e
func	- -- Float
macro	2.718281828459045
title	Euler's number

desc
Euler's number, the natural logarithm base.
end

example
e -- 2.718281828459045
end
*/

/*
oper	exp
func	ExpFloat p0:Float -- Float
title	Natural exponential

desc
Natural exponential of *p0*.
end

example
2 exp -- 7.38905609893065
end
*/
func ExpFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Exp(a0)
	zc.PushFloat(c, r0)
}

/*
oper	floor
func	NoOp p0:BigInt -- BigInt
func	FloorDecimal p0:Decimal -- Decimal
func	FloorFloat p0:Float -- Float
title	Floor

desc
The nearest integer value less than or equal to *p0*.
end

example
6.12 -- 6.12
floor -- 6
end
*/
func FloorDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Floor()
	zc.PushDecimal(c, r0)
}

func FloorFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Floor(a0)
	zc.PushFloat(c, r0)
}

/*
oper	log
func	LogFloat p0:Float -- Float
title	Natural logarithm

desc
Natural logarithm of *p0*.
end

example
8 log -- 2.0794415416798357
end
*/
func LogFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Log(a0)
	zc.PushFloat(c, r0)
}

/*
oper	log10
func	Log10Float p0:Float -- Float
title	Decimal logarithm

desc
Decimal logarithm of *p0*.
end

example
50 log10 -- 1.6989700043360187
end
*/
func Log10Float(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Log10(a0)
	zc.PushFloat(c, r0)
}

/*
oper	log2
func	Log2Float p0:Float -- Float
title	Binary logarithm

desc
Binary logarithm of *p0*.
end

example
250 log2 -- 7.965784284662087
end
*/
func Log2Float(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Log2(a0)
	zc.PushFloat(c, r0)
}

/*
oper	pi
func	- -- Float
macro	3.141592653589793
alias	π
title	Circumference to diameter ratio

desc
Circumference to diameter ratio of a circle
end

example
pi -- 3.141592653589793
end
*/

/*
oper	sin
func	SinFloat p0:Float -- Float
title	Sine

desc
Sine of *p0* in radians
end

example
2 sin -- 0.9092974268256816
end
*/
func SinFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Sin(a0)
	zc.PushFloat(c, r0)
}

/*
oper	sinh
func	SinhFloat p0:Float -- Float
title 	Hyperbolic sine

desc
Hyperbolic sine of *p0* in radians.
end

example
2 sinh -- 3.626860407847019
end
*/
func SinhFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Sinh(a0)
	zc.PushFloat(c, r0)
}

/*
oper	tan
func	TanFloat p0:Float -- Float
title	Tangent

desc
Tangent of *p0* in radians
end

example
2 tan -- -2.185039863261519
end
*/
func TanFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Tan(a0)
	zc.PushFloat(c, r0)
}

/*
oper 	tanh
func	TanhFloat p0:Float -- Float
title	Hyperbolic tangent

desc
Hyperbolic tangent of *p0* in radians.
end

example
2 tanh -- 0.9640275800758169
end
*/
func TanhFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := math.Tanh(a0)
	zc.PushFloat(c, r0)
}
