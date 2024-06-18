package ops

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/fn"
)

var (
	TuckBigInt = zc.Op{
		Params: []zc.Type{zc.Any, zc.BigInt},
		Func:   fn.TuckBigInt,
	}
)
