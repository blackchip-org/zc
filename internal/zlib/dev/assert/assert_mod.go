package assert

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "assert",
		ScriptPath: "zc:zlib/dev/assert/assert.zc",
		Natives: map[string]zc.CalcFunc{
			"bigint":  BigInt,
			"bool":    BoolAssert,
			"decimal": Decimal,
			"float":   Float,
			"int":     Int,
			"int64":   Int64,
			"int32":   Int32,
		},
	}
)
