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

func (s *Stack[T]) Get(i int) T {
	if i >= s.pos {
		panic(ErrStackUnderflow)
	}
	return s.items[s.pos-i-1]
}

func (s *Stack[T]) Len() int {
	return s.pos
}

func (s *Stack[T]) Next() T {
	if s.pos < 2 {
		panic(ErrStackUnderflow)
	}
	return s.items[s.pos-2]
}

func (s *Stack[T]) Pop() T {
	if s.pos == 0 {
		panic(ErrStackEmpty)
	}
	s.pos--
	return s.items[s.pos]
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

func (s *Stack[T]) Rotate() {
	if s.pos < 3 {
		panic(ErrNotEnoughArgs(s.pos, 3))
	}
	rot := s.items[0]
	s.items[0] = s.items[1]
	s.items[1] = s.items[2]
	s.items[2] = rot
}

func (s *Stack[T]) Set(i int, v T) {
	if i >= s.pos {
		panic(ErrStackUnderflow)
	}
	s.items[s.pos-i-1] = v
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

func (s *Stack[T]) TopRef() *T {
	if s.pos == 0 {
		panic(ErrStackEmpty)
	}
	return &s.items[s.pos-1]
}
