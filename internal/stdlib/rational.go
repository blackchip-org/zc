package stdlib

import (
	"fmt"
	"math/big"
	sl "math/big"

	"github.com/blackchip-org/zc/types"
)

type slRat struct {
	val *sl.Rat
}

func unwrapR(r types.Rational) *sl.Rat {
	return r.(slRat).val
}

func wrapR(r *sl.Rat) slRat {
	return slRat{val: r}
}

func (r slRat) Add(r2 types.Rational) types.Rational {
	var z big.Rat
	return wrapR(z.Add(r.val, unwrapR(r2)))
}

func (r slRat) Type() types.Type {
	return types.RationalType{}
}

func (r slRat) Value() types.Value {
	return r
}

func (r slRat) Native() any {
	return r
}

func (r slRat) String() string {
	return r.val.String()
}

func ParseRational(s string) (types.Rational, error) {
	var r big.Rat
	z, ok := r.SetString(s)
	if !ok {
		return wrapR(z), fmt.Errorf("unable to parse: %v", s)
	}
	return wrapR(z), nil
}

func NewRational(a int, b int) types.Rational {
	return wrapR(big.NewRat(int64(a), int64(b)))
}

func UseRational() {
	types.UseRational(types.RationalImpl{
		Parse: ParseRational,
		New:   NewRational,
	})
}
