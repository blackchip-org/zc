package zc

import (
	"github.com/blackchip-org/zc/ops"
	"github.com/blackchip-org/zc/types"
)

func FuncFloat1(fn func(float64) float64) func(*Env) error {
	return func(e *Env) error {
		x, err := e.Stack.PopFloat()
		if err != nil {
			return err
		}
		z := fn(x)
		e.Stack.PushFloat(z)
		return nil
	}
}

func FuncGeneric(op ops.Def) func(*Env) error {
	return func(e *Env) error {
		var args []types.Generic
		for i := 0; i < op.NArg; i++ {
			s, err := e.Stack.Pop()
			if err != nil {
				return err
			}
			args = append([]types.Generic{types.Parse(s)}, args...)
		}
		result, err := ops.Eval(op, args)
		if err != nil {
			return err
		}
		for _, r := range result {
			e.Stack.Push(r.Format())
		}
		return nil
	}
}
