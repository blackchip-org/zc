package types

import (
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6/calc/state"
	"github.com/cockroachdb/apd/v3"
)

func mustParseBigFloat(s string) *big.Float {
	var f big.Float
	_, ok := f.SetString(s)
	if !ok {
		panic("invalid float: " + s)
	}
	return &f
}

func TestFloatDecimalTo(t *testing.T) {
	tests := []struct {
		from *apd.Decimal
		to   *big.Float
	}{
		{from: apd.New(1, 0), to: big.NewFloat(1)},
		{from: apd.New(1, 1), to: big.NewFloat(10)},
		{from: apd.New(12, -1), to: big.NewFloat(1.2)},
		{from: apd.New(1, 100000), to: mustParseBigFloat("1e+100000")},
	}

	for _, test := range tests {
		t.Run(test.from.String(), func(t *testing.T) {
			toAny, ok := Float.To(state.New(), test.from)
			if !ok {
				t.Fatalf("conversion not supported")
			}
			to, ok := toAny.(*big.Float)
			if !ok {
				t.Fatalf("wrong type")
			}
			if to.Cmp(test.to) != 0 {
				t.Fatalf("\n have: %v \n want: %v", to, test.to)
			}
		})
	}
}
