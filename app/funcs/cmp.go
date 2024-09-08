package funcs

import (
	"cmp"
	"slices"

	"github.com/blackchip-org/zc/v6"
)

// ----------------------------------------------------------------------------
func EqBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) == 0)
}

func EqBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) == 0)
}

func EqDecimal(c zc.Calc) {
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) == 0)
}

func EqRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) == 0)
}

func EqComplex(c zc.Calc) {
	y := zc.Complex.Pop(c)
	x := zc.Complex.Pop(c)
	zc.Bool.Push(c, x == y)
}

func EqString(c zc.Calc) {
	y := zc.String.Pop(c)
	x := zc.String.Pop(c)
	zc.Bool.Push(c, x == y)
}

// ----------------------------------------------------------------------------
func GtBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) > 0)
}

func GtBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) > 0)
}

func GtDecimal(c zc.Calc) {
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) > 0)
}

func GtRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) > 0)
}

func GtString(c zc.Calc) {
	y := zc.String.Pop(c)
	x := zc.String.Pop(c)
	zc.Bool.Push(c, x > y)
}

// ----------------------------------------------------------------------------
func GteBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) >= 0)
}

func GteBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) >= 0)
}

func GteDecimal(c zc.Calc) {
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) >= 0)
}

func GteRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) >= 0)
}

func GteString(c zc.Calc) {
	y := zc.String.Pop(c)
	x := zc.String.Pop(c)
	zc.Bool.Push(c, x >= y)
}

// ----------------------------------------------------------------------------
func LtBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) < 0)
}

func LtBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) < 0)
}

func LtDecimal(c zc.Calc) {
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) < 0)
}

func LtRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) < 0)
}

func LtString(c zc.Calc) {
	y := zc.String.Pop(c)
	x := zc.String.Pop(c)
	zc.Bool.Push(c, x < y)
}

// ----------------------------------------------------------------------------
func LteBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) <= 0)
}

func LteBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) <= 0)
}

func LteDecimal(c zc.Calc) {
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) <= 0)
}

func LteRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) <= 0)
}

func LteString(c zc.Calc) {
	y := zc.String.Pop(c)
	x := zc.String.Pop(c)
	zc.Bool.Push(c, x <= y)
}

// ----------------------------------------------------------------------------
func NeqBigInt(c zc.Calc) {
	y := zc.BigInt.Pop(c)
	x := zc.BigInt.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) != 0)
}

func NeqBigFloat(c zc.Calc) {
	y := zc.BigFloat.Pop(c)
	x := zc.BigFloat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) != 0)
}

func NeqDecimal(c zc.Calc) {
	y := zc.Decimal.Pop(c)
	x := zc.Decimal.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) != 0)
}

func NeqRat(c zc.Calc) {
	y := zc.Rat.Pop(c)
	x := zc.Rat.Pop(c)
	zc.Bool.Push(c, x.Cmp(y) != 0)
}

func NeqComplex(c zc.Calc) {
	y := zc.Complex.Pop(c)
	x := zc.Complex.Pop(c)
	zc.Bool.Push(c, x != y)
}

func NeqString(c zc.Calc) {
	y := zc.String.Pop(c)
	x := zc.String.Pop(c)
	zc.Bool.Push(c, x != y)
}

// ----------------------------------------------------------------------------

func SortBigInt(c zc.Calc) {
	s := c.Stack()
	slices.SortStableFunc(s, func(a, b zc.Item) int {
		return zc.BigInt.As(a.TypeVal).Cmp(zc.BigInt.As(b.TypeVal))
	})
	c.SetStack(s)
}

func SortBigFloat(c zc.Calc) {
	s := c.Stack()
	slices.SortStableFunc(s, func(a, b zc.Item) int {
		return zc.BigFloat.As(a.TypeVal).Cmp(zc.BigFloat.As(b.TypeVal))
	})
	c.SetStack(s)
}

func SortDecimal(c zc.Calc) {
	s := c.Stack()
	slices.SortStableFunc(s, func(a, b zc.Item) int {
		return zc.Decimal.As(a.TypeVal).Cmp(zc.Decimal.As(b.TypeVal))
	})
	c.SetStack(s)
}

func SortRat(c zc.Calc) {
	s := c.Stack()
	slices.SortStableFunc(s, func(a, b zc.Item) int {
		return zc.Rat.As(a.TypeVal).Cmp(zc.Rat.As(b.TypeVal))
	})
	c.SetStack(s)
}

func SortString(c zc.Calc) {
	s := c.Stack()
	slices.SortStableFunc(s, func(a, b zc.Item) int {
		return cmp.Compare(zc.String.As(a.TypeVal), zc.String.As(b.TypeVal))
	})
	c.SetStack(s)
}
