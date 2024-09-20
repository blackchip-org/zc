package msg

import (
	"fmt"
	"reflect"
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/scan"
)

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

func Quote(str string) string {
	quote := ""
	s := scan.NewScannerFromString("", str)

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

func goName(v any) string {
	if v == nil {
		return "nil"
	}
	var name strings.Builder
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
		name.WriteRune('*')
	}
	name.WriteString(t.Name())
	return name.String()
}
