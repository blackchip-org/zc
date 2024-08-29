package calc

import "testing"

func TestDecimalAbs(t *testing.T) {
	c := NewDecimal()
	c.PushString("-4.2")
	c.Abs()

	have := c.PopString()
	want := "4.2"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalAdd(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.1")
	c.PushString("2.2")
	c.Add()

	have := c.PopString()
	want := "3.3"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalCbrt(t *testing.T) {
	c := NewDecimal()
	c.PushString("10.648")
	c.Cbrt()
	c.Reduce()

	have := c.PopString()
	want := "2.2"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalCeil(t *testing.T) {
	c := NewDecimal()
	c.PushString("4.2")
	c.Ceil()

	have := c.PopString()
	want := "5"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalCmp(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.1")
	c.PushString("2.2")

	have := c.Cmp()
	want := -1
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalCmpTotal(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.1")
	c.PushString("2.2")

	have := c.CmpTotal()
	want := -1
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalExp(t *testing.T) {
	c := NewDecimal()
	c.PushString("2")
	c.Exp()

	have := c.PopString()
	want := "7.389056098930650227230427461"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalFloor(t *testing.T) {
	c := NewDecimal()
	c.PushString("4.2")
	c.Floor()

	have := c.PopString()
	want := "4"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalLn(t *testing.T) {
	c := NewDecimal()
	c.PushString("7.389056098930650227230427461")
	c.Ln()
	c.Reduce()

	have := c.PopString()
	want := "2"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalLog10(t *testing.T) {
	c := NewDecimal()
	c.PushString("100")
	c.Log10()
	c.Reduce()

	have := c.PopString()
	want := "2"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalMul(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.1")
	c.PushString("2.2")
	c.Mul()

	have := c.PopString()
	want := "2.42"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalNeg(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.1")
	c.Neg()

	have := c.PopString()
	want := "-1.1"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalPopFloat64(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.1")

	have, _ := c.PopFloat64()
	want := 1.1
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalPopInt64(t *testing.T) {
	c := NewDecimal()
	c.PushString("123")

	have, _ := c.PopInt64()
	want := int64(123)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalPopInt64Err(t *testing.T) {
	c := NewDecimal()
	c.PushString("123.45")

	_, err := c.PopInt64()
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestDecimalPow(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.1")
	c.PushString("2.2")
	c.Pow()

	have := c.PopString()
	want := "1.233286300554662510989000588"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalPushFloat64(t *testing.T) {
	c := NewDecimal()
	c.PushFloat64(1.1)

	have := c.PopString()
	want := "1.1"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalPushInt(t *testing.T) {
	c := NewDecimal()
	c.PushInt(1)

	have := c.PopString()
	want := "1"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalPushInt64(t *testing.T) {
	c := NewDecimal()
	c.PushInt64(1)

	have := c.PopString()
	want := "1"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalQuantize(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.6666666666666")
	c.Quantize(-3)

	have := c.PopString()
	want := "1.667"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalQuo(t *testing.T) {
	c := NewDecimal()
	c.PushString("2.42")
	c.PushString("2.2")
	c.Quo()
	c.Reduce()

	have := c.PopString()
	want := "1.1"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalInteger(t *testing.T) {
	c := NewDecimal()
	c.PushString("2.42")
	c.PushString("2.2")
	c.QuoInteger()

	have := c.PopString()
	want := "1"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalReduce(t *testing.T) {
	c := NewDecimal()
	c.PushString("2.42000")
	c.Reduce()

	have := c.PopString()
	want := "2.42"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalRem(t *testing.T) {
	c := NewDecimal()
	c.PushString("8")
	c.PushString("3")
	c.Rem()

	have := c.PopString()
	want := "2"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalRound(t *testing.T) {
	c := NewDecimal()
	c.PushString("1.2345")
	c.Ctx = c.Ctx.WithPrecision(3)
	c.Round()

	have := c.PopString()
	want := "1.23"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalRoundToIntegralExact(t *testing.T) {
	c := NewDecimal()
	c.PushString("2.9")
	c.RoundToIntegeralExact()

	have := c.PopString()
	want := "3"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalRoundToIntegralValue(t *testing.T) {
	c := NewDecimal()
	c.PushString("2.9")
	c.RoundToIntegeralValue()

	have := c.PopString()
	want := "3"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalSqrt(t *testing.T) {
	c := NewDecimal()
	c.PushString("4.84")
	c.Sqrt()
	c.Reduce()

	have := c.PopString()
	want := "2.2"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalSub(t *testing.T) {
	c := NewDecimal()
	c.PushString("3.3")
	c.PushString("1.1")
	c.Sub()

	have := c.PopString()
	want := "2.2"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestDecimalTrunc(t *testing.T) {
	c := NewDecimal()
	c.PushString("3.9")
	c.Trunc()

	have := c.PopString()
	want := "3"
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}
