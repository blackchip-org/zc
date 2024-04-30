package test

import (
	"testing"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/calc"
	"github.com/blackchip-org/zc/v6/ops"
)

func TestIntAdd(t *testing.T) {
	c := calc.New()

	c.Push(2)
	c.Push(3)
	c.Do(ops.AddInt)

	if c.Err != nil {
		t.Fatalf("unexpected error: %v", c.Err)
	}

	have := c.PopString()
	want := "5"

	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func TestFib(t *testing.T) {
	c := fib(10)
	if c.Err != nil {
		t.Fatalf("unexpected error: %v", c.Err)
	}

	have := c.PopString()
	want := "55"

	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func fib(n int) *zc.Calc {
	c := calc.New()

	c.Push(1)
	c.Push(1)
	for i := 3; i <= n; i++ {
		c.Do(ops.Dup, ops.Down, ops.AddInt)
	}
	return c
}

func BenchmarkFib100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(100)
	}
}

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
