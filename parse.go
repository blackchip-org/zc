package zc

import (
	"unicode"

	"github.com/blackchip-org/scan"
)

const (
	TokenValue = "value"
	TokenName  = "name"
)

type valueRule struct{}

func (r valueRule) Eval(s *scan.Scanner) bool {
	if !IsValuePrefix(s.This, s.Next) {
		return false
	}
	s.Type = TokenValue
	s.Keep()
	scan.While(s, scan.Not(scan.IsSpace), s.Keep)
	return true
}

type slashValueRule struct{}

func (r slashValueRule) Eval(s *scan.Scanner) bool {
	if s.This != '/' {
		return false
	}
	s.Type = TokenValue
	s.Skip()
	scan.While(s, scan.Not(scan.IsSpace), s.Keep)
	return true
}

var rules = scan.NewRuleSet(
	scan.SkipSpaceRule,
	scan.StrDoubleQuoteRule.
		WithType(TokenValue).
		WithOptionalTerminator(true),
	scan.StrSingleQuoteRule.WithType(TokenValue).
		WithType(TokenValue).
		WithOptionalTerminator(true),
	scan.NewStrRule('[', ']').
		WithType(TokenValue).
		WithOptionalTerminator(true),
	valueRule{},
	slashValueRule{},
	scan.NewWhileRule(scan.Not(scan.IsSpace), TokenName),
).WithNoMatchFunc(scan.UnexpectedUntil(scan.IsSpace))

func ScanWords(line string) []scan.Token {
	s := scan.NewScannerFromString("", line)
	runner := scan.NewRunner(s, rules)
	return runner.All()
}

func isDecoration(r rune) bool {
	if r == ',' || r == '_' || r == ' ' {
		return true
	}
	// Currency symbols
	if unicode.Is(unicode.Sc, r) {
		return true
	}
	return false
}

func isAltExponent(s *scan.Scanner) bool {
	return (s.This == 'x' || s.This == '×') && s.Next == '1' && s.Peek(2) == '0'
}

func PreParseInt(str string) string {
	s := scan.NewScannerFromString("", str)
	for s.HasMore() {
		switch {
		case isDecoration(s.This):
			s.Skip()
		default:
			s.Keep()
		}
	}
	return s.Emit().Val
}

func PreParseDecimal(str string) string {
	s := scan.NewScannerFromString("", str)
	for s.HasMore() {
		switch {
		case isDecoration(s.This):
			s.Skip()
		case isAltExponent(s):
			scan.Repeat(s.Skip, 3)
			s.Val.WriteRune('e')
		default:
			s.Keep()
		}
	}
	return s.Emit().Val
}

func IsValuePrefix(ch rune, next rune) bool {
	switch {
	case unicode.IsDigit(ch):
		return true
	// Currency symbols
	case unicode.Is(unicode.Sc, ch):
		return true
	case (ch == '-' || ch == '+' || ch == '.') && unicode.IsDigit(next):
		return true
	}
	return false
}
