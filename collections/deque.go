package collections

type Deque[T any] struct {
	data []T
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (q *Deque[T]) Len() int {
	return len(q.data)
}

func (q *Deque[T]) At(index int) T {
	if index < 0 {
		index = len(q.data) + index
	}
	return q.data[index]
}

func (q *Deque[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *Deque[T]) PushAll(v []T) {
	q.data = append(q.data, v...)
}

func (q *Deque[T]) Pop() (T, bool) {
	var r T
	if len(q.data) == 0 {
		return r, false
	}
	i := len(q.data) - 1
	q.data, r = q.data[:i], q.data[i]
	return r, true
}

func (q *Deque[T]) MustPop() T {
	r, ok := q.Pop()
	if !ok {
		panic("empty")
	}
	return r
}
