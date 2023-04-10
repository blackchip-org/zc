package bool_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
)

func And(env *zc.Env) error { return env.Stack.Eval(ops.And, 2) }
func Eq(env *zc.Env) error  { return env.Stack.Eval(ops.Eq, 2) }
func Neq(env *zc.Env) error { return env.Stack.Eval(ops.Neq, 2) }
func Gt(env *zc.Env) error  { return env.Stack.Eval(ops.Gt, 2) }
func Gte(env *zc.Env) error { return env.Stack.Eval(ops.Gte, 2) }
func Lt(env *zc.Env) error  { return env.Stack.Eval(ops.Lt, 2) }
func Lte(env *zc.Env) error { return env.Stack.Eval(ops.Lte, 2) }
func Or(env *zc.Env) error  { return env.Stack.Eval(ops.Or, 2) }
func Not(env *zc.Env) error { return env.Stack.Eval(ops.Not, 1) }
