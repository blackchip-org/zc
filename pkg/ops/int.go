package ops

import (
	"math"
	"math/big"

	"github.com/blackchip-org/zc/v5/pkg/zc"
)

func AddBigInt(c zc.Calc) {
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)
	x.Add(x, y)
	zc.PushBigInt(c, x)
}

/*
oper	add-ia
func	AddIntArch 	Int Int -- Int
title	add-id
*/
func AddIntArch(c zc.Calc) {
	y := zc.PopInt(c)
	x := zc.PopInt(c)
	zc.PushInt(c, x+y)
}

/*
oper	add-i64
func	AddInt64 	Int64 Int64 -- Int64
title	add-i64
*/
func AddInt64(c zc.Calc) {
	y := zc.PopInt64(c)
	x := zc.PopInt64(c)
	zc.PushInt64(c, x+y)
}

/*
oper	add-i32
func	AddInt32 	Int32 Int32 -- Int32
title	add-i32
*/
func AddInt32(c zc.Calc) {
	y := zc.PopInt32(c)
	x := zc.PopInt32(c)
	zc.PushInt32(c, x+y)
}

/*
oper	add-i16
func	AddInt16 	Int16 Int16 -- Int16
title	add-i16
*/
func AddInt16(c zc.Calc) {
	y := zc.PopInt16(c)
	x := zc.PopInt16(c)
	zc.PushInt16(c, x+y)
}

/*
oper	add-i8
func	AddInt8 	Int8 Int8 -- Int18
title	add-i8
*/
func AddInt8(c zc.Calc) {
	y := zc.PopInt8(c)
	x := zc.PopInt8(c)
	zc.PushInt8(c, x+y)
}

/*
oper	add-ua
func	AddUintArch 	Uint Uint -- Uint
title	add-ua
*/
func AddUintArch(c zc.Calc) {
	y := zc.PopUint(c)
	x := zc.PopUint(c)
	zc.PushUint(c, x+y)
}

/*
oper	add-u64
func	AddUint64 	Uint64 Uint64 -- Uint64
title	add-u64
*/
func AddUint64(c zc.Calc) {
	y := zc.PopUint64(c)
	x := zc.PopUint64(c)
	zc.PushUint64(c, x+y)
}

/*
oper	add-u32
func	AddUint32 	Uint32 Uint32 -- Uint32
title	add-u32
*/
func AddUint32(c zc.Calc) {
	y := zc.PopUint32(c)
	x := zc.PopUint32(c)
	zc.PushUint32(c, x+y)
}

/*
oper	add-u16
func	AddUint16 	Uint16 Uint16 -- Uint16
title	add-u16
*/
func AddUint16(c zc.Calc) {
	y := zc.PopUint16(c)
	x := zc.PopUint16(c)
	zc.PushUint16(c, x+y)
}

/*
oper	add-u8
func	AddUint8 	Uint8 Uint8 -- Uint18
title	add-u8
*/
func AddUint8(c zc.Calc) {
	y := zc.PopUint8(c)
	x := zc.PopUint8(c)
	zc.PushUint8(c, x+y)
}

/*
oper 		div-int
func		DivBigInt 	BigInt BigInt -- BigInt
alias		d-int
title		Euclidean Division
*/
func DivBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)

	if y.Cmp(&zero) == 0 {
		zc.ErrDivisionByZero(c)
		return
	}
	x.Div(x, y)
	zc.PushBigInt(c, x)
}

/*
oper 		div-mod-int
func		DivModBigInt 	BigInt BigInt -- BigInt BigInt
alias		dm-int
title		Euclidean Division and modulus
*/
func DivModBigInt(c zc.Calc) {
	var m, zero big.Int
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)

	if y.Cmp(&zero) == 0 {
		zc.ErrDivisionByZero(c)
		return
	}
	x.DivMod(x, y, &m)
	zc.PushBigInt(c, x)
	zc.PushBigInt(c, &m)
	zc.Annotate(c, "modulus")
}

