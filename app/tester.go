package app

import (
	"reflect"
	"testing"

	"github.com/blackchip-org/zc/v6"
)

type CalcTester struct {
	Calc *Calc
	t    *testing.T
}

func NewCalcTester(t *testing.T) *CalcTester {
	return &CalcTester{
		Calc: NewCalc(),
		t:    t,
	}
}

func (c *CalcTester) Eval(line string) {
	c.t.Helper()
	c.t.Logf("%v> %v\n", zc.ProgName, line)
	c.Calc.Eval(line)

	if c.Calc.Err != nil {
		c.t.Logf("(!) %v\n", c.Calc.Err)
	}
	if c.Calc.Notice != "" {
		c.t.Logf("(?) %v", c.Calc.Notice)
	}

	c.t.Logf("%v\n", c.Calc.String())
}

func (c *CalcTester) AssertStack(vals ...any) {
	c.t.Helper()

	fmtWant := zc.FormatList(zc.Strings(vals...))
	fmtHave := c.Calc.String()

	if c.Calc.Err != nil {
		c.t.Fatalf("(FAIL) unexpected error")
	}
	if c.Calc.Notice != "" {
		c.t.Fatalf("(FAIL) unexpected notice")
	}
	if !reflect.DeepEqual(fmtHave, fmtWant) {
		c.t.Fatalf("(FAIL) expected: %v", fmtWant)
	}
}

func (c *CalcTester) AssertError(err string) {
	c.t.Helper()
	if c.Calc.Err == nil || c.Calc.Err.Error() != err {
		c.t.Fatalf("(FAIL) expected error: %v", err)
	}
}

func (c *CalcTester) AssertNotice(notice string) {
	c.t.Helper()
	if c.Calc.Notice != notice {
		c.t.Fatalf("(FAIL) expected notice: %v", notice)
	}
}
