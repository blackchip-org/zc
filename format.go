package zc

import (
	"fmt"
	"strings"

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
		strs[i] = item.Val()
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
	for s.HasMore() {
		switch {
		case s.This == ' ':
			quote = `"`
			s.Keep()
		case s.This == '"':
			quote = `"`
			s.Val.WriteString(`\"`)
			s.Skip()
		default:
			s.Val.WriteString(Escape(s.This))
			s.Skip()
		}
	}
	return quote + s.Emit().Val + quote
}
