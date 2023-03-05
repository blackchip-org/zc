package zlib

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/zc"
)

func AutoCurrency(env *zc.Env) error {
	b, err := env.Stack.PopBool()
	if err != nil {
		return err
	}
	env.Calc.AutoCurrency = b
	env.Calc.Info = fmt.Sprintf("auto-currency set to %v", b)
	return nil
}

func AutoCurrencyGet(env *zc.Env) error {
	env.Stack.PushBool(env.Calc.AutoCurrency)
	return nil
}

func AutoFormat(env *zc.Env) error {
	b, err := env.Stack.PopBool()
	if err != nil {
		return err
	}
	env.Calc.AutoFormat = b
	env.Calc.Info = fmt.Sprintf("auto-format set to %v", b)
	return nil
}

func AutoFormatGet(env *zc.Env) error {
	env.Stack.PushBool(env.Calc.AutoFormat)
	return nil
}

func Format_(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	z := env.Calc.ApplyLayout(a)
	env.Stack.Push(z)
	return nil
}

func IntLayout(env *zc.Env) error {
	format, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	for _, ch := range format {
		if strings.ContainsRune(zc.ValidSeparators, ch) {
			continue
		}
		if ch == '0' {
			continue
		}
		return fmt.Errorf("invalid character in format: %v", string(ch))
	}
	env.Calc.IntLayout = format
	env.Calc.Info = fmt.Sprintf("int-layout set to '%v'", format)
	return nil
}

func IntLayoutGet(env *zc.Env) error {
	env.Stack.Push(env.Calc.IntLayout)
	return nil
}

func MinDigits(env *zc.Env) error {
	digits, err := env.Stack.PopUint()
	if err != nil {
		return err
	}
	env.Calc.MinDigits = digits
	env.Calc.Info = fmt.Sprintf("min-digits set to %v", digits)
	return nil
}

func MinDigitsGet(env *zc.Env) error {
	env.Stack.PushUint(env.Calc.MinDigits)
	return nil
}

func Precision(env *zc.Env) error {
	prec, err := env.Stack.PopInt32()
	if err != nil {
		return err
	}
	if prec < 0 {
		return fmt.Errorf("invalid precision: %v", prec)
	}
	env.Calc.Precision = prec
	env.Calc.Info = fmt.Sprintf("precision set to %v", prec)
	return nil
}

func PrecisionGet(env *zc.Env) error {
	env.Stack.PushInt32(env.Calc.Precision)
	return nil
}

func Point(env *zc.Env) error {
	point, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	if utf8.RuneCountInString(point) != 1 {
		return fmt.Errorf("invalid decimal point: %v", point)
	}
	ch, _ := utf8.DecodeRuneInString(point)
	if !strings.ContainsRune(zc.ValidPoints, ch) {
		return fmt.Errorf("invalid decimal point: %v", point)
	}
	env.Calc.Point = ch
	env.Calc.Info = fmt.Sprintf("point set to '%v'", string(ch))
	return nil
}

func PointGet(env *zc.Env) error {
	env.Stack.Push(string(env.Calc.Point))
	return nil
}

func Round(env *zc.Env) error {
	places, err := env.Stack.PopInt32()
	if err != nil {
		return err
	}
	value, err := env.Stack.PopFixed()
	if err != nil {
		return err
	}
	fn, ok := zc.RoundingFuncsFix[env.Calc.RoundingMode]
	if !ok {
		return fmt.Errorf("invalid rounding mode: %v", env.Calc.RoundingMode)
	}
	r := fn(value, places)
	env.Stack.PushFixed(r)
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
	env.Calc.RoundingMode = mode
	env.Calc.Info = fmt.Sprintf("rounding-mode set to %v", zc.Quote(a))
	return err
}

func RoundingModeGet(env *zc.Env) error {
	env.Stack.Push(env.Calc.RoundingMode.String())
	return nil
}
