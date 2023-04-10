package format

import (
	"fmt"

	"github.com/blackchip-org/zc"
)

func Round(env *zc.Env) error {
	places, err := env.Stack.PopInt32()
	if err != nil {
		return err
	}
	value, err := env.Stack.PopDecimal()
	if err != nil {
		return err
	}
	fn, ok := zc.RoundingFuncsFix[env.Calc.Config().RoundingMode]
	if !ok {
		return fmt.Errorf("invalid rounding mode: %v", env.Calc.Config().RoundingMode)
	}
	r := fn(value, places)
	env.Stack.PushDecimal(r)
	return nil
}

func RoundingMode(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	mode, ok := zc.ParseRoundingMode(a)
	if !ok {
		return fmt.Errorf("invalid rounding mode: %v", a)
	}
	env.Calc.Config().RoundingMode = mode
	env.Calc.SetInfo("rounding-mode set to %v", zc.Quote(a))
	return err
}

func RoundingModeGet(env *zc.Env) error {
	env.Stack.Push(env.Calc.Config().RoundingMode.String())
	return nil
}
