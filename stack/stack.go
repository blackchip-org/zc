package stack

import "strconv"

type Stack[T any] interface {
	Push(T)
	Pop() (T, bool)
	Len() int
	Items() []T
	SetItems([]T)
}

func At[T any](s Stack[T], index int) T {
	n := s.Len()
	if index < -n || index >= n {
		panic("index out of bounds: " + strconv.Itoa(index))
	}
	items := s.Items()
	if index < 0 {
		index = n + index
	}
	return items[index]
}

func Dup[T any](s Stack[T]) bool {
	item, ok := Top(s)
	if ok {
		s.Push(item)
	}
	return ok
}

func Top[T any](s Stack[T]) (T, bool) {
	var zero T
	if s.Len() == 0 {
		return zero, false
	}
	return At(s, -1), true
}

func Up[T any](s Stack[T]) {
	n := s.Len()
	if n == 0 {
		return
	}
	items := s.Items()
	items = append([]T{items[n-1]}, items[:n-1]...)
	s.SetItems(items)
}
