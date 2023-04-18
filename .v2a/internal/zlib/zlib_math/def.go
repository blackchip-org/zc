package zlib_math

import (
	"github.com/blackchip-org/zc"
)

/*
var Def = zc.ModuleDef{
	Name:   "math",
	Script: "zlib_math/math.zc",
	GenOps: map[string]zc.FuncDecl{
		"add": zc.GenOpsDecl(types.BigInt, types.BigInt),
	},
}
*/

var Def = zc.NewModuleDef("math", "zlib_math/math.zc")
