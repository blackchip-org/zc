package zlib

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc"
)

var zeroBigInt big.Int

func AbsBigInt(e zc.Env) {
	var r0 big.Int
	a0 := zc.PopBigInt(e)
	r0.Abs(a0)
	zc.PushBigInt(e, &r0)
}

func AddBigInt(e zc.Env) {
	var r0 big.Int
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0.Add(a0, a1)
	zc.PushBigInt(e, &r0)
}

func AndBigInt(e zc.Env) {
	var r0 big.Int
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0.And(a0, a1)
	zc.PushBigInt(e, &r0)
}

func EqBigInt(e zc.Env) {
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(e, r0)
}

func GtBigInt(e zc.Env) {
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(e, r0)
}

func GteBigInt(e zc.Env) {
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(e, r0)
}

func HexBigInt(e zc.Env) {
	a0 := zc.PopBigInt(e)
	r0 := fmt.Sprintf("0x%x", a0)
	e.Push(r0)
}

func LtBigInt(e zc.Env) {
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(e, r0)
}

func LteBigInt(e zc.Env) {
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(e, r0)
}

func ModBigInt(e zc.Env) {
	var r0 big.Int
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)

	if a1.Cmp(&zeroBigInt) == 0 {
		e.Error(zc.ErrDivisionByZero)
		return
	}

	r0.Mod(a0, a1)
	zc.PushBigInt(e, &r0)
}

func MulBigInt(e zc.Env) {
	var r0 big.Int
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0.Mul(a0, a1)
	zc.PushBigInt(e, &r0)
}

func NeqBigInt(e zc.Env) {
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0 := a0.Cmp(a1) != 0
	zc.PushBool(e, r0)
}

func NegBigInt(e zc.Env) {
	var r0 big.Int
	a0 := zc.PopBigInt(e)
	r0.Neg(a0)
	zc.PushBigInt(e, &r0)
}

func NotBigInt(e zc.Env) {
	var r0 big.Int
	a0 := zc.PopBigInt(e)
	r0.Not(a0)
	zc.PushBigInt(e, &r0)
}

func OrBigInt(e zc.Env) {
	var r0 big.Int
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0.Or(a0, a1)
	zc.PushBigInt(e, &r0)
}

func PowBigInt(e zc.Env) {
	var r0 big.Int
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0.Exp(a0, a1, nil)
	zc.PushBigInt(e, &r0)
}

func RemBigInt(e zc.Env) {
	var r0 big.Int
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0.Rem(a0, a1)
	zc.PushBigInt(e, &r0)
}

func SignBigInt(e zc.Env) {
	a0 := zc.PopBigInt(e)
	r0 := a0.Sign()
	zc.PushInt(e, r0)
}

func SubBigInt(e zc.Env) {
	var r0 big.Int
	a1 := zc.PopBigInt(e)
	a0 := zc.PopBigInt(e)
	r0.Sub(a0, a1)
	zc.PushBigInt(e, &r0)
}
