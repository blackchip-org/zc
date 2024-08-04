package test

import (
	"testing"

	"github.com/blackchip-org/zc/v6/app"
)

func TestDecimalOverflow(t *testing.T) {
	c := app.NewCalcTester(t)

	c.Eval("1e100001 sci")
	c.AssertError("scientific.notation: overflow: 1e100001")
}
