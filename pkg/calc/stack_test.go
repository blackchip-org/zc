package calc

import (
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6"
)

func TestStackPushPop(t *testing.T) {
	var c BigInt

	c.Push(big.NewInt(1))
	c.Push(big.NewInt(2))
	c.Push(big.NewInt(3))
	c.Push(big.NewInt(4))
	c.Push(big.NewInt(5))

	v := c.Pop()
	have := v.Int64()
	want := int64(5)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	c.Pop()
	v = c.Pop()
	have = v.Int64()
	want = int64(3)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	c.Push(big.NewInt(6))
	c.Push(big.NewInt(7))
	v = c.Pop()
	have = v.Int64()
	want = int64(7)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	c.Pop()
	c.Pop()
	v = c.Pop()
	have = v.Int64()
	want = int64(1)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	defer func() {
		if have := recover(); have != nil {
			want := zc.ErrStackEmpty
			if have != want {
				t.Fatalf("\n have panic: %v \n want panic: %v", have, want)
			}
		} else {
			t.Fatal("expected panic")
		}
	}()
	c.Pop()
}

func TestStackRecycle(t *testing.T) {
	var c BigInt

	c.Push(big.NewInt(1))
	c.Push(big.NewInt(2))
	c.Push(big.NewInt(3))

	three := c.Pop()
	have := three.Int64()
	want := int64(3)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	c.Pop()
	i, ok := c.Recycle()
	if !ok {
		t.Fatal("expected ok")
	}
	i.SetInt64(4)

	i, ok = c.Recycle()
	if !ok {
		t.Fatal("expected ok")
	}
	i.SetInt64(5)

	_, ok = c.Recycle()
	if ok {
		t.Fatal("expectd not ok")
	}

	v := c.Pop()
	have = v.Int64()
	want = int64(5)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	have = three.Int64()
	want = int64(3)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}
