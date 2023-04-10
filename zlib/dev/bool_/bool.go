package bool_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
)

func And(env *zc.Env) error { return env.Stack.Eval(ops.And) }
func Eq(env *zc.Env) error  { return env.Stack.Eval(ops.Eq) }
func Neq(env *zc.Env) error { return env.Stack.Eval(ops.Neq) }
func Gt(env *zc.Env) error  { return env.Stack.Eval(ops.Gt) }
func Gte(env *zc.Env) error { return env.Stack.Eval(ops.Gte) }
func Lt(env *zc.Env) error  { return env.Stack.Eval(ops.Lt) }
func Lte(env *zc.Env) error { return env.Stack.Eval(ops.Lte) }
func Or(env *zc.Env) error  { return env.Stack.Eval(ops.Or) }
func Not(env *zc.Env) error { return env.Stack.Eval(ops.Not) }
