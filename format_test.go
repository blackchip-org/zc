package zc

import (
	"testing"
)

func TestFormatList(t *testing.T) {
	have := FormatList([]string{"1", "23", "45"})
	want := "1 | 23 | 45"
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
