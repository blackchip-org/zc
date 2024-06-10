package types

import (
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6"
)

func TestString(t *testing.T) {
	tests := []struct {
		val  any
		str  string
		name string
	}{
		{big.NewInt(1234), "1234", "BigInt"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			str := zc.String(test.val)
			if str != test.str {
				t.Fatalf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}
