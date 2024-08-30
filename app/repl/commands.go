package repl

import (
	"fmt"
	"slices"
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6"
)

type Cmd func(*Repl, []scan.Token) error

var cmds map[string]Cmd

func init() {
	cmds = map[string]Cmd{
		"def":   def,
		"drop":  drop,
		"redo":  redo,
		"u":     undo,
		"quit":  quit,
		"quote": quote,
		"undo":  undo,
	}
}

func def(r *Repl, toks []scan.Token) error {
	if len(toks) == 0 {
		return fmt.Errorf("no macro name")
	}

	name := toks[0].Val
	if _, exists := cmds[name]; exists {
		return fmt.Errorf("invalid name")
	}

	toks = toks[1:]

	if len(toks) == 0 {
		if _, exists := r.macros[name]; !exists {
			return fmt.Errorf("macro not defined: %v", name)
		}
		delete(r.macros, name)
		r.notice = fmt.Sprintf("macro '%v' undefined", name)
		return nil
	}

	if _, exists := r.macros[name]; exists {
		r.notice = fmt.Sprintf("macro '%v' redefined", name)
	} else if _, exists := r.ops[name]; exists {
		r.notice = fmt.Sprintf("macro '%v' overrides", name)
	} else {
		r.notice = fmt.Sprintf("macro '%v' defined", name)
	}
	r.macros[name] = toks
	return nil
}

func drop(r *Repl, _ []scan.Token) error {
	if r.EndQuote == "" {
		r.Calc.Pop()
	}
	return nil
}

func redo(r *Repl, _ []scan.Token) error {
	if len(r.redoStack) == 0 {
		return fmt.Errorf("redo stack is empty")
	}
	r.undoStack = append([][]zc.Item{slices.Clone(r.Calc.Stack())}, r.undoStack...)
	r.Calc.SetStack(r.redoStack[0])
	r.redoStack = r.redoStack[1:]
	return nil
}

func quit(_ *Repl, _ []scan.Token) error {
	return errQuit
}

func quote(r *Repl, toks []scan.Token) error {
	if len(toks) == 0 {
		return fmt.Errorf("expected text to be used as a delimiter")
	}
	if len(toks) > 1 {
		return fmt.Errorf("expected single word to be used as a delimiter")
	}
	r.EndQuote = strings.TrimSpace(toks[0].Val)
	return nil
}

func undo(r *Repl, _ []scan.Token) error {
	if len(r.undoStack) == 0 {
		return fmt.Errorf("undo stack is empty")
	}
	r.redoStack = append([][]zc.Item{slices.Clone(r.Calc.Stack())}, r.redoStack...)
	r.Calc.SetStack(r.undoStack[0])
	r.undoStack = r.undoStack[1:]
	return nil
}
