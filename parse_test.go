package zc

import "testing"

func TestPreParseNumber(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1234", "1234"},
		{"1,234", "1234"},
		{"1_234", "1234"},
		{"1 234", "1234"},
		{"$1234", "1234"},
		{"0x10ab", "0x10ab"},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := PreParseNumber(test.in)
			if out != test.out {
				t.Fatalf("\n have: %v \n want: %v", out, test.out)
			}
		})
	}
}

func TestPreParseFloat(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1234", "1234"},
		{"1,234", "1234"},
		{"1_234", "1234"},
		{"1 234", "1234"},
		{"$1234", "1234"},
		{"1234x1056", "1234e56"},
		{"3.041409320×1064", "3.041409320e64"},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := PreParseFloat(test.in)
			if out != test.out {
				t.Fatalf("\n have: %v \n want: %v", out, test.out)
			}
		})
	}
}
