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

func AssertStack(t *testing.T, c *Calc, want ...any) {
	t.Helper()

	fmtWant := FormatList(want...)
	fmtHave := FormatStack(c.Stack)

	if c.Err != nil {
		t.Fatalf("unexpected error: %v\nstack: %v", c.Err, fmtHave)
	}
	if c.Info != "" {
		t.Fatalf("unexpected info: %v\n", c.Info)
	}
	if !reflect.DeepEqual(fmtHave, fmtWant) {
		t.Fatalf("\n have: %v \n want: %v", fmtHave, fmtWant)
	}
}

func AssertError(t *testing.T, c *Calc, want string) {
	t.Helper()
	if c.Err == nil {
		t.Fatalf("expected error: %v", want)
	}
	if c.Err.Error() != want {
		t.Fatalf("\n have error: %v \n want error: %v", c.Err.Error(), want)
	}
}

func AssertInfo(t *testing.T, c *Calc, info string) {
	t.Helper()
	if c.Info != info {
		t.Fatalf("\n have info: %v \n want info: %v", c.Info, info)
	}
}
