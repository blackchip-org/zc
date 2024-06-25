package bench

import (
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app"
	"github.com/blackchip-org/zc/v6/pkg/calc"
	"github.com/cockroachdb/apd/v3"
)

func distTestZcEval(c *app.Calc) *apd.Decimal {
	c.Eval("5 2 sub 2 pow")
	c.Eval("7 3 sub 2 pow")
	c.Eval("add sqrt")
	return zc.Decimal.As(c.Pop().Val)
}

// func distTestZc(c *app.Calc) *big.Int {
// 	c.PushVal(5, 2)
// 	c.Do(ops.SubBigInt)
// 	c.PushVal(2)
// 	c.Do(ops.PowBigInt)

// 	c.PushVal(7, 3)
// 	c.Do(ops.SubBigInt)
// 	c.PushVal(2)
// 	c.Do(ops.PowBigInt)

// 	c.Do(ops.AddBigInt)
// 	c.Do(ops.Sqrt)

// 	return zc.BigInt.As(c.Pop().Val)
// }

func distTestCalc(c *calc.BigInt) *big.Int {
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

func TestDistZcEval(t *testing.T) {
	c := app.NewCalc()
	r := zc.Format(distTestZcEval(c))
	if r != "5" {
		t.Fatalf("\n have: %v \n want: %v", r, "5")
	}
}

// func TestDistZc(t *testing.T) {
// 	c := app.NewCalc()
// 	r := distTestZc(c).String()
// 	if r != "5" {
// 		t.Fatalf("\n have: %v \n want: %v", r, "5")
// 	}
// }

func TestDistCalc(t *testing.T) {
	c := calc.NewBigInt()
	r := distTestCalc(c).String()
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

func BenchmarkDistZcEval(b *testing.B) {
	c := app.NewCalc()
	for i := 0; i < b.N; i++ {
		distTestZcEval(c)
	}
}

// func BenchmarkDistZc(b *testing.B) {
// 	c := app.NewCalc()
// 	for i := 0; i < b.N; i++ {
// 		distTestZc(c)
// 	}
// }

func BenchmarkDistCalc(b *testing.B) {
	c := calc.NewBigInt()
	for i := 0; i < b.N; i++ {
		distTestCalc(c)
	}
}

func BenchmarkDistNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		distTestNative()
	}
}
