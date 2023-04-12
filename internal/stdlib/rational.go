package stdlib

import (
	"errors"
	"math/big"
	sl "math/big"

	"github.com/blackchip-org/zc/number"
)

type slRat struct {
	val *sl.Rat
}

func unwrapR(r number.Rational) *sl.Rat {
	return r.(slRat).val
}

func wrapR(r *sl.Rat) slRat {
	return slRat{val: r}
}

func (r slRat) Add(r2 number.Rational) number.Rational {
	var z big.Rat
	return wrapR(z.Add(r.val, unwrapR(r2)))
}

func (r slRat) String() string {
	return r.val.String()
}

func ParseRational(s string) (number.Rational, error) {
	var r big.Rat
	z, ok := r.SetString(s)
	if !ok {
		return wrapR(z), errors.New("unable to parse")
	}
	return wrapR(z), nil
}

func NewRational(a int, b int) number.Rational {
	return wrapR(big.NewRat(int64(a), int64(b)))
}

func UseRational() {
	number.UseRational(number.RationalImpl{
		Parse: ParseRational,
		New:   NewRational,
	})
}
