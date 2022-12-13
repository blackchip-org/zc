package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
	"github.com/shopspring/decimal"
)

func opEqFix(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) == 0, nil }
func opNeqFix(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) != 0, nil }
func opGtFix(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) > 0, nil }
func opGteFix(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) >= 0, nil }
func opLtFix(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) < 0, nil }
func opLteFix(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) <= 0, nil }

func EqFix(calc *zc.Calc) error  { return funcs.EvalCompareFix(calc, opEqFix) }
func GtFix(calc *zc.Calc) error  { return funcs.EvalCompareFix(calc, opGtFix) }
func GteFix(calc *zc.Calc) error { return funcs.EvalCompareFix(calc, opGteFix) }
func NeqFix(calc *zc.Calc) error { return funcs.EvalCompareFix(calc, opNeqFix) }
func LtFix(calc *zc.Calc) error  { return funcs.EvalCompareFix(calc, opLtFix) }
func LteFix(calc *zc.Calc) error { return funcs.EvalCompareFix(calc, opLteFix) }
