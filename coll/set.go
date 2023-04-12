package coll

type Set[T comparable] interface {
	Add(...T)
	Remove(...T)
	Contains(T) bool
	Items() []T
	Len() int
}

type MapSet[T comparable] struct {
	xs map[T]struct{}
}

func NewMapSet[T comparable]() Set[T] {
	return &MapSet[T]{xs: make(map[T]struct{})}
}

func (s *MapSet[T]) Add(xs ...T) {
	for _, x := range xs {
		s.xs[x] = struct{}{}
	}
}

func (s *MapSet[T]) Remove(xs ...T) {
	for _, x := range xs {
		delete(s.xs, x)
	}
}

func (s *MapSet[T]) Contains(x T) bool {
	_, exists := s.xs[x]
	return exists
}

func (s *MapSet[T]) Items() []T {
	var r []T
	for x := range s.xs {
		r = append(r, x)
	}
	return r
}

func (s *MapSet[T]) Len() int {
	return len(s.xs)
}
