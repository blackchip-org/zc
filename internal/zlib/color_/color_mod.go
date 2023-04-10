package color_

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "color",
		ScriptPath: "zc:zlib/color_/color.zc",
		Natives: map[string]zc.CalcFunc{
			"cmyk-rgb": CMYKToRGB,
			"hsl-rgb":  HSLtoRGB,
			"rgb-cmyk": RBGToCMYK,
			"rgb-hsl":  RGBToHSL,
			"sample":   Sample,
		},
	}
)
