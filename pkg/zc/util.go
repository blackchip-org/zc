package zc

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/v5/pkg/scanner"
	"golang.org/x/exp/constraints"
)

const (
	AnnotationMarker = "#!(anno)"
	AnnotationSep    = "#"
)

func Annotate(c Calc, format string, a ...any) {
	if os.Getenv("ZC_NO_ANNO") != "" || c.Error() != nil || format == "" {
		return
	}
	anno := fmt.Sprintf(format, a...)
	v := c.MustPop()
	c.Push(fmt.Sprintf("%v %v %v", v, AnnotationMarker, anno))
}

func FormatStackItem(v string) string {
	return strings.Replace(v, AnnotationMarker, AnnotationSep, 1)
}

func RemoveAnnotation(v string) string {
	if i := strings.Index(v, AnnotationMarker); i >= 0 {
		v = v[:i-1]
	}
	return v
}

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
	case ch == '/':
		return true
	}
	return false
}

func StackString(c Calc) string {
	var items []string
	for _, item := range c.Stack() {
		items = append(items, FormatStackItem(item))
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
