package math_

import (
	"github.com/blackchip-org/zc"
)

func Abs(env *zc.Env) error     { return env.Stack.Eval("abs", 1) }
func Add(env *zc.Env) error     { return env.Stack.Eval("add", 2) }
func Ceil(env *zc.Env) error    { return env.Stack.Eval("ceil", 1) }
func Div(env *zc.Env) error     { return env.Stack.Eval("div", 2) }
func Floor(env *zc.Env) error   { return env.Stack.Eval("floor", 1) }
func Modulus(env *zc.Env) error { return env.Stack.Eval("mod", 2) }
func Mul(env *zc.Env) error     { return env.Stack.Eval("mul", 2) }
func Neg(env *zc.Env) error     { return env.Stack.Eval("neg", 1) }
func Pow(env *zc.Env) error     { return env.Stack.Eval("pow", 2) }
func Rem(env *zc.Env) error     { return env.Stack.Eval("rem", 2) }
func Sign(env *zc.Env) error    { return env.Stack.Eval("sign", 1) }
func Sqrt(env *zc.Env) error    { return env.Stack.Eval("sqrt", 1) }
func Sub(env *zc.Env) error     { return env.Stack.Eval("sub", 2) }

func Sum(env *zc.Env) error {
	for env.Stack.Len() > 1 {
		if err := Add(env); err != nil {
			return err
		}
	}
	return nil
}
