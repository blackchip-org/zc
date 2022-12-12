package zc

import (
	"errors"
	"math/big"

	"github.com/shopspring/decimal"
)

type FuncBigIntNumOp2 func(*big.Int, *big.Int, *big.Int)
type FuncDecimalNumOp2 func(decimal.Decimal, decimal.Decimal) decimal.Decimal

type FuncsNumOp2 struct {
	BigInt  FuncBigIntNumOp2
	Decimal FuncDecimalNumOp2
}

type FuncBigIntCompOp func(*big.Int, *big.Int) bool
type FuncDecimalCompOp func(decimal.Decimal, decimal.Decimal) bool
type FuncStringCompOp func(string, string) bool

type FuncsCompOp struct {
	BigInt  FuncBigIntCompOp
	Decimal FuncDecimalCompOp
	String  FuncStringCompOp
}

type FuncBoolOp func(bool, bool) bool

func resolveRadix(rx int, ry int) int {
	switch {
	case rx == 16 || ry == 16:
		return 16
	case rx == 8 || ry == 8:
		return 8
	case rx == 2 || ry == 2:
		return 2
	}
	return 10
}

func BigIntNumOp2(calc *Calc, fn FuncBigIntNumOp2) error {
	x, y, err := calc.Pop2()
	if err != nil {
		return err
	}

	a, err := ParseBigInt(x)
	if err != nil {
		return err
	}

	b, err := ParseBigInt(y)
	if err != nil {
		return err
	}

	var c big.Int
	fn(&c, a, b)

	radix := resolveRadix(ParseRadix(x), ParseRadix(y))
	calc.Stack.Push(FormatBigIntBase(&c, radix))
	return nil
}

func BigIntCompOp(calc *Calc, fn FuncBigIntCompOp) error {
	a, b, err := calc.PopBigInt2()
	if err != nil {
		return err
	}
	r := fn(a, b)

	calc.Stack.Push(FormatBool(r))
	return nil
}

func BoolOp2(calc *Calc, fn FuncBoolOp) error {
	a, b, err := calc.PopBool2()
	if err != nil {
		return err
	}

	r := fn(a, b)
	calc.Stack.Push(FormatBool(r))
	return nil
}

func DecNumOp2(calc *Calc, fn func(a decimal.Decimal, b decimal.Decimal) decimal.Decimal) (err error) {
	defer func() {
		if p := recover(); p != nil {
			msg, ok := p.(string)
			if !ok {
				panic(p)
			}
			if msg == "decimal division by 0" {
				err = errors.New("division by zero")
			} else {
				panic(p)
			}
		}
	}()

	a, b, err := calc.PopDecimal2()
	if err != nil {
		return
	}

	r := fn(a, b)
	calc.Stack.Push(FormatDecimal(r))
	return nil
}

func DecCompOp(calc *Calc, fn func(a decimal.Decimal, b decimal.Decimal) bool) (err error) {
	a, b, err := calc.PopDecimal2()
	if err != nil {
		return
	}

	r := fn(a, b)
	calc.Stack.Push(FormatBool(r))
	return nil
}

func StringCompOp(calc *Calc, fn func(a string, b string) bool) error {
	a, b, err := calc.Pop2()
	if err != nil {
		return err
	}

	r := fn(a, b)
	calc.Stack.Push(FormatBool(r))
	return nil
}

func NumOp2(calc *Calc, ops FuncsNumOp2) error {
	a, b, err := calc.Peek2()
	if err != nil {
		return err
	}

	switch {
	case IsBigInt(a) && IsBigInt(b):
		return BigIntNumOp2(calc, ops.BigInt)
	default:
		return DecNumOp2(calc, ops.Decimal)
	}
}

func CompOp(calc *Calc, ops FuncsCompOp) error {
	a, b, err := calc.Pop2()
	if err != nil {
		return err
	}

	var result bool
	switch {
	case IsBigInt(a) && IsBigInt(b):
		x, y := MustParseBigInt(a), MustParseBigInt(b)
		result = ops.BigInt(x, y)
	case IsDecimal(a) && IsDecimal(b):
		x, y := MustParseDecimal(a), MustParseDecimal(b)
		result = ops.Decimal(x, y)
	default:
		result = ops.String(a, b)
	}

	calc.Stack.Push(FormatBool(result))
	return nil
}
