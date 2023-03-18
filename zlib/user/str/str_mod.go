package str

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "str",
		ScriptPath: "zc:zlib/user/str/str.zc",
		Natives: map[string]zc.CalcFunc{
			"join":        Join,
			"left":        Left,
			"len":         Len,
			"lower":       LowerStr,
			"right":       Right,
			"split":       Split,
			"starts-with": StartsWith,
			"upper":       UpperStr,
		},
	}
)
