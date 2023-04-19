package zlib

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc"
)

var zeroBigInt big.Int

func AbsBigInt(c zc.Calc) {
	var r0 big.Int
	a0 := zc.PopBigInt(c)
	r0.Abs(a0)
	zc.PushBigInt(c, &r0)
}

func AddBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Add(a0, a1)
	zc.PushBigInt(c, &r0)
}

func AndBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.And(a0, a1)
	zc.PushBigInt(c, &r0)
}

func EqBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) == 0
	zc.PushBool(c, r0)
}

func GtBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) > 0
	zc.PushBool(c, r0)
}

func GteBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) >= 0
	zc.PushBool(c, r0)
}

func HexBigInt(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	r0 := fmt.Sprintf("0x%x", a0)
	c.Push(r0)
}

func LtBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) < 0
	zc.PushBool(c, r0)
}

func LteBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) <= 0
	zc.PushBool(c, r0)
}

func ModBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)

	if a1.Cmp(&zeroBigInt) == 0 {
		c.SetError(zc.ErrDivisionByZero)
		return
	}

	r0.Mod(a0, a1)
	zc.PushBigInt(c, &r0)
}

func MulBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Mul(a0, a1)
	zc.PushBigInt(c, &r0)
}

func NeqBigInt(c zc.Calc) {
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Cmp(a1) != 0
	zc.PushBool(c, r0)
}

func NegBigInt(c zc.Calc) {
	var r0 big.Int
	a0 := zc.PopBigInt(c)
	r0.Neg(a0)
	zc.PushBigInt(c, &r0)
}

func NotBigInt(c zc.Calc) {
	var r0 big.Int
	a0 := zc.PopBigInt(c)
	r0.Not(a0)
	zc.PushBigInt(c, &r0)
}

func OrBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Or(a0, a1)
	zc.PushBigInt(c, &r0)
}

func PowBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Exp(a0, a1, nil)
	zc.PushBigInt(c, &r0)
}

func RemBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Rem(a0, a1)
	zc.PushBigInt(c, &r0)
}

func SignBigInt(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	r0 := a0.Sign()
	zc.PushInt(c, r0)
}

func SubBigInt(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Sub(a0, a1)
	zc.PushBigInt(c, &r0)
}
