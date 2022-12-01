package zc

import "errors"

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
		return "", errors.New("stack empty")
	}
	var top string
	top, s.data = s.data[0], s.data[1:]
	return top, nil
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *Stack) Set(v string) {
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
	return s.data[0], nil
}
