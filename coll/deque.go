package coll

type Deque[T any] interface {
	PushBack(T)
	PopBack() (T, bool)
	PushFront(T)
	PopFront() (T, bool)
	Len() int
	Items() []T
	Clear()
}

type DequeSlice[T any] struct {
	xs []T
}

func NewDequeSlice[T any](items ...T) Deque[T] {
	d := &DequeSlice[T]{}
	if len(items) > 0 {
		d.xs = make([]T, len(items))
		copy(d.xs, items)
	}
	return d
}

func (d *DequeSlice[T]) Len() int {
	return len(d.xs)
}

func (d *DequeSlice[T]) Items() []T {
	var items []T
	items = append(items, d.xs...)
	return items
}

func (d *DequeSlice[T]) Clear() {
	d.xs = nil
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

func Push[T any](d Deque[T], xs ...T) {
	for _, x := range xs {
		d.PushBack(x)
	}
}

func Pop[T any](d Deque[T]) (T, bool) {
	return d.PopBack()
}

func Peek[T any](d Deque[T]) (T, bool) {
	var r T
	if d.Len() == 0 {
		return r, false
	}
	r = d.Items()[d.Len()-1]
	return r, false
}

/*
func Enqueue[T any](d Deque[T], e T) {
	d.PushFront(e)
}

func Dequeue[T any](d Deque[T]) (T, bool) {
	return d.PopBack()
}
*/
