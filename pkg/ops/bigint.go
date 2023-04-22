package ops

import (
	"fmt"
	"math"
	"math/big"

	"github.com/blackchip-org/zc/pkg/zc"
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

func Bin(c zc.Calc) {
	var zero big.Int
	var r0 string
	a0 := zc.PopBigInt(c)
	if a0.Cmp(&zero) < 0 {
		var t0 big.Int
		t0.Abs(a0)
		r0 = fmt.Sprintf("-0b%b", &t0)
	} else {
		r0 = fmt.Sprintf("0b%b", a0)
	}
	zc.PushString(c, r0)
}

func Bit(c zc.Calc) {
	i := zc.PopInt(c)
	a0 := zc.PopBigInt(c)
	r0 := a0.Bit(i)
	zc.PushUint(c, r0)
}

func Bits(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	r0 := a0.BitLen()
	zc.PushInt(c, r0)
}

func Bytes(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	r0 := int(math.Ceil(float64(a0.BitLen()) / 8.0))
	zc.PushInt(c, r0)
}

func Dec(c zc.Calc) {
	a0 := zc.PopBigInt(c)
	zc.PushBigInt(c, a0)
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
	var zero big.Int
	var r0 string
	a0 := zc.PopBigInt(c)
	if a0.Cmp(&zero) < 0 {
		var t0 big.Int
		t0.Abs(a0)
		r0 = fmt.Sprintf("-0x%x", &t0)
	} else {
		r0 = fmt.Sprintf("0x%x", a0)
	}
	zc.PushString(c, r0)
}

func Lsh(c zc.Calc) {
	var r0 big.Int
	n := zc.PopUint(c)
	a0 := zc.PopBigInt(c)
	r0.Lsh(a0, n)
	zc.PushBigInt(c, &r0)
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
		zc.ErrDivisionByZero(c)
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

func Oct(c zc.Calc) {
	var zero big.Int
	var r0 string
	a0 := zc.PopBigInt(c)
	if a0.Cmp(&zero) < 0 {
		var t0 big.Int
		t0.Abs(a0)
		r0 = fmt.Sprintf("-0o%o", &t0)
	} else {
		r0 = fmt.Sprintf("0o%o", a0)
	}
	zc.PushString(c, r0)
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

func Rsh(c zc.Calc) {
	var r0 big.Int
	n := zc.PopUint(c)
	a0 := zc.PopBigInt(c)
	r0.Rsh(a0, n)
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

func Xor(c zc.Calc) {
	var r0 big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	r0.Xor(a0, a1)
	zc.PushBigInt(c, &r0)
}
