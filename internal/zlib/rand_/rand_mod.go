package rand_

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "rand",
		ScriptPath: "zc:zlib/rand_/rand.zc",
		Init:       InitRand,
		Natives: map[string]zc.CalcFunc{
			"float":   FloatRand,
			"int":     IntRand,
			"seed":    Seed,
			"seed=":   SeedGet,
			"shuffle": Shuffle,
		},
	}
)
