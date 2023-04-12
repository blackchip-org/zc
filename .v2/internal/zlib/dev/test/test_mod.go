package test

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "test",
		ScriptPath: "zc:zlib/dev/test/test.zc",
		Natives: map[string]zc.CalcFunc{
			"file": TestFile,
		},
	}
)
