package types

import (
	"testing"

	"github.com/cockroachdb/apd/v3"
)

func TestDecFormat(t *testing.T) {
	c := apd.BaseContext.WithPrecision(16)

	var d66, d2, d33 apd.Decimal
	d66.SetString("6.6")
	d2.SetString("2")
	c.Quo(&d33, &d66, &d2)

	tests := []struct {
		num *apd.Decimal
		str string
	}{
		{&d33, "3.3"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			str := Dec.Format(test.num)
			if str != test.str {
				t.Errorf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}
