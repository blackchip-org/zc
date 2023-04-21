package ptime

import (
	"fmt"
	"time"

	"github.com/blackchip-org/zc/pkg/ptime/locale"
)

type P struct {
	Locale *locale.Locale
	Parser *Parser
}

func For(loc *locale.Locale) *P {
	return &P{
		Locale: loc,
		Parser: NewParser(loc),
	}
}

func ForLocale(name string) (*P, error) {
	loc, ok := locale.Lookup(name)
	if !ok {
		return nil, fmt.Errorf("unknown locale: %v", name)
	}
	return For(loc), nil
}

func (p *P) Parse(text string) (Parsed, error) {
	return p.Parser.Parse(text)
}

func (p *P) ParseDate(text string) (Parsed, error) {
	return p.Parser.ParseDate(text)
}

func (p *P) ParseTime(text string) (Parsed, error) {
	return p.Parser.ParseTime(text)
}

func (p *P) Time(parsed Parsed, now time.Time) (time.Time, error) {
	return Time(p.Locale, parsed, now)
}

func (p *P) Format(layout string, t time.Time) string {
	return Format(p.Locale, layout, t)
}

func FormatOffset(offset int, sep string) string {
	sign := "+"
	if offset < 0 {
		sign = "-"
	}
	h := offset / 3600
	if h < 0 {
		h = h * -1
	}
	m := offset / 60 % 60
	return fmt.Sprintf("%v%02d%v%02d", sign, h, sep, m)
}
