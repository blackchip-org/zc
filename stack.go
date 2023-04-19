package zc

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func PopBigInt(e Env) *big.Int         { return BigInt.MustParse(e.MustPop()) }
func PopBool(e Env) bool               { return Bool.MustParse(e.MustPop()) }
func PopComplex(e Env) complex128      { return Complex.MustParse(e.MustPop()) }
func PopDecimal(e Env) decimal.Decimal { return Decimal.MustParse(e.MustPop()) }
func PopFloat(e Env) float64           { return Float.MustParse(e.MustPop()) }
func PopInt(e Env) int                 { return Int.MustParse(e.MustPop()) }

func PushBigInt(e Env, r *big.Int)         { e.Push(BigInt.Format(r)) }
func PushBool(e Env, r bool)               { e.Push(Bool.Format(r)) }
func PushComplex(e Env, r complex128)      { e.Push(Complex.Format(r)) }
func PushDecimal(e Env, r decimal.Decimal) { e.Push(Decimal.Format(r)) }
func PushFloat(e Env, r float64)           { e.Push(Float.Format(r)) }
func PushInt(e Env, r int)                 { e.Push(Int.Format(r)) }
