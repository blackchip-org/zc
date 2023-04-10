package math_

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "math",
		Include:    true,
		ScriptPath: "zc:zlib/user/math_/math.zc",
		Natives: map[string]zc.CalcFunc{
			"abs":   Abs,
			"add":   Add,
			"ceil":  Ceil,
			"div":   Div,
			"floor": Floor,
			"mod":   Modulus,
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
	/*
		ModBigInt = zc.ModuleDef{
			Name:       "math.bigint",
			ScriptPath: "zc:zlib/user/math_/bigint.zc",
			Natives: map[string]zc.CalcFunc{
				"abs":   AbsBigInt,
				"add":   AddBigInt,
				"ceil":  CeilBigInt,
				"div":   DivBigInt,
				"floor": FloorBigInt,
				"mod":   ModulusBigInt,
				"mul":   MulBigInt,
				"neg":   NegBigInt,
				"pow":   PowBigInt,
				"rem":   RemBigInt,
				"sign":  SignBigInt,
				"sub":   SubBigInt,
			},
		}
		ModDecimal = zc.ModuleDef{
			Name:       "math.decimal",
			ScriptPath: "zc:zlib/user/math_/decimal.zc",
			Natives: map[string]zc.CalcFunc{
				"abs":   AbsDecimal,
				"add":   AddDecimal,
				"ceil":  CeilDecimal,
				"div":   DivDecimal,
				"floor": FloorDecimal,
				"mod":   ModulusDecimal,
				"mul":   MulDecimal,
				"neg":   NegDecimal,
				"pow":   PowDecimal,
				"rem":   RemDecimal,
				"sign":  SignDecimal,
				"sub":   SubDecimal,
			},
		}
		ModFloat = zc.ModuleDef{
			Name:       "math.float",
			ScriptPath: "zc:zlib/user/math_/float.zc",
			Natives: map[string]zc.CalcFunc{
				"sqrt": SqrtFloat,
			},
		}
	*/
)
