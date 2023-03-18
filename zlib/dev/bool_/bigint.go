package bool_

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

func EqBigInt(env *zc.Env) error  { return funcs.EvalCompareBigInt(env, opEqBigInt) }
func GtBigInt(env *zc.Env) error  { return funcs.EvalCompareBigInt(env, opGtBigInt) }
func GteBigInt(env *zc.Env) error { return funcs.EvalCompareBigInt(env, opGteBigInt) }
func NeqBigInt(env *zc.Env) error { return funcs.EvalCompareBigInt(env, opNeqBigInt) }
func LtBigInt(env *zc.Env) error  { return funcs.EvalCompareBigInt(env, opLtBigInt) }
func LteBigInt(env *zc.Env) error { return funcs.EvalCompareBigInt(env, opLteBigInt) }
