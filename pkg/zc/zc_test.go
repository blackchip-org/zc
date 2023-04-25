package zc

import (
	"fmt"
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

func TestClamp(t *testing.T) {
	tests := []struct {
		src  int
		want int
	}{
		{-5, 0},
		{0, 0},
		{5, 5},
		{10, 10},
		{15, 10},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.src), func(t *testing.T) {
			have := Clamp(test.src, 0, 10)
			if have != test.want {
				t.Errorf("\n have: %v \n want: %v", have, test.want)
			}
		})
	}
}
