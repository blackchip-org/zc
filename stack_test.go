package zc

import (
	"testing"
)

func TestStackPushPop(t *testing.T) {
	var s Stack[int]

	s.Push(1, 2, 3, 4, 5)
	have := s.Pop()
	want := 5
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	s.Pop()
	have = s.Pop()
	want = 3
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	s.Push(6, 7)
	have = s.Pop()
	want = 7
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	s.Pop()
	s.Pop()
	have = s.Pop()
	want = 1
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	defer func() {
		if have := recover(); have != nil {
			want := ErrStackEmpty
			if have != want {
				t.Fatalf("\n have panic: %v \n want panic: %v", have, want)
			}
		} else {
			t.Fatal("expected panic")
		}
	}()
	s.Pop()
}

/*
func TestStackRecycle(t *testing.T) {
	var s Stack[*big.Int]

	s.Push(big.NewInt(1))
	s.Push(big.NewInt(2))
	s.Push(big.NewInt(3))

	three := s.Pop()
	have := three.Int64()
	want := int64(3)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	s.Pop()
	i, ok := s.Recycle()
	if !ok {
		t.Fatal("expected ok")
	}
	i.SetInt64(4)

	i, ok = s.Recycle()
	if !ok {
		t.Fatal("expected ok")
	}
	i.SetInt64(5)

	_, ok = s.Recycle()
	if ok {
		t.Fatal("expectd not ok")
	}

	v := s.Pop()
	have = v.Int64()
	want = int64(5)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	// Check that this got recycled
	have = three.Int64()
	want = int64(5)
	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}
*/
