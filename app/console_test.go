package app

import (
	"testing"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/ansi"
)

func TestUndo(t *testing.T) {
	ansi.Enabled = false
	calc, _ := zc.NewCalc(&zc.Config{})
	c := NewConsole(calc)

	c.Eval("1")
	c.Eval("2")
	c.Eval("3")
	// 1 2 3
	c.undo()
	// 1 2
	c.undo()
	// 1
	top, _ := calc.Env.Stack.Peek()
	if top != "1" {
		t.Fatalf("\n have: %v want: 1", top)
	}
	c.undo()
	// empty
	if err := c.undo(); err == nil {
		t.Fatalf("expected error")
	}
	c.redo()
	// 1
	top, _ = calc.Env.Stack.Peek()
	if top != "1" {
		t.Fatalf("\n have: %v want: 1", top)
	}
	c.redo()
	// 1 2
	c.redo()
	// 1 2 3
	top, _ = calc.Env.Stack.Peek()
	if top != "3" {
		t.Fatalf("\n have: %v want: 3", top)
	}
	if err := c.redo(); err == nil {
		t.Fatalf("expected error")
	}
}
