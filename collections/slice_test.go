package collections

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	in := []int{1, 2, 3, 4}
	want := []int{4, 3, 2, 1}
	have := Reverse(in)
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}
