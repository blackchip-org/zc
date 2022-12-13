package zc

import (
	"math/big"
	"testing"
)

func TestParseBigInt(t *testing.T) {
	tests := []struct {
		s    string
		want *big.Int
	}{
		{"1234", new(big.Int).SetInt64(1234)},
		{"0xffd2", new(big.Int).SetInt64(65490)},
	}

	c := NewCalc()
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			have, err := c.ParseBigInt(test.s)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if have.Cmp(test.want) != 0 {
				t.Errorf("\n have: %v \n want: %v", have, test.want)
			}
		})
	}
}

func TestParseBigIntInvalid(t *testing.T) {
	tests := []string{
		"abcd",
	}

	c := NewCalc()
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			_, err := c.ParseBigInt(test)
			if err == nil {
				t.Errorf("expecting error")
			}
		})
	}
}

func TestParseRadix(t *testing.T) {
	tests := []struct {
		s     string
		radix int
	}{
		{"1234", 10},
		{"abcd", 10},
		{"0xabcd", 16},
		{"0XABCD", 16},
		{"0b0101", 2},
		{"0o755", 8},
	}

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			radix := ParseRadix(test.s)
			if radix != test.radix {
				t.Errorf("\n have: %v \n want: %v", radix, test.radix)
			}
		})
	}
}
