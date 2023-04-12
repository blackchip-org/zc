package sci

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "sci",
		Include:    true,
		ScriptPath: "zc:zlib/user/sci/sci.zc",
		Natives: map[string]zc.CalcFunc{
			"acos":                Acos,
			"acosh":               Acosh,
			"asin":                Asin,
			"asinh":               Asinh,
			"atan":                Atan,
			"atanh":               Atanh,
			"cos":                 Cos,
			"cosh":                Cosh,
			"exp":                 Exp,
			"log":                 Log,
			"log10":               Log10,
			"log2":                Log2,
			"scientific-notation": ScientificNotation,
			"sin":                 Sin,
			"sinh":                Sinh,
			"tan":                 Tan,
			"tanh":                Tanh,
		},
	}
)
