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

func EqStr(calc *zc.Calc) error  { return funcs.EvalCompareStr(calc, opEqString) }
func GtStr(calc *zc.Calc) error  { return funcs.EvalCompareStr(calc, opGtString) }
func GteStr(calc *zc.Calc) error { return funcs.EvalCompareStr(calc, opGteString) }
func NeqStr(calc *zc.Calc) error { return funcs.EvalCompareStr(calc, opNeqString) }
func LtStr(calc *zc.Calc) error  { return funcs.EvalCompareStr(calc, opLtString) }
func LteStr(calc *zc.Calc) error { return funcs.EvalCompareStr(calc, opLteString) }
