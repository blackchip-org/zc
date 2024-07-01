package zc

import (
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
		id   string
		in   any
		out  any
		from Type
		to   Type
		ok   bool
	}{
		// BigFloat
		{"fl01", fl("42.42"), fl("42.42"), BigFloat, BigFloat, true},
		{"fl02", in("42"), fl("42"), BigInt, BigFloat, true},
		{"fl03", complex128(42.42), fl("42.42"), Complex, BigFloat, true},
		{"fl04", complex128(42 + 42i), nil, Complex, BigFloat, false},
		{"fl05", dc("42.42"), fl("42.42"), Decimal, BigFloat, true},
		{"fl06", float64(42.42), fl("42.42"), Float64, BigFloat, true},
		{"fl07", int(42), fl("42"), Int, BigFloat, true},
		{"fl08", int8(42), fl("42"), Int8, BigFloat, true},
		{"fl09", int16(42), fl("42"), Int16, BigFloat, true},
		{"fl10", int32(42), fl("42"), Int32, BigFloat, true},
		{"fl11", int64(42), fl("42"), Int64, BigFloat, true},
		{"fl12", rt("1/2"), fl("0.5"), Rat, BigFloat, true},
		{"fl13", rt("1/3"), fl("0.333333333333333315"), Rat, BigFloat, true},
		{"fl14", uint(42), fl("42"), Uint, BigFloat, true},
		{"fl15", uint8(42), fl("42"), Uint8, BigFloat, true},
		{"fl16", uint16(42), fl("42"), Uint16, BigFloat, true},
		{"fl17", uint32(42), fl("42"), Uint32, BigFloat, true},
		{"fl18", uint64(42), fl("42"), Uint64, BigFloat, true},
		{"fl19", "42.42", fl("42.42"), String, BigFloat, true},
		{"fl20", "$42_000.42", fl("42000.42"), String, BigFloat, true},
		{"fl21", "x", nil, String, BigFloat, false},

		// BigInt
		{"in01", big.NewInt(42), in("42"), BigInt, BigInt, true},
		{"in02", big.NewFloat(42), in("42"), BigFloat, BigInt, true},
		{"in03", complex128(42), in("42"), Complex, BigInt, true},
		{"in04", complex128(42 + 1i), nil, Complex, BigInt, false},
		{"in05", dc("42"), in("42"), Decimal, BigInt, true},
		{"in06", dc("42.2"), nil, Decimal, BigInt, false},
		{"in07", float64(42), in("42"), Float64, BigInt, true},
		{"in08", float64(42.2), nil, Float64, BigInt, false},
		{"in09", float64(4.2e1), in("42"), Float64, BigInt, true},
		{"in10", float64(math.MaxInt64), nil, Float64, BigInt, false},
		{"in11", float64(math.MinInt64), nil, Float64, BigInt, false},
		{"in12", int(42), in("42"), Int, BigInt, true},
		{"in13", int8(42), in("42"), Int8, BigInt, true},
		{"in14", int16(42), in("42"), Int16, BigInt, true},
		{"in15", int32(42), in("42"), Int32, BigInt, true},
		{"in16", int64(42), in("42"), Int64, BigInt, true},
		{"in17", big.NewRat(42, 1), in("42"), Rat, BigInt, true},
		{"in18", big.NewRat(42, 10), nil, Rat, BigInt, false},
		{"in19", uint(42), in("42"), Uint, BigInt, true},
		{"in20", uint8(42), in("42"), Uint8, BigInt, true},
		{"in21", uint16(42), in("42"), Uint16, BigInt, true},
		{"in22", uint32(42), in("42"), Uint32, BigInt, true},
		{"in23", uint64(42), in("42"), Uint64, BigInt, true},
		{"in24", "42", in("42"), String, BigInt, true},
		{"in25", "x", nil, String, BigInt, false},

		// Complex
		{"cp01", complex128(42), complex128(42), Complex, Complex, true},
		{"cp02", big.NewInt(42), complex128(42), BigInt, Complex, true},
		{"cp03", in(big1), nil, BigInt, Complex, false},
		{"cp04", dc("42"), complex128(42), Decimal, Complex, true},
		{"cp05", dc("42.42"), complex128(42.42), Decimal, Complex, true},
		{"cp06", dc(big1), complex128(1e100), Decimal, Complex, true},
		{"cp07", float64(42.42), complex128(42.42), Float64, Complex, true},
		{"cp08", int(42), complex128(42), Int, Complex, true},
		{"cp09", int8(42), complex128(42), Int8, Complex, true},
		{"cp10", int16(42), complex128(42), Int16, Complex, true},
		{"cp11", int32(42), complex128(42), Int32, Complex, true},
		{"cp12", int64(42), complex128(42), Int64, Complex, true},
		{"cp13", big.NewRat(42, 1), complex128(42), Rat, Complex, true},
		{"cp14", rt(big1), complex128(1e100), Rat, Complex, true},
		{"cp15", uint(42), complex128(42), Uint, Complex, true},
		{"cp16", uint8(42), complex128(42), Uint8, Complex, true},
		{"cp17", uint16(42), complex128(42), Uint16, Complex, true},
		{"cp18", uint32(42), complex128(42), Uint32, Complex, true},
		{"cp19", uint64(42), complex128(42), Uint64, Complex, true},
		{"cp20", "42", complex128(42), String, Complex, true},
		{"cp21", "42+42i", complex(42, 42), String, Complex, true},
		{"cp22", "x", nil, String, Complex, false},

		// Decimal
		{"dc01", dc("42.42"), dc("42.42"), Decimal, Decimal, true},
		{"dc02", in("42"), dc("42"), BigInt, Decimal, true},
		{"dc03", complex128(42.42), dc("42.42"), Complex, Decimal, true},
		{"dc04", complex128(42 + 42i), nil, Complex, Decimal, false},
		{"dc05", float64(42.42), dc("42.42"), Float64, Decimal, true},
		{"dc06", int(42), dc("42"), Int, Decimal, true},
		{"dc07", int8(42), dc("42"), Int8, Decimal, true},
		{"dc08", int16(42), dc("42"), Int16, Decimal, true},
		{"dc09", int32(42), dc("42"), Int32, Decimal, true},
		{"dc10", int64(42), dc("42"), Int64, Decimal, true},
		{"dc11", rt("1/2"), dc("0.5"), Rat, Decimal, true},
		{"dc12", rt("1/3"), dc("0.33333333333333333333"), Rat, Decimal, true},
		{"dc13", uint(42), dc("42"), Uint, Decimal, true},
		{"dc14", uint8(42), dc("42"), Uint8, Decimal, true},
		{"dc15", uint16(42), dc("42"), Uint16, Decimal, true},
		{"dc16", uint32(42), dc("42"), Uint32, Decimal, true},
		{"dc17", uint64(42), dc("42"), Uint64, Decimal, true},
		{"dc18", "42.42", dc("42.42"), String, Decimal, true},
		{"dc19", "$42,000.42", dc("42000.42"), String, Decimal, true},
		{"dc20", "x", nil, String, Decimal, false},

		// Float64
		{"6f01", float64(42.42), float64(42.42), Float64, Float64, true},
		{"6f02", fl("42.42"), float64(42.42), BigFloat, Float64, true},
		{"6f03", fl("42.42e100000"), float64(42.42), BigFloat, Float64, false},
		{"6f04", complex128(42.42), float64(42.42), Complex, Float64, true},
		{"6f05", complex128(42 + 42i), nil, Complex, Float64, false},
		{"6f06", int(42), float64(42), Int, Float64, true},
		{"6f07", int8(42), float64(42), Int8, Float64, true},
		{"6f08", int16(42), float64(42), Int16, Float64, true},
		{"6f09", int32(42), float64(42), Int32, Float64, true},
		{"6f10", int64(42), float64(42), Int64, Float64, true},
		{"6f11", rt("1/2"), float64(0.5), Rat, Float64, true},
		{"6f12", rt("1/3"), float64(0.3333333333333333), Rat, Float64, true},
		{"6f13", uint(42), float64(42), Uint, Float64, true},
		{"6f14", uint8(42), float64(42), Uint8, Float64, true},
		{"6f15", uint16(42), float64(42), Uint16, Float64, true},
		{"6f16", uint32(42), float64(42), Uint32, Float64, true},
		{"6f17", uint64(42), float64(42), Uint64, Float64, true},
		{"6f18", "42.42", float64(42.42), String, Float64, true},
		{"6f19", "$42_000.42", float64(42000.42), String, Float64, true},
		{"6f20", "x", nil, String, BigFloat, false},

		// int
		{"is01", int(42), int(42), Int, Int, true},
		{"is02", fl("42"), int(42), BigFloat, Int, true},
		{"is03", fl("42.42"), nil, BigFloat, Int, false},
		{"is04", fl("1e42"), nil, BigFloat, Int, false},
		{"is05", in("42"), int(42), BigInt, Int, true},
		{"is06", in("9223372036854775808"), nil, BigInt, Int, false},
		{"is07", complex128(42), int(42), Complex, Int, true},
		{"is08", complex128(42 + 42i), nil, Complex, Int, false},
		{"is09", complex128(42.42), nil, Complex, Int, false},
		{"is10", complex128(1e42), nil, Complex, Int, false},
		{"is11", dc("42"), int(42), Decimal, Int, true},
		{"is12", dc("42.42"), nil, Decimal, Int, false},
		{"is13", dc("42e42"), nil, Decimal, Int, false},
		{"is14", int8(42), int(42), Int8, Int, true},
		{"is15", int16(42), int(42), Int16, Int, true},
		{"is16", int32(42), int(42), Int32, Int, true},
		{"is17", int64(42), int(42), Int64, Int, true},
		{"is18", rt("84/2"), int(42), Rat, Int, true},
		{"is19", rt("84/11"), nil, Rat, Int, false},
		{"is20", rt("9223372036854775808"), nil, Rat, Int, false},
		{"is21", uint(42), int(42), Uint8, Int, true},
		{"is22", uint8(42), int(42), Uint8, Int, true},
		{"is23", uint16(42), int(42), Uint16, Int, true},
		{"is24", uint32(42), int(42), Uint32, Int, true},
		{"is25", uint64(42), int(42), Uint64, Int, true},
		{"is26", uint64(18446744073709551615), nil, Uint64, Int, false},
		{"is27", "42", int(42), String, Int, true},
		{"is28", "$42,000", int(42000), String, Int, true},
		{"is29", "42.42", nil, String, Int, false},
		{"is30", "9223372036854775808", nil, String, Int, false},

		// int8
		{"sa01", int8(-42), int8(-42), Int8, Int8, true},
		{"sa02", fl("-42"), int8(-42), BigFloat, Int8, true},
		{"sa03", fl("-42.42"), nil, BigFloat, Int8, false},
		{"sa04", fl("-129"), nil, BigFloat, Int8, false},
		{"sa05", in("-42"), int8(-42), BigInt, Int8, true},
		{"sa06", in("-129"), nil, BigInt, Int8, false},
		{"sa07", complex128(-42), int8(-42), Complex, Int8, true},
		{"sa08", complex128(-42.42), nil, Complex, Int8, false},
		{"sa09", complex128(-42 + 42i), nil, Complex, Int8, false},
		{"sa10", complex128(-129), nil, Complex, Int8, false},
		{"sa11", dc("-42"), int8(-42), Decimal, Int8, true},
		{"sa12", dc("-42.42"), nil, Decimal, Int8, false},
		{"sa13", dc("-129"), nil, Decimal, Int8, false},
		{"sa14", float64(-42), int8(-42), Float64, Int8, true},
		{"sa15", float64(-42.42), nil, Float64, Int8, false},
		{"sa16", float64(-129), nil, Float64, Int8, false},
		{"sa17", int(-42), int8(-42), Int, Int8, true},
		{"sa18", int16(-42), int8(-42), Int16, Int8, true},
		{"sa19", int32(-42), int8(-42), Int32, Int8, true},
		{"sa20", int64(-42), int8(-42), Int64, Int8, true},
		{"sa21", int(-129), nil, Int, Int8, false},
		{"sa22", int16(-129), nil, Int16, Int8, false},
		{"sa23", int32(-129), nil, Int32, Int8, false},
		{"sa24", int64(-129), nil, Int64, Int8, false},
		{"sa25", rt("-84/2"), int8(-42), Rat, Int8, true},
		{"sa26", rt("-84/11"), nil, Rat, Int8, false},
		{"sa27", rt("-129"), nil, Rat, Int8, false},
		{"sa28", uint(42), int8(42), Uint, Int8, true},
		{"sa29", uint8(42), int8(42), Uint8, Int8, true},
		{"sa30", uint16(42), int8(42), Uint16, Int8, true},
		{"sa31", uint32(42), int8(42), Uint32, Int8, true},
		{"sa32", uint64(42), int8(42), Uint64, Int8, true},
		{"sa33", uint(129), nil, Uint, Int8, false},
		{"sa34", uint8(129), nil, Uint8, Int8, false},
		{"sa35", uint16(129), nil, Uint16, Int8, false},
		{"sa36", uint32(129), nil, Uint32, Int8, false},
		{"sa37", uint64(129), nil, Uint64, Int8, false},
		{"sa38", "-42", int8(-42), String, Int8, true},
		{"sa39", "$42", int8(42), String, Int8, true},
		{"sa40", "42.42", nil, String, Int8, false},
		{"sa41", "-129", nil, String, Int8, false},

		// int16
		{"sb01", int16(-42), int16(-42), Int16, Int16, true},
		{"sb02", fl("-42"), int16(-42), BigFloat, Int16, true},
		{"sb03", fl("-42.42"), nil, BigFloat, Int16, false},
		{"sb04", fl("-32769"), nil, BigFloat, Int16, false},
		{"sb05", in("-42"), int16(-42), BigInt, Int16, true},
		{"sb06", in("-32769"), nil, BigInt, Int16, false},
		{"sb07", complex128(-42), int16(-42), Complex, Int16, true},
		{"sb08", complex128(-42.42), nil, Complex, Int16, false},
		{"sb09", complex128(-42 + 42i), nil, Complex, Int16, false},
		{"sb10", complex128(-32769), nil, Complex, Int16, false},
		{"sb11", dc("-42"), int16(-42), Decimal, Int16, true},
		{"sb12", dc("-42.42"), nil, Decimal, Int16, false},
		{"sb13", dc("-32769"), nil, Decimal, Int16, false},
		{"sb14", float64(-42), int16(-42), Float64, Int16, true},
		{"sb15", float64(-42.42), nil, Float64, Int16, false},
		{"sb16", float64(-32769), nil, Float64, Int16, false},
		{"sb17", int(-42), int16(-42), Int, Int16, true},
		{"sb18", int8(-42), int16(-42), Int8, Int16, true},
		{"sb19", int32(-42), int16(-42), Int32, Int16, true},
		{"sb20", int64(-42), int16(-42), Int64, Int16, true},
		{"sb21", int(-32769), nil, Int, Int16, false},
		{"sb22", int32(-32769), nil, Int32, Int16, false},
		{"sb23", int64(-32769), nil, Int64, Int16, false},
		{"sb24", rt("-84/2"), int16(-42), Rat, Int16, true},
		{"sb25", rt("-84/11"), nil, Rat, Int16, false},
		{"sb26", rt("-32769"), nil, Rat, Int16, false},
		{"sb27", uint(42), int16(42), Uint, Int16, true},
		{"sb28", uint8(42), int16(42), Uint8, Int16, true},
		{"sb29", uint16(42), int16(42), Uint16, Int16, true},
		{"sb30", uint32(42), int16(42), Uint32, Int16, true},
		{"sb31", uint64(42), int16(42), Uint64, Int16, true},
		{"sb32", uint(32768), nil, Uint, Int16, false},
		{"sb33", uint16(32768), nil, Uint16, Int16, false},
		{"sb34", uint32(32768), nil, Uint32, Int16, false},
		{"sb35", uint64(32768), nil, Uint64, Int16, false},
		{"sb36", "-42", int16(-42), String, Int16, true},
		{"sb37", "$4,200", int16(4200), String, Int16, true},
		{"sb38", "42.42", nil, String, Int16, false},
		{"sb39", "-32769", nil, String, Int16, false},

		// int32
		{"sc01", int32(-42), int32(-42), Int32, Int32, true},
		{"sc02", fl("-42"), int32(-42), BigFloat, Int32, true},
		{"sc03", fl("-42.42"), nil, BigFloat, Int32, false},
		{"sc04", fl("-2147483649"), nil, BigFloat, Int32, false},
		{"sc05", in("-42"), int32(-42), BigInt, Int32, true},
		{"sc06", in("-2147483649"), nil, BigInt, Int32, false},
		{"sc07", complex128(-42), int32(-42), Complex, Int32, true},
		{"sc08", complex128(-42.42), nil, Complex, Int32, false},
		{"sc09", complex128(-42 + 42i), nil, Complex, Int32, false},
		{"sc10", complex128(-2147483649), nil, Complex, Int32, false},
		{"sc11", dc("-42"), int32(-42), Decimal, Int32, true},
		{"sc12", dc("-42.42"), nil, Decimal, Int32, false},
		{"sc13", dc("-2147483649"), nil, Decimal, Int32, false},
		{"sc14", float64(-42), int32(-42), Float64, Int32, true},
		{"sc15", float64(-42.42), nil, Float64, Int32, false},
		{"sc16", float64(-2147483649), nil, Float64, Int32, false},
		{"sc17", int(-42), int32(-42), Int, Int32, true},
		{"sc18", int8(-42), int32(-42), Int8, Int32, true},
		{"sc19", int16(-42), int32(-42), Int16, Int32, true},
		{"sc20", int64(-42), int32(-42), Int64, Int32, true},
		{"sc21", int(-2147483649), nil, Int, Int32, false},
		{"sc22", int64(-2147483649), nil, Int64, Int32, false},
		{"sc23", rt("-84/2"), int32(-42), Rat, Int32, true},
		{"sc24", rt("-84/11"), nil, Rat, Int32, false},
		{"sc25", rt("-2147483649"), nil, Rat, Int32, false},
		{"sc26", uint(42), int32(42), Uint, Int32, true},
		{"sc27", uint8(42), int32(42), Uint8, Int32, true},
		{"sc28", uint16(42), int32(42), Uint16, Int32, true},
		{"sc29", uint32(42), int32(42), Uint32, Int32, true},
		{"sc30", uint64(42), int32(42), Uint64, Int32, true},
		{"sc31", uint(2147483648), nil, Uint, Int32, false},
		{"sc32", uint32(2147483648), nil, Uint32, Int32, false},
		{"sc33", uint64(2147483648), nil, Uint64, Int32, false},
		{"sc34", "-42", int32(-42), String, Int32, true},
		{"sc35", "$42,000", int32(42000), String, Int32, true},
		{"sc36", "42.42", nil, String, Int32, false},
		{"sc37", "-2147483649", nil, String, Int32, false},

		// int64
		{"sd01", int64(-42), int64(-42), Int64, Int64, true},
		{"sd02", fl("-42"), int64(-42), BigFloat, Int64, true},
		{"sd03", fl("-42.42"), nil, BigFloat, Int64, false},
		{"sd04", fl("-9223372036854775809"), nil, BigFloat, Int64, false},
		{"sd05", in("-42"), int64(-42), BigInt, Int64, true},
		{"sd06", in("-9223372036854775809"), nil, BigInt, Int64, false},
		{"sd07", complex128(-42), int64(-42), Complex, Int64, true},
		{"sd08", complex128(-42.42), nil, Complex, Int64, false},
		{"sd09", complex128(-42 + 42i), nil, Complex, Int64, false},
		{"sd10", complex128(-9223372036854775809), nil, Complex, Int64, false},
		{"sd11", dc("-42"), int64(-42), Decimal, Int64, true},
		{"sd12", dc("-42.42"), nil, Decimal, Int64, false},
		{"sd13", dc("-9223372036854775809"), nil, Decimal, Int64, false},
		{"sd14", float64(-42), int64(-42), Float64, Int64, true},
		{"sd15", float64(-42.42), nil, Float64, Int64, false},
		{"sd16", float64(-9223372036854775809), nil, Float64, Int64, false},
		{"sd17", int(-42), int64(-42), Int, Int64, true},
		{"sd18", int8(-42), int64(-42), Int8, Int64, true},
		{"sd19", int16(-42), int64(-42), Int16, Int64, true},
		{"sd20", int32(-42), int64(-42), Int32, Int64, true},
		{"sd23", rt("-84/2"), int64(-42), Rat, Int64, true},
		{"sd24", rt("-84/11"), nil, Rat, Int64, false},
		{"sd25", rt("-9223372036854775809"), nil, Rat, Int64, false},
		{"sd26", uint(42), int64(42), Uint, Int64, true},
		{"sd27", uint8(42), int64(42), Uint8, Int64, true},
		{"sd28", uint16(42), int64(42), Uint16, Int64, true},
		{"sd29", uint32(42), int64(42), Uint32, Int64, true},
		{"sd30", uint64(42), int64(42), Uint64, Int64, true},
		{"sd31", uint(9223372036854775808), nil, Uint, Int64, false},
		{"sd33", uint64(9223372036854775808), nil, Uint64, Int64, false},
		{"sd34", "-42", int64(-42), String, Int64, true},
		{"sd35", "$42,000", int64(42000), String, Int64, true},
		{"sd36", "42.42", nil, String, Int64, false},
		{"sd37", "-9223372036854775809", nil, String, Int64, false},

		// Rat
		{"rt01", rt("42"), rt("42"), Rat, Rat, true},
		{"rt02", fl("42.42"), rt("42 21/50"), BigFloat, Rat, true},
		{"rt03", fl("1e100"), rt(big1 + "/1"), BigFloat, Rat, true},
		{"rt04", in("42"), rt("42"), BigInt, Rat, true},
		{"rt05", complex128(42), rt("42"), Complex, Rat, true},
		{"rt06", complex128(42.42), rt("42 21/50"), Complex, Rat, true},
		{"rt07", complex128(42 + 42i), nil, Complex, Rat, false},
		{"rt08", dc("42.42"), rt("42 21/50"), Decimal, Rat, true},
		{"rt09", dc("1e100"), rt(big1 + "/1"), Decimal, Rat, true},
		{"rt10", float64(42.42), rt("42 21/50"), Float64, Rat, true},
		{"rt11", float64(1e100), rt(big1 + "/1"), Float64, Rat, true},
		{"rt12", int(42), rt("42"), Int, Rat, true},
		{"rt13", int8(42), rt("42"), Int8, Rat, true},
		{"rt14", int16(42), rt("42"), Int16, Rat, true},
		{"rt15", int32(42), rt("42"), Int32, Rat, true},
		{"rt16", int64(42), rt("42"), Int64, Rat, true},
	}

	for _, test := range tests {
		t.Run(test.id, func(t *testing.T) {
			out, from, ok := test.to.From(s, test.in)
			if (ok && !Equal(out, test.out)) || from != test.from || ok != test.ok {
				t.Fatalf("\n have: %v %v %v \n want: %v %v %v", Format(out), from.AppName(), ok, Format(test.out), test.from.AppName(), test.ok)
			}
		})
	}
}

func TestRatParse(t *testing.T) {
	s := state.New()
	rt := func(v string) *big.Rat { return Rat.MustParse(s, v) }

	tests := []struct {
		in  string
		out *big.Rat
		ok  bool
	}{
		{"2", rt("2"), true},
		{"0.5", rt("1/2"), true},
		{"1/2", rt("1/2"), true},
		{"-1/2", rt("-1/2"), true},
		{"2 1/2", rt("5/2"), true},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out, ok := Rat.Parse(s, test.in)
			if (ok && !Equal(out, test.out)) || ok != test.ok {
				t.Fatalf("\n have: %v %v \n want: %v %v", Format(out), ok, Format(test.out), test.ok)
			}
		})
	}
}
