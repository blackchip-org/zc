package zlib

import (
	"image/color"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/ansi"
)

func CMYKToRGB(env *zc.Env) error {
	k, err := env.Stack.PopUint8()
	if err != nil {
		return err
	}
	y, err := env.Stack.PopUint8()
	if err != nil {
		return err
	}
	m, err := env.Stack.PopUint8()
	if err != nil {
		return err
	}
	c, err := env.Stack.PopUint8()
	if err != nil {
		return err
	}

	r, g, b := color.CMYKToRGB(c, m, y, k)
	env.Stack.PushUint8(r)
	env.Stack.PushUint8(g)
	env.Stack.PushUint8(b)
	return nil
}

func RBGToCMYK(env *zc.Env) error {
	b, err := env.Stack.PopUint8()
	if err != nil {
		return err
	}
	g, err := env.Stack.PopUint8()
	if err != nil {
		return err
	}
	r, err := env.Stack.PopUint8()
	if err != nil {
		return err
	}

	c, m, y, k := color.RGBToCMYK(r, g, b)
	env.Stack.PushUint8(c)
	env.Stack.PushUint8(m)
	env.Stack.PushUint8(y)
	env.Stack.PushUint8(k)
	return nil
}

func Sample(env *zc.Env) error {
	b, err := env.Stack.PopUint8()
	if err != nil {
		return nil
	}
	g, err := env.Stack.PopUint8()
	if err != nil {
		return nil
	}
	r, err := env.Stack.PopUint8()
	if err != nil {
		return nil
	}
	s := "#raw:" + ansi.BgColor24(r, g, b) + "     " + ansi.Reset
	env.Stack.Push(s)
	return nil
}
