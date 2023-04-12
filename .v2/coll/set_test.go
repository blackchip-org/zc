package coll

import (
	"reflect"
	"sort"
	"testing"
)

func TestAddRemoveMapSet(t *testing.T) {
	s := NewMapSet[string]()

	want := []string{"bar", "baz", "foo"}
	s.Add(want...)
	have := s.Items()
	sort.Strings(have)

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("\n have: %v \n want: %v", have, want)
	}

	s.Remove("baz")
	want2 := []string{"bar", "foo"}
	have2 := s.Items()
	sort.Strings(have2)

	if !reflect.DeepEqual(have2, want2) {
		t.Fatalf("\n have: %v \n want: %v", have2, want2)
	}
}

func TestContainsMapSet(t *testing.T) {
	s := NewMapSet[string]()

	s.Add("foo")
	if !s.Contains("foo") {
		t.Fatalf("expected foo")
	}

	if s.Contains("bar") {
		t.Fatalf("unexpected bar")
	}
}
