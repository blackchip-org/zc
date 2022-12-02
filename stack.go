package zc

import (
	"errors"
	"log"
	"os"
)

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Items() []string {
	items := make([]string, len(s.data))
	copy(items, s.data)
	return items
}

func (s *Stack) Len() int {
	return len(s.data)
}

func (s *Stack) Pop() (string, error) {
	if len(s.data) == 0 {
		//panic("stack empty")
		return "", errors.New("stack empty")
	}
	var top string
	len := len(s.data)
	top, s.data = s.data[len-1], s.data[:len-1]
	s.trace("pop: %v", top)
	return top, nil
}

func (s *Stack) MustPop() string {
	val, err := s.Pop()
	if err != nil {
		panic(err)
	}
	return val
}

func (s *Stack) Push(v string) {
	s.trace("push: %v", v)
	s.data = append(s.data, v)
}

func (s *Stack) Set(v string) {
	s.trace("set: %v", v)
	if len(s.data) == 0 {
		s.data = append(s.data, v)
	} else {
		s.data[0] = v
	}
}

func (s *Stack) Get() (string, error) {
	if len(s.data) == 0 {
		return "", errors.New("undefined")
	}
	s.trace("get: %v", s.data[0])
	return s.data[0], nil
}

func (s *Stack) Clear() {
	s.trace("clear")
	s.data = nil
}

var traceStack bool

func init() {
	traceStack = os.Getenv("ZC_TRACE_STACK") == "true"
}

func (s *Stack) trace(format string, a ...any) {
	if traceStack {
		log.Printf("stack: "+format, a...)
	}
}
