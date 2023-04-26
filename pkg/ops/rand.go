package ops

import (
	"math/rand"
	"time"
	"unicode"

	"github.com/blackchip-org/zc/pkg/scanner"
	"github.com/blackchip-org/zc/pkg/zc"
)

type randState struct {
	rand *rand.Rand
	seed int64
}

func getRandState(c zc.Calc) *randState {
	s, ok := c.State("rand")
	if !ok {
		seed := time.Now().UnixMilli()
		s = &randState{
			seed: seed,
			rand: rand.New(rand.NewSource(seed)),
		}
		c.NewState("rand", s)
	}
	return s.(*randState)
}

/*
oper	rand
func	Rand -- Float
title	Random float

desc
Random number between 0 and 1.
end

example
0 rand-seed -- *seed set to 0*
rand -- 0.9451961492941164
end
*/
func Rand(c zc.Calc) {
	s := getRandState(c)
	r0 := s.rand.Float64()
	zc.PushFloat(c, r0)
}

/*
oper	rand-choice
func	RandChoice ... a0:Val ... -- a0:Val
title	Randomly select item on stack

desc
Randomly select an item on the stack.
end

example
2 rand-seed -- *seed set to 2*
1 2 3 4 5 6 -- 1 | 2 | 3 | 4 | 5 | 6
rand-choice -- 5
end
*/
func RandChoice(c zc.Calc) {
	s := getRandState(c)
	n := c.StackLen()
	i := s.rand.Intn(n)
	c.SetStack([]string{c.Stack()[i]})
}

/*
oper	rand-int
func	RandInt n:Int -- Int
title	Random integer

desc
Random integer between 1 and *n*.
end

example
0 rand-seed -- *seed set to 0*
10 rand-int -- 5
end
*/
func RandInt(c zc.Calc) {
	s := getRandState(c)
	max := zc.PopInt(c)
	if max < 1 {
		zc.ErrInvalidArgs(c)
		return
	}
	r0 := s.rand.Intn(max) + 1
	zc.PushInt(c, r0)
}

/*
oper 	rand-seed
func	RandSeed seed:Int64 --
title	Set the random number seed

desc
Sets the random number seed to *seed*.
end

example
1 rand-seed -- *seed set to 1*
10 rand-int -- 2
end
*/
func RandSeed(c zc.Calc) {
	s := getRandState(c)
	s.seed = zc.PopInt64(c)
	s.rand = rand.New(rand.NewSource(s.seed))
	c.SetInfo("seed set to %v", s.seed)
}

/*
oper	rand-seed=
func	RandSeedGet -- Int64
title	Get the random number seed

desc
Gets the random number seed.
end

example
3 rand-seed -- *seed set to 3*
rand-seed= -- 3
end
*/
func RandSeedGet(c zc.Calc) {
	s := getRandState(c)
	zc.PushInt64(c, s.seed)
}

/*
oper	roll
func	Roll dice:Str -- Int*
title	Dice roller

desc
Rolls dice as specified by *dice* in standard dice notation. The argument
*dice* may start with the number of dice to roll, followed by the literal
character `d`, and then the number of faces found on each die. For example,
use `3d6` to roll three six sided dice.
end

example
99 rand-seed -- *seed set to 99*
3d6 roll -- 6 | 2 | 1
sum -- `9
end
*/
func Roll(c zc.Calc) {
	state := getRandState(c)
	var s scanner.Scanner
	a0 := zc.PopString(c)
	s.SetString(a0)

	nTok := s.Scan(scanner.UInt)
	if unicode.ToLower(s.Ch) != 'd' {
		zc.ErrInvalidArgs(c)
		return
	}
	s.Next()
	sidesTok := s.Scan(scanner.UInt)
	if sidesTok == "" || !s.End() {
		zc.ErrInvalidArgs(c)
		return
	}
	if nTok == "" {
		nTok = "1"
	}
	n := zc.Int.MustParse(nTok)
	sides := zc.Int.MustParse(sidesTok)

	for i := 0; i < n; i++ {
		r := state.rand.Intn(sides) + 1
		zc.PushInt(c, r)
	}
}

/*
oper	shuffle
func	Shuffle ... -- ...
title	Shuffle the stack

desc
Shuffle the stack.
end

example
0 rand-seed -- *seed set to 0*
1 2 3 4 5 6 -- 1 | 2 | 3 | 4 | 5 | 6
shuffle -- 5 | 4 | 1 | 3 | 2 | 6`
end
*/
func Shuffle(c zc.Calc) {
	s := getRandState(c)
	as := c.Stack()
	s.rand.Shuffle(len(as), func(i, j int) {
		as[i], as[j] = as[j], as[i]
	})
	c.SetStack(as)
}
