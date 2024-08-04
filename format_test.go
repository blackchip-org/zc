package zc

import (
	"testing"
)

func TestEscapeString(t *testing.T) {
	tests := []struct {
		str string
		esc string
	}{
		{"1", "1"},
		{"\a", "\\a"},
		{"\b", "\\b"},
		{"\f", "\\f"},
		{"\n", "\\n"},
		{"\r", "\\r"},
		{"\t", "\\t"},
		{"\v", "\\v"},
		{"\x01", "\\x01"},
		{"\ufffe", "\\ufffe"},
		{"\U00100000", "\\U00100000"},
		{"1\a\b\f2", "1\\a\\b\\f2"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			esc := EscapeString(test.str)
			if esc != test.esc {
				t.Fatalf("\n have: %v \n want: %v", esc, test.esc)
			}
		})
	}
}

func TestFormatList(t *testing.T) {
	have := FormatList([]string{"1", "23", "45"})
	want := "1 | 23 | 45"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestQuote(t *testing.T) {
	tests := []struct {
		str    string
		quoted string
	}{
		{`123`, `123`},
		{`12 34`, `"12 34"`},
		{"12\n34", `12\n34`},
	}

	for _, test := range tests {
		t.Run(test.quoted, func(t *testing.T) {
			quoted := Quote(test.str)
			if quoted != test.quoted {
				t.Fatalf("\n have: %v \n want: %v", quoted, test.quoted)
			}
		})
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
