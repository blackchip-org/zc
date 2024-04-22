package ops

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/kinds"
)

var (
	AddInt = zc.OpDecl{
		Name:    "add-int",
		Params:  []zc.Kind{kinds.Int, kinds.Int},
		Returns: []zc.Kind{kinds.Int},
		Func: func(c *zc.OpContext) {
			y := kinds.Int.As(c.Args[0])
			x := kinds.Int.As(c.Args[1])
			x.Add(x, y)
			c.Returns = []any{x}
		},
	}
)
