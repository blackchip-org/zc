package ops

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/kinds"
)

var (
	Down = zc.OpDecl{
		Name:       "down",
		Params:     []zc.Kind{kinds.Val},
		VarParams:  true,
		Returns:    []zc.Kind{kinds.Val},
		VarReturns: true,
		Func: func(c *zc.OpContext) {
			top := len(c.Args) - 1
			c.Returns = append(
				[]any{c.Args[top]},
				c.Args[:top]...,
			)
		},
	}
	Dup = zc.OpDecl{
		Name:    "dup",
		Params:  []zc.Kind{kinds.Val},
		Returns: []zc.Kind{kinds.Val, kinds.Val},
		Func: func(c *zc.OpContext) {
			x := c.Catalog.Copy(c.Args[0])
			c.Returns = []any{x, c.Args[0]}
		},
	}
)
