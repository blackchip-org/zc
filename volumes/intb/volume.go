package intb

import (
	"github.com/blackchip-org/zc/v6"
)

var Volume = zc.Volume{
	Name:  "int",
	Kinds: []zc.Kind{IntKind},
	Ops: []zc.Op{
		Add,
	},
}
