package zc

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/cockroachdb/apd/v3"
)

func bi(v string) *big.Int {
	i := new(big.Int)
	_, ok := i.SetString(v, 0)
	if !ok {
		panic(v)
	}
	return i
}

func bf(v string) *big.Float {
	f := new(big.Float)
	f.SetPrec(60)
	_, ok := f.SetString(v)
	if !ok {
		panic(v)
	}
	return f
}

func br(v string) *big.Rat {
	r := new(big.Rat)
	_, ok := r.SetString(v)
	if !ok {
		panic(v)
	}
	return r
}

func d(v string) *apd.Decimal {
	d, _, err := apd.NewFromString(v)
	if err != nil {
		panic(v)
	}
	return d
}

func TestFrom(t *testing.T) {
	big1 := "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	tests := []struct {
		in   any
		out  any
		from Type
		to   Type
		ok   bool
	}{
		// Decimal
		{d("42.42"), d("42.42"), Decimal, Decimal, true},
		{bi("42"), d("42"), BigInt, Decimal, true},
		{complex128(42.42), d("42.42"), Complex, Decimal, true},
		{complex128(42 + 42i), nil, Complex, Decimal, false},
		{float64(42.42), d("42.42"), Float64, Decimal, true},
		{int(42), d("42"), Int, Decimal, true},
		{int8(42), d("42"), Int8, Decimal, true},
		{int16(42), d("42"), Int16, Decimal, true},
		{int32(42), d("42"), Int32, Decimal, true},
		{int64(42), d("42"), Int64, Decimal, true},
		{br("1/2"), d("0.5"), Rat, Decimal, true},
		{br("1/3"), d("0.33333333333333333333"), Rat, Decimal, true},

		// BigFloat
		{bf("42.42"), bf("42.42"), BigFloat, BigFloat, true},
		{bi("42"), bf("42"), BigInt, BigFloat, true},
		{complex128(42.42), bf("42.42"), Complex, BigFloat, true},
		{complex128(42 + 42i), nil, Complex, BigFloat, false},
		{d("42.42"), bf("42.42"), Decimal, BigFloat, true},
		{float64(42.42), bf("42.42"), Float64, BigFloat, true},
		{int(42), bf("42"), Int, BigFloat, true},
		{int8(42), bf("42"), Int8, BigFloat, true},
		{int16(42), bf("42"), Int16, BigFloat, true},
		{int32(42), bf("42"), Int32, BigFloat, true},
		{int64(42), bf("42"), Int64, BigFloat, true},
		{br("1/2"), bf("0.5"), Rat, BigFloat, true},
		{br("1/3"), bf("0.333333333333333315"), Rat, BigFloat, true},
		{uint(42), bf("42"), Uint, BigFloat, true},
		{uint8(42), bf("42"), Uint8, BigFloat, true},
		{uint16(42), bf("42"), Uint16, BigFloat, true},
		{uint32(42), bf("42"), Uint32, BigFloat, true},
		{uint64(42), bf("42"), Uint64, BigFloat, true},
		{"42.42", bf("42.42"), String, BigFloat, true},
		{"$42_000.42", bf("42000.42"), String, BigFloat, true},
		{"x", nil, String, BigFloat, false},

		// BigInt
		{big.NewInt(42), bi("42"), BigInt, BigInt, true},
		{big.NewFloat(42), bi("42"), BigFloat, BigInt, true},
		{complex128(42), bi("42"), Complex, BigInt, true},
		{complex128(42 + 1i), nil, Complex, BigInt, false},
		{d("42"), bi("42"), Decimal, BigInt, true},
		{d("42.2"), nil, Decimal, BigInt, false},
		{float64(42), bi("42"), Float64, BigInt, true},
		{float64(42.2), nil, Float64, BigInt, false},
		{float64(4.2e1), bi("42"), Float64, BigInt, true},
		{float64(math.MaxInt64), nil, Float64, BigInt, false},
		{float64(math.MinInt64), nil, Float64, BigInt, false},
		{int(42), bi("42"), Int, BigInt, true},
		{int8(42), bi("42"), Int8, BigInt, true},
		{int16(42), bi("42"), Int16, BigInt, true},
		{int32(42), bi("42"), Int32, BigInt, true},
		{int64(42), bi("42"), Int64, BigInt, true},
		{big.NewRat(42, 1), bi("42"), Rat, BigInt, true},
		{big.NewRat(42, 10), nil, Rat, BigInt, false},
		{uint(42), bi("42"), Uint, BigInt, true},
		{uint8(42), bi("42"), Uint8, BigInt, true},
		{uint16(42), bi("42"), Uint16, BigInt, true},
		{uint32(42), bi("42"), Uint32, BigInt, true},
		{uint64(42), bi("42"), Uint64, BigInt, true},
		{"42", bi("42"), String, BigInt, true},
		{"x", nil, String, BigInt, false},

		// Complex
		{complex128(42), complex128(42), Complex, Complex, true},
		{big.NewInt(42), complex128(42), BigInt, Complex, true},
		{bi(big1), nil, BigInt, Complex, false},
		{d("42"), complex128(42), Decimal, Complex, true},
		{d("42.42"), complex128(42.42), Decimal, Complex, true},
		{d(big1), complex128(1e100), Decimal, Complex, true},
		{float64(42.42), complex128(42.42), Float64, Complex, true},
		{int(42), complex128(42), Int, Complex, true},
		{int8(42), complex128(42), Int8, Complex, true},
		{int16(42), complex128(42), Int16, Complex, true},
		{int32(42), complex128(42), Int32, Complex, true},
		{int64(42), complex128(42), Int64, Complex, true},
		{big.NewRat(42, 1), complex128(42), Rat, Complex, true},
		{br(big1), complex128(1e100), Rat, Complex, true},
		{uint(42), complex128(42), Uint, Complex, true},
		{uint8(42), complex128(42), Uint8, Complex, true},
		{uint16(42), complex128(42), Uint16, Complex, true},
		{uint32(42), complex128(42), Uint32, Complex, true},
		{uint64(42), complex128(42), Uint64, Complex, true},
		{"42", complex128(42), String, Complex, true},
		{"42+42i", complex(42, 42), String, Complex, true},
		{"x", nil, String, Complex, false},
	}

	s := state.New()
	conf := state.ForConf(s)
	conf.DecPrec = 20
	conf.FloatPrec = 60

	for _, test := range tests {
		name := fmt.Sprintf("%v_%v_%v", test.in, test.from.GoName(), test.to.GoName())
		t.Run(name, func(t *testing.T) {
			out, from, ok := test.to.From(s, test.in)
			if (ok && !Equal(out, test.out)) || from != test.from || ok != test.ok {
				t.Errorf("\n have: %v %v %v \n want: %v %v %v", out, from.Name(), ok, test.out, test.from.Name(), test.ok)
			}
		})
	}
}
