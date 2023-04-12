package zc

import (
	"fmt"

	"github.com/blackchip-org/zc/coll"
	"github.com/blackchip-org/zc/types"
)

type CalcFunc func(Env, []types.Value) ([]types.Value, error)

type FuncDecl struct {
	Name   string
	Params []types.Type
}

func funcDecl(name string, params []types.Type) string {
	paramTypes := coll.Map(params, func(t types.Type) string { return t.String() })
	return fmt.Sprintf("%v(%v)", name, paramTypes)
}

type Library struct {
	genOps map[string]CalcFunc
}

func (l *Library) RegisterGenOp(name string, fn CalcFunc, params ...types.Type) {
	l.genOps[funcDecl(name, params)] = fn
}

func (l *Library) EvalGenOp(env Env, name string, args []types.Value) ([]types.Value, error) {
	argTypes := coll.Map(args, func(v types.Value) types.Type { return v.Type() })
	decl := funcDecl(name, argTypes)
	fn, ok := l.genOps[decl]
	if !ok {
		return nil, fmt.Errorf("no operation for %v", decl)
	}
	return fn(env, args)
}
