package stack

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "stack",
		Include:    true,
		ScriptPath: "zc:zlib/user/stack/stack.zc",
		Natives: map[string]zc.CalcFunc{
			"at":      At,
			"reverse": Reverse,
			"n":       N,
		},
	}
)
