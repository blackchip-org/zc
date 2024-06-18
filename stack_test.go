package zc

import (
	"testing"

	"github.com/blackchip-org/zc/v6/pkg/assert"
)

func TestClear(t *testing.T) {
	var s Stack[int]

	s.Push(1, 2, 3)
	s.Clear()

	if s.Len() != 0 {
		t.Fatalf("expected empty stack")
	}
}

func TestGet(t *testing.T) {
	var s Stack[int]

	s.Push(1, 2, 3, 4, 5)
	have := s.Get(3)
	want := 2

	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestGetPanic(t *testing.T) {
	var s Stack[int]

	defer assert.Panic(t, ErrStackUnderflow.Error())
	s.Push(1, 2, 3, 4, 5)
	s.Get(5)
}

func TestLen(t *testing.T) {
	var s Stack[int]
	s.Push(1, 2, 3, 4, 5)

	have := s.Len()
	want := 5

	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestNext(t *testing.T) {
	var s Stack[int]
	s.Push(1, 2)

	have := s.Next()
	want := 1

	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestNextPanic(t *testing.T) {
	var s Stack[int]
	s.Push(1)

	defer assert.Panic(t, ErrStackUnderflow.Error())
	s.Next()
}

func TestRotate(t *testing.T) {
	var s Stack[int]

	s.Push(1, 2, 3)
	s.Rotate()

	have := s.String()
	want := FormatList(2, 3, 1)

	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestRotatePanic(t *testing.T) {
	var s Stack[int]

	defer assert.Panic(t, ErrNotEnoughArgs(1, 3).Error())
	s.Push(1)
	s.Rotate()
}

func TestSet(t *testing.T) {
	var s Stack[int]

	s.Push(1, 2, 3, 4, 5)
	s.Set(3, 6)

	have := s.String()
	want := FormatList(1, 6, 3, 4, 5)

	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestSetPanic(t *testing.T) {
	var s Stack[int]

	defer assert.Panic(t, ErrStackUnderflow.Error())
	s.Push(1, 2, 3, 4, 5)
	s.Set(5, 6)
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

	defer assert.Panic(t, ErrStackEmpty.Error())
	s.Pop()
}

func TestTop(t *testing.T) {
	var s Stack[int]
	s.Push(1, 2)

	have := s.Top()
	want := 2

	if have != want {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func TestTopPanic(t *testing.T) {
	var s Stack[int]

	defer assert.Panic(t, ErrStackEmpty.Error())
	s.Top()
}
