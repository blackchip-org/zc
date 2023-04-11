package coll

type Deque[T any] interface {
	PushBack(T)
	PopBack() (T, bool)
	PushFront(T)
	PopFront() (T, bool)
	Len() int
	At(int) T
}

type DequeSlice[T any] struct {
	xs []T
}

func NewDequeSlice[T any]() Deque[T] {
	return &DequeSlice[T]{}
}

func (d *DequeSlice[T]) Len() int {
	return len(d.xs)
}

func (d *DequeSlice[T]) At(i int) T {
	if i < 0 {
		i = len(d.xs) + i
	}
	return d.xs[i]
}

func (d *DequeSlice[T]) PushBack(x T) {
	d.xs = append(d.xs, x)
}

func (d *DequeSlice[T]) PopBack() (T, bool) {
	var x T
	n := len(d.xs)
	if n == 0 {
		return x, false
	}
	x, d.xs = d.xs[n-1], d.xs[:n-1]
	return x, true
}

func (d *DequeSlice[T]) PushFront(e T) {
	d.xs = append([]T{e}, d.xs...)
}

func (d *DequeSlice[T]) PopFront() (T, bool) {
	var r T
	n := len(d.xs)
	if n == 0 {
		return r, false
	}
	r, d.xs = d.xs[0], d.xs[1:]
	return r, true
}

func Push[T any](d Deque[T], x T) {
	d.PushBack(x)
}

func PushAll[T any](d Deque[T], xs ...T) {
	for _, x := range xs {
		d.PushBack(x)
	}
}

func Pop[T any](d Deque[T]) (T, bool) {
	return d.PopBack()
}

/*
func Enqueue[T any](d Deque[T], e T) {
	d.PushFront(e)
}

func Dequeue[T any](d Deque[T]) (T, bool) {
	return d.PopBack()
}
*/
