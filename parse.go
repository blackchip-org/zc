package zc

import (
	"strings"
	"unicode"
)

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
