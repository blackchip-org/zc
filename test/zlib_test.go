package test

import (
	"testing"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/app"
)

func TestZlib(t *testing.T) {
	t.Skip("needs refactoring")
	c := app.NewDefaultCalc()
	//c.Trace = true
	fileName := "zc:test/suite.zc"
	src, err := zc.LoadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Eval(fileName, src); err != nil {
		t.Fatal(err)
	}
	b, err := c.Env.Stack.PopBool()
	if err != nil {
		t.Fatal(err)
	}
	if !b {
		t.Fatalf("test suite failed")
	}
}
