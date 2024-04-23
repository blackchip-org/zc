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
	l := s.Len()
	if index < -l || index >= l {
		panic("index out of bounds: " + strconv.Itoa(index))
	}
	items := s.Items()
	if index < 0 {
		index = l + index
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

func PopN[T any](s Stack[T], n int) ([]T, bool) {
	var ret []T
	if s.Len() < n {
		return ret, false
	}
	l := s.Len()
	items := s.Items()
	ret, items = items[l-n:], items[:l-n]
	s.SetItems(items)
	return ret, true
}

func PushN[T any](s Stack[T], add []T) {
	items := append(s.Items(), add...)
	s.SetItems(items)
}

func Top[T any](s Stack[T]) (T, bool) {
	var zero T
	if s.Len() == 0 {
		return zero, false
	}
	return At(s, -1), true
}

func Up[T any](s Stack[T]) {
	l := s.Len()
	if l == 0 {
		return
	}
	items := s.Items()
	items = append([]T{items[l-1]}, items[:l-1]...)
	s.SetItems(items)
}
