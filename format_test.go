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
