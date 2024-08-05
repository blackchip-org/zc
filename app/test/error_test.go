package test

import (
	"testing"

	"github.com/blackchip-org/zc/v6/app"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		expr string
		err  string
	}{
		{"add", "add: not enough arguments"},
		{"42 add", "add: not enough arguments"},
		{"1 atanh", "atanh: +infinity"},
		{"-1 atanh", "atanh: -infinity"},
		{"2 atanh", "atanh: not a number"},
	}

	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			c := app.NewCalcTester(t)
			c.Eval(test.expr)
			c.AssertError(test.err)
		})
	}
}
