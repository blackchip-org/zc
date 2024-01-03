package ops

import (
	"math/big"

	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper 	average
func	Average BigInt*   -- BigInt
func	-       Decimal*  -- Decimal
func	-       Float*    -- Float
func	-       Rational* -- Rational
func	-       Complex*  -- Complex
alias	avg
title	Average

desc
The average of all items on the stack.
end

example
0 100 25 75 avg -- 50
end
*/
func Average(c zc.Calc) {
	n := c.StackLen()
	if n > 0 {
		c.Eval("sum %v div", n)
	}
}

/*
oper 	factorial
func 	Factorial n:BigInt -- BigInt
alias	fact
title	Factorial

desc
The product of all positive integers less than or equal to *n*. If *n* is
negative, an invalid argument error is displayed.
end

example
c 3 fact -- 6
c 10 fact -- 3628800
end
*/
func Factorial(c zc.Calc) {
	zero := big.NewInt(0)
	one := big.NewInt(1)

	n := zc.PopBigInt(c)
	switch n.Cmp(zero) {
	case -1:
		zc.ErrInvalidArgs(c, "cannot be negative")
	case 0:
		zc.PushBigInt(c, one)
	case 1:
		acc := big.NewInt(1)
		i := big.NewInt(1)
		for i.Cmp(n) <= 0 {
			acc.Mul(acc, i)
			i.Add(i, one)
		}
		zc.PushBigInt(c, acc)
	}
}

/*
oper	prod
func	- BigInt*   -- BigInt
func	- Decimal*  -- Decimal
func	- Float*    -- Float
func	- Rational* -- Rational
func	- Complex*  -- Complex
macro	[mul] fold
title 	Product

desc
The product of all items on the stack.
end

example
1 2 3 4 5 -- 1 | 2 | 3 | 4 | 5
prod -- 120
end
*/

/*
oper	sum
func	- BigInt*   -- BigInt
func	- Decimal*  -- Decimal
func	- Float*    -- Float
func	- Rational* -- Rational
func	- Complex*  -- Complex
macro	[add] fold
title 	Summation

desc
The sum of all items on the stack.
end

example
1 2 3 4 5 -- 1 | 2 | 3 | 4 | 5
sum -- 15
end
*/
