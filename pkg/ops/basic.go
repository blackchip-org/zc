package ops

import (
	"math"
	"math/big"
	"math/cmplx"

	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper	add
func	AddBigInt   p0:BigInt   p1:BigInt   -- BigInt
func	AddDecimal  p0:Decimal  p1:Decimal  -- Decimal
func	AddFloat    p0:Float    p1:Float    -- Float
func    AddRational p0:Rational p1:Rational -- Rational
func	AddComplex  p0:Complex  p1:Complex  -- Complex
alias	a
alias	+
title 	Addition

desc
Adds the value of *p1* to *p0*.
end

example
6 -- 6
2 -- 6 | 2
a -- 8
end
*/
func AddBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Add(a0, a1)
	zc.PushBigInt(c, &r0)
}

func AddDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Add(a1)
	zc.PushDecimal(c, r0)
}

func AddFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 + a1
	zc.PushFloat(c, r0)
}

func AddRational(c zc.Calc) {
	var r0 big.Rat
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0.Add(a0, a1)
	zc.PushRational(c, &r0)
}

func AddComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 + a1
	zc.PushComplex(c, r0)
}

/*
oper	div
func	DivDecimal  p0:Decimal  p1:Decimal  -- Decimal
func	DivFloat    p0:Float    p1:Float    -- Float
func    DivRational p0:Rational p1:Rational -- Rational
func	DivComplex  p0:Complex  p1:Complex  -- Complex
alias	d
alias	/
title 	Division

desc
Divides the value of *p0* by *p1*. If *p1* is zero, a 'division by zero' error
is raised.
end

example
6 -- 6
2 -- 6 | 2
d -- 3
end
*/
func DivDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)

	if a1.IsZero() {
		zc.ErrDivisionByZero(c)
		return
	}

	r0 := a0.Div(a1).Truncate(19)
	zc.PushDecimal(c, r0)
}

func DivFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)

	if a1 == 0 {
		zc.ErrDivisionByZero(c)
		return
	}

	r0 := a0 / a1
	zc.PushFloat(c, r0)
}

func DivRational(c zc.Calc) {
	var r0 big.Rat
	var zero big.Rat
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)

	if a1.Cmp(&zero) == 0 {
		zc.ErrDivisionByZero(c)
		return
	}

	r0.Quo(a0, a1)
	zc.PushRational(c, &r0)
}

func DivComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)

	if real(a1) == 0 && imag(a1) == 0 {
		zc.ErrDivisionByZero(c)
		return
	}

	r0 := a0 / a1
	zc.PushComplex(c, r0)
}

/*
oper	mod
func	ModBigInt  p0:BigInt  p1:BigInt  -- BigInt
func	ModDecimal p0:Decimal p1:Decimal -- Decimal
func	ModFloat   p0:Float   p1:Float   -- Float
title	Modulus

desc
The modulus when *p0* is divided by *p1*. If *p1* is zero, a 'division by zero'
error is raised.
end

example
-7 2 mod -- 1
end
*/
func ModBigInt(c zc.Calc) {
	var r0 big.Int
	var zero big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)

	if a1.Cmp(&zero) == 0 {
		zc.ErrDivisionByZero(c)
		return
	}

	r0.Mod(a0, a1)
	zc.PushBigInt(c, &r0)
}

func ModDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)

	if a1.IsZero() {
		zc.ErrDivisionByZero(c)
		return
	}

	r0 := a0.Mod(a1)
	zc.PushDecimal(c, r0)
}

func ModFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)

	if a1 == 0 {
		zc.ErrDivisionByZero(c)
		return
	}

	r0 := math.Mod(a0, a1)
	zc.PushFloat(c, r0)
}

/*
oper	mul
func	MulBigInt   p0:BigInt   p1:BigInt   -- BigInt
func	MulDecimal  p0:Decimal  p1:Decimal  -- Decimal
func 	MulFloat    p0:Float    p1:Float    -- Float
func 	MulRational p0:Rational p1:Rational -- Rational
func	MulComplex  p0:Complex  p1:Complex  -- Complex
alias	m
alias	*
title	Multiplication

desc
Multiplies *p0* by *p1*.
end

example
6 -- 6
2 -- 6 | 2
m -- 12
end
*/
func MulBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Mul(a0, a1)
	zc.PushBigInt(c, &r0)
}

func MulDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Mul(a1).Round(17).Truncate(16)
	zc.PushDecimal(c, r0)
}

func MulFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 * a1
	zc.PushFloat(c, r0)
}

func MulRational(c zc.Calc) {
	var r0 big.Rat
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0.Mul(a0, a1)
	zc.PushRational(c, &r0)
}

func MulComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 * a1
	zc.PushComplex(c, r0)
}

