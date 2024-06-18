package app

import (
	"reflect"
	"testing"

	"github.com/blackchip-org/zc/v6"
)

type CalcTester struct {
	calc *Calc
	t    *testing.T
}

func NewCalcTester(t *testing.T) *CalcTester {
	return &CalcTester{
		calc: NewCalc(),
		t:    t,
	}
}

func (c *CalcTester) Eval(line string) {
	c.t.Helper()
	c.t.Logf("%v> %v\n", zc.ProgName, line)
	c.calc.Eval(line)
	if c.calc.Err != nil {
		c.t.Logf("(!) %v\n", c.calc.Err)
	}
	if c.calc.Info != "" {
		c.t.Logf("(info) %v", c.calc.Info)
	}
	c.t.Logf("%v\n", c.calc.Stack.String())
}

func (c *CalcTester) AssertStack(vals ...any) {
	c.t.Helper()

	fmtWant := zc.FormatList(vals...)
	fmtHave := c.calc.Stack.String()

	if c.calc.Err != nil {
		c.t.Fatalf("(FAIL) unexpected error")
	}
	if c.calc.Info != "" {
		c.t.Fatalf("(FAIL) unexpected info")
	}
	if !reflect.DeepEqual(fmtHave, fmtWant) {
		c.t.Fatalf("(FAIL) expected: %v", fmtWant)
	}
}

func (c *CalcTester) AssertError(err string) {
	c.t.Helper()
	if c.calc.Err == nil || c.calc.Err.Error() != err {
		c.t.Fatalf("(FAIL) expected error: %v", err)
	}
}

func (c *CalcTester) AssertInfo(info string) {
	c.t.Helper()
	if c.calc.Info != info {
		c.t.Fatalf("(FAIL) expected info: %v", info)
	}
}
