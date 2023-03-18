package dev

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "dev",
		ScriptPath: "zc:zlib/dev/dev.zc",
		Include:    true,
		Natives: map[string]zc.CalcFunc{
			"abort":     Abort,
			"exit":      Exit,
			"nothing":   Nothing,
			"trace":     Trace,
			"trace-off": TraceOff,
			"quote":     Quote,
			"undef":     Undef,
		},
	}
)
