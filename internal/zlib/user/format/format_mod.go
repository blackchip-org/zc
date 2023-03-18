package format

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "format",
		Include:    true,
		ScriptPath: "zc:zlib/user/format/format.zc",
		Natives: map[string]zc.CalcFunc{
			"auto-currency":  AutoCurrency,
			"auto-currency=": AutoCurrencyGet,
			"auto-format":    AutoFormat,
			"auto-format=":   AutoFormatGet,
			"format":         Format_,
			"int-layout":     IntLayout,
			"int-layout=":    IntLayoutGet,
			"min-digits":     MinDigits,
			"min-digits=":    MinDigitsGet,
			"point":          Point,
			"point=":         PointGet,
			"round":          Round,
			"rounding-mode":  RoundingMode,
			"rounding-mode=": RoundingModeGet,
		},
	}
)