/*
oper	int?
func	IsInt		Val -- Bool
title 	int?
*/
func IsInt(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.BigInt.Is(x))
}

/*
oper	ia?
func	IsIntArch	Val -- Bool
title 	ia?
*/
func IsIntArch(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Int.Is(x))
}

/*
oper	i64?
func	IsInt64		Val -- Bool
title 	i64?
*/
func IsInt64(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Int64.Is(x))
}

/*
oper	i32?
func	IsInt32		Val -- Bool
title 	i32?
*/
func IsInt32(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Int32.Is(x))
}

/*
oper	i16?
func	IsInt16		Val -- Bool
title 	i16?
*/
func IsInt16(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Int16.Is(x))
}

/*
oper	i8?
func	IsInt8		Val -- Bool
title 	i8?
*/
func IsInt8(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Int8.Is(x))
}

/*
oper	ua?
func	IsUintArch	Val -- Bool
title 	ia?
*/
func IsUintArch(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Uint.Is(x))
}

/*
oper	u64?
func	IsUint64		Val -- Bool
title 	u64?
*/
func IsUint64(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Uint64.Is(x))
}

/*
oper	u32?
func	IsUint32		Val -- Bool
title 	u32?
*/
func IsUint32(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Uint32.Is(x))
}

/*
oper	u16?
func	IsUint16		Val -- Bool
title 	u16?
*/
func IsUint16(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Uint16.Is(x))
}

/*
oper	u8?
func	IsUint8		Val -- Bool
title 	u8?
*/
func IsUint8(c zc.Calc) {
	x := zc.PopString(c)
	zc.PushBool(c, zc.Uint8.Is(x))
}

/*
oper	ia-max
func	MaxIntArch -- Int
title	ia-max
*/
func MaxIntArch(c zc.Calc) {
	zc.PushInt(c, math.MaxInt)
}

/*
oper	i64-max
macro	9223372036854775807
title	i64-max
*/

/*
oper	i32-max
macro	2147483647
title	i32-max
*/

/*
oper	i16-max
macro	32767
title	i16-max
*/

/*
oper	i8-max
macro	127
title	i8-max
*/

/*
oper	ia-min
func	MinIntArch -- Int
title	ia-min
*/
func MinIntArch(c zc.Calc) {
	zc.PushInt(c, math.MinInt)
}

/*
oper	i64-min
macro	-9223372036854775808
title	i64-min
*/

/*
oper	i32-min
macro	-2147483648
title	i32-min
*/

/*
oper	i16-min
macro	-32768
title	i16-min
*/

/*
oper	i8-min
macro	-128
title	i8-min
*/

/*
oper	ua-max
func	MaxUintArch -- UInt
title	ua-max
*/
func MaxUintArch(c zc.Calc) {
	zc.PushUint(c, math.MaxUint)
}

/*
oper	u64-max
macro	18446744073709551615
title	i64-max
*/

/*
oper	u32-max
macro	4294967295
title	i32-max
*/

/*
oper	u16-max
macro	65535
title	i16-max
*/

/*
oper	u8-max
macro	255
title	u8-max
*/

/*
oper	int
func	IntBigInt	BigInt -- BigInt
func	IntBigFloat	BigFloat -- BigInt
func	IntDecimal	Decimal -- BigInt
func 	IntRational Rational -- BigInt
title	 int
*/
func IntBigInt(c zc.Calc) {
	x := zc.PopBigInt(c)
	zc.PushBigInt(c, x)
}

func IntBigFloat(c zc.Calc) {
	var r big.Int
	x := zc.PopBigFloat(c)
	x.Int(&r)
	zc.PushBigInt(c, &r)

}

func IntDecimal(c zc.Calc) {
	x := zc.PopDecimal(c)
	zc.PushBigInt(c, x.BigInt())
}

func IntRational(c zc.Calc) {
	x := zc.PopRational(c)
	n := x.Num()
	d := x.Denom()
	n.Quo(n, d)
	zc.PushBigInt(c, n)
}

func ModBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)

	if y.Cmp(&zero) == 0 {
		zc.ErrDivisionByZero(c)
		return
	}

	x.Mod(x, y)
	zc.PushBigInt(c, x)
}

func MulBigInt(c zc.Calc) {
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)
	x.Mul(x, y)
	zc.PushBigInt(c, x)
}

/*
oper	mul-ia
func	MulIntArch 	Int Int -- Int
title	mul-ia
*/
func MulIntArch(c zc.Calc) {
	y := zc.PopInt(c)
	x := zc.PopInt(c)
	zc.PushInt(c, x*y)
}

/*
oper	mul-i64
func	MulInt64 	Int64 Int64 -- Int64
title	mul-i64
*/
func MulInt64(c zc.Calc) {
	y := zc.PopInt64(c)
	x := zc.PopInt64(c)
	zc.PushInt64(c, x*y)
}

/*
oper	mul-i32
func	MulInt32 	Int32 Int32 -- Int32
title	mul-i32
*/
func MulInt32(c zc.Calc) {
	y := zc.PopInt32(c)
	x := zc.PopInt32(c)
	zc.PushInt32(c, x*y)
}

/*
oper	mul-i16
func	MulInt16 	Int16 Int16 -- Int16
title	mul-i16
*/
func MulInt16(c zc.Calc) {
	y := zc.PopInt16(c)
	x := zc.PopInt16(c)
	zc.PushInt16(c, x*y)
}

/*
oper	mul-i8
func	MulInt8 	Int8 Int8 -- Int18
title	mul-i8
*/
func MulInt8(c zc.Calc) {
	y := zc.PopInt8(c)
	x := zc.PopInt8(c)
	zc.PushInt8(c, x*y)
}

/*
oper	mul-ua
func	MulUintArch 	Uint Uint -- Uint
title	mul-ua
*/
func MulUintArch(c zc.Calc) {
	y := zc.PopUint(c)
	x := zc.PopUint(c)
	zc.PushUint(c, x*y)
}

/*
oper	mul-u64
func	MulUint64 	Uint64 Uint64 -- Uint64
title	mul-u64
*/
func MulUint64(c zc.Calc) {
	y := zc.PopUint64(c)
	x := zc.PopUint64(c)
	zc.PushUint64(c, x*y)
}

/*
oper	mul-u32
func	MulUint32 	Uint32 Uint32 -- Uint32
title	mul-u32
*/
func MulUint32(c zc.Calc) {
	y := zc.PopUint32(c)
	x := zc.PopUint32(c)
	zc.PushUint32(c, x*y)
}

/*
oper	mul-u16
func	MulUint16 	Uint16 Uint16 -- Uint16
title	mul-u16
*/
func MulUint16(c zc.Calc) {
	y := zc.PopUint16(c)
	x := zc.PopUint16(c)
	zc.PushUint16(c, x*y)
}

/*
oper	mul-u8
func	MulUint8 	Uint8 Uint8 -- Uint18
title	mul-u8
*/
func MulUint8(c zc.Calc) {
	y := zc.PopUint8(c)
	x := zc.PopUint8(c)
	zc.PushUint8(c, x*y)
}

func NegBigInt(c zc.Calc) {
	x := zc.PopBigInt(c)
	x.Neg(x)
	zc.PushBigInt(c, x)
}

func PowBigInt(c zc.Calc) {
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)
	x.Exp(x, y, nil)
	zc.PushBigInt(c, x)
}

func RemBigInt(c zc.Calc) {
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)
	x.Rem(x, y)
	zc.PushBigInt(c, x)
}

func SignBigInt(c zc.Calc) {
	x := zc.PopBigInt(c)
	r := x.Sign()
	zc.PushInt(c, r)
}

/*
oper		sqrt-int
func 		SqrtBigInt		BigInt -- BigInt
title		sqrt-int
*/
func SqrtBigInt(c zc.Calc) {
	x := zc.PopBigInt(c)
	x.Sqrt(x)
	zc.PushBigInt(c, x)
}

