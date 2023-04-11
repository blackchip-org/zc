package coll

import "testing"

func TestPushPop(t *testing.T) {
	d := NewDequeSlice[int]()

	in := []int{0, 1, 2, 3}
	for i, v := range in {
		if i != d.Len() {
			t.Fatalf("\n have len: %v \n want len: %v", d.Len(), i)
		}
		Push(d, v)
	}

	out := Reverse(in)
	for _, want := range out {
		have, ok := Pop(d)
		if !ok {
			t.Fatalf("unexpected !ok")
		}
		if have != want {
			t.Fatalf("\n have: %v \n want: %v", have, want)
		}
		if d.Len() != have {
			t.Fatalf("\n have len: %v \n want len: %v", d.Len(), have)
		}
	}

	_, ok := Pop(d)
	if ok {
		t.Fatalf("unexpected !ok")
	}
}

func TestAt(t *testing.T) {
	d := NewDequeSlice[int]()

	in := []int{0, 1, 2, 3}
	PushAll(d, in...)

	for i := 0; i < d.Len(); i++ {
		v := d.At(i)
		if i != v {
			t.Fatalf("\n i: %v \n v: %v", i, v)
		}
		if i > 0 {
			v2 := d.At(-i)
			i2 := len(in) - i
			if i2 != v2 {
				t.Fatalf("\n i2: %v \n v2: %v", i2, v2)
			}
		}
	}
}
