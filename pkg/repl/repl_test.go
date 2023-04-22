package repl

import (
	"testing"

	"github.com/blackchip-org/zc/pkg/ansi"
	"github.com/blackchip-org/zc/pkg/calc"
)

func TestUndo(t *testing.T) {
	ansi.Enabled = false
	c := calc.New()
	repl := New(c)

	repl.Eval("1")
	repl.Eval("2")
	repl.Eval("3")
	// 1 2 3
	repl.Eval("undo")
	// 1 2
	repl.Eval("undo")
	// 1
	top, _ := c.Peek(0)
	if top != "1" {
		t.Fatalf("\n have: %v \n want: 1", top)
	}
	repl.Eval("undo")
	// empty
	repl.Eval("undo")
	if repl.Calc.Error() == nil {
		t.Fatalf("expected error")
	}
	repl.Eval("redo")
	// 1
	top, _ = c.Peek(0)
	if top != "1" {
		t.Fatalf("\n have: %v want: 1", top)
	}
	repl.Eval("redo")
	// 1 2
	repl.Eval("redo")
	// 1 2 3
	top, _ = c.Peek(0)
	if top != "3" {
		t.Fatalf("\n have: %v want: 3", top)
	}
	repl.Eval("redo")
	if repl.Calc.Error() == nil {
		t.Fatalf("expected error")
	}
}
