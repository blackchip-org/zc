package zlib

import (
	"fmt"

	"github.com/blackchip-org/zc"
)

func Places(env *zc.Env) error {
	places, err := env.Stack.PopInt32()
	if err != nil {
		return err
	}
	if places < 0 {
		return fmt.Errorf("invalid number of places: %v", places)
	}
	env.Calc.Places = places
	env.Calc.Info = "ok"
	return nil
}

func PlacesGet(env *zc.Env) error {
	env.Stack.PushInt32(env.Calc.Places)
	return nil
}

func RoundMode(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	mode, ok := zc.ParseRoundingMode(a)
	if !ok {
		return fmt.Errorf("invalid rounding mode: %v", a)
	}
	env.Calc.RoundingMode = mode
	env.Calc.Info = "ok"
	return err
}

func RoundModeGet(env *zc.Env) error {
	env.Stack.Push(env.Calc.RoundingMode.String())
	return nil
}
