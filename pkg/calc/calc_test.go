package calc

import (
	"math/big"
	"testing"

	"github.com/blackchip-org/zc/v6"
)

func TestStackPushPop(t *testing.T) {
	e := newBigIntEnv()

	e.Push(big.NewInt(1))
	e.Push(big.NewInt(2))
	e.Push(big.NewInt(3))
	e.Push(big.NewInt(4))
	e.Push(big.NewInt(5))

	v := e.Pop()
	have := v.Int64()
	want := int64(5)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	e.Pop()
	v = e.Pop()
	have = v.Int64()
	want = int64(3)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	e.Push(big.NewInt(6))
	e.Push(big.NewInt(7))
	v = e.Pop()
	have = v.Int64()
	want = int64(7)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	e.Pop()
	e.Pop()
	v = e.Pop()
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
	e.Pop()
}

func TestStackRecycle(t *testing.T) {
	e := newBigIntEnv()

	e.Push(big.NewInt(1))
	e.Push(big.NewInt(2))
	e.Push(big.NewInt(3))

	three := e.Pop()
	have := three.Int64()
	want := int64(3)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	e.Pop()
	i, ok := e.Recycle()
	if !ok {
		t.Fatal("expected ok")
	}
	i.SetInt64(4)

	i, ok = e.Recycle()
	if !ok {
		t.Fatal("expected ok")
	}
	i.SetInt64(5)

	_, ok = e.Recycle()
	if ok {
		t.Fatal("expectd not ok")
	}

	v := e.Pop()
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
