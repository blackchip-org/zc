package zc

import (
	"fmt"
	"slices"
	"testing"
)

type Event interface {
	Type() string
	String() string
}

type Listener func(Event)

type StackEvent struct {
	type_ string
	stack []Item
}

func NewStackEvent(c Calc, type_ string) StackEvent {
	return StackEvent{
		type_: type_,
		stack: slices.Clone(c.Stack()),
	}
}

func (e StackEvent) Type() string {
	return e.type_
}

func (e StackEvent) String() string {
	return fmt.Sprintf("%4s: %v", e.type_, e.stack)
}

type OpEvent struct {
	name string
}

func NewOpEvent(name string) OpEvent {
	return OpEvent{name: name}
}

func (e OpEvent) Type() string { return "op" }

func (e OpEvent) String() string {
	return fmt.Sprintf("  op: %v", e.name)
}

func TestLogger(t *testing.T) func(Event) {
	return func(e Event) {
		t.Log(e.String())
	}
}

func ConsoleLogger(e Event) {
	fmt.Println(e)
}
