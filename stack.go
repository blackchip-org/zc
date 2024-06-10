package zc

import (
	"strings"
)

type Meta struct {
	Unit  string
	Label string
}

type Item struct {
	Val  any
	Type TypeID
	Meta *Meta
}

func (i Item) String() string {
	if i.Meta == nil {
		return String(i.Val)
	}
	var label string
	if i.Meta.Label != "" {
		label = " # " + i.Meta.Label
	}
	return String(i.Val) + i.Meta.Unit + label
}

type Stack struct {
	items []Item
	pos   int
}

func (s *Stack) Push(item Item) {
	if s.pos < len(s.items) {
		s.items[s.pos] = item
	} else {
		s.items = append(s.items, item)
	}
	s.pos++
}

func (s *Stack) Pop() (Item, bool) {
	var item Item
	if s.pos == 0 {
		return item, false
	}
	s.pos--
	return s.items[s.pos], true
}

func (s *Stack) Items() []Item {
	return s.items[:s.pos]
}

func (s *Stack) String() string {
	var strs []string
	for _, i := range s.items[:s.pos] {
		strs = append(strs, i.String())
	}
	return strings.Join(strs, " | ")
}
