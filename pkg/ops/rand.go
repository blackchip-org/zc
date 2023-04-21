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

func Choice(c zc.Calc) {
	s := getRandState(c)
	n := c.StackLen()
	i := s.rand.Intn(n)
	c.SetStack([]string{c.Stack()[i]})
}

func Rand(c zc.Calc) {
	s := getRandState(c)
	r0 := s.rand.Float64()
	zc.PushFloat(c, r0)
}

func RandInt(c zc.Calc) {
	s := getRandState(c)
	max := zc.PopInt(c)
	if max < 1 {
		zc.ErrInvalidArgument(c, c.Op(), max)
		return
	}
	r0 := s.rand.Intn(max) + 1
	zc.PushInt(c, r0)
}

func Roll(c zc.Calc) {
	state := getRandState(c)
	var s scanner.Scanner
	a0 := zc.PopString(c)
	s.SetString(a0)

	nTok := s.Scan(scanner.UInt)
	if unicode.ToLower(s.Ch) != 'd' {
		zc.ErrInvalidArgument(c, c.Op(), a0)
		return
	}
	s.Next()
	sidesTok := s.Scan(scanner.UInt)
	if sidesTok == "" || !s.End() {
		zc.ErrInvalidArgument(c, c.Op(), a0)
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

func Seed(c zc.Calc) {
	s := getRandState(c)
	s.seed = zc.PopInt64(c)
	s.rand = rand.New(rand.NewSource(s.seed))
	c.SetInfo("seed set to %v", s.seed)
}

func SeedGet(c zc.Calc) {
	s := getRandState(c)
	zc.PushInt64(c, s.seed)
}

func Shuffle(c zc.Calc) {
	s := getRandState(c)
	as := c.Stack()
	s.rand.Shuffle(len(as), func(i, j int) {
		as[i], as[j] = as[j], as[i]
	})
	c.SetStack(as)
}
