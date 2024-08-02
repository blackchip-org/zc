package app

import (
	"testing"

	"github.com/blackchip-org/zc/v6"
)

func TestStack(t *testing.T) {
	c := NewCalc(nil)
	tests := []struct {
		op    func()
		stack string
	}{
		{func() { zc.String.Push(c, "1") }, "1"},
		{func() { zc.String.Push(c, "2") }, "1 | 2"},
		{func() { zc.String.Push(c, "3") }, "1 | 2 | 3"},
		{func() { zc.String.Pop(c) }, "1 | 2"},
		{func() { zc.String.Push(c, "4") }, "1 | 2 | 4"},
		{func() { zc.String.Pop(c) }, "1 | 2"},
		{func() { zc.String.Pop(c) }, "1"},
		{func() { zc.String.Pop(c) }, ""},
	}

	for _, test := range tests {
		test.op()
		stack := c.String()
		if stack != test.stack {
			t.Fatalf("\n have: %v \n want: %v", stack, test.stack)
		}
	}
}
