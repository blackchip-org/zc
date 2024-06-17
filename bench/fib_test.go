package bench

import (
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6/pkg/calc"
)

const fib1000 = "43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875"

func fibTest(n int) *big.Int {
	c := calc.NewBigInt()
	c.PushInt(1)
	c.PushInt(1)
	for i := 3; i <= n; i++ {
		c.Dup()
		c.Rotate()
		c.Add()
	}
	return c.Pop()
}

func fibTestNative(n int) *big.Int {
	var f0, f1, f2 big.Int
	f0.SetInt64(int64(1))
	f1.SetInt64(int64(1))
	for i := 3; i <= n; i++ {
		f2.Add(&f0, &f1)
		f0.Set(&f1)
		f1.Set(&f2)
	}
	return &f2
}

func TestFib(t *testing.T) {
	result := fibTest(1000).String()
	if result != fib1000 {
		t.Fatalf("\n have: %v \n want: %v", result, fib1000)
	}
}

func TestFibNative(t *testing.T) {
	result := fibTestNative(1000).String()
	if result != fib1000 {
		t.Fatalf("\n have: %v \n want: %v", result, fib1000)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibTest(1000)
	}
}

func BenchmarkFibNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibTestNative(1000)
	}
}
