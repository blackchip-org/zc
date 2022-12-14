package zlib

import (
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
	return decimal.Zero, zc.UnsupportedError{Name: "fix-rem"}
}

func AbsFixed(calc *zc.Calc) error   { return funcs.EvalUnaryFixed(calc, opAbsFixed) }
func AddFixed(calc *zc.Calc) error   { return funcs.EvalBinaryFixed(calc, opAddFixed) }
func CeilFixed(calc *zc.Calc) error  { return funcs.EvalUnaryFixed(calc, opCeilFixed) }
func DivFixed(calc *zc.Calc) error   { return funcs.EvalBinaryFixed(calc, opDivFixed) }
func FloorFixed(calc *zc.Calc) error { return funcs.EvalUnaryFixed(calc, opCeilFixed) }
func ModFixed(calc *zc.Calc) error   { return funcs.EvalBinaryFixed(calc, opModFixed) }
func MulFixed(calc *zc.Calc) error   { return funcs.EvalBinaryFixed(calc, opMulFixed) }
func NegFixed(calc *zc.Calc) error   { return funcs.EvalUnaryFixed(calc, opNegFixed) }
func PowFixed(calc *zc.Calc) error   { return funcs.EvalBinaryFixed(calc, opPowFixed) }
func RemFixed(calc *zc.Calc) error   { return funcs.EvalBinaryFixed(calc, opRemFixed) }
func SignFixed(calc *zc.Calc) error  { return funcs.EvalUnaryFixed(calc, opSignFixed) }
func SubFixed(calc *zc.Calc) error   { return funcs.EvalBinaryFixed(calc, opSubFixed) }
