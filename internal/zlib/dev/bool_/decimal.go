package bool_

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
	"github.com/shopspring/decimal"
)

func opEqDecimal(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) == 0, nil }
func opNeqDecimal(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) != 0, nil }
func opGtDecimal(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) > 0, nil }
func opGteDecimal(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) >= 0, nil }
func opLtDecimal(a decimal.Decimal, b decimal.Decimal) (bool, error)  { return a.Cmp(b) < 0, nil }
func opLteDecimal(a decimal.Decimal, b decimal.Decimal) (bool, error) { return a.Cmp(b) <= 0, nil }

func EqDecimal(env *zc.Env) error  { return funcs.EvalCompareDecimal(env, opEqDecimal) }
func GtDecimal(env *zc.Env) error  { return funcs.EvalCompareDecimal(env, opGtDecimal) }
func GteDecimal(env *zc.Env) error { return funcs.EvalCompareDecimal(env, opGteDecimal) }
func NeqDecimal(env *zc.Env) error { return funcs.EvalCompareDecimal(env, opNeqDecimal) }
func LtDecimal(env *zc.Env) error  { return funcs.EvalCompareDecimal(env, opLtDecimal) }
func LteDecimal(env *zc.Env) error { return funcs.EvalCompareDecimal(env, opLteDecimal) }
