package repl

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

type Cmd func(*REPL, *scan.Scanner) error

var cmds map[string]Cmd

func init() {
	cmds = map[string]Cmd{
		"def":   def,
		"":      pop,
		"redo":  redo,
		"u":     undo,
		"quit":  quit,
		"quote": quote,
		"undo":  undo,
	}
}

func def(r *REPL, s *scan.Scanner) error {
	scan.Space(s)
	if !s.HasMore() {
		return fmt.Errorf("expected macro name")
	}
	if zc.IsValuePrefix(s.This, s.Next) {
		return fmt.Errorf("invalid name")
	}

	name := scan.Word(s)
	if _, exists := cmds[name]; exists {
		return fmt.Errorf("invalid name")
	}

	scan.Space(s)
	expr := scan.All(s)

	if expr == "" {
		if _, exists := r.macros[name]; !exists {
			return fmt.Errorf("macro not defined: %v", name)
		}
		delete(r.macros, name)
		r.info = fmt.Sprintf("macro %v undefined", zc.Quote(name))
		return nil
	}

	if _, exists := r.macros[name]; exists {
		r.info = fmt.Sprintf("macro %v redefined", zc.Quote(name))
	} else if _, exists := r.ops[name]; exists {
		r.info = fmt.Sprintf("macro %v overrides", zc.Quote(name))
	} else {
		r.info = fmt.Sprintf("macro %v defined", zc.Quote(name))
	}
	r.macros[name] = expr
	return nil
}

func pop(r *REPL, _ *scan.Scanner) error {
	r.Calc.Pop()
	return nil
}

func redo(r *REPL, _ *scan.Scanner) error {
	if len(r.redoStack) == 0 {
		return fmt.Errorf("redo stack is empty")
	}
	r.undoStack = append([][]string{r.Calc.Stack()}, r.undoStack...)
	r.Calc.SetStack(r.redoStack[0])
	r.redoStack = r.redoStack[1:]
	return nil
}

func quit(_ *REPL, _ *scan.Scanner) error {
	return errQuit
}

func quote(r *REPL, s *scan.Scanner) error {
	scan.Space(s)
	if !s.HasMore() {
		return fmt.Errorf("expected text to be used as a delimiter")
	}
	r.quoteEnd = strings.TrimSpace(scan.All(s))
	return nil
}

func undo(r *REPL, _ *scan.Scanner) error {
	if len(r.undoStack) == 0 {
		return fmt.Errorf("undo stack is empty")
	}
	r.redoStack = append([][]string{r.Calc.Stack()}, r.redoStack...)
	r.Calc.SetStack(r.undoStack[0])
	r.undoStack = r.undoStack[1:]
	return nil
}
