package prog

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "prog",
		Include:    true,
		ScriptPath: "zc:zlib/user/prog/prog.zc",
		Natives: map[string]zc.CalcFunc{
			"and":   And,
			"bin":   Bin,
			"bit":   Bit,
			"bits":  Bits,
			"bytes": Bytes,
			"dec":   Dec,
			"hex":   Hex,
			"lsh":   Lsh,
			"not":   Not,
			"oct":   Oct,
			"or":    Or,
			"rsh":   Rsh,
			"xor":   Xor,
		},
	}
)
