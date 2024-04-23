package zcalc

import "github.com/blackchip-org/zc/v6"

var Volume = zc.Volume{
	Name: "zcalc",
	Ops: []zc.Op{
		Down,
		Dup,
	},
}
