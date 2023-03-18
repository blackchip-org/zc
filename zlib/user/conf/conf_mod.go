package conf

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "conf",
		ScriptPath: "zc:zlib/user/conf/conf.zc",
		Natives: map[string]zc.CalcFunc{
			"locale":  Locale,
			"locale=": LocaleGet,
		},
	}
)
