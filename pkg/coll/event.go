package coll

import (
	"fmt"
	"testing"
)

type Event struct {
	Type    string
	Message string
}

func (e Event) String() string {
	if e.Message == "" {
		return e.Type
	}
	return e.Type + ": " + e.Message
}

func NewEvent(t string, format string, args ...any) Event {
	e := Event{Type: t}
	if format != "" {
		e.Message = fmt.Sprintf(format, args...)
	}
	return e
}

type Listener func(Event)

type Dispatcher struct {
	Listener Listener
}

func (d *Dispatcher) Emit(t string) {
	if d.Listener != nil {
		d.Listener(NewEvent(t, ""))
	}
}

func (d *Dispatcher) Emitf(t string, format string, args ...any) {
	if d.Listener != nil {
		d.Listener(NewEvent(t, format, args...))
	}
}

func TestLogger(t *testing.T) func(Event) {
	return func(e Event) {
		t.Log(e.String())
	}
}

func ConsoleLogger(e Event) {
	fmt.Println(e)
}
