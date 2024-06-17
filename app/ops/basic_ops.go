package ops

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/fn"
	"github.com/blackchip-org/zc/v6/app/types"
)

var (
	AddBigInt = zc.Op{
		Params: []zc.Type{types.BigInt, types.BigInt},
		Func:   fn.AddBigInt,
	}
	MulBigInt = zc.Op{
		Params: []zc.Type{types.BigInt, types.BigInt},
		Func:   fn.MulBigInt,
	}
	PowBigInt = zc.Op{
		Params: []zc.Type{types.BigInt, types.BigInt},
		Func:   fn.PowBigInt,
	}
	SqrtBigInt = zc.Op{
		Params: []zc.Type{types.BigInt},
		Func:   fn.SqrtBigInt,
	}
	SubBigInt = zc.Op{
		Params: []zc.Type{types.BigInt, types.BigInt},
		Func:   fn.SubBigInt,
	}
)
