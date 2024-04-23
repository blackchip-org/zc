package ints

import "github.com/blackchip-org/zc/v6"

var Volume = zc.Volume{
	Name:  "ints",
	Kinds: []zc.Kind{IntArchKind},
	Ops: []zc.Op{
		Add,
	},
}
