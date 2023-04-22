package zc

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func PopBigInt(c Calc) *big.Int         { return BigInt.MustParse(c.MustPop()) }
func PopBool(c Calc) bool               { return Bool.MustParse(c.MustPop()) }
func PopComplex(c Calc) complex128      { return Complex.MustParse(c.MustPop()) }
func PopDecimal(c Calc) decimal.Decimal { return Decimal.MustParse(c.MustPop()) }
func PopInt(c Calc) int                 { return Int.MustParse(c.MustPop()) }
func PopInt64(c Calc) int64             { return Int64.MustParse(c.MustPop()) }
func PopInt32(c Calc) int32             { return Int32.MustParse(c.MustPop()) }
func PopRational(c Calc) *big.Rat       { return Rational.MustParse(c.MustPop()) }
func PopRune(c Calc) rune               { return Rune.MustParse(c.MustPop()) }
func PopString(c Calc) string           { return c.MustPop() }
func PopUint(c Calc) uint               { return Uint.MustParse(c.MustPop()) }
func PopUint8(c Calc) uint8             { return Uint8.MustParse(c.MustPop()) }

func PushBigInt(c Calc, r *big.Int)         { c.Push(BigInt.Format(r)) }
func PushBool(c Calc, r bool)               { c.Push(Bool.Format(r)) }
func PushComplex(c Calc, r complex128)      { c.Push(Complex.Format(r)) }
func PushDecimal(c Calc, r decimal.Decimal) { c.Push(Decimal.Format(r)) }
func PushInt(c Calc, r int)                 { c.Push(Int.Format(r)) }
func PushInt64(c Calc, r int64)             { c.Push(Int64.Format(r)) }
func PushInt32(c Calc, r int32)             { c.Push(Int32.Format(r)) }
func PushRational(c Calc, r *big.Rat)       { c.Push(Rational.Format(r)) }
func PushRune(c Calc, r rune)               { c.Push(Rune.Format(r)) }
func PushString(c Calc, r string)           { c.Push(r) }
func PushUint(c Calc, r uint)               { c.Push(Uint.Format(r)) }
func PushUint8(c Calc, r uint8)             { c.Push(Uint8.Format(r)) }
