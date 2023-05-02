package ops

import "github.com/blackchip-org/zc/pkg/zc"

/*
oper	sum
func	- BigInt* -- BigInt
func	- Decimal* -- Decimal
func	- Float* -- Float
func	- Rational* -- Rational
func	- Complex* -- Complex
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

/*
oper 	average
func	Average BigInt* -- BigInt
func	- Decimal* -- Decimal
func	- Float* -- Float
func	- Rational* -- Rational
func	- Complex* -- Complex
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
	c.Eval("sum %v div", n)
}
