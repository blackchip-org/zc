package types

import "testing"

func TestNumbers(t *testing.T) {
	tests := []struct {
		text  string
		type_ Type
	}{
		{"1234", BigInt},
		{"0x1234", BigInt},
		{"12.34", Decimal},
		{"12.34e5", Float},
		{"12.34e5d", Decimal},
		{"1.234f", Float},
		{"12+34i", Complex},
		{"1/2", Rational},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			g := Parse(test.text)
			if g.Type() != test.type_ {
				t.Errorf("\n have: %v \n want: %v", g.Type(), test.type_)
			}
		})
	}
}
