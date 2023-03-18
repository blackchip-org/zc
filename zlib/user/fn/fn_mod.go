package fn

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "fn",
		ScriptPath: "zc:zlib/user/fn/fn.zc",
		Include:    true,
		Natives: map[string]zc.CalcFunc{
			"eval":   Eval,
			"filter": Filter,
			"fold":   Fold,
			"map":    Map,
			"repeat": Repeat,
		},
	}
)
