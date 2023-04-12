package math_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
)

var (
	Abs   = zc.OpsFunc(ops.Abs)
	Add   = zc.OpsFunc(ops.Add)
	Ceil  = zc.OpsFunc(ops.Ceil)
	Div   = zc.OpsFunc(ops.Div)
	Floor = zc.OpsFunc(ops.Floor)
	Mod   = zc.OpsFunc(ops.Mod)
	Mul   = zc.OpsFunc(ops.Mul)
	Neg   = zc.OpsFunc(ops.Neg)
	Pow   = zc.OpsFunc(ops.Pow)
	Rem   = zc.OpsFunc(ops.Rem)
	Sign  = zc.OpsFunc(ops.Sign)
	Sqrt  = zc.OpsFunc(ops.Sqrt)
	Sub   = zc.OpsFunc(ops.Sub)
)

func Sum(env *zc.Env) error {
	for env.Stack.Len() > 1 {
		if err := Add(env); err != nil {
			return err
		}
	}
	return nil
}
