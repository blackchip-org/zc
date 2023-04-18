package calc

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/blackchip-org/zc"
)

type calc struct {
	stack []string
	err   error
}

func New() zc.Calc {
	return &calc{}
}

func (c *calc) Stack() []string {
	s := make([]string, len(c.stack))
	copy(s, c.stack)
	return s
}

func (c *calc) Eval(s string) error {
	c.err = nil
	words := c.parseWords(s)
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
	return nil
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
	op(&env{calc: c})
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
