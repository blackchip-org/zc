package calc

import (
	"fmt"
	"math/big"
	"strings"
)

type Env[T any] struct {
	items []T
	pos   int
	clone func(T) T
}

func NewEnv[T any](cloneFunc func(T) T) *Env[T] {
	return &Env[T]{clone: cloneFunc}
}

func (e *Env[T]) Push(item T) {
	if e.pos < len(e.items) {
		e.items[e.pos] = item
	} else {
		e.items = append(e.items, item)
	}
	e.pos++
}

func (e *Env[T]) Pop() T {
	if e.pos == 0 {
		panic("stack empty")
	}
	e.pos--
	return e.clone(e.items[e.pos])
}

func (e *Env[T]) Release() T {
	if e.pos == 0 {
		panic("stack empty")
	}
	e.pos--
	return e.items[e.pos]
}

func (e *Env[T]) Recycle() (T, bool) {
	var item T
	if e.pos >= len(e.items) {
		return item, false
	}
	item = e.items[e.pos]
	e.pos++
	return item, true
}

func (e *Env[T]) Top() T {
	if e.pos == 0 {
		panic("stack empty")
	}
	return e.items[e.pos-1]
}

func (e *Env[T]) Down() {
	if e.pos == 0 {
		return
	}
	rotate := e.items[e.pos-1]
	for i := e.pos - 1; i > 0; i++ {
		e.items[i] = e.items[i-1]
	}
	e.items[0] = rotate
}

func (e *Env[T]) String() string {
	items := e.items[:e.pos]
	var strs []string
	for _, item := range items {
		strs = append(strs, fmt.Sprintf("%v", item))
	}
	return strings.Join(strs, " | ")
}

var (
	CloneBigInt = func(src *big.Int) *big.Int {
		var dest big.Int
		dest.Set(src)
		return &dest
	}
)

func NewBigIntEnv() *Env[*big.Int] {
	return NewEnv(CloneBigInt)
}
