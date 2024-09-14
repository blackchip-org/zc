package funcs

import (
	"math"

	"github.com/blackchip-org/zc/v6"
)

func AddInt(c zc.Calc) {
	y := zc.Int.Pop(c)
	x := zc.Int.Pop(c)
	zc.Int.Push(c, x+y)
}

func AddInt8(c zc.Calc) {
	y := zc.Int8.Pop(c)
	x := zc.Int8.Pop(c)
	zc.Int8.Push(c, x+y)
}

func AddInt16(c zc.Calc) {
	y := zc.Int16.Pop(c)
	x := zc.Int16.Pop(c)
	zc.Int16.Push(c, x+y)
}

func AddInt32(c zc.Calc) {
	y := zc.Int32.Pop(c)
	x := zc.Int32.Pop(c)
	zc.Int32.Push(c, x+y)
}

func AddInt64(c zc.Calc) {
	y := zc.Int64.Pop(c)
	x := zc.Int64.Pop(c)
	zc.Int64.Push(c, x+y)
}

func AddUint(c zc.Calc) {
	y := zc.Uint.Pop(c)
	x := zc.Uint.Pop(c)
	zc.Uint.Push(c, x+y)
}

func AddUint8(c zc.Calc) {
	y := zc.Uint8.Pop(c)
	x := zc.Uint8.Pop(c)
	zc.Uint8.Push(c, x+y)
}

func AddUint16(c zc.Calc) {
	y := zc.Uint16.Pop(c)
	x := zc.Uint16.Pop(c)
	zc.Uint16.Push(c, x+y)
}

func AddUint32(c zc.Calc) {
	y := zc.Uint32.Pop(c)
	x := zc.Uint32.Pop(c)
	zc.Uint32.Push(c, x+y)
}

func AddUint64(c zc.Calc) {
	y := zc.Uint64.Pop(c)
	x := zc.Uint64.Pop(c)
	zc.Uint64.Push(c, x+y)
}

func BigIntIs(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.BigInt.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func DataToInt(c zc.Calc) {
	x := zc.Data.Pop(c)
	i := zc.BigInt.New()
	i.SetBytes(x.Bytes())
	zc.BigInt.Push(c, i)
	zc.Data.Recycle(x)
}

func IntIs(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Int.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func Int8Is(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Int8.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func Int16Is(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Int16.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func Int32Is(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Int32.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func Int64Is(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Int64.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func IntToData(c zc.Calc) {
	x := zc.BigInt.Pop(c)
	d := zc.Data.New()
	d.Write(x.Bytes())
	zc.Data.Push(c, d)
	zc.BigInt.Recycle(x)
}

func MaxInt(c zc.Calc) {
	zc.Int.Push(c, math.MaxInt)
}

func MaxUint(c zc.Calc) {
	zc.Uint.Push(c, math.MaxUint)
}

func MinInt(c zc.Calc) {
	zc.Int.Push(c, math.MinInt)
}

func SubInt(c zc.Calc) {
	y := zc.Int.Pop(c)
	x := zc.Int.Pop(c)
	zc.Int.Push(c, x-y)
}

func SubInt8(c zc.Calc) {
	y := zc.Int8.Pop(c)
	x := zc.Int8.Pop(c)
	zc.Int8.Push(c, x-y)
}

func SubInt16(c zc.Calc) {
	y := zc.Int16.Pop(c)
	x := zc.Int16.Pop(c)
	zc.Int16.Push(c, x-y)
}

func SubInt32(c zc.Calc) {
	y := zc.Int32.Pop(c)
	x := zc.Int32.Pop(c)
	zc.Int32.Push(c, x-y)
}

func SubInt64(c zc.Calc) {
	y := zc.Int64.Pop(c)
	x := zc.Int64.Pop(c)
	zc.Int64.Push(c, x-y)
}

func SubUint(c zc.Calc) {
	y := zc.Uint.Pop(c)
	x := zc.Uint.Pop(c)
	zc.Uint.Push(c, x-y)
}

func SubUint8(c zc.Calc) {
	y := zc.Uint8.Pop(c)
	x := zc.Uint8.Pop(c)
	zc.Uint8.Push(c, x-y)
}

func SubUint16(c zc.Calc) {
	y := zc.Uint16.Pop(c)
	x := zc.Uint16.Pop(c)
	zc.Uint16.Push(c, x-y)
}

func SubUint32(c zc.Calc) {
	y := zc.Uint32.Pop(c)
	x := zc.Uint32.Pop(c)
	zc.Uint32.Push(c, x-y)
}

func SubUint64(c zc.Calc) {
	y := zc.Uint64.Pop(c)
	x := zc.Uint64.Pop(c)
	zc.Uint64.Push(c, x-y)
}

func UintIs(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Uint.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func Uint8Is(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Uint8.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func Uint16Is(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Uint16.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func Uint32Is(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Uint32.Parse(c, x)
	zc.Bool.Push(c, ok)
}

func Uint64Is(c zc.Calc) {
	x := zc.String.Pop(c)
	_, ok, _ := zc.Uint64.Parse(c, x)
	zc.Bool.Push(c, ok)
}
