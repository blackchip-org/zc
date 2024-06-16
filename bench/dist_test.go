package bench

import (
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6/pkg/calc"
)

func distTest(c *calc.BigInt) *big.Int {
	c.PushInt(5, 2)
	c.Sub()
	c.PushInt(2)
	c.Pow()

	c.PushInt(7, 3)
	c.Sub()
	c.PushInt(2)
	c.Pow()

	c.Add()
	c.Sqrt()

	return c.Pop()
}

func distTestNative() *big.Int {
	var d1, d2, r big.Int

	two := big.NewInt(2)

	x := big.NewInt(5)
	y := big.NewInt(2)
	d1.Sub(x, y)
	d1.Exp(&d1, two, nil)

	x.SetInt64(7)
	y.SetInt64(3)
	d2.Sub(x, y)
	d2.Exp(&d2, two, nil)

	r.Add(&d1, &d2)
	r.Sqrt(&r)
	return &r
}

func TestDist(t *testing.T) {
	var c calc.BigInt
	r := distTest(&c).String()
	if r != "5" {
		t.Fatalf("\n have: %v \n want: %v", r, "5")
	}
}

func TestDistNative(t *testing.T) {
	r := distTestNative().String()
	if r != "5" {
		t.Fatalf("\n have: %v \n want: %v", r, "5")
	}
}

func BenchmarkDist(b *testing.B) {
	var c calc.BigInt
	for i := 0; i < b.N; i++ {
		distTest(&c)
	}
}

func BenchmarkDistNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		distTestNative()
	}
}
