package ops

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6/pkg/zc"
)

/*
oper	join-bits
func	JoinBits BigInt* n:Int -- BigInt*
alias	joinb
title	Join a binary number by bit width

desc
Join all numbers in the stack, which contain *n* bits into a single number.
If a value on the stack contains more than *n* bits, an error is raised.
end

example
c 0xa 0xb 0xc 0xd 4 joinb hex -- 0xabcd
c 0xab 0xcd 8 joinb hex -- 0xabcd
end
*/
func JoinBits(c zc.Calc) {
	n := zc.PopInt(c)
	if n < 0 {
		zc.ErrInvalidArgs(c, "bit count cannot be negative")
		return
	}

	var b big.Int
	bi := 0
	for c.StackLen() > 0 {
		s := zc.PopString(c)
		a, ok := zc.BigInt.Parse(s)
		if !ok {
			zc.ErrInvalidArgs(c, fmt.Sprintf("expected integer but got '%v'", s))
			return
		}
		if a.BitLen() > n {
			zc.ErrInvalidArgs(c, fmt.Sprintf("'%v' exceeded the bit width", s))
		}
		for ai := 0; ai < n; ai++ {
			b.SetBit(&b, bi, a.Bit(ai))
			bi++
		}
	}
	c.SetStack([]string{})
	zc.PushBigInt(c, &b)
}

/*
oper	split-bits
func	SplitBits p0:BigInt n:Int -- BigInt*
alias	splitb
title	Split a binary number by bit width

desc
Split *p0* every *n* bits starting with the least significant bit.
end

example
c 0xabcd 4 splitb /hex map -- 0xa | 0xb | 0xc | 0xd
c 0xabcd 8 splitb /hex map -- 0xab | 0xcd
end
*/
func SplitBits(c zc.Calc) {
	n := zc.PopInt(c)
	if n < 0 {
		zc.ErrInvalidArgs(c, "bit count cannot be negative")
		return
	}

	var results []*big.Int
	a := zc.PopBigInt(c)
	b := big.NewInt(0)
	bi := 0

	for ai := 0; ai < a.BitLen(); ai++ {
		b.SetBit(b, bi, a.Bit(ai))
		bi++
		if bi >= n {
			bi = 0
			results = append(results, b)
			b = big.NewInt(0)
		}
	}
	if bi > 0 {
		results = append(results, b)
	}
	for i := len(results) - 1; i >= 0; i-- {
		zc.PushBigInt(c, results[i])
	}
}
