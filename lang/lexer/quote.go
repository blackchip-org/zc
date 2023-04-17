package lexer

import (
	"strings"
	"unicode"
)

func isValue(ch rune, next rune) bool {
	switch {
	case unicode.IsDigit(ch), unicode.Is(unicode.Sc, ch):
		return true
	case (ch == '-' || ch == '+' || ch == '.') && unicode.IsDigit(next):
		return true
	}
	return false
}

func Quote(v string) string {
	required := false
	runes := []rune(v)
	for i, ch := range runes {
		if i == 0 {
			var next rune
			if len(runes) > 1 {
				next = runes[i+1]
			}
			if !isValue(ch, next) {
				required = true
				break
			}
		}
		if unicode.IsSpace(ch) {
			required = true
			break
		}
	}
	if !required {
		return v
	}
	var ret strings.Builder
	ret.WriteRune('\'')
	for _, ch := range runes {
		if ch == '\'' {
			ret.WriteString("\\'")
		} else {
			ret.WriteRune(ch)
		}
	}
	ret.WriteRune('\'')
	return ret.String()
}
