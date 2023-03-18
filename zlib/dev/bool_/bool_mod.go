package bool_

import "github.com/blackchip-org/zc"

var (
	Mod = zc.ModuleDef{
		Name:       "bool",
		Include:    true,
		ScriptPath: "zc:zlib/dev/bool_/bool.zc",
		Natives: map[string]zc.CalcFunc{
			"and": And,
			"eq":  Eq,
			"neq": Neq,
			"gt":  Gt,
			"gte": Gte,
			"lt":  Lt,
			"lte": Lte,
			"not": Not,
			"or":  Or,
		},
	}
	ModBigInt = zc.ModuleDef{
		Name:       "bool.bigint",
		ScriptPath: "zc:zlib/dev/bool_/bigint.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqBigInt,
			"gt":  GtBigInt,
			"gte": GteBigInt,
			"neq": NeqBigInt,
			"lt":  LtBigInt,
			"lte": LteBigInt,
		},
	}
	ModDecimal = zc.ModuleDef{
		Name:       "bool.decimal",
		ScriptPath: "zc:zlib/dev/bool_/decimal.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqDecimal,
			"gt":  GtDecimal,
			"gte": GteDecimal,
			"neq": NeqDecimal,
			"lt":  LtDecimal,
			"lte": LteDecimal,
		},
	}
	ModStr = zc.ModuleDef{
		Name:       "str-bool",
		ScriptPath: "zc:zlib/dev/bool_/str.zc",
		Natives: map[string]zc.CalcFunc{
			"eq":  EqStr,
			"gt":  GtStr,
			"gte": GteStr,
			"neq": NeqStr,
			"lt":  LtStr,
			"lte": LteStr,
		},
	}
)
