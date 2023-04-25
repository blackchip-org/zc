package ops

import (
	"image/color"
	"math"

	"github.com/blackchip-org/zc/pkg/ansi"
	"github.com/blackchip-org/zc/pkg/zc"
)

func CMYKToRGB(calc zc.Calc) {
	k := zc.PopUint8(calc)
	y := zc.PopUint8(calc)
	m := zc.PopUint8(calc)
	c := zc.PopUint8(calc)

	r, g, b := color.CMYKToRGB(c, m, y, k)

	zc.PushUint8(calc, r)
	zc.PushUint8(calc, g)
	zc.PushUint8(calc, b)
}

func ColorSample(calc zc.Calc) {
	b := zc.PopUint8(calc)
	g := zc.PopUint8(calc)
	r := zc.PopUint8(calc)

	r0 := "#raw:" + ansi.BgColor24(r, g, b) + "     " + ansi.Reset
	zc.PushString(calc, r0)
}

func HSLtoRGB(calc zc.Calc) {
	l := zc.PopFloat(calc)
	s := zc.PopFloat(calc)
	h := zc.PopFloat(calc)

	l = zc.Clamp(l, 0.0, 1.0)
	s = zc.Clamp(s, 0.0, 1.0)

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

	zc.PushUint8(calc, r)
	zc.PushUint8(calc, g)
	zc.PushUint8(calc, b)
}

func RBGToCMYK(calc zc.Calc) {
	b := zc.PopUint8(calc)
	g := zc.PopUint8(calc)
	r := zc.PopUint8(calc)

	c, m, y, k := color.RGBToCMYK(r, g, b)

	zc.PushUint8(calc, c)
	zc.PushUint8(calc, m)
	zc.PushUint8(calc, y)
	zc.PushUint8(calc, k)
}

func RGBToHSL(calc zc.Calc) {
	b := zc.PopUint8(calc)
	g := zc.PopUint8(calc)
	r := zc.PopUint8(calc)

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

	zc.PushFloat(calc, h)
	zc.PushFloat(calc, s)
	zc.PushFloat(calc, l)
}
