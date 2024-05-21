package test

import (
	"testing"

	"github.com/blackchip-org/zc/v6/calc"
	"github.com/blackchip-org/zc/v6/ops"
)

func BenchmarkAddNative(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := 0
		for i := 0; i < 100; i++ {
			a = a + 2
		}
		if a != 200 {
			b.Error(a)
		}
	}
}

func BenchmarkAddIntArch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c := calc.New()
		c.Push(0)
		for i := 0; i < 100; i++ {
			c.Push(2)
			c.Do(ops.AddIntArch)
		}
		a := c.PopString()
		if a != "200" {
			b.Error(a)
		}
	}
}

func BenchmarkAddInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		c := calc.New()
		c.Push(0)
		for i := 0; i < 100; i++ {
			c.Push(2)
			c.Do(ops.AddInt)
		}
		a := c.PopString()
		if a != "200" {
			b.Error(a)
		}
	}
}

func fib(b *testing.B, max int, ans string) {
	for n := 0; n < b.N; n++ {
		c := calc.New()

		c.Push(1)
		c.Push(1)
		for i := 3; i <= max; i++ {
			c.Do(ops.Dup, ops.Down, ops.AddInt)
		}
		a := c.PopString()
		if a != ans {
			b.Error(a)
		}
	}
}

func BenchmarkFib10(b *testing.B) {
	fib(b, 10, "55")
}

func BenchmarkFib100(b *testing.B) {
	fib(b, 100, "354224848179261915075")
}
func BenchmarkFib1000(b *testing.B) {
	fib(b, 1000, "43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875")
}
