package zc

import (
	"fmt"
	"testing"

	"github.com/blackchip-org/zc/v6/pkg/stack"
)

type Event interface {
	Type() string
	String() string
}

type Listener func(Event)

type StackEvent struct {
	type_ string
	stack stack.Stack[Item]
}

func NewStackEvent(type_ string, stack stack.Stack[Item]) StackEvent {
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
