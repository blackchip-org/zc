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

	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			have, err := ParseBigInt(test.s)
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

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			_, err := ParseBigInt(test)
			if err == nil {
				t.Errorf("expecting error")
			}
		})
	}
}

func TestParseRadix(t *testing.T) {
	tests := []struct {
		s    string
		want int
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
			have := ParseRadix(test.s)
			if have != test.want {
				t.Errorf("\n have: %v \n want: %v", have, test.want)
			}
		})
	}
}
