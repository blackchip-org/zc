package zlib

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func opEqString(a string, b string) (bool, error)  { return a == b, nil }
func opNeqString(a string, b string) (bool, error) { return a != b, nil }
func opGtString(a string, b string) (bool, error)  { return a > b, nil }
func opGteString(a string, b string) (bool, error) { return a >= b, nil }
func opLtString(a string, b string) (bool, error)  { return a < b, nil }
func opLteString(a string, b string) (bool, error) { return a <= b, nil }

func EqStr(env *zc.Env) error  { return funcs.EvalCompareStr(env, opEqString) }
func GtStr(env *zc.Env) error  { return funcs.EvalCompareStr(env, opGtString) }
func GteStr(env *zc.Env) error { return funcs.EvalCompareStr(env, opGteString) }
func NeqStr(env *zc.Env) error { return funcs.EvalCompareStr(env, opNeqString) }
func LtStr(env *zc.Env) error  { return funcs.EvalCompareStr(env, opLtString) }
func LteStr(env *zc.Env) error { return funcs.EvalCompareStr(env, opLteString) }
