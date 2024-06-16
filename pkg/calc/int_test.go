package calc

import (
	"testing"
)

// func TestBigIntAbs(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(-5)
// 	c.Abs()

// 	have := c.Top().Int64()
// 	want := int64(5)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

func TestBigIntAdd(t *testing.T) {
	c := NewBigInt()
	c.PushInt(6, 3)
	c.Add()

	have := c.Top().Int64()
	want := int64(9)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

// func TestBigIntAnd(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(0b1100)
// 	c.PushInt(0b1010)
// 	c.And()

// 	have := c.Top().Int64()
// 	want := int64(0b1000)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntAndNot(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(0b1100)
// 	c.PushInt(0b1010)
// 	c.AndNot()

// 	have := c.Top().Int64()
// 	want := int64(0b0100)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntBinomial(t *testing.T) {
// 	var c BigInt
// 	c.Binomial(9, 5)

// 	have := c.Top().Int64()
// 	want := int64(126)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntCmp(t *testing.T) {
// 	var c BigInt

// 	c.PushInt(1, 2)
// 	have := c.Cmp()
// 	want := -1
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntCmpAbs(t *testing.T) {
// 	var c BigInt

// 	c.PushInt(1, -2)
// 	have := c.CmpAbs()
// 	want := -1
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntDiv(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(-7)
// 	c.PushInt(2)
// 	c.Div()

// 	have := c.Top().Int64()
// 	want := int64(-4)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntDivMod(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(-7, 2)
// 	c.DivMod()

// 	have := c.Top().Int64()
// 	want := int64(1)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}

// 	have = c.Next().Int64()
// 	want = int64(-4)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntDup(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(3)
// 	c.Dup()
// 	c.Add()

// 	have := c.Top().Int64()
// 	want := int64(6)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntExp(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(6, 2, 0)
// 	c.Exp()

// 	have := c.Top().Int64()
// 	want := int64(36)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntGCD(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(1, 1, 8, 12)
// 	c.GCD()

// 	have := c.Top().Int64()
// 	want := int64(4)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntGCD2(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(8, 12)
// 	c.GCD2()

// 	have := c.Top().Int64()
// 	want := int64(4)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntLsh(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(0b10)
// 	c.Lsh(2)

// 	have := c.Top().Int64()
// 	want := int64(0b1000)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntMod(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(-7, 2)
// 	c.Mod()

// 	have := c.Top().Int64()
// 	want := int64(1)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntModInverse(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(7, 11)
// 	c.ModInverse()

// 	have := c.Top().Int64()
// 	want := int64(8)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntModSqrt(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(223, 17)
// 	c.ModSqrt()

// 	have := c.Top().Int64()
// 	want := int64(6)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

func TestBigIntMul(t *testing.T) {
	var c BigInt
	c.PushInt(6, 2)
	c.Mul()

	have := c.Top().Int64()
	want := int64(12)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

// func TestBigIntMulRange(t *testing.T) {
// 	var c BigInt
// 	c.MulRange(1, 10)

// 	have := c.Top().Int64()
// 	want := int64(3628800)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntNeg(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(6)
// 	c.Neg()

// 	have := c.Top().Int64()
// 	want := int64(-6)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntNot(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(0b101)
// 	c.Not()

// 	have := c.Top().Int64()
// 	want := int64(-0b110)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntOr(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(0b1100, 0b1010)
// 	c.Or()

// 	have := c.Top().Int64()
// 	want := int64(0b1110)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntQuo(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(-7)
// 	c.PushInt(3)
// 	c.Quo()

// 	have := c.Top().Int64()
// 	want := int64(-2)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntQuoRem(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(-7)
// 	c.PushInt(3)
// 	c.QuoRem()

// 	have := c.Top().Int64()
// 	want := int64(-1)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}

// 	have = c.Next().Int64()
// 	want = int64(-2)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntPow(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(6)
// 	c.PushInt(2)
// 	c.Pow()

// 	have := c.Top().Int64()
// 	want := int64(36)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntRem(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(-7)
// 	c.PushInt(3)
// 	c.Rem()

// 	have := c.Top().Int64()
// 	want := int64(-1)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntRsh(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(0b1000)
// 	c.Rsh(2)

// 	have := c.Top().Int64()
// 	want := int64(0b10)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntSetBit(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(0b101)
// 	c.SetBit(1, 1)

// 	have := c.Top().Int64()
// 	want := int64(0b111)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }

// func TestBigIntSqrt(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(256)
// 	c.Sqrt()

// 	have := c.Top().Int64()
// 	want := int64(16)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntSub(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(6, 2)
// 	c.Sub()

// 	have := c.Top().Int64()
// 	want := int64(4)
// 	if have != want {
// 		t.Fatalf("\n have: %v \n want: %v", have, want)
// 	}
// }

// func TestBigIntXor(t *testing.T) {
// 	var c BigInt
// 	c.PushInt(0b1100, 0b1010)
// 	c.Xor()

// 	have := c.Top().Int64()
// 	want := int64(0b110)
// 	if have != want {
// 		t.Fatalf("\n have: %b \n want: %b", have, want)
// 	}
// }
