package zc

import (
	"reflect"
	"testing"
)

type Expect struct {
	Input  string   `yaml:"i"`
	Output []string `yaml:"o"`
	Error  string   `yaml:"error"`
	Info   string   `yaml:"info"`
}

type Test struct {
	Name string   `yaml:"name"`
	Test []Expect `yaml:"test"`
}

func NoOp(_ *OpEnv) {}

type CalcTester struct {
	calc *Calc
	t    *testing.T
}

func NewCalcTester(calc *Calc, t *testing.T) CalcTester {
	return CalcTester{
		calc: calc,
		t:    t,
	}
}

func (c *CalcTester) Eval(line string) {
	c.t.Helper()
	c.t.Logf("%v > %v\n", ProgName, line)
	c.calc.Eval(line)
	if c.calc.Err != nil {
		c.t.Logf("(!) %v\n", c.calc.Err)
	}
	if c.calc.Info != "" {
		c.t.Logf("(info) %v", c.calc.Info)
	}
	c.t.Logf("%v\n", FormatStack(c.calc.Stack))
}

func (c *CalcTester) AssertStack(vals ...any) {
	c.t.Helper()

	fmtWant := FormatList(vals...)
	fmtHave := FormatStack(c.calc.Stack)

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
