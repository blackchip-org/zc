package test

import (
	"testing"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/modules"
)

func TestSuites(t *testing.T) {
	calc := zc.NewCalc(modules.Prelude)
	tests, _ := zc.Static.ReadFile("test/suites.zc")
	err := calc.Eval(tests)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = calc.Global("__testing.failed").Get()
	if err == nil {
		t.Fatalf("test suite failed")
	}
}