func SubBigInt(c zc.Calc) {
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)
	x.Sub(x, y)
	zc.PushBigInt(c, x)
}

/*
oper	sub-ia
func	SubIntArch 	Int Int -- Int
title	sub-id
*/
func SubIntArch(c zc.Calc) {
	y := zc.PopInt(c)
	x := zc.PopInt(c)
	zc.PushInt(c, x-y)
}

/*
oper	sub-i64
func	SubInt64 	Int64 Int64 -- Int64
title	sub-i64
*/
func SubInt64(c zc.Calc) {
	y := zc.PopInt64(c)
	x := zc.PopInt64(c)
	zc.PushInt64(c, x-y)
}

/*
oper	sub-i32
func	SubInt32 	Int32 Int32 -- Int32
title	sub-i32
*/
func SubInt32(c zc.Calc) {
	y := zc.PopInt32(c)
	x := zc.PopInt32(c)
	zc.PushInt32(c, x-y)
}

/*
oper	sub-i16
func	SubInt16 	Int16 Int16 -- Int16
title	sub-i16
*/
func SubInt16(c zc.Calc) {
	y := zc.PopInt16(c)
	x := zc.PopInt16(c)
	zc.PushInt16(c, x-y)
}

/*
oper	sub-i8
func	SubInt8 	Int8 Int8 -- Int18
title	sub-i8
*/
func SubInt8(c zc.Calc) {
	y := zc.PopInt8(c)
	x := zc.PopInt8(c)
	zc.PushInt8(c, x-y)
}

/*
oper	sub-ua
func	SubUintArch 	Uint Uint -- Uint
title	sub-ua
*/
func SubUintArch(c zc.Calc) {
	y := zc.PopUint(c)
	x := zc.PopUint(c)
	zc.PushUint(c, x-y)
}

/*
oper	sub-u64
func	SubUint64 	Uint64 Uint64 -- Uint64
title	sub-u64
*/
func SubUint64(c zc.Calc) {
	y := zc.PopUint64(c)
	x := zc.PopUint64(c)
	zc.PushUint64(c, x-y)
}

/*
oper	sub-u32
func	SubUint32 	Uint32 Uint32 -- Uint32
title	sub-u32
*/
func SubUint32(c zc.Calc) {
	y := zc.PopUint32(c)
	x := zc.PopUint32(c)
	zc.PushUint32(c, x-y)
}

/*
oper	sub-u16
func	SubUint16 	Uint16 Uint16 -- Uint16
title	sub-u16
*/
func SubUint16(c zc.Calc) {
	y := zc.PopUint16(c)
	x := zc.PopUint16(c)
	zc.PushUint16(c, x-y)
}

/*
oper	sub-u8
func	SubUint8 	Uint8 Uint8 -- Uint18
title	sub-u8
*/
func SubUint8(c zc.Calc) {
	y := zc.PopUint8(c)
	x := zc.PopUint8(c)
	zc.PushUint8(c, x-y)
}

/*
oper		quo-int
func		QuoBigInt		BigInt BigInt -- BigInt
alias		q-int
title		quo-int
*/
func QuoBigInt(c zc.Calc) {
	var zero big.Int
	y := zc.PopBigInt(c)
	x := zc.PopBigInt(c)

	if y.Cmp(&zero) == 0 {
		zc.ErrDivisionByZero(c)
		return
	}

	x.Quo(x, y)
	zc.PushBigInt(c, x)
}

/*
oper		quo-rem-int
func		QuoRemBigInt	BigInt BigInt -- BigInt BigInt
alias		qr-int
title		Truncated division and remainder
*/
func QuoRemBigInt(c zc.Calc) {
	var q, r big.Int
	a1 := zc.PopBigInt(c)
	a0 := zc.PopBigInt(c)
	q.QuoRem(a0, a1, &r)
	zc.PushBigInt(c, &q)
	zc.PushBigInt(c, &r)
	zc.Annotate(c, "remainder")
}
