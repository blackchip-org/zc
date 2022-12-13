package zc

import (
	"log"
	"testing"
)

func TestInterpolate(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{"`answer`", "42"},
		{"`list`", "one  two"},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			calc, err := NewCalcWithConfig(Config{})
			if err != nil {
				log.Panicf("unexpected error: %v", err)
			}
			calc.Define("answer").Set("42")
			stack := calc.Define("list")
			stack.Push("one")
			stack.Push("two")

			have, err := calc.Interpolate(test.src)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if have != test.want {
				t.Errorf("\n have %v \n want %v", have, test.want)
			}
		})

	}

}
