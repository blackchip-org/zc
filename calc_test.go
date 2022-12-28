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

	calc, _ := NewCalc(Config{})
	for _, test := range tests {
		t.Run(test.s, func(t *testing.T) {
			have, err := calc.ParseBigInt(test.s)
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

	calc, _ := NewCalc(Config{})
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			_, err := calc.ParseBigInt(test)
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

func TestFormatNumberString(t *testing.T) {
	confDefault := Config{
		IntPat: ",000",
		Point:  '.',
	}
	confFR := Config{
		IntPat: ".000",
		Point:  ',',
	}
	confBin := Config{
		IntPat: "__0000_0000",
	}
	confEmpty := Config{}

	tests := []struct {
		in   string
		want string
		conf Config
	}{
		{"1", "1", confDefault},
		{"123", "123", confDefault},
		{"1234", "1,234", confDefault},
		{"123456", "123,456", confDefault},
		{"1234567", "1,234,567", confDefault},
		{"-123", "-123", confDefault},
		{".123", ".123", confDefault},
		{".12345", ".12345", confDefault},
		{"1234567.8901", "1,234,567.8901", confDefault},
		{"1234567.8901", "1234567.8901", confEmpty},
		{"1234567.8901", "1.234.567,8901", confFR},
		{"11110000", "1111_0000", confBin},
		{"1111000011110000", "1111_0000__1111_0000", confBin},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			calc, _ := NewCalc(test.conf)
			have := calc.FormatNumberString(test.in)
			if have != test.want {
				t.Errorf("\n have: %v \n want: %v", have, test.want)
			}
		})
	}
}
