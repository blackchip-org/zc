package zc

type Pool[T any] struct {
	items []*T
	pos   int
}

func NewPool[T any](n int) *Pool[T] {
	return &Pool[T]{items: make([]*T, n)}
}

func (p *Pool[T]) New() *T {
	if p.pos == 0 {
		return new(T)
	}
	p.pos--
	return p.items[p.pos]
}

func (p *Pool[T]) Recycle(v *T) {
	if p.pos >= len(p.items) {
		return
	}
	p.items[p.pos] = v
	p.pos++
}
