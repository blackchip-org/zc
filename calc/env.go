package calc

import "github.com/blackchip-org/zc"

type env struct {
	calc *calc
}

func (e *env) Peek(i int) (string, bool) {
	n := len(e.calc.stack)
	stackI := n - i
	if stackI < 0 || stackI >= n {
		return "", false
	}
	return e.calc.stack[stackI], true
}

func (e *env) Pop() (string, bool) {
	n := len(e.calc.stack)
	if n == 0 {
		return "", false
	}
	var item string
	e.calc.stack, item = e.calc.stack[:n-1], e.calc.stack[n-1]
	return item, true
}

func (e *env) MustPop() string {
	item, ok := e.Pop()
	if !ok {
		panic(zc.ErrStackEmpty)
	}
	return item
}

func (e *env) Push(item string) {
	e.calc.stack = append(e.calc.stack, item)
}

func (e *env) Error(err error) {
	if e.calc.err == nil {
		e.calc.err = err
	}
}