/*
oper	neg
func	NegBigInt   p0:BigInt   -- BigInt
func	NegDecimal  p0:Decimal  -- Decimal
func	NegFloat    p0:Float    -- Float
func	NegRational p0:Rational -- Rational
title	Negation

desc
Changes the sign of `p0`.
end

example
-6 -- -6
neg -- 6
neg -- -6
end
*/
func NegBigInt(c zc.Calc) {
	var r0 big.Int
	a0 := zc.PopBigInt(c)
	r0.Neg(a0)
	zc.PushBigInt(c, &r0)
}

func NegDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Neg()
	zc.PushDecimal(c, r0)
}

func NegFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	r0 := -a0
	zc.PushFloat(c, r0)
}

func NegRational(c zc.Calc) {
	var r0 big.Rat
	a0 := zc.PopRational(c)
	r0.Neg(a0)
	zc.PushRational(c, &r0)
}

/*
oper	pow
func	PowBigInt  p0:BigInt  p1:BigInt  -- BigInt
func	PowFloat   p0:Float   p1:Float   -- Float
func	PowComplex p0:Complex p1:Complex -- Complex
alias	**
alias	^
title	Exponentiation

desc
Raises *p0* to the power of *p1*.
end

example
6 -- 6
2 -- 6 | 2
pow -- 36
end
*/
func PowBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Exp(a0, a1, nil)
	zc.PushBigInt(c, &r0)
}

func PowFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := math.Pow(a0, a1)
	zc.PushFloat(c, r0)
}

func PowComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := cmplx.Pow(a0, a1)
	zc.PushComplex(c, r0)
}

/*
oper	rem
func	RemBigInt p0:BigInt p1:BigInt -- BigInt
func	RemFloat  p0:Float  p1:Float  -- Float
title	Remainder

desc
The remainder when *p0* is divided by *p1*. If *p1* is zero, a
'division by zero' error is raised.
end

example
-7 -- -7
2 -- -7 | 2
rem -- -1
end
*/
func RemBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Rem(a0, a1)
	zc.PushBigInt(c, &r0)
}

func RemFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := math.Remainder(a0, a1)
	zc.PushFloat(c, r0)
}

/*
oper	sign
func	SignBigInt   p0:BigInt   -- Int
func	SignDecimal  p0:Decimal  -- Int
func	SignFloat    p0:Float    -- Int
func 	SignRational p0:Rational -- Int
title	Sign

desc
Returns `-1` if *p0* is negative, `1` if *p0* is positive, or `0` if *p0*
is zero.
end

example
c -6 sign -- -1
c 6 sign -- 1
c 0 sign -- 0
end
*/
func SignBigInt(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	r0 := a0.Sign()
	zc.PushInt(c, r0)
}

func SignDecimal(c zc.Calc) {
	a0 := zc.PopDecimal(c)
	r0 := a0.Sign()
	zc.PushInt(c, r0)
}

func SignFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)

	r0 := 0
	if a0 > 0 {
		r0 = 1
	}
	if a0 < 0 {
		r0 = -1
	}

	zc.PushInt(c, r0)
}

func SignRational(c zc.Calc) {
	a0 := zc.PopRational(c)
	r0 := a0.Sign()
	zc.PushInt(c, r0)
}

/*
oper	sqrt
func	SqrtFloat   p0:Float   -- Float
func	-           p0:Float   -- Complex
func 	SqrtComplex p0:Complex -- Complex
alias	square-root
title	Square root

desc
The square root of *p0*. If *p0* is a positive or zero then a Float is
returned. If *p0* is negative, a Complex is returned.
end

example
256 -- 256
sqrt -- 16
end
*/
func SqrtFloat(c zc.Calc) {
	a0 := zc.PopFloat(c)
	if a0 < 0 {
		r0 := cmplx.Sqrt(complex(a0, 0))
		zc.PushComplex(c, r0)
	} else {
		r0 := math.Sqrt(a0)
		zc.PushFloat(c, r0)
	}
}

func SqrtComplex(c zc.Calc) {
	a0 := zc.PopComplex(c)
	r0 := cmplx.Sqrt(a0)
	zc.PushComplex(c, r0)
}

/*
oper	sub
func	SubBigInt   p0:BigInt   p1:BigInt   -- BigInt
func 	SubDecimal  p0:Decimal  p1:Decimal  -- Decimal
func	SubFloat    p0:Float    p1:Float    -- Float
func	SubRational p0:Rational p1:Rational -- Rational
func	SubComplex  p0:Complex  p1:Complex  -- Complex
alias	s
alias	-
title	Subtraction

desc
Subtract *p1* from *p0*.
end

example
6 -- 6
2 -- 6 | 2
s -- 4
end
*/
func SubBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Sub(a0, a1)
	zc.PushBigInt(c, &r0)
}

func SubDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Sub(a1)
	zc.PushDecimal(c, r0)
}

func SubFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 - a1
	zc.PushFloat(c, r0)
}

func SubRational(c zc.Calc) {
	var r0 big.Rat
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0.Sub(a0, a1)
	zc.PushRational(c, &r0)
}

func SubComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 - a1
	zc.PushComplex(c, r0)
}
