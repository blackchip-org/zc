package zlib

import (
	"math"
	"strconv"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func Acos(env *zc.Env) error  { return funcs.EvalUnaryFloatSafe(env, math.Acos) }
func Acosh(env *zc.Env) error { return funcs.EvalUnaryFloatSafe(env, math.Acosh) }
func Asin(env *zc.Env) error  { return funcs.EvalUnaryFloatSafe(env, math.Asin) }
func Asinh(env *zc.Env) error { return funcs.EvalUnaryFloatSafe(env, math.Asinh) }
func Atan(env *zc.Env) error  { return funcs.EvalUnaryFloatSafe(env, math.Atan) }
func Atanh(env *zc.Env) error { return funcs.EvalUnaryFloatSafe(env, math.Atanh) }
func Cos(env *zc.Env) error   { return funcs.EvalUnaryFloatSafe(env, math.Cos) }
func Cosh(env *zc.Env) error  { return funcs.EvalUnaryFloatSafe(env, math.Cosh) }
func Exp(env *zc.Env) error   { return funcs.EvalUnaryFloatSafe(env, math.Exp) }
func Log(env *zc.Env) error   { return funcs.EvalUnaryFloatSafe(env, math.Log) }
func Log10(env *zc.Env) error { return funcs.EvalUnaryFloatSafe(env, math.Log10) }
func Log2(env *zc.Env) error  { return funcs.EvalUnaryFloatSafe(env, math.Log2) }
func Sin(env *zc.Env) error   { return funcs.EvalUnaryFloatSafe(env, math.Sin) }
func Sinh(env *zc.Env) error  { return funcs.EvalUnaryFloatSafe(env, math.Sinh) }
func Tan(env *zc.Env) error   { return funcs.EvalUnaryFloatSafe(env, math.Tan) }
func Tanh(env *zc.Env) error  { return funcs.EvalUnaryFloatSafe(env, math.Tanh) }

func ScientificNotation(env *zc.Env) error {
	f, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	env.Stack.Push(strconv.FormatFloat(f, 'e', -1, 64))
	return nil
}
