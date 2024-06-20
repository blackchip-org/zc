package funcs

import (
	"fmt"
	"math"
	"math/big"

	"github.com/blackchip-org/zc/v6"
)

func AndBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.And(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}

func Bin(e *zc.OpEnv) {
	var zero big.Int
	var b string

	x := zc.BigInt.Pop(e)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0b%b", x)
	} else {
		b = fmt.Sprintf("0b%b", x)
	}
	zc.String.Push(e, b)
	zc.BigInt.Recycle(x)
}

func Bit(e *zc.OpEnv) {
	i := zc.Int.Pop(e)
	if i < 0 {
		e.Err = zc.ErrInvalidArg(e, "%v < 0", i)
		return
	}
	x := zc.BigInt.Pop(e)
	b := x.Bit(i)
	zc.Uint.Push(e, b)
	zc.BigInt.Recycle(x)
}

func Bits(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	b := x.BitLen()
	zc.Int.Push(e, b)
	zc.BigInt.Recycle(x)
}

func Bytes(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	b := int(math.Ceil(float64(x.BitLen()) / 8.0))
	zc.Int.Push(e, b)
	zc.BigInt.Recycle(x)
}

func DecBigInt(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	zc.BigInt.Push(e, x)
}

func Hex(e *zc.OpEnv) {
	var zero big.Int
	var b string

	x := zc.BigInt.Pop(e)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0x%x", x)
	} else {
		b = fmt.Sprintf("0x%x", x)
	}
	zc.String.Push(e, b)
	zc.BigInt.Recycle(x)
}

func Lsh(e *zc.OpEnv) {
	n := zc.Uint.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Lsh(x, n)
	zc.BigInt.Push(e, x)
}

func NotBigInt(e *zc.OpEnv) {
	x := zc.BigInt.Pop(e)
	x.Not(x)
	zc.BigInt.Push(e, x)
}

func Oct(e *zc.OpEnv) {
	var zero big.Int
	var b string

	x := zc.BigInt.Pop(e)
	if x.Cmp(&zero) < 0 {
		x.Abs(x)
		b = fmt.Sprintf("-0o%o", x)
	} else {
		b = fmt.Sprintf("0o%o", x)
	}
	zc.String.Push(e, b)
	zc.BigInt.Recycle(x)
}

func OrBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Or(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}

func Rsh(e *zc.OpEnv) {
	n := zc.Uint.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Rsh(x, n)
	zc.BigInt.Push(e, x)
}

func XorBigInt(e *zc.OpEnv) {
	y := zc.BigInt.Pop(e)
	x := zc.BigInt.Pop(e)
	x.Xor(x, y)
	zc.BigInt.Push(e, x)
	zc.BigInt.Recycle(y)
}
