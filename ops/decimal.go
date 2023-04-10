package ops

import (
	t "github.com/blackchip-org/zc/types"
	"github.com/shopspring/decimal"
)

func op1Decimal(fn func(decimal.Decimal) (decimal.Decimal, error)) Func {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.Decimal.Value(args[0])
		z, err := fn(x)
		if err != nil {
			return []t.Generic{}, err
		}
		return []t.Generic{t.Decimal.Generic(z)}, nil
	}
}

func op2Decimal(fn func(decimal.Decimal, decimal.Decimal) (decimal.Decimal, error)) Func {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.Decimal.Value(args[0])
		y := t.Decimal.Value(args[1])
		z, err := fn(x, y)
		if err != nil {
			return []t.Generic{}, err
		}
		return []t.Generic{t.Decimal.Generic(z)}, nil
	}
}

func cmpDecimalFn(fn func(x decimal.Decimal, y decimal.Decimal) bool) Func {
	return func(args []t.Generic) ([]t.Generic, error) {
		x := t.Decimal.Value(args[0])
		y := t.Decimal.Value(args[1])
		z := fn(x, y)
		return []t.Generic{t.Bool.Generic(z)}, nil
	}
}

func divDecimalFn(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) {
	if y.IsZero() {
		return decimal.Zero, ErrDivisionByZero
	}
	return x.Div(y), nil
}

func modDecimalFn(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) {
	if y.IsZero() {
		return decimal.Zero, ErrDivisionByZero
	}
	return x.Mod(y), nil
}

var (
	absDecimal   = op1Decimal(func(x decimal.Decimal) (decimal.Decimal, error) { return x.Abs(), nil })
	addDecimal   = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Add(y), nil })
	ceilDecimal  = op1Decimal(func(x decimal.Decimal) (decimal.Decimal, error) { return x.Ceil(), nil })
	divDecimal   = op2Decimal(divDecimalFn)
	eqDecimal    = cmpDecimalFn(func(x decimal.Decimal, y decimal.Decimal) bool { return x.Cmp(y) == 0 })
	floorDecimal = op1Decimal(func(x decimal.Decimal) (decimal.Decimal, error) { return x.Floor(), nil })
	gtDecimal    = cmpDecimalFn(func(x decimal.Decimal, y decimal.Decimal) bool { return x.Cmp(y) > 0 })
	gteDecimal   = cmpDecimalFn(func(x decimal.Decimal, y decimal.Decimal) bool { return x.Cmp(y) >= 0 })
	ltDecimal    = cmpDecimalFn(func(x decimal.Decimal, y decimal.Decimal) bool { return x.Cmp(y) < 0 })
	lteDecimal   = cmpDecimalFn(func(x decimal.Decimal, y decimal.Decimal) bool { return x.Cmp(y) <= 0 })
	modDecimal   = op2Decimal(modDecimalFn)
	mulDecimal   = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Mul(y), nil })
	negDecimal   = op1Decimal(func(x decimal.Decimal) (decimal.Decimal, error) { return x.Neg(), nil })
	neqDecimal   = cmpDecimalFn(func(x decimal.Decimal, y decimal.Decimal) bool { return x.Cmp(y) != 0 })
	signDecimal  = op1Decimal(func(x decimal.Decimal) (decimal.Decimal, error) { return decimal.NewFromInt(int64(x.Sign())), nil })
	subDecimal   = op2Decimal(func(x decimal.Decimal, y decimal.Decimal) (decimal.Decimal, error) { return x.Sub(y), nil })
)
