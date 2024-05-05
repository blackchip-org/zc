package test

import (
	"testing"

	"github.com/blackchip-org/zc/v6/calc"
	"github.com/blackchip-org/zc/v6/ops"
)

func TestNotEnoughArgs(t *testing.T) {
	c := calc.New()
	c.Push(1)
	c.Do(ops.AddInt)

	if c.Err == nil {
		t.Fatal("expected error")
	}
	have := c.Err.Error()
	want := "not enough arguments, expected 2"

	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func TestArgsWrongKind(t *testing.T) {
	c := calc.New()
	c.Push(1)
	c.Push("foo")
	c.Do(ops.AddInt)

	if c.Err == nil {
		t.Fatal("expected error")
	}
	have := c.Err.Error()
	want := "invalid arguments, expected Int | Int"

	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func TestCannotConvert(t *testing.T) {
	c := calc.New()
	c.Push("foo")
	var ret int

	err := c.Pop(&ret)
	if err == nil {
		t.Fatal("expected error")
	}
	have := c.Err.Error()
	want := "cannot convert foo from Text to *int"

	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}
