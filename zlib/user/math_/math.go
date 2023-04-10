package math_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
)

var (
	Abs   = zc.FuncGeneric(ops.Abs)
	Add   = zc.FuncGeneric(ops.Add)
	Ceil  = zc.FuncGeneric(ops.Ceil)
	Div   = zc.FuncGeneric(ops.Div)
	Floor = zc.FuncGeneric(ops.Floor)
	Mod   = zc.FuncGeneric(ops.Mod)
	Mul   = zc.FuncGeneric(ops.Mul)
	Neg   = zc.FuncGeneric(ops.Neg)
	Pow   = zc.FuncGeneric(ops.Pow)
	Rem   = zc.FuncGeneric(ops.Rem)
	Sign  = zc.FuncGeneric(ops.Sign)
	Sqrt  = zc.FuncGeneric(ops.Sqrt)
	Sub   = zc.FuncGeneric(ops.Sub)
)

func Sum(env *zc.Env) error {
	for env.Stack.Len() > 1 {
		if err := Add(env); err != nil {
			return err
		}
	}
	return nil
}
