package repl

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/v5/pkg/scanner"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

type Cmd func(*REPL, *scanner.Scanner) error

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

func def(r *REPL, s *scanner.Scanner) error {
	s.ScanWhile(unicode.IsSpace)
	if s.End() {
		return fmt.Errorf("expected macro name")
	}
	if zc.IsValuePrefix(s.Ch, s.Lookahead) {
		return fmt.Errorf("invalid name")
	}

	name := s.Scan(scanner.Word)
	if _, exists := cmds[name]; exists {
		return fmt.Errorf("invalid name")
	}

	s.ScanWhile(unicode.IsSpace)
	expr := s.Scan(scanner.Remaining)

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

func pop(r *REPL, _ *scanner.Scanner) error {
	r.Calc.Pop()
	return nil
}

func redo(r *REPL, _ *scanner.Scanner) error {
	if len(r.redoStack) == 0 {
		return fmt.Errorf("redo stack is empty")
	}
	r.undoStack = append([][]string{r.Calc.Stack()}, r.undoStack...)
	r.Calc.SetStack(r.redoStack[0])
	r.redoStack = r.redoStack[1:]
	return nil
}

func quit(_ *REPL, _ *scanner.Scanner) error {
	return errQuit
}

func quote(r *REPL, s *scanner.Scanner) error {
	s.ScanWhile(unicode.IsSpace)
	if s.End() {
		return fmt.Errorf("expected text to be used as a delimiter")
	}
	for !s.End() {
		s.Keep()
	}
	r.quoteEnd = strings.TrimSpace(s.Token())
	return nil
}

func undo(r *REPL, _ *scanner.Scanner) error {
	if len(r.undoStack) == 0 {
		return fmt.Errorf("undo stack is empty")
	}
	r.redoStack = append([][]string{r.Calc.Stack()}, r.redoStack...)
	r.Calc.SetStack(r.undoStack[0])
	r.undoStack = r.undoStack[1:]
	return nil
}
