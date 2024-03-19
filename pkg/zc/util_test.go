package zc

import "testing"

func TestRemoveTrailingZeros(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"12.345", "12.345"},
		{"12.340", "12.34"},
		{"12.300", "12.3"},
		{"12.000", "12"},
		{"10", "10"},
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
