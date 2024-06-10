package test

import (
	"testing"

	calc5 "github.com/blackchip-org/zc/v5/pkg/calc"
	"github.com/blackchip-org/zc/v6/calc"
)

func BenchmarkFactorialV5Native(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := calc5.New()
		c.Eval("1000 fact")
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := calc.New()
		c.Eval("1000 fact")
	}
}

func BenchmarkFactorialV5WithCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := calc5.New()
		c.Eval("1 1 1")
		for i := 0; i < 1000; i++ {
			c.Eval("mul up 1 add dup down")
		}
		c.Eval("drop")
	}
}
