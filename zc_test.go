package zc

import (
	"testing"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{"1234", "1234"},
		{"abcd", "'abcd'"},
		{"12 34", "'12 34'"},
		{"a", "'a'"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			have := Quote(test.src)
			if have != test.want {
				t.Errorf("\n have: %v \n want: %v", have, test.want)
			}
		})
	}
}
