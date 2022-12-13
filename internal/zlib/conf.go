package zlib

import (
	"fmt"

	"github.com/blackchip-org/zc"
)

func Places(calc *zc.Calc) error {
	places, err := calc.PopInt32()
	if err != nil {
		return err
	}
	if places < 0 {
		return fmt.Errorf("invalid number of places: %v", places)
	}
	calc.Settings.Places = places
	calc.Info = "ok"
	return nil
}

func PlacesGet(calc *zc.Calc) error {
	calc.Stack.Push(calc.FormatInt(int(calc.Settings.Places)))
	return nil
}

func RoundMode(calc *zc.Calc) error {
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	mode, ok := zc.ParseRoundingMode(a)
	if !ok {
		return fmt.Errorf("invalid rounding mode: %v", a)
	}
	calc.Settings.RoundingMode = mode
	calc.Print("ok")
	return err
}

func RoundModeGet(calc *zc.Calc) error {
	calc.Stack.Push(calc.Settings.RoundingMode.String())
	return nil
}
