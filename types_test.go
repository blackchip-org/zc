package zc

import (
	"math/big"
	"reflect"
	"testing"
)

func TestFrom(t *testing.T) {
	bi42 := big.NewInt(42)

	tests := []struct {
		in   any
		out  any
		from Type
		to   Type
		ok   bool
	}{
		{big.NewInt(42), bi42, BigInt, BigInt, true},
		{big.NewFloat(42), bi42, BigFloat, BigInt, true},
		{int(42), bi42, Int, BigInt, true},
		{int8(42), bi42, Int8, BigInt, true},
		{int16(42), bi42, Int16, BigInt, true},
		{int32(42), bi42, Int32, BigInt, true},
		{int64(42), bi42, Int64, BigInt, true},
		{uint(42), bi42, Uint, BigInt, true},
		{uint8(42), bi42, Uint8, BigInt, true},
		{uint16(42), bi42, Uint16, BigInt, true},
		{uint32(42), bi42, Uint32, BigInt, true},
		{uint64(42), bi42, Uint64, BigInt, true},
	}

	for _, test := range tests {
		t.Run(goName(test.in), func(t *testing.T) {
			out, from, ok := test.to.From(test.in)
			if !reflect.DeepEqual(out, test.out) || from != test.from || ok != test.ok {
				t.Errorf("\n have: %v %v %v \n want: %v %v %v", out, from.Name(), ok, test.out, test.from.Name(), test.ok)
			}
		})
	}
}
