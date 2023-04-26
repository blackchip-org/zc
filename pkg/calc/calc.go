package calc

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/blackchip-org/zc/pkg/zc"
)

type calc struct {
	stack []string
	err   error
	info  string
	state map[string]any
	op    string
	args  []string
}

func New() zc.Calc {
	return &calc{state: make(map[string]any)}
}

func (c *calc) Stack() []string {
	s := make([]string, len(c.stack))
	copy(s, c.stack)
	return s
}

func (c *calc) StackLen() int {
	return len(c.stack)
}

func (c *calc) SetStack(s []string) {
	c.stack = c.stack[:len(s)]
	copy(c.stack, s)
}

func (c *calc) Info() string {
	return c.info
}

func (c *calc) SetInfo(format string, args ...any) {
	c.info = fmt.Sprintf(format, args...)
}

func (c *calc) Eval(s string) error {
	c.err = nil
	c.info = ""
	c.op = ""
	c.args = nil

	lines := strings.Split(s, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		words := c.parseWords(line)
		for _, word := range words {
			ch, _ := utf8.DecodeRuneInString(word)
			if isValue(ch) {
				c.stack = append(c.stack, strings.TrimPrefix(word, "\""))
			} else {
				c.evalOp(word)
			}
			if c.err != nil {
				return c.err
			}
		}
	}
	return nil
}

func (c *calc) Peek(i int) (string, bool) {
	n := len(c.stack)
	stackI := n - 1 - i
	if stackI < 0 || stackI >= n {
		return "", false
	}
	return c.stack[stackI], true
}

func (c *calc) Pop() (string, bool) {
	n := len(c.stack)
	if n == 0 {
		return "", false
	}
	var item string
	c.stack, item = c.stack[:n-1], c.stack[n-1]
	return item, true
}

func (c *calc) MustPop() string {
	item, ok := c.Pop()
	if !ok {
		panic("stack empty")
	}
	return item
}

func (c *calc) Push(item string) {
	c.stack = append(c.stack, item)
}

func (c *calc) SetError(err error) {
	if c.err == nil {
		c.err = err
	}
}

func (c *calc) Error() error {
	return c.err
}

func (c *calc) Derive() zc.Calc {
	return New()
}

func (c *calc) NewState(name string, s any) {
	c.state[name] = s
}

func (c *calc) State(name string) (any, bool) {
	s, ok := c.state[name]
	return s, ok
}

func (c *calc) SetOp(op string) {
	c.op = op
}

func (c *calc) SetArgs(args []string) {
	c.args = args
}

func (c *calc) Op() zc.OpCall {
	return zc.OpCall{Name: c.op, Args: c.args}
}

func (c *calc) OpNames() []string {
	var os []string
	for name := range opsTable {
		os = append(os, name)
	}
	return os
}

func (c *calc) parseWords(str string) []string {
	var words []string
	var word strings.Builder

	var inWord, inQuote bool
	var endQuote rune

	for _, ch := range str {
		if !inWord {
			if unicode.IsSpace(ch) {
				continue
			}
			word.Reset()
			inWord = true
			switch ch {
			case '"':
				inQuote = true
				endQuote = '"'
				word.WriteRune('"')
			case '\'':
				inQuote = true
				endQuote = '\''
				word.WriteRune('"')
			case '[':
				inQuote = true
				endQuote = ']'
				word.WriteRune('"')
			default:
				word.WriteRune(ch)
			}
		} else {
			if (unicode.IsSpace(ch) && !inQuote) || ch == endQuote {
				inWord = false
				inQuote = false
				words = append(words, word.String())
			} else {
				word.WriteRune(ch)
			}
		}
	}
	if inWord {
		words = append(words, word.String())
	}
	return words
}

func (c *calc) evalOp(name string) {
	op, ok := opsTable[name]
	if !ok {
		c.err = zc.ErrUnknownOp(name)
		return
	}
	op(c)
}

func isValue(ch rune) bool {
	switch {
	case unicode.IsNumber(ch):
		return true
	case unicode.Is(unicode.Sc, ch):
		return true
	case ch == '+' || ch == '-':
		return true
	case ch == '.':
		return true
	case ch == '"':
		return true
	}
	return false
}

func evalOp(op zc.OpDecl) zc.CalcFunc {
	return func(c zc.Calc) {
		c.SetOp(op.Name)
		if op.Macro != "" {
			c.Eval(op.Macro)
			return
		}
		switch len(op.Funcs) {
		case 0:
			panic("no functions for operation")
		case 1:
			evalOpSingle(c, op)
		default:
			evalOpDispatch(c, op)
		}
	}
}

func evalOpSingle(c zc.Calc, op zc.OpDecl) {
	fn := op.Funcs[0]
	var args []string
	n := len(fn.Params)
	for i, p := range fn.Params {
		v, ok := c.Peek(n - i - 1)
		if !ok {
			zc.ErrNotEnoughArgs(c, op.Name, len(fn.Params))
			return
		}
		if !p.Is(v) {
			zc.ErrExpectedType(c, p, v)
			return
		}
		args = append(args, v)
	}
	c.SetArgs(args)
	fn.Func(c)
}

func evalOpDispatch(c zc.Calc, op zc.OpDecl) {
	for _, decl := range op.Funcs {
		if isOpMatch(c, decl) {
			nArgs := len(decl.Params)
			n := c.StackLen()
			args := c.Stack()[n-nArgs:]
			c.SetArgs(args)
			decl.Func(c)
			return
		}
	}

	// For now, we are going to check the first decl to
	// determine the number of arguments.
	nArgs := len(op.Funcs[0].Params)
	var args []string
	for i := nArgs - 1; i >= 0; i-- {
		v, ok := c.Peek(i)
		if !ok {
			zc.ErrNotEnoughArgs(c, op.Name, nArgs)
			return
		}
		args = append(args, v)
	}
	zc.ErrNoOpFor(c, op.Name, args)
}

func isOpMatch(c zc.Calc, decl zc.FuncDecl) bool {
	for i, param := range decl.Params {
		arg, ok := c.Peek(len(decl.Params) - i - 1)
		if !ok {
			return false
		}
		if !param.Is(arg) {
			return false
		}
	}
	return true
}
