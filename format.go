package zc

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/scan"
)

var escapeMap = map[rune]string{
	'\a': `\a`,
	'\b': `\b`,
	'\f': `\f`,
	'\n': `\n`,
	'\r': `\r`,
	'\t': `\t`,
	'\v': `\v`,
}

func Escape(r rune) string {
	es, ok := escapeMap[r]
	if ok {
		return es
	}
	switch {
	case scan.IsPrintable(r):
		return string(r)
	case r <= 0xff:
		return fmt.Sprintf(`\x%02x`, r)
	case r <= 0xffff:
		return fmt.Sprintf(`\u%04x`, r)
	default:
		return fmt.Sprintf(`\U%08x`, r)
	}
}

func EscapeString(str string) string {
	s := scan.NewScannerFromString("", str)
	for s.HasMore() {
		s.Val.WriteString(Escape(s.This))
		s.Skip()
	}
	return s.Emit().Val
}

func FormatItems(items []Item) []string {
	strs := make([]string, len(items))
	for i, item := range items {
		strs[i] = item.String()
	}
	return strs
}

func FormatList(vals []string) string {
	var s scan.Scanner
	var strs []string
	for _, val := range vals {
		s.InitFromString("", val)
		for s.HasMore() {
			switch {
			case s.This == '|':
				s.Val.WriteString(`\|`)
				s.Skip()
			default:
				s.Val.WriteString(Escape(s.This))
				s.Skip()
			}
		}
		strs = append(strs, s.Emit().Val)
	}
	return strings.Join(strs, " | ")
}

func Quote(str string) string {
	quote := ""
	s := scan.NewScannerFromString("", str)

	if !IsValuePrefix(s.This, s.Next) {
		quote = `'`
	}
	for s.HasMore() {
		switch {
		case s.This == ' ':
			quote = `'`
			s.Keep()
		case s.This == '\'':
			quote = `'`
			s.Val.WriteString(`\'`)
			s.Skip()
		default:
			s.Val.WriteString(Escape(s.This))
			s.Skip()
		}
	}
	return quote + s.Emit().Val + quote
}

func Strings(xs ...any) []string {
	strs := make([]string, len(xs))
	for i, x := range xs {
		strs[i] = fmt.Sprint(x)
	}
	return strs
}

func FormatExponent(str string) string {
	s := scan.NewScannerFromString("", str)

	// Keep everything before the exponent
	scan.Until(s, scan.Rune('E', 'e'), s.Keep)

	// If no more, we didn't see the exponent
	if !s.HasMore() {
		return str
	}

	// We did see the exponent. Always write this out in lower case.
	s.Val.WriteRune('e')
	s.Skip()

	// Omit positive signs but keep the negative ones
	if s.This == '+' {
		s.Skip()
	}

	// Remove all leading zeros
	for s.This == '0' && s.Next != scan.EndOfText {
		s.Skip()
	}

	// Actual digits of the exponent
	scan.While(s, scan.IsAny, s.Keep)
	return s.Emit().Val
}

func Abbr(str string) string {
	if len(str) < 80 {
		return str
	}
	var abbr strings.Builder
	i := 0
	for i < 80 {
		r, w := utf8.DecodeRuneInString(str[i:])
		abbr.WriteRune(r)
		i += w
	}
	return abbr.String() + "…"
}
