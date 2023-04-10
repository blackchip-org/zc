package sci

import (
	"math"
	"strconv"

	"github.com/blackchip-org/zc"
)

var (
	Acos  = zc.FuncFloat1(math.Acos)
	Acosh = zc.FuncFloat1(math.Acosh)
	Asin  = zc.FuncFloat1(math.Asin)
	Asinh = zc.FuncFloat1(math.Asinh)
	Atan  = zc.FuncFloat1(math.Atan)
	Atanh = zc.FuncFloat1(math.Atanh)
	Cos   = zc.FuncFloat1(math.Cos)
	Cosh  = zc.FuncFloat1(math.Cosh)
	Exp   = zc.FuncFloat1(math.Exp)
	Log   = zc.FuncFloat1(math.Log)
	Log10 = zc.FuncFloat1(math.Log10)
	Log2  = zc.FuncFloat1(math.Log2)
	Sin   = zc.FuncFloat1(math.Sin)
	Sinh  = zc.FuncFloat1(math.Sinh)
	Tan   = zc.FuncFloat1(math.Tan)
	Tanh  = zc.FuncFloat1(math.Tanh)
)

func ScientificNotation(env *zc.Env) error {
	f, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	env.Stack.Push(strconv.FormatFloat(f, 'e', -1, 64))
	return nil
}
