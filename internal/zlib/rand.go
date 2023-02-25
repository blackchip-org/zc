package zlib

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/blackchip-org/zc"
)

type randState struct {
	rand *rand.Rand
	seed int64
}

func InitRand(env *zc.Env) error {
	seed := time.Now().UnixMilli()
	env.Calc.States["rand"] = &randState{
		seed: seed,
		rand: rand.New(rand.NewSource(seed)),
	}
	return nil
}

func getRandState(env *zc.Env) *randState {
	return env.Calc.States["rand"].(*randState)
}

func FloatRand(env *zc.Env) error {
	s := getRandState(env)
	r := s.rand.Float64()
	env.Stack.PushFloat(r)
	return nil
}

func IntRand(env *zc.Env) error {
	s := getRandState(env)
	max, err := env.Stack.PopInt()
	if err != nil {
		return err
	}
	if max < 1 {
		return fmt.Errorf("'%v' must be greater than zero", max)
	}
	r := s.rand.Intn(max) + 1
	env.Stack.PushInt(r)
	return nil
}

func Seed(env *zc.Env) error {
	s := getRandState(env)
	seed, err := env.Stack.PopInt64()
	if err != nil {
		return err
	}
	s.seed = seed
	s.rand = rand.New(rand.NewSource(seed))
	env.Calc.Info = fmt.Sprintf("seed set to %v", seed)
	return nil
}

func SeedGet(env *zc.Env) error {
	s := getRandState(env)
	env.Stack.PushInt64(s.seed)
	return nil
}

func Shuffle(env *zc.Env) error {
	s := getRandState(env)
	items := env.Stack.Items()
	s.rand.Shuffle(len(items), func(i, j int) {
		items[i], items[j] = items[j], items[i]
	})
	env.Stack.Clear()
	env.Stack.PushAll(items)
	return nil
}
