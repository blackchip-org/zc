package zc

import (
	"testing"
)

func TestRotate(t *testing.T) {
	var s Stack[int]

	s.Push(1, 2, 3)
	s.Rotate()
	have := s.String()
	want := "2 | 3 | 1"

	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

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
