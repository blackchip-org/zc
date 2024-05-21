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
