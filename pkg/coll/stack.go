package coll

import (
	"fmt"
	"slices"
	"strings"
)

type Stack[T any] struct {
	Dispatcher
	items []T
	pos   int
}

func (s *Stack[T]) Clear() {
	s.pos = 0
	s.Emit("clear")
}

func (s *Stack[T]) Clone() Stack[T] {
	items := slices.Clone(s.items)[:s.pos]
	return Stack[T]{items: items, pos: len(items)}
}

func (s *Stack[T]) Drop() {
	if s.pos == 0 {
		panic(ErrStackEmpty)
	}
	s.pos--
	s.Emitf("drop", "%v", s.items[s.pos])
}

func (s *Stack[T]) Get(i int) T {
	if i >= s.pos {
		panic(ErrIndexOutOfBounds(i))
	}
	return s.items[s.pos-i-1]
}

func (s *Stack[T]) Items() []T {
	return s.items[:s.pos]
}

func (s *Stack[T]) SetItems(items []T) {
	s.items = items
	s.pos = len(s.items)
	s.Emit("set-items")
}

func (s *Stack[T]) Len() int {
	return s.pos
}

func (s *Stack[T]) Pop() T {
	if s.pos == 0 {
		panic(ErrStackEmpty)
	}
	s.pos--
	s.Emitf("pop", "%v", s.items[s.pos])
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
		s.Emitf("push", "%v", val)
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
	s.Emit("rotate")
}

func (s *Stack[T]) Set(i int, v T) {
	if i >= s.pos {
		panic(ErrIndexOutOfBounds(i))
	}
	s.items[s.pos-i-1] = v
	s.Emitf("set", "%v: %v", i, v)
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

func (s *Stack[T]) Debug(on bool) {
	if on {
		lastStack := ""
		s.Listener = func(e Event) {
			fmt.Println(e)
			stack := s.String()
			if stack != "" && lastStack != stack {
				fmt.Print("\t")
				fmt.Println(s)
			}
			lastStack = stack
		}
	} else {
		s.Listener = nil
	}
}
