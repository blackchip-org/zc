package test

import (
	"testing"

	"github.com/blackchip-org/zc/app"
)

func TestZlib(t *testing.T) {
	c := app.NewDefaultCalc()
	fileName := "zc:test/suite.zc"
	src, err := c.LoadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Eval(fileName, src); err != nil {
		t.Fatal(err)
	}
	b, err := c.PopBool()
	if err != nil {
		t.Fatal(err)
	}
	if !b {
		t.Fatalf("test suite failed")
	}
}
