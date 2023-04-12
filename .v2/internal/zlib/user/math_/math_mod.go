package math_

import "github.com/blackchip-org/zc"

var (
	Module = zc.ModuleDef{
		Name:       "math",
		Include:    true,
		ScriptPath: "zc:zlib/user/math_/math.zc",
		Natives: map[string]zc.CalcFunc{
			"abs":   Abs,
			"add":   Add,
			"ceil":  Ceil,
			"div":   Div,
			"floor": Floor,
			"mod":   Mod,
			"mul":   Mul,
			"neg":   Neg,
			"pow":   Pow,
			"rem":   Rem,
			"sign":  Sign,
			"sqrt":  Sqrt,
			"sub":   Sub,
			"sum":   Sum,
		},
	}
)
