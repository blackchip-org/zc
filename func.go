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
type FuncBoolCompOp func(bool, bool) bool
type FuncDecimalCompOp func(decimal.Decimal, decimal.Decimal) bool
type FuncStringCompOp func(string, string) bool

type FuncsCompOp struct {
	BigInt  FuncBigIntCompOp
	Bool    FuncBoolCompOp
	Decimal FuncDecimalCompOp
	String  FuncStringCompOp
}

func BigIntNumOp2(calc *Calc, fn FuncBigIntNumOp2) error {
	a, b, err := calc.PopBigInt2()
	if err != nil {
		return err
	}
	var r big.Int
	fn(&r, a, b)

	calc.Stack.Push(FormatBigInt(&r))
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

func NumOp2(calc *Calc, ops FuncsNumOp2) error {
	a, b, err := calc.Pop2()
	if err != nil {
		return err
	}

	if IsBigInt(a) && IsBigInt(b) {
		bi, err := ParseBigInt(b)
		if err != nil {
			return err
		}
		ai, err := ParseBigInt(a)
		if err != nil {
			return err
		}
		var zi big.Int
		ops.BigInt(&zi, ai, bi)

		calc.Stack.Push(FormatBigInt(&zi))
		return nil
	}

	bd, err := ParseDecimal(b)
	if err != nil {
		return err
	}
	ad, err := ParseDecimal(a)
	if err != nil {
		return err
	}

	zd := ops.Decimal(ad, bd)
	calc.Stack.Push(FormatDecimal(zd))
	return nil
}

func CompOp(calc *Calc, ops FuncsCompOp) error {
	a, b, err := calc.Pop2()
	if err != nil {
		return err
	}

	var result bool
	switch {
	case IsBool(a) && IsBool(b):
		x, y := MustParseBool(a), MustParseBool(b)
		result = ops.Bool(x, y)
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
