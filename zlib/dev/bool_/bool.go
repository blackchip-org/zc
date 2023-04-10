package bool_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
)

var (
	And = zc.FuncGeneric(ops.And)
	Eq  = zc.FuncGeneric(ops.Eq)
	Neq = zc.FuncGeneric(ops.Neq)
	Gt  = zc.FuncGeneric(ops.Gt)
	Gte = zc.FuncGeneric(ops.Gte)
	Lt  = zc.FuncGeneric(ops.Lt)
	Lte = zc.FuncGeneric(ops.Lte)
	Or  = zc.FuncGeneric(ops.Or)
	Not = zc.FuncGeneric(ops.Not)
)
