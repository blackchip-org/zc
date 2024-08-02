package test

import (
	"testing"

	"github.com/blackchip-org/zc/v6/app"
)

func TestRepr(t *testing.T) {
	c := app.NewCalcTester(t)

	c.Eval("42 hex")
	c.AssertStack("0x2a")

	c.Eval("2 add")
	c.AssertStack("44")
}
