package zc

import (
	"math/big"
	"testing"
)

func TestItemString(t *testing.T) {
	tests := []struct {
		item Item
		str  string
		name string
	}{
		{Item{Val: "foo"}, "foo", "value"},
		{Item{Val: 123, Meta: &Meta{Unit: "m"}}, "123m", "unit"},
		{Item{Val: 123, Meta: &Meta{Label: "label"}}, "123 # label", "label"},
		{Item{Val: 123, Meta: &Meta{Unit: "m", Label: "label"}}, "123m # label", "label"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			str := test.item.String()
			if str != test.str {
				t.Fatalf("\n have: %v \n want: %v", str, test.str)
			}
		})
	}
}

func TestStackPushPop(t *testing.T) {
	var s Stack

	AssertStack(t, s)

	s.Push(Item{Val: 1})
	AssertStack(t, s, 1)

	s.Push(Item{Val: 2})
	AssertStack(t, s, 1, 2)

	s.Pop()
	AssertStack(t, s, 1)

	s.Push(Item{Val: 3})
	s.Push(Item{Val: 4})
	AssertStack(t, s, 1, 3, 4)

	s.Pop()
	s.Pop()
	s.Pop()
	AssertStack(t, s)

	s.Push(Item{Val: 5})
	AssertStack(t, s, 5)

	s.Pop()
	_, ok := s.Pop()
	if ok {
		t.Fatalf("expected not ok")
	}
}

func BenchmarkTest(b *testing.B) {
	one := big.NewInt(1)
	for i := 0; i < b.N; i++ {
		var s Stack
		s.Push(Item{Val: big.NewInt(1)})
		s.Push(Item{Val: big.NewInt(1)})
		for j := 0; j < 1000; j++ {
			iy, _ := s.Pop()
			ix, _ := s.Pop()

			y := iy.Val.(*big.Int)
			x := ix.Val.(*big.Int)

			x.Mul(x, y)
			y.Add(y, one)
			s.Push(Item{Val: x})
			s.Push(Item{Val: y})
		}
		s.Pop()
	}
}
