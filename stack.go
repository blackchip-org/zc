package zc

import (
	"slices"
)

type Item struct {
	Value any
	Type  Type
	Anno  string
}

func (i Item) String() string {
	return ToString(i.Value)
}

func (i Item) StringWithAnno() string {
	if i.Anno == "" {
		return ToString(i.Value)
	}
	return ToString(i.Value) + " # " + i.Anno
}

type Stack struct {
	Items []Item
}

func (s *Stack) Push(item Item) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() (Item, bool) {
	var item Item
	n := len(s.Items)
	if n == 0 {
		return item, false
	}
	item, s.Items = s.Items[n-1], s.Items[:n-1]
	return item, true
}

func (s *Stack) PopN(n int) ([]Item, bool) {
	var ret []Item
	l := len(s.Items)
	if l < n {
		return nil, false
	}
	ret, s.Items = s.Items[l-n:], s.Items[:l-n]
	return ret, true
}

func (s *Stack) At(idx int) Item {
	return s.Items[idx]
}

func (s *Stack) Len() int {
	return len(s.Items)
}

func (s *Stack) Clone() Stack {
	var sc Stack
	sc.Items = slices.Clone(s.Items)
	return sc
}

func (s *Stack) Clear() {
	s.Items = nil
}
