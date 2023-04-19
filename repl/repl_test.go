package repl

import (
	"testing"

	"github.com/blackchip-org/zc/calc"
	"github.com/blackchip-org/zc/internal/ansi"
)

func TestUndo(t *testing.T) {
	ansi.Enabled = false
	c := calc.New()
	repl := New(c)

	repl.Eval("1")
	repl.Eval("2")
	repl.Eval("3")
	// 1 2 3
	repl.undo()
	// 1 2
	repl.undo()
	// 1
	top, _ := c.Peek(0)
	if top != "1" {
		t.Fatalf("\n have: %v \n want: 1", top)
	}
	repl.undo()
	// empty
	if err := repl.undo(); err == nil {
		t.Fatalf("expected error")
	}
	repl.redo()
	// 1
	top, _ = c.Peek(0)
	if top != "1" {
		t.Fatalf("\n have: %v want: 1", top)
	}
	repl.redo()
	// 1 2
	repl.redo()
	// 1 2 3
	top, _ = c.Peek(0)
	if top != "3" {
		t.Fatalf("\n have: %v want: 3", top)
	}
	if err := repl.redo(); err == nil {
		t.Fatalf("expected error")
	}
}
