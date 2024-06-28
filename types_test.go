package zc

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/cockroachdb/apd/v3"
)

func TestFrom(t *testing.T) {
	big1 := "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	s := state.New()
	conf := state.ForConf(s)
	conf.DecPrec = 20
	conf.FloatPrec = 60

	fl := func(v string) *big.Float { return BigFloat.MustParse(s, v) }
	in := func(v string) *big.Int { return BigInt.MustParse(s, v) }
	rt := func(v string) *big.Rat { return Rat.MustParse(s, v) }
	dc := func(v string) *apd.Decimal { return Decimal.MustParse(s, v) }

	tests := []struct {
		in   any
		out  any
		from Type
		to   Type
		ok   bool
	}{
		// Decimal
		{dc("42.42"), dc("42.42"), Decimal, Decimal, true},
		{in("42"), dc("42"), BigInt, Decimal, true},
		{complex128(42.42), dc("42.42"), Complex, Decimal, true},
		{complex128(42 + 42i), nil, Complex, Decimal, false},
		{float64(42.42), dc("42.42"), Float64, Decimal, true},
		{int(42), dc("42"), Int, Decimal, true},
		{int8(42), dc("42"), Int8, Decimal, true},
		{int16(42), dc("42"), Int16, Decimal, true},
		{int32(42), dc("42"), Int32, Decimal, true},
		{int64(42), dc("42"), Int64, Decimal, true},
		{rt("1/2"), dc("0.5"), Rat, Decimal, true},
		{rt("1/3"), dc("0.33333333333333333333"), Rat, Decimal, true},

		// BigFloat
		{fl("42.42"), fl("42.42"), BigFloat, BigFloat, true},
		{in("42"), fl("42"), BigInt, BigFloat, true},
		{complex128(42.42), fl("42.42"), Complex, BigFloat, true},
		{complex128(42 + 42i), nil, Complex, BigFloat, false},
		{dc("42.42"), fl("42.42"), Decimal, BigFloat, true},
		{float64(42.42), fl("42.42"), Float64, BigFloat, true},
		{int(42), fl("42"), Int, BigFloat, true},
		{int8(42), fl("42"), Int8, BigFloat, true},
		{int16(42), fl("42"), Int16, BigFloat, true},
		{int32(42), fl("42"), Int32, BigFloat, true},
		{int64(42), fl("42"), Int64, BigFloat, true},
		{rt("1/2"), fl("0.5"), Rat, BigFloat, true},
		{rt("1/3"), fl("0.333333333333333315"), Rat, BigFloat, true},
		{uint(42), fl("42"), Uint, BigFloat, true},
		{uint8(42), fl("42"), Uint8, BigFloat, true},
		{uint16(42), fl("42"), Uint16, BigFloat, true},
		{uint32(42), fl("42"), Uint32, BigFloat, true},
		{uint64(42), fl("42"), Uint64, BigFloat, true},
		{"42.42", fl("42.42"), String, BigFloat, true},
		{"$42_000.42", fl("42000.42"), String, BigFloat, true},
		{"x", nil, String, BigFloat, false},

		// BigInt
		{big.NewInt(42), in("42"), BigInt, BigInt, true},
		{big.NewFloat(42), in("42"), BigFloat, BigInt, true},
		{complex128(42), in("42"), Complex, BigInt, true},
		{complex128(42 + 1i), nil, Complex, BigInt, false},
		{dc("42"), in("42"), Decimal, BigInt, true},
		{dc("42.2"), nil, Decimal, BigInt, false},
		{float64(42), in("42"), Float64, BigInt, true},
		{float64(42.2), nil, Float64, BigInt, false},
		{float64(4.2e1), in("42"), Float64, BigInt, true},
		{float64(math.MaxInt64), nil, Float64, BigInt, false},
		{float64(math.MinInt64), nil, Float64, BigInt, false},
		{int(42), in("42"), Int, BigInt, true},
		{int8(42), in("42"), Int8, BigInt, true},
		{int16(42), in("42"), Int16, BigInt, true},
		{int32(42), in("42"), Int32, BigInt, true},
		{int64(42), in("42"), Int64, BigInt, true},
		{big.NewRat(42, 1), in("42"), Rat, BigInt, true},
		{big.NewRat(42, 10), nil, Rat, BigInt, false},
		{uint(42), in("42"), Uint, BigInt, true},
		{uint8(42), in("42"), Uint8, BigInt, true},
		{uint16(42), in("42"), Uint16, BigInt, true},
		{uint32(42), in("42"), Uint32, BigInt, true},
		{uint64(42), in("42"), Uint64, BigInt, true},
		{"42", in("42"), String, BigInt, true},
		{"x", nil, String, BigInt, false},

		// Complex
		{complex128(42), complex128(42), Complex, Complex, true},
		{big.NewInt(42), complex128(42), BigInt, Complex, true},
		{in(big1), nil, BigInt, Complex, false},
		{dc("42"), complex128(42), Decimal, Complex, true},
		{dc("42.42"), complex128(42.42), Decimal, Complex, true},
		{dc(big1), complex128(1e100), Decimal, Complex, true},
		{float64(42.42), complex128(42.42), Float64, Complex, true},
		{int(42), complex128(42), Int, Complex, true},
		{int8(42), complex128(42), Int8, Complex, true},
		{int16(42), complex128(42), Int16, Complex, true},
		{int32(42), complex128(42), Int32, Complex, true},
		{int64(42), complex128(42), Int64, Complex, true},
		{big.NewRat(42, 1), complex128(42), Rat, Complex, true},
		{rt(big1), complex128(1e100), Rat, Complex, true},
		{uint(42), complex128(42), Uint, Complex, true},
		{uint8(42), complex128(42), Uint8, Complex, true},
		{uint16(42), complex128(42), Uint16, Complex, true},
		{uint32(42), complex128(42), Uint32, Complex, true},
		{uint64(42), complex128(42), Uint64, Complex, true},
		{"42", complex128(42), String, Complex, true},
		{"42+42i", complex(42, 42), String, Complex, true},
		{"x", nil, String, Complex, false},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v_%v_%v", test.in, test.from.GoName(), test.to.GoName())
		t.Run(name, func(t *testing.T) {
			out, from, ok := test.to.From(s, test.in)
			if (ok && !Equal(out, test.out)) || from != test.from || ok != test.ok {
				t.Errorf("\n have: %v %v %v \n want: %v %v %v", Format(out), from.Name(), ok, Format(test.out), test.from.Name(), test.ok)
			}
		})
	}
}
