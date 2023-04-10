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

func OpsFunc(op ops.Def) func(*Env) error {
	return func(e *Env) error {
		var args []types.Value
		for i := 0; i < op.NArg; i++ {
			s, err := e.Stack.Pop()
			if err != nil {
				return err
			}
			args = append([]types.Value{types.Parse(s)}, args...)
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

func Func(fn func(types.Value) types.Value, param types.Type) CalcFunc {
	return func(e *Env) error {
		s, err := e.Stack.Pop()
		if err != nil {
			return err
		}
		g := types.Parse(s)
		cg, err := types.To(g, param)
		if err != nil {
			return err
		}
		r := fn(cg)
		e.Stack.Push(r.Format())
		return nil
	}
}

func FuncNN(params []types.Type, fn func([]types.Value) []types.Value) CalcFunc {
	return func(e *Env) error {
		var args []types.Value
		for _, param := range params {
			s, err := e.Stack.Pop()
			if err != nil {
				return err
			}
			g, err := param.ParseValue(s)
			if err != nil {
				return err
			}
			args = append(args, g)
		}
		rs := fn(args)
		for _, r := range rs {
			e.Stack.Push(r.Format())
		}
		return nil
	}
}
