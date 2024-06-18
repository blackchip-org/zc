package zc

import "github.com/blackchip-org/scan"

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
