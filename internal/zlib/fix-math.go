package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
	"github.com/shopspring/decimal"
)

func fixSign(a decimal.Decimal) decimal.Decimal {
	return decimal.NewFromInt(int64(a.Sign()))
}

func opAbsFix(a decimal.Decimal) (decimal.Decimal, error)                    { return a.Abs(), nil }
func opAddFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Add(b), nil }
func opCeilFix(a decimal.Decimal) (decimal.Decimal, error)                   { return a.Ceil(), nil }
func opDivFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Div(b), nil }
func opFloorFix(a decimal.Decimal) (decimal.Decimal, error)                  { return a.Floor(), nil }
func opMulFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Mul(b), nil }
func opNegFix(a decimal.Decimal) (decimal.Decimal, error)                    { return a.Neg(), nil }
func opModFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Mod(b), nil }
func opPowFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Pow(b), nil }
func opSignFix(a decimal.Decimal) (decimal.Decimal, error)                   { return fixSign(a), nil }
func opSubFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) { return a.Sub(b), nil }

func opRemFix(a decimal.Decimal, b decimal.Decimal) (decimal.Decimal, error) {
	return decimal.Zero, zc.ErrUnsupported("fix-rem")
}

func AbsFix(calc *zc.Calc) error   { return funcs.EvalUnaryFix(calc, opAbsFix) }
func AddFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opAddFix) }
func CeilFix(calc *zc.Calc) error  { return funcs.EvalUnaryFix(calc, opCeilFix) }
func DivFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opDivFix) }
func FloorFix(calc *zc.Calc) error { return funcs.EvalUnaryFix(calc, opCeilFix) }
func ModFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opModFix) }
func MulFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opMulFix) }
func NegFix(calc *zc.Calc) error   { return funcs.EvalUnaryFix(calc, opNegFix) }
func PowFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opPowFix) }
func RemFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opRemFix) }
func SignFix(calc *zc.Calc) error  { return funcs.EvalUnaryFix(calc, opSignFix) }
func SubFix(calc *zc.Calc) error   { return funcs.EvalBinaryFix(calc, opSubFix) }
