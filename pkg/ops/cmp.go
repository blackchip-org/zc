package ops

import (
	"sort"

	"github.com/blackchip-org/zc/v5/pkg/zc"
)

/*
oper	eq
func	EqBigInt   p0:BigInt   p1:BigInt   -- BigInt
func 	EqDecimal  p0:Decimal  p1:Decimal  -- Decimal
func	EqFloat    p0:Float    p1:Float    -- Float
func	EqRational p0:Rational p1:Rational -- Rational
func	EqComplex  p0:Complex  p1:Complex  -- Complex
func	Is         p0:Str      p1:Str      -- Str
title	Equal

desc
`true` if *p0* and *p1* are equal, otherwise `false`.
end

example
c 1234.56 1,234.56   eq -- true
c 1234.56 1234.56000 eq -- true
c 1234.56 $1,234.56  eq -- true
c 1234.56 +1,234.56  eq -- true
end
*/
func EqBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(c, r0)
}

func EqDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(c, r0)
}

func EqFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 == a1
	zc.PushBool(c, r0)
}

func EqRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(c, r0)
}

func EqComplex(c zc.Calc) {
	a1 := zc.PopComplex(c)
	a0 := zc.PopComplex(c)
	r0 := a0 == a1
	zc.PushBool(c, r0)
}

/*
oper	gt
func	GtBigInt   p0:BigInt   p1:BigInt   -- Bool
func	GtDecimal  p0:Decimal  p1:Decimal  -- Bool
func	GtRational p0:Rational p1:Rational -- Bool
func	GtFloat    p0:Float    p1:Float    -- Bool
func 	GtStr      p0:Str      p1:Str      -- Bool
alias	greater-than
title	Greater than

desc
`true` if *p0* is greater than *p1*, otherwise `false`.
end

example
c 1  0 gt -- true
c 0  0 gt -- false
c -1 0 gt -- false
end
*/
func GtBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(c, r0)
}

func GtDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(c, r0)
}

func GtFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 > a1
	zc.PushBool(c, r0)
}

func GtRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(c, r0)
}

func GtStr(c zc.Calc) {
	a1 := zc.PopString(c)
	a0 := zc.PopString(c)
	r0 := a0 > a1
	zc.PushBool(c, r0)
}

/*
oper	gte
func	GteBigInt   p0:BigInt   p1:BigInt   -- Bool
func	GteDecimal  p0:Decimal  p1:Decimal  -- Bool
func	GteRational p0:Rational p1:Rational -- Bool
func	GteFloat    p0:Float    p1:Float    -- Bool
func 	GteStr      p0:Str      p1:Str      -- Bool
alias	greater-than-or-equal
title	Greater than or equal

desc
`true` if *p0* is greater than or equal *p1*, otherwise `false`.
end

example
c 1  0 gte -- true
c 0  0 gte -- true
c -1 0 gte -- false
end
*/
func GteBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(c, r0)
}

func GteDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(c, r0)
}

func GteFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 >= a1
	zc.PushBool(c, r0)
}

func GteRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(c, r0)
}

func GteStr(c zc.Calc) {
	a1 := zc.PopString(c)
	a0 := zc.PopString(c)
	r0 := a0 >= a1
	zc.PushBool(c, r0)
}

/*
oper	is
func	Is p0:Str p1:Str -- Bool
title	Byte equality

desc
`true` if *p0* and *p1* are the same, byte by byte.
end

example
c 1.2 1.20 is -- false
c 1.2 1.2 is  -- true
end
*/
func Is(c zc.Calc) {
	a1 := zc.PopString(c)
	a0 := zc.PopString(c)
	r0 := a0 == a1
	zc.PushBool(c, r0)
}

/*
oper	lt
func	LtBigInt   p0:BigInt   p1:BigInt   -- Bool
func	LtDecimal  p0:Decimal  p1:Decimal  -- Bool
func	LtRational p0:Rational p1:Rational -- Bool
func	LtFloat    p0:Float    p1:Float    -- Bool
func 	LtStr      p0:Str      p1:Str      -- Bool
alias	less-than
title	Less than

desc
`true` if *p0* is less than *p1*, otherwise `false`.
end

example
c 1  0 lt -- false
c 0  0 lt -- false
c -1 0 lt -- true
end
*/
func LtBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(c, r0)
}

func LtDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(c, r0)
}

func LtFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 < a1
	zc.PushBool(c, r0)
}

func LtRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(c, r0)
}

func LtStr(c zc.Calc) {
	a1 := zc.PopString(c)
	a0 := zc.PopString(c)
	r0 := a0 < a1
	zc.PushBool(c, r0)
}

/*
oper	lte
func	LteBigInt   p0:BigInt   p1:BigInt   -- Bool
func	LteDecimal  p0:Decimal  p1:Decimal  -- Bool
func	LteRational p0:Rational p1:Rational -- Bool
func	LteFloat    p0:Float    p1:Float    -- Bool
func 	LteStr      p0:Str      p1:Str      -- Bool
alias	less-than-or-equal
title	Less than or equal

desc
`true` if *p0* is less than or equal to *p1*, otherwise `false`.
end

example
c 1  0 lte -- false
c 0  0 lte -- true
c -1 0 lte -- true
end
*/
func LteBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(c, r0)
}

func LteDecimal(c zc.Calc) {
	a1 := zc.PopDecimal(c)
	a0 := zc.PopDecimal(c)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(c, r0)
}

func LteFloat(c zc.Calc) {
	a1 := zc.PopFloat(c)
	a0 := zc.PopFloat(c)
	r0 := a0 <= a1
	zc.PushBool(c, r0)
}

func LteRational(c zc.Calc) {
	a1 := zc.PopRational(c)
	a0 := zc.PopRational(c)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(c, r0)
}

func LteStr(c zc.Calc) {
	a1 := zc.PopString(c)
	a0 := zc.PopString(c)
	r0 := a0 <= a1
	zc.PushBool(c, r0)
}

/*
oper	sort
func	Sort        Val* -- Val*
title	Type-aware stack sort

desc
Sorts the stack in a type-aware way. When two elements on the stack are
compared for their sort order, a common type between those values is found and
the comparison operation uses that type. If a common type cannot be found,
the comparison is performed as plain text. Use `sort-str` to treat all
values on the stack as strings.
end

example
c 3 1 2 sort -- 1 | 2 | 3
c 003 01.1 2 sort -- 01.1 | 2 | 003
end
*/
func Sort(c zc.Calc) {
	vals := c.Stack()
	sort.Sort(zc.SortInterface(vals))
	c.SetStack(vals)
}

/*
oper	sort-str
func	SortStr        Val* -- Val*
title	String based stack sort

desc
Sorts the stack by their values as strings. Use `sort` for a type-aware
sort.
end

example
c 003 01.1 2 sort-str -- 003 | 01.1 | 2
end
*/
func SortStr(c zc.Calc) {
	vals := c.Stack()
	sort.Strings(vals)
	c.SetStack(vals)
}
