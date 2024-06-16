package zc

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	items []T
	pos   int
}

func (s *Stack[T]) Clear() {
	s.pos = 0
}

func (s *Stack[T]) Push(vals ...T) {
	for _, val := range vals {
		if s.pos < len(s.items) {
			s.items[s.pos] = val
		} else {
			s.items = append(s.items, val)
		}
		s.pos++
	}
}

func (s *Stack[T]) Recycle() (T, bool) {
	var item T
	if s.pos >= len(s.items) {
		return item, false
	}
	item = s.items[s.pos]
	s.pos++
	return item, true
}

func (s *Stack[T]) Drop() T {
	if s.pos == 0 {
		panic(ErrStackEmpty)
	}
	s.pos--
	return s.items[s.pos]
}

func (s *Stack[T]) Next() T {
	if s.pos < 2 {
		panic(ErrStackUnderflow)
	}
	return s.items[s.pos-2]
}

func (s *Stack[T]) RotateDown() {
	if s.pos < 3 {
		panic(ErrNotEnoughArgs(s.pos, 3))
	}
	rot := s.items[0]
	s.items[0] = s.items[1]
	s.items[1] = s.items[2]
	s.items[2] = rot
}

func (s *Stack[T]) String() string {
	items := s.items[:s.pos]
	var strs []string
	for _, item := range items {
		strs = append(strs, fmt.Sprintf("%v", item))
	}
	return strings.Join(strs, " | ")
}

func (s *Stack[T]) Top() T {
	if s.pos == 0 {
		panic(ErrStackEmpty)
	}
	return s.items[s.pos-1]
}
