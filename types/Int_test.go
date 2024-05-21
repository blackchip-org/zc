package types

import (
	"fmt"
	"math/big"
	"testing"
)

func TestIntDup(t *testing.T) {
	one := big.NewInt(1)
	a := big.NewInt(2)
	b := Int.Dup(a).(*big.Int)
	b.Add(b, one)

	if b.Int64() != 3 {
		t.Errorf("\n have: %v \n want: %v", b.Int64(), 3)
	}
	if a.Int64() != 2 {
		t.Errorf("\n have: %v \n want: %v", b.Int64(), 2)
	}
}

func TestTo(t *testing.T) {
	tests := []struct {
		src  any
		want *big.Int
	}{
		{"12", big.NewInt(12)},
		{12, big.NewInt(12)},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.want)
		t.Run(name, func(t *testing.T) {
			conv, ok := Int.To(test.src)
			if !ok {
				t.Fatalf("conversion failed")
			}
			have := Int.As(conv)
			if have.Cmp(test.want) != 0 {
				t.Errorf("\n have: %v \n want: %v", have, test.want)
			}
		})
	}
}
