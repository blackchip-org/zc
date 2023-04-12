package geo

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "geo",
		ScriptPath: "zc:zlib/geo/geo.zc",
		/*
			Natives: map[string]zc.CalcFunc{
				"decimal-degrees":  DecimalDegrees,
				"degrees-minutes":  DegreesMinutes,
				"round-coordinate": RoundCoordinate,
				"transform":        Transform,
			},
		*/
	}
)
