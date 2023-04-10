package format

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "format",
		Include:    true,
		ScriptPath: "zc:zlib/user/format/format.zc",
		Natives: map[string]zc.CalcFunc{
			"round":          Round,
			"rounding-mode":  RoundingMode,
			"rounding-mode=": RoundingModeGet,
		},
	}
)
