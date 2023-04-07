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
		{"12+34i", Complex},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			v, ok := ParseNumber(test.text)
			if !ok {
				t.Fatalf("unable to parse: %v", test.text)
			}
			if v.Type() != test.type_ {
				t.Errorf("\n have: %v \n want: %v", v.Type(), test.type_)
			}
		})
	}
}
