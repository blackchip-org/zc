package dms

import "github.com/blackchip-org/scan"

const (
	IntType   = scan.IntType
	RealType  = scan.RealType
	DegType   = "deg"
	MinType   = "min"
	SecType   = "sec"
	EastType  = "e"
	NorthType = "n"
	SouthType = "s"
	WestType  = "w"
)

var (
	Sign  = NewClassRule(scan.Sign)
	Deg   = NewClassRule(scan.Rune('d', '°')).WithType(DegType)
	Min   = NewClassRule(scan.Rune('m', '\'', '′')).WithType(MinType)
	Sec   = NewClassRule(scan.Rune('s', '"', '″')).WithType(SecType)
	East  = NewClassRule(scan.Rune('e', 'E')).WithType(EastType)
	North = NewClassRule(scan.Rune('n', 'N')).WithType(NorthType)
	South = NewClassRule(scan.Rune('s', 'S')).WithType(SouthType)
	West  = NewClassRule(scan.Rune('w', 'W')).WithType(WestType)
)

type ClassRule struct {
	class scan.Class
	type_ string
}

func NewClassRule(c scan.Class) ClassRule {
	return ClassRule{class: c}
}

func (r ClassRule) WithType(t string) ClassRule {
	r.type_ = t
	return r
}

func (r ClassRule) Eval(s *scan.Scanner) bool {
	if !s.Is(r.class) {
		return false
	}
	s.Keep()
	if r.type_ != "" {
		s.Type = r.type_
	}
	return true
}

type Context struct {
	RuleSet scan.RuleSet
}

func NewContext() *Context {
	c := &Context{}
	c.RuleSet = scan.NewRuleSet(
		scan.NewSpaceRule(scan.Whitespace),
		scan.Real,
		Sign,
		Deg, Min, Sec,
		East, North, South, West,
	)
	return c
}
