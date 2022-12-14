package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
	"github.com/shopspring/decimal"
)

func opEqFixed(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) == 0, nil }
func opNeqFixed(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) != 0, nil }
func opGtFixed(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) > 0, nil }
func opGteFixed(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) >= 0, nil }
func opLtFixed(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) < 0, nil }
func opLteFixed(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) <= 0, nil }

func EqFixed(calc *zc.Calc) error  { return funcs.EvalCompareFixed(calc, opEqFixed) }
func GtFixed(calc *zc.Calc) error  { return funcs.EvalCompareFixed(calc, opGtFixed) }
func GteFixed(calc *zc.Calc) error { return funcs.EvalCompareFixed(calc, opGteFixed) }
func NeqFixed(calc *zc.Calc) error { return funcs.EvalCompareFixed(calc, opNeqFixed) }
func LtFixed(calc *zc.Calc) error  { return funcs.EvalCompareFixed(calc, opLtFixed) }
func LteFixed(calc *zc.Calc) error { return funcs.EvalCompareFixed(calc, opLteFixed) }
