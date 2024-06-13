package calc

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/zc/v6"
)

type Calc[T any] struct {
	items []T
	pos   int
	err   error
	Clone func(T) T
}

func (s *Calc[T]) AssertArgs(n int) {
	if s.pos < n {
		panic(zc.ErrNotEnoughArgs(s.pos, 3))
	}
}

func (s *Calc[T]) Clear() {
	s.err = nil
	s.pos = 0
}

func (s *Calc[T]) Err() error {
	return s.err
}

func (s *Calc[T]) Pop() T {
	if s.err != nil {
		panic(s.err)
	}
	if s.pos == 0 {
		panic(zc.ErrStackEmpty)
	}
	s.pos--
	return s.Clone(s.items[s.pos])
}

func (s *Calc[T]) Push(vals ...T) {
	for _, val := range vals {
		if s.pos < len(s.items) {
			s.items[s.pos] = val
		} else {
			s.items = append(s.items, val)
		}
		s.pos++
	}
}

func (s *Calc[T]) Recycle() (T, bool) {
	var item T
	if s.pos >= len(s.items) {
		return item, false
	}
	item = s.items[s.pos]
	s.pos++
	return item, true
}

func (s *Calc[T]) Release() T {
	if s.pos == 0 {
		panic(zc.ErrStackEmpty)
	}
	s.pos--
	return s.items[s.pos]
}

func (s *Calc[T]) RotateDown() {
	if s.pos < 3 {
		panic(zc.ErrNotEnoughArgs(s.pos, 3))
	}
	rot := s.items[0]
	s.items[0] = s.items[1]
	s.items[1] = s.items[2]
	s.items[2] = rot
}

func (s *Calc[T]) String() string {
	items := s.items[:s.pos]
	var strs []string
	for _, item := range items {
		strs = append(strs, fmt.Sprintf("%v", item))
	}
	return strings.Join(strs, " | ")
}

func (s *Calc[T]) Top() T {
	if s.pos == 0 {
		panic(zc.ErrStackEmpty)
	}
	return s.items[s.pos-1]
}
