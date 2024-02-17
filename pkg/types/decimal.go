package types

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func init() {
	decimal.DivisionPrecision = 20
}

type Decimal struct {
	v decimal.Decimal
}

var DecimalZero = Decimal{}

func NewDecimalFromFloat(v float64) Decimal {
	return Decimal{v: decimal.NewFromFloat(v)}
}

func NewDecimalFromInt(v int64) Decimal {
	return Decimal{v: decimal.NewFromInt(v)}
}

func NewDecimalFromString(v string) (Decimal, error) {
	d, err := decimal.NewFromString(v)
	if err != nil {
		return DecimalZero, err
	}
	return Decimal{v: d}, nil
}

func (d Decimal) Abs() Decimal {
	return Decimal{v: d.v.Abs()}
}

func (d Decimal) Add(d2 Decimal) Decimal {
	return Decimal{v: d.v.Add(d2.v)}
}

func (d Decimal) Ceil() Decimal {
	return Decimal{v: d.v.Ceil()}
}

func (d Decimal) Cmp(d2 Decimal) int {
	return d.v.Cmp(d2.v)
}

func (d Decimal) Coefficient() *big.Int {
	return d.v.Coefficient()
}

func (d Decimal) Div(d2 Decimal) Decimal {
	return Decimal{v: d.v.Div(d2.v)}
}

func (d Decimal) Equal(d2 Decimal) bool {
	return d.v.Equal(d2.v)
}

func (d Decimal) Floor() Decimal {
	return Decimal{v: d.v.Floor()}
}

func (d Decimal) Float() float64 {
	f, _ := d.v.Float64()
	return f
}

func (d Decimal) GreaterThan(d2 Decimal) bool {
	return d.v.GreaterThan(d2.v)
}

func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	return d.v.GreaterThanOrEqual(d2.v)
}

func (d Decimal) Int() Decimal {
	return NewDecimalFromInt(d.v.IntPart())
}

func (d Decimal) IsInteger() bool {
	return d.v.IsInteger()
}

func (d Decimal) IsNegative() bool {
	return d.v.IsNegative()
}

func (d Decimal) IsPositive() bool {
	return d.v.IsPositive()
}

func (d Decimal) IsZero() bool {
	return d.v.IsZero()
}

func (d Decimal) LessThan(d2 Decimal) bool {
	return d.v.LessThan(d2.v)
}

func (d Decimal) LessThanOrEqual(d2 Decimal) bool {
	return d.v.LessThanOrEqual(d2.v)
}

func (d Decimal) Mod(d2 Decimal) Decimal {
	return Decimal{v: d.v.Mod(d2.v)}
}

func (d Decimal) Mul(d2 Decimal) Decimal {
	return Decimal{v: d.v.Mul(d2.v)}
}

func (d Decimal) Neg() Decimal {
	return Decimal{v: d.v.Neg()}
}

func (d Decimal) Pow(d2 Decimal) Decimal {
	return Decimal{v: d.v.Pow(d2.v)}
}

func (d Decimal) QuoRem(d2 Decimal, prec int32) (Decimal, Decimal) {
	q, r := d.v.QuoRem(d2.v, prec)
	return Decimal{q}, Decimal{r}
}

func (d Decimal) Round(places int32) Decimal {
	return Decimal{v: d.v.Round(places)}
}

func (d Decimal) RoundBank(places int32) Decimal {
	return Decimal{v: d.v.RoundBank(places)}
}

func (d Decimal) RoundCeil(places int32) Decimal {
	return Decimal{v: d.v.RoundCeil(places)}
}

func (d Decimal) RoundDown(places int32) Decimal {
	return Decimal{v: d.v.RoundDown(places)}
}

func (d Decimal) RoundFloor(places int32) Decimal {
	return Decimal{v: d.v.RoundFloor(places)}
}

func (d Decimal) RoundUp(places int32) Decimal {
	return Decimal{v: d.v.RoundUp(places)}
}

func (d Decimal) Sign() int {
	return d.v.Sign()
}

func (d Decimal) String() string {
	return d.v.String()
}

func (d Decimal) StringRound(places int) string {
	if places < 0 {
		return d.String()
	}
	return d.v.StringFixed(int32(places))
}

func (d Decimal) Sub(d2 Decimal) Decimal {
	return Decimal{v: d.v.Sub(d2.v)}
}

func (d Decimal) Truncate(precision int32) Decimal {
	return Decimal{v: d.v.Truncate(precision)}
}
