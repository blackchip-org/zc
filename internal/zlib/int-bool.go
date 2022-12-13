package zlib

import (
	"math/big"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/funcs"
)

func opEqBigInt(a *big.Int, b *big.Int) (bool, error)  { return a.Cmp(b) == 0, nil }
func opNeqBigInt(a *big.Int, b *big.Int) (bool, error) { return a.Cmp(b) != 0, nil }
func opGtBigInt(a *big.Int, b *big.Int) (bool, error)  { return a.Cmp(b) > 0, nil }
func opGteBigInt(a *big.Int, b *big.Int) (bool, error) { return a.Cmp(b) >= 0, nil }
func opLtBigInt(a *big.Int, b *big.Int) (bool, error)  { return a.Cmp(b) < 0, nil }
func opLteBigInt(a *big.Int, b *big.Int) (bool, error) { return a.Cmp(b) <= 0, nil }

func EqBigInt(calc *zc.Calc) error  { return funcs.EvalCompareBigInt(calc, opEqBigInt) }
func GtBigInt(calc *zc.Calc) error  { return funcs.EvalCompareBigInt(calc, opGtBigInt) }
func GteBigInt(calc *zc.Calc) error { return funcs.EvalCompareBigInt(calc, opGteBigInt) }
func NeqBigInt(calc *zc.Calc) error { return funcs.EvalCompareBigInt(calc, opNeqBigInt) }
func LtBigInt(calc *zc.Calc) error  { return funcs.EvalCompareBigInt(calc, opLtBigInt) }
func LteBigInt(calc *zc.Calc) error { return funcs.EvalCompareBigInt(calc, opLteBigInt) }
