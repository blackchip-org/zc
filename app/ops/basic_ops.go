package ops

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/fn"
	"github.com/blackchip-org/zc/v6/app/types"
)

var (
	BigIntAdd = zc.Op{
		Params: []zc.Type{types.BigInt, types.BigInt},
		Func:   fn.BigIntAdd,
	}
	BigIntMul = zc.Op{
		Params: []zc.Type{types.BigInt, types.BigInt},
		Func:   fn.BigIntMul,
	}
)
