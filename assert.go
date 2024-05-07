package zc

import (
	"reflect"
	"testing"
)

func AssertStack(t *testing.T, c *Calc, want ...string) {
	t.Helper()
	if c.Err != nil {
		t.Fatalf("unexpected error: %v\nstack: %v", c.Err, FormatStack(c.Stack()))
	}
	have := c.StackStrings()
	if !reflect.DeepEqual(have, want) {
		t.Fatalf("\n have: %v \n want: %v", FormatStackValues(have), FormatStackValues(want))
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
