package zc

import (
	"errors"
	"math/big"

	"github.com/shopspring/decimal"
)

type FuncBigInt2 func(*big.Int, *big.Int, *big.Int)
type FuncDec2 func(decimal.Decimal, decimal.Decimal) decimal.Decimal

type NumOp2 struct {
	BigInt2 FuncBigInt2
	Dec2    FuncDec2
}

func BigInt2(calc *Calc, fn FuncBigInt2) error {
	b, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	a, err := calc.Stack.Pop()
	if err != nil {
		return err
	}

	bi, err := ParseBigInt(b)
	if err != nil {
		return err
	}
	ai, err := ParseBigInt(a)
	if err != nil {
		return err
	}

	var zi big.Int
	fn(&zi, ai, bi)

	calc.Stack.Push(FormatBigInt(&zi))
	return nil
}

func Dec2(calc *Calc, fn func(a decimal.Decimal, b decimal.Decimal) decimal.Decimal) (err error) {
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

	b, err := calc.Stack.Pop()
	if err != nil {
		return
	}
	a, err := calc.Stack.Pop()
	if err != nil {
		return
	}

	bd, err := ParseDecimal(b)
	if err != nil {
		return
	}
	ad, err := ParseDecimal(a)
	if err != nil {
		return
	}

	zd := fn(ad, bd)
	calc.Stack.Push(FormatDecimal(zd))
	return nil
}

func Num2(calc *Calc, ops NumOp2) error {
	b, err := calc.Stack.Pop()
	if err != nil {
		return err
	}
	a, err := calc.Stack.Pop()
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
		ops.BigInt2(&zi, ai, bi)

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

	zd := ops.Dec2(ad, bd)
	calc.Stack.Push(FormatDecimal(zd))
	return nil
}
