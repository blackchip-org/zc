package stack

import (
	"reflect"
	"strconv"
	"testing"
)

func testPush(t *testing.T, s Stack[int]) {
	s.Push(1)
	s.Push(2)
	s.Push(3)

	have := s.Items()
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func testPop(t *testing.T, s Stack[int]) {
	s.Push(1)
	s.Push(2)
	s.Pop()
	s.Push(3)

	have := s.Items()
	want := []int{1, 3}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func testPopEmpty(t *testing.T, s Stack[int]) {
	_, ok := s.Pop()
	if ok {
		t.Errorf("expected empty stack")
	}
}

func testLen(t *testing.T, s Stack[int]) {
	for i := 0; i < 10; i++ {
		s.Push(1)
		s.Push(2)
		s.Pop()
	}

	have := s.Len()
	want := 10

	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func testAt(t *testing.T, s Stack[int]) {
	s.Push(123)
	s.Push(234)
	s.Push(345)
	s.Push(456)

	tests := []struct {
		at   int
		want int
		err  bool
	}{
		{-5, 0, true},
		{-4, 123, false},
		{-1, 456, false},
		{0, 123, false},
		{3, 456, false},
		{4, 0, true},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.at), func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && !test.err {
					t.Fatal(err)
				}
			}()

			have := At(s, test.at)
			if test.err {
				t.Fatalf("expected error")
			}
			if have != test.want {
				t.Errorf("\n have: %v \n want: %v", have, test.want)
			}
		})
	}
}

func testDup(t *testing.T, s Stack[int]) {
	s.Push(1)
	s.Push(2)
	Dup(s)

	have := s.Items()
	want := []int{1, 2, 2}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func testPopN(t *testing.T, s Stack[int]) {
	s.Push(1)
	s.Push(2)
	s.Push(3)

	have, ok := PopN(s, 2)
	if !ok {
		t.Fatalf("unexpected empty stack")
	}
	want := []int{2, 3}

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	have = s.Items()
	want = []int{1}

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func testPushN(t *testing.T, s Stack[int]) {
	s.Push(1)
	PushN(s, []int{2, 3})

	have := s.Items()
	want := []int{1, 2, 3}

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}
}

func testTop(t *testing.T, s Stack[int]) {
	s.Push(1)
	s.Push(2)
	s.Push(3)

	have, ok := Top(s)
	if !ok {
		t.Fatalf("unexpected empty stack")
	}
	want := 3

	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}

func testUp(t *testing.T, s Stack[int]) {
	s.Push(1)
	s.Push(2)
	s.Push(3)
	Up(s)

	have := s.Items()
	want := []int{3, 1, 2}

	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}
