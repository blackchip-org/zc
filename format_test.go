package zc

import (
	"math/big"
	"testing"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		val any
		str string
	}{
		{1, "1"},
		{big.NewInt(2), "2"},
		{"3", "3"},
		{struct {
			A string
			B string
		}{A: "a", B: "b"}, "{a b}"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := Format(test.val)
			if str != test.str {
				t.Fatalf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

func TestFormatList(t *testing.T) {
	have := FormatList(1, "2", 3.3)
	want := "1 | 2 | 3.3"

	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestFormatExponent(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1", "1"},
		{"1e1", "1e1"},
		{"1E1", "1e1"},
		{"2e+1", "2e1"},
		{"1e+22", "1e22"},
		{"1e-22", "1e-22"},
		{"1e+09", "1e9"},
		{"1e0", "1e0"},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := FormatExponent(test.in)
			if out != test.out {
				t.Errorf("\n have: %v \n want: %v", out, test.out)
			}
		})
	}
}

func TestRemoveTrailingZeros(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"0", "0"},
		{"100", "100"},
		{"1.100", "1.1"},
		{"1.000", "1"},
		{"1.0001", "1.0001"},
		{"1.0001000", "1.0001"},
		{"1.1e+00", "1.1e+00"},
		{"1.100e+00", "1.1e+00"},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := RemoveTrailingZeros(test.in)
			if out != test.out {
				t.Errorf("\n have: %v \n want: %v", out, test.out)
			}
		})
	}
}
