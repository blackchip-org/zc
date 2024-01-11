package ops

import (
	"math/big"

	"github.com/blackchip-org/zc/v5/pkg/zc"
)

/*
oper 	div-rem
func	DivRemBigInt p0:Decimal p1:Decimal -- r:Decimal q:Decimal
alias	dr
title	Division with remainder

desc
Divides *p0* by *p1* and returns the quotient *q* and remainder *r*.
end

example
1234 100 div-rem -- 34 | 12
end
*/
func DivRemBigInt(c zc.Calc) {
	var q, r big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	q.QuoRem(a0, a1, &r)
	zc.PushBigInt(c, &r)
	zc.Annotate(c, "remainder")
	zc.PushBigInt(c, &q)
}
