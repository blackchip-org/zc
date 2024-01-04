package calc

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/blackchip-org/zc/v5/pkg/zc"
)

type Calc struct {
	Trace bool
	stack []string
	err   error
	info  string
	state map[string]any
	op    string
	args  []string
}

func New() *Calc {
	return &Calc{state: make(map[string]any)}
}

func (c *Calc) Stack() []string {
	s := make([]string, len(c.stack))
	copy(s, c.stack)
	return s
}

func (c *Calc) StackLen() int {
	return len(c.stack)
}

func (c *Calc) SetStack(s []string) {
	c.info = ""
	c.stack = slices.Clone(s)
	c.ptrace("set : %v", zc.StackString(c))
}

func (c *Calc) Info() string {
	return c.info
}

func (c *Calc) SetInfo(format string, args ...any) {
	c.info = fmt.Sprintf(format, args...)
}

func (c *Calc) Eval(s string, args ...any) error {
	c.err = nil
	c.info = ""
	c.op = ""
	c.args = nil

	eval := fmt.Sprintf(s, args...)
	lines := strings.Split(eval, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		words := c.parseWords(line)
		for _, word := range words {
			ch, size := utf8.DecodeRuneInString(word)
			next, _ := utf8.DecodeRuneInString(word[size:])
			if isValue(ch, next) {
				c.Push(strings.TrimPrefix(word, "\""))
			} else {
				c.ptrace("oper: %v", word)
				c.evalOp(word)
			}
			if c.err != nil {
				return c.err
			}
		}
	}
	return nil
}

func (c *Calc) MustEval(s string, args ...any) {
	if err := c.Eval(s, args...); err != nil {
		panic(err)
	}
}

func (c *Calc) Peek(i int) (string, bool) {
	n := len(c.stack)
	stackI := n - 1 - i
	if stackI < 0 || stackI >= n {
		return "", false
	}
	return zc.RemoveAnnotation(c.stack[stackI]), true
}

func (c *Calc) Pop() (string, bool) {
	c.info = ""
	n := len(c.stack)
	if n == 0 {
		c.ptrace("pop: <stack empty?")
		return "", false
	}
	var item string
	c.stack, item = c.stack[:n-1], c.stack[n-1]
	if len(c.stack) != 0 {
		c.ptrace("pop : %v => %v", zc.StackString(c), item)
	} else {
		c.ptrace("pop : => %v", item)
	}
	return zc.RemoveAnnotation(item), true
}

func (c *Calc) MustPop() string {
	item, ok := c.Pop()
	if !ok {
		panic("stack empty")
	}
	return item
}

func (c *Calc) Push(item string) {
	c.info = ""
	if len(c.stack) != 0 {
		c.ptrace("push: %v <= %v", zc.StackString(c), item)
	} else {
		c.ptrace("push: <= %v", item)
	}
	c.stack = append(c.stack, item)
}

func (c *Calc) SetError(err error) {
	if c.err == nil {
		c.err = err
	}
	if err != nil {
		c.info = ""
	}
}

func (c *Calc) Error() error {
	return c.err
}

func (c *Calc) Derive() zc.Calc {
	sub := New()
	sub.state = maps.Clone(c.state)
	return sub
}

func (c *Calc) NewState(name string, s any) {
	c.state[name] = s
}

func (c *Calc) State(name string) (any, bool) {
	s, ok := c.state[name]
	return s, ok
}

func (c *Calc) SetOp(op string) {
	c.op = op
}

func (c *Calc) SetArgs(args []string) {
	c.args = args
}

func (c *Calc) Op() zc.OpCall {
	return zc.OpCall{Name: c.op, Args: c.args}
}

func (c *Calc) OpNames() []string {
	var os []string
	for name := range opsTable {
		os = append(os, name)
	}
	return os
}

func (c *Calc) parseWords(str string) []string {
	var words []string
	var word strings.Builder

	var inQuote int
	var inWord bool
	var beginQuote, endQuote rune

	runes := []rune(str)
	for i, ch := range runes {
		var next rune
		if i < len(runes)-1 {
			next = runes[i+1]
		}
		if !inWord {
			if unicode.IsSpace(ch) {
				continue
			}
			word.Reset()
			inWord = true
			switch {
			case ch == '"':
				inQuote++
				beginQuote = '"'
				endQuote = '"'
				word.WriteRune('"')
			case ch == '\'':
				inQuote++
				beginQuote = '\''
				endQuote = '\''
				word.WriteRune('"')
			case ch == '[':
				inQuote++
				beginQuote = '['
				endQuote = ']'
				word.WriteRune('"')
			case ch == '/' && !unicode.IsSpace(next) && next != 0:
				word.WriteRune('"')
			default:
				word.WriteRune(ch)
			}
		} else {
			if (unicode.IsSpace(ch) && inQuote == 0) || ch == endQuote {
				if inQuote > 0 {
					inQuote--
				}
				if inQuote == 0 {
					inWord = false
					words = append(words, word.String())
				} else {
					word.WriteRune(ch)
				}
			} else if ch == beginQuote {
				inQuote++
				word.WriteRune(ch)
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

func (c *Calc) evalOp(name string) {
	op, ok := opsTable[name]
	if !ok {
		c.err = zc.ErrUnknownOp(name)
		return
	}
	op(c)
}

func (c *Calc) ptrace(format string, args ...any) {
	if c.Trace {
		fmt.Printf(format, args...)
		fmt.Println()
	}
}

func isValue(ch rune, next rune) bool {
	switch {
	case unicode.IsNumber(ch):
		return true
	case unicode.Is(unicode.Sc, ch):
		return true
	case (ch == '+' || ch == '-') && unicode.IsNumber(next):
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
