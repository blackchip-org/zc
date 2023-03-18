package math_

import (
	"math"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func opSqrtFloat(a float64) (float64, error) { return math.Sqrt(a), nil }

func SqrtFloat(env *zc.Env) error { return funcs.EvalUnaryFloat(env, opSqrtFloat) }
