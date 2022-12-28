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

func EqFixed(env *zc.Env) error  { return funcs.EvalCompareFixed(env, opEqFixed) }
func GtFixed(env *zc.Env) error  { return funcs.EvalCompareFixed(env, opGtFixed) }
func GteFixed(env *zc.Env) error { return funcs.EvalCompareFixed(env, opGteFixed) }
func NeqFixed(env *zc.Env) error { return funcs.EvalCompareFixed(env, opNeqFixed) }
func LtFixed(env *zc.Env) error  { return funcs.EvalCompareFixed(env, opLtFixed) }
func LteFixed(env *zc.Env) error { return funcs.EvalCompareFixed(env, opLteFixed) }
