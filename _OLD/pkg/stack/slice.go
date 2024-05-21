package stack

import "slices"

type sliceStack[T any] struct {
	items []T
}

func NewSlice[T any]() Stack[T] {
	return &sliceStack[T]{}
}

func (s *sliceStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *sliceStack[T]) Pop() (T, bool) {
	var item T
	n := len(s.items)
	if n == 0 {
		return item, false
	}
	item, s.items = s.items[n-1], s.items[:n-1]
	return item, true
}

func (s *sliceStack[T]) Len() int {
	return len(s.items)
}

func (s *sliceStack[T]) Items() []T {
	return s.items
}

func (s *sliceStack[T]) SetItems(items []T) {
	s.items = items
}

func (s *sliceStack[T]) Clear() {
	s.items = nil
}

func (s *sliceStack[T]) Clone() Stack[T] {
	s2 := NewSlice[T]()
	s2.SetItems(slices.Clone(s.Items()))
	return s2
}
