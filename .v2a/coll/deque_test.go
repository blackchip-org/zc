package coll

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushPop(t *testing.T) {
	d := NewDequeSlice[int]()

	in := []int{0, 1, 2, 3}
	for i, v := range in {
		assert.Equal(t, i, d.Len())
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
