package zc_

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "zc",
		Include:    true,
		ScriptPath: "zc:zlib/user/zc_/zc.zc",
		Natives: map[string]zc.CalcFunc{
			"version": Version,
		},
	}
)
