package bool_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
)

var (
	And = zc.OpsFunc(ops.And)
	Eq  = zc.OpsFunc(ops.Eq)
	Neq = zc.OpsFunc(ops.Neq)
	Gt  = zc.OpsFunc(ops.Gt)
	Gte = zc.OpsFunc(ops.Gte)
	Lt  = zc.OpsFunc(ops.Lt)
	Lte = zc.OpsFunc(ops.Lte)
	Or  = zc.OpsFunc(ops.Or)
	Not = zc.OpsFunc(ops.Not)
)
