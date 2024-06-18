package ops

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/fn"
)

var (
	AddBigInt = zc.Op{
		Params: []zc.Type{zc.BigInt, zc.BigInt},
		Func:   fn.AddBigInt,
	}
	MulBigInt = zc.Op{
		Params: []zc.Type{zc.BigInt, zc.BigInt},
		Func:   fn.MulBigInt,
	}
	PowBigInt = zc.Op{
		Params: []zc.Type{zc.BigInt, zc.BigInt},
		Func:   fn.PowBigInt,
	}
	SqrtBigInt = zc.Op{
		Params: []zc.Type{zc.BigInt},
		Func:   fn.SqrtBigInt,
	}
	SubBigInt = zc.Op{
		Params: []zc.Type{zc.BigInt, zc.BigInt},
		Func:   fn.SubBigInt,
	}
)
