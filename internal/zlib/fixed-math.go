package zlib

import (
	"errors"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
	"github.com/shopspring/decimal"
)

func fixedSign(a decimal.Decimal) decimal.Decimal {
	return decimal.NewFromInt(int64(a.Sign()))
}

func opAbsFixed(a decimal.Decimal) (decimal.Decimal, error)                    { return a.Abs(), nil }
func opAddFixed(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Add(b), nil }
func opCeilFixed(a decimal.Decimal) (decimal.Decimal, error)                   { return a.Ceil(), nil }
func opDivFixed(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Div(b), nil }
func opFloorFixed(a decimal.Decimal) (decimal.Decimal, error)                  { return a.Floor(), nil }
func opMulFixed(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Mul(b), nil }
func opNegFixed(a decimal.Decimal) (decimal.Decimal, error)                    { return a.Neg(), nil }
func opModFixed(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Mod(b), nil }
func opPowFixed(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Pow(b), nil }
func opSignFixed(a decimal.Decimal) (decimal.Decimal, error)                   { return fixedSign(a), nil }
func opSubFixed(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Sub(b), nil }

func opRemFixed(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return decimal.Zero, errors.New("unsupported operation: fix-rem")
}

func AbsFixed(env *zc.Env) error   { return funcs.EvalUnaryFixed(env, opAbsFixed) }
func AddFixed(env *zc.Env) error   { return funcs.EvalBinaryFixed(env, opAddFixed) }
func CeilFixed(env *zc.Env) error  { return funcs.EvalUnaryFixed(env, opCeilFixed) }
func DivFixed(env *zc.Env) error   { return funcs.EvalBinaryFixed(env, opDivFixed) }
func FloorFixed(env *zc.Env) error { return funcs.EvalUnaryFixed(env, opFloorFixed) }
func ModFixed(env *zc.Env) error   { return funcs.EvalBinaryFixed(env, opModFixed) }
func MulFixed(env *zc.Env) error   { return funcs.EvalBinaryFixed(env, opMulFixed) }
func NegFixed(env *zc.Env) error   { return funcs.EvalUnaryFixed(env, opNegFixed) }
func PowFixed(env *zc.Env) error   { return funcs.EvalBinaryFixed(env, opPowFixed) }
func RemFixed(env *zc.Env) error   { return funcs.EvalBinaryFixed(env, opRemFixed) }
func SignFixed(env *zc.Env) error  { return funcs.EvalUnaryFixed(env, opSignFixed) }
func SubFixed(env *zc.Env) error   { return funcs.EvalBinaryFixed(env, opSubFixed) }
