package zc

import "testing"

func testStack() *Stack {
	c, err := NewCalc(Config{})
	if err != nil {
		panic(err)
	}
	s := NewStack(c, "test")
	return s
}

func TestPush(t *testing.T) {
	s := testStack()
	s.Push("1")
	s.Push("2")

	v1 := s.MustPop()
	if v1 != "2" {
		t.Errorf("\n have: %v \n want: 2", v1)
	}
	v2 := s.MustPop()
	if v2 != "1" {
		t.Errorf("\n have: %v \n want: 1", v1)
	}
	_, err := s.Pop()
	if err == nil {
		t.Errorf("expected empty stack error")
	}
}

func TestString(t *testing.T) {
	s := testStack()
	s.Push("1")
	s.Push("2")
	have := s.String()
	want := "1 | 2"
	if have != want {
		t.Errorf("\n have: %v \n want %v", have, want)
	}
}

func TestPeek(t *testing.T) {
	s := testStack()
	s.Push("1")
	s.Push("2")
	have, err := s.Peek()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "2"
	if have != want {
		t.Errorf("\n have: %v \n want %v", have, want)
	}
}
