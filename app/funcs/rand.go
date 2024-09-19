package funcs

import (
	"strconv"
	"unicode"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vars"
	"github.com/blackchip-org/zc/v6/msg"
)

func Rand(c zc.Calc) {
	v := vars.ForRand(c)
	zc.Float64.Push(c, v.Rand.Float64())
}

func RandInt(c zc.Calc) {
	v := vars.ForRand(c)
	max := zc.Int.Pop(c)
	if max < 1 {
		c.Raise(msg.ErrInvalidArg("%v < 1", max))
		return
	}
	zc.Int.Push(c, v.Rand.IntN(max)+1)
}

func RandSeed(c zc.Calc) {
	v := vars.ForRand(c)
	v.Seed = zc.BigInt.Pop(c)
	lo, hi := v.SplitSeed()
	v.Source.Seed(lo, hi)
	c.Notify(msg.SeedSet(v.Seed))
}

func RandSeedGet(c zc.Calc) {
	v := vars.ForRand(c)
	zc.BigInt.Push(c, v.Seed)
}

func RandTake(c zc.Calc) {
	v := vars.ForRand(c)
	n := c.Len()
	i := v.Rand.IntN(n)
	c.SetStack([]zc.Item{c.Stack()[i]})
}

func Roll(c zc.Calc) {
	v := vars.ForRand(c)
	var s scan.Scanner

	spec := zc.String.Pop(c)
	s.InitFromString("", spec)

	var num, sides int64
	var err error

	tok, ok := s.Eval(scan.IntRule)
	if !ok {
		num = 1
	} else {
		num, err = strconv.ParseInt(tok.Val, 10, 64)
		if err != nil {
			c.Raise(msg.ErrInvalidArg(spec))
			return
		}
	}

	if unicode.ToLower(s.This) != 'd' {
		c.Raise(msg.ErrInvalidArg(spec))
		return
	}
	s.Discard()

	tok, ok = s.Eval(scan.IntRule)
	if !ok {
		c.Raise(msg.ErrInvalidArg(spec))
		return
	}
	sides, err = strconv.ParseInt(tok.Val, 10, 64)
	if err != nil {
		c.Raise(msg.ErrInvalidArg(spec))
		return
	}

	for i := int64(0); i < num; i++ {
		r := v.Rand.IntN(int(sides)) + 1
		zc.Int.Push(c, r)
	}
}

func Shuffle(c zc.Calc) {
	v := vars.ForRand(c)
	items := c.Stack()
	v.Rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	c.SetStack(items)
}
