package ops

import (
	"math/big"

	"github.com/blackchip-org/zc/v5/pkg/zc"
)

/*
oper 	fibonacci
func 	Fibonacci n:BigInt -- BigInt*
alias	fib
title	Fibonacci sequence

desc
Calculates the *n*th element in the Fibonacci sequence. The value of *n*
must be equal to or greater than zero.
end

example
1 5 seq /fib map -- 1 | 1 | 2 | 3 | 5
end
*/
func Fibonacci(c zc.Calc) {
	var zero big.Int
	one := big.NewInt(1)

	n := zc.PopBigInt(c)
	if n.Cmp(&zero) < 0 {
		zc.ErrInvalidArgs(c, "element index cannot be negative")
		return
	}

	if n.Cmp(&zero) == 0 {
		zc.PushBigInt(c, &zero)
		return
	}
	if n.Cmp(one) == 0 {
		zc.PushBigInt(c, one)
		return
	}

	var acc big.Int
	f0 := big.NewInt(0)
	f1 := big.NewInt(1)
	i := big.NewInt(1)

	for i.Cmp(n) < 0 {
		i.Add(i, one)
		acc.Add(f0, f1)
		f0.Set(f1)
		f1.Set(&acc)
	}
	zc.PushBigInt(c, f1)
}

/*
oper	sequence
func	Sequence p0:BigInt p1:BigInt -- BigInt*
alias	seq
title	Sequence of integers

desc
Adds the integers from *p0* to *p1* to the stack. If *p0* is greater than
*p1*, the list of integers is in decreasing order
end

example
4 8 seq -- 4 | 5 | 6 | 7 | 8
c 8 4 seq -- 8 | 7 | 6 | 5 | 4
end
*/
func Sequence(c zc.Calc) {
	to := zc.PopBigInt(c)
	from := zc.PopBigInt(c)
	one := big.NewInt(1)

	if from.Cmp(to) <= 0 {
		i := from
		for i.Cmp(to) <= 0 {
			zc.PushBigInt(c, i)
			i.Add(i, one)
		}
	} else {
		i := from
		for i.Cmp(to) >= 0 {
			zc.PushBigInt(c, i)
			i.Sub(i, one)
		}
	}
}
