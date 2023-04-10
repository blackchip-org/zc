package io_

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "io",
		ScriptPath: "zc:zlib/dev/io_/io.zc",
		Natives: map[string]zc.CalcFunc{
			"print": Print,
		},
	}
)
