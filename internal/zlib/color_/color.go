package color_

import (
	"image/color"
	"math"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/ansi"
)

/*
var CMYKToRGB = zc.FuncNN([]types.Type{
}


	func(args []types.Value) []types.Value {
	c := types.Uint8.Native(args[0])
	m := types.Uint8.Native(args[1])
	y := types.Uint8.Native(args[2])
	k := types.Uint8.Native(args[3])

	r, g, b := color.CMYKToRGB(c, m, y, k)

	return []types.Value{
		types.Unit8.Value(r),
		types.Unit8.Value(g),
		types.Unit8.Value(b),
	}
}, types.Uint8, types.Uint8, types.Uint8)
*/

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

func HSLtoRGB(env *zc.Env) error {
	l, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	s, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	h, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}

	// Written by ChatGPT
	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var rp, gp, bp float64
	if h >= 0 && h < 60 {
		rp, gp, bp = c, x, 0
	} else if h >= 60 && h < 120 {
		rp, gp, bp = x, c, 0
	} else if h >= 120 && h < 180 {
		rp, gp, bp = 0, c, x
	} else if h >= 180 && h < 240 {
		rp, gp, bp = 0, x, c
	} else if h >= 240 && h < 300 {
		rp, gp, bp = x, 0, c
	} else {
		rp, gp, bp = c, 0, x
	}

	r := uint8((rp + m) * 255)
	g := uint8((gp + m) * 255)
	b := uint8((bp + m) * 255)
	// End ChatGPT

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

func RGBToHSL(env *zc.Env) error {
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
	// Written by ChatGPT
	// Convert RGB values from 0-255 range to 0-1 range
	rf := float64(r) / 255.0
	gf := float64(g) / 255.0
	bf := float64(b) / 255.0

	// Find the minimum and maximum values of the RGB components
	cmax := math.Max(math.Max(rf, gf), bf)
	cmin := math.Min(math.Min(rf, gf), bf)

	// Calculate lightness
	l := (cmax + cmin) / 2.0

	var s float64
	// If the minimum and maximum values are equal, then the color is gray and saturation is 0
	if cmax == cmin {
		s = 0.0
	} else {
		// Calculate saturation
		if l < 0.5 {
			s = (cmax - cmin) / (cmax + cmin)
		} else {
			s = (cmax - cmin) / (2.0 - cmax - cmin)
		}
	}

	// Calculate hue
	rc := (cmax - rf) / (cmax - cmin)
	gc := (cmax - gf) / (cmax - cmin)
	bc := (cmax - bf) / (cmax - cmin)

	var h float64
	if rf == cmax {
		h = bc - gc
	} else if gf == cmax {
		h = 2.0 + rc - bc
	} else {
		h = 4.0 + gc - rc
	}
	h *= 60.0
	if h < 0 {
		h += 360.0
	}
	// End ChatGPT

	env.Stack.PushFloat(h)
	env.Stack.PushFloat(s)
	env.Stack.PushFloat(l)
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
