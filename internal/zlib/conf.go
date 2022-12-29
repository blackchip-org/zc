package zlib

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/zc"
)

func IntFormat(env *zc.Env) error {
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
	env.Calc.IntFormat = format
	return nil
}

func IntFormatGet(env *zc.Env) error {
	env.Stack.Push(env.Calc.IntFormat)
	return nil
}

func MinDigits(env *zc.Env) error {
	digits, err := env.Stack.PopUint()
	if err != nil {
		return err
	}
	env.Calc.MinDigits = digits
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
	env.Calc.Info = "ok"
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
	return nil
}

func PointGet(env *zc.Env) error {
	env.Stack.Push(string(env.Calc.Point))
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
