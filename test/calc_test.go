package test

import (
	"testing"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/calc"
)

func TestIntAdd(t *testing.T) {
	c := calc.New()

	c.Push(2)
	c.Push(3)
	c.Do("addi")

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
		c.Do("dup")
		c.Do("down")
		c.Do("addi")
	}
	return c
}

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fib(100)
	}
}
