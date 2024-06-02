package types

import "testing"

func TestGoName(t *testing.T) {
	var (
		a int
		b *int
		c **int
	)

	tests := []struct {
		v    any
		name string
	}{
		{a, "int"},
		{b, "*int"},
		{c, "**int"},
	}

	for _, test := range tests {
		name := GoName(test.v)
		if name != test.name {
			t.Fatalf("\n have: %v \n want: %v", name, test.name)
		}
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
