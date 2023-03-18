package runtime

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "runtime",
		ScriptPath: "zc:zlib/dev/runtime/runtime.zc",
		Natives: map[string]zc.CalcFunc{
			"exports": Exports,
			"funcs":   Funcs,
		},
	}
)
