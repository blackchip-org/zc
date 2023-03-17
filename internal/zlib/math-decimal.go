package zlib

import (
	"errors"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
	"github.com/shopspring/decimal"
)

func DecimalSign(a decimal.Decimal) decimal.Decimal {
	return decimal.NewFromInt(int64(a.Sign()))
}

func opAbsDecimal(a decimal.Decimal) (decimal.Decimal, error) { return a.Abs(), nil }
func opAddDecimal(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return a.Add(b), nil
}
func opCeilDecimal(a decimal.Decimal) (decimal.Decimal, error) { return a.Ceil(), nil }
func opDivDecimal(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return a.Div(b), nil
}
func opFloorDecimal(a decimal.Decimal) (decimal.Decimal, error) { return a.Floor(), nil }
func opMulDecimal(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return a.Mul(b), nil
}
func opNegDecimal(a decimal.Decimal) (decimal.Decimal, error) { return a.Neg(), nil }
func opModDecimal(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return a.Mod(b), nil
}
func opPowDecimal(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return a.Pow(b), nil
}
func opSignDecimal(a decimal.Decimal) (decimal.Decimal, error) { return DecimalSign(a), nil }
func opSubDecimal(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return a.Sub(b), nil
}

func opRemDecimal(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return decimal.Zero, errors.New("unsupported operation: fix-rem")
}

func AbsDecimal(env *zc.Env) error   { return funcs.EvalUnaryDecimal(env, opAbsDecimal) }
func AddDecimal(env *zc.Env) error   { return funcs.EvalBinaryDecimal(env, opAddDecimal) }
func CeilDecimal(env *zc.Env) error  { return funcs.EvalUnaryDecimal(env, opCeilDecimal) }
func DivDecimal(env *zc.Env) error   { return funcs.EvalBinaryDecimal(env, opDivDecimal) }
func FloorDecimal(env *zc.Env) error { return funcs.EvalUnaryDecimal(env, opFloorDecimal) }
func ModDecimal(env *zc.Env) error   { return funcs.EvalBinaryDecimal(env, opModDecimal) }
func MulDecimal(env *zc.Env) error   { return funcs.EvalBinaryDecimal(env, opMulDecimal) }
func NegDecimal(env *zc.Env) error   { return funcs.EvalUnaryDecimal(env, opNegDecimal) }
func PowDecimal(env *zc.Env) error   { return funcs.EvalBinaryDecimal(env, opPowDecimal) }
func RemDecimal(env *zc.Env) error   { return funcs.EvalBinaryDecimal(env, opRemDecimal) }
func SignDecimal(env *zc.Env) error  { return funcs.EvalUnaryDecimal(env, opSignDecimal) }
func SubDecimal(env *zc.Env) error   { return funcs.EvalBinaryDecimal(env, opSubDecimal) }
