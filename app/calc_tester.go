package app

import (
	"fmt"
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
	if c.Calc.Info != "" {
		c.t.Logf("(info) %v", c.Calc.Info)
	}
	c.t.Logf("%v\n", c.Calc.Stack.String())
	c.t.Logf("\t%v\n", c.StackTypes())
}

func (c *CalcTester) AssertStack(vals ...any) {
	c.t.Helper()

	fmtWant := zc.FormatList(vals...)
	fmtHave := c.Calc.Stack.String()

	if c.Calc.Err != nil {
		c.t.Fatalf("(FAIL) unexpected error")
	}
	if c.Calc.Info != "" {
		c.t.Fatalf("(FAIL) unexpected info")
	}
	if !reflect.DeepEqual(fmtHave, fmtWant) {
		fmt.Printf("\n have: %v \n want: %v\n", fmtHave, fmtWant)
		c.t.Fatalf("(FAIL) expected: %v", fmtWant)
	}
}

func (c *CalcTester) AssertError(err string) {
	c.t.Helper()
	if c.Calc.Err == nil || c.Calc.Err.Error() != err {
		c.t.Fatalf("(FAIL) expected error: %v", err)
	}
}

func (c *CalcTester) AssertInfo(info string) {
	c.t.Helper()
	if c.Calc.Info != info {
		c.t.Fatalf("(FAIL) expected info: %v", info)
	}
}

func (c *CalcTester) StackTypes() string {
	var types []any
	items := c.Calc.Items()
	for _, item := range items {
		types = append(types, reflect.TypeOf(item.Val).String())
	}
	return zc.FormatList(types...)
}
