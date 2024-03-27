package ops

import (
	"math/big"

	"github.com/blackchip-org/zc/v6/pkg/zc"
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
oper standard-deviation-pop
func StandardDeviationPop BigInt*   -- BigInt
func -                    Decimal*  -- Decimal
func -                    Float*    -- Float
func -                    Rational* -- Rational
alias stdev-p
macro var-p sqrt
title Population standard deviation

desc
Standard deviation of the stack where it contains the entire population
end

example
2 4 4 4 5 5 7 9 -- 2 | 4 | 4 | 4 | 5 | 5 | 7 | 9
stdev-p -- 2
end
*/

/*
oper standard-deviation-samp
func StandardDeviationSamp BigInt*   -- BigInt
func -                     Decimal*  -- Decimal
func -                     Float*    -- Float
func -                     Rational* -- Rational
alias stdev-s
macro var-s sqrt
title Sample standard deviation

desc
Standard deviation of the stack where it contains a sample of the population
end

example
2 4 4 4 5 5 7 9 -- 2 | 4 | 4 | 4 | 5 | 5 | 7 | 9
stdev-s 2 round -- 2.14
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

/*
oper variance-pop
func VariancePop BigInt*   -- BigInt
func -           Decimal*  -- Decimal
func -           Float*    -- Float
func -           Rational* -- Rational
alias var-p
title Population variance

desc
Variance of the stack where it contains the entire population
end

example
2 4 4 4 5 5 7 9 -- 2 | 4 | 4 | 4 | 5 | 5 | 7 | 9
var-p -- 4
end
*/
func VariancePop(c zc.Calc) {
	variance(c, 0)
}

/*
oper variance-samp
func VarianceSamp BigInt*   -- BigInt
func -            Decimal*  -- Decimal
func -            Float*    -- Float
func -            Rational* -- Rational
alias var-s
title Sample variance

desc
Variance of the stack where it contains a sample of the population
end

example
2 4 4 4 5 5 7 9 -- 2 | 4 | 4 | 4 | 5 | 5 | 7 | 9
var-s 2 round -- 4.57
end
*/
func VarianceSamp(c zc.Calc) {
	variance(c, -1)
}

func variance(c zc.Calc, nadj int) {
	n := c.StackLen()
	if n == 0 {
		return
	}
	data := c.Stack()
	if err := c.Eval("average"); err != nil {
		return
	}
	mean := zc.PopInt(c)
	c.SetStack(data)
	zc.PushInt(c, mean)
	c.MustEval("[sub square] [map] 2 apply sum")
	zc.PushInt(c, n+nadj)
	c.MustEval("div")
}
