package zc

import (
	"strings"
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

func isFormatting(ch rune) bool {
	if ch == ',' || ch == '_' || ch == ' ' {
		return true
	}
	if unicode.Is(unicode.Sc, ch) {
		return true
	}
	return false
}

func cleanNumber(str string) string {
	var res strings.Builder
	for _, ch := range str {
		if !isFormatting(ch) {
			res.WriteRune(ch)
		}
	}
	return res.String()
}

func PreParseNumber(s string) string {
	s = cleanNumber(s)
	s = strings.Replace(s, "×10", "e", 1)
	s = strings.Replace(s, "x10", "e", 1)
	return s
}
