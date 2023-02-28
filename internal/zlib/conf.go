package zlib

import (
	"github.com/blackchip-org/zc"
)

func Locale(env *zc.Env) error {
	locale, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	return env.Calc.SetLocale(locale)
}

func LocaleGet(env *zc.Env) error {
	env.Stack.Push(env.Calc.Locale)
	return nil
}
