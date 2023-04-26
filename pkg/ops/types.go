package ops

import (
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
