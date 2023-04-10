package math_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/ops"
)

func Abs(env *zc.Env) error     { return env.Stack.Eval(ops.Abs) }
func Add(env *zc.Env) error     { return env.Stack.Eval(ops.Add) }
func Ceil(env *zc.Env) error    { return env.Stack.Eval(ops.Ceil) }
func Div(env *zc.Env) error     { return env.Stack.Eval(ops.Div) }
func Floor(env *zc.Env) error   { return env.Stack.Eval(ops.Floor) }
func Modulus(env *zc.Env) error { return env.Stack.Eval(ops.Mod) }
func Mul(env *zc.Env) error     { return env.Stack.Eval(ops.Mul) }
func Neg(env *zc.Env) error     { return env.Stack.Eval(ops.Neg) }
func Pow(env *zc.Env) error     { return env.Stack.Eval(ops.Pow) }
func Rem(env *zc.Env) error     { return env.Stack.Eval(ops.Rem) }
func Sign(env *zc.Env) error    { return env.Stack.Eval(ops.Sign) }
func Sqrt(env *zc.Env) error    { return env.Stack.Eval(ops.Sqrt) }
func Sub(env *zc.Env) error     { return env.Stack.Eval(ops.Sub) }

func Sum(env *zc.Env) error {
	for env.Stack.Len() > 1 {
		if err := Add(env); err != nil {
			return err
		}
	}
	return nil
}
