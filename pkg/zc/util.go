package zc

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/pkg/scanner"
	"golang.org/x/exp/constraints"
)

func Clamp[T constraints.Ordered](v T, min T, max T) T {
	if v > max {
		return max
	}
	if v < min {
		return min
	}
	return v
}

func IsValuePrefix(ch rune, next rune) bool {
	switch {
	case unicode.IsDigit(ch), unicode.Is(unicode.Sc, ch):
		return true
	case (ch == '-' || ch == '+' || ch == '.') && unicode.IsDigit(next):
		return true
	}
	return false
}

func StackString(c Calc) string {
	var items []string
	for _, item := range c.Stack() {
		fmt.Printf("** QUOTING: %v\n", item)
		items = append(items, Quote(item))
	}
	return strings.Join(items, " | ")
}

func Quote(v string) string {
	var s scanner.Scanner
	s.SetString(v)

	needsQuotes := false
	if !IsValuePrefix(s.Ch, s.Lookahead) {
		needsQuotes = true
	} else {
		s.ScanUntil(unicode.IsSpace)
		if !s.End() {
			needsQuotes = true
		}
	}

	if !needsQuotes {
		return v
	}

	s.SetString(v)
	s.Text.WriteRune('\'')
	for s.Ok() {
		if s.Ch == '\'' {
			s.Text.WriteString("\\'")
			s.Next()
		} else {
			s.Keep()
		}
	}
	s.Text.WriteRune('\'')
	return s.Token()
}
