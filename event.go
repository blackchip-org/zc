package zc

import (
	"fmt"
	"testing"
)

type Event interface {
	Type() string
	String() string
}

type Listener func(Event)

type StackEvent struct {
	type_ string
	stack *Stack
}

func NewStackEvent(type_ string, stack Stack) StackEvent {
	return StackEvent{
		type_: type_,
		stack: stack.Clone(),
	}
}

func (e StackEvent) Type() string {
	return e.type_
}

func (e StackEvent) String() string {
	return fmt.Sprintf("%4s: %v", e.type_, FormatStack(e.stack))
}

type OpEvent struct {
	name string
}

func NewOpEvent(env *OpEnv) OpEvent {
	return OpEvent{name: env.Op.Name}
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
