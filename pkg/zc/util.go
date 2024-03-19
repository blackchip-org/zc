package zc

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/blackchip-org/scan"
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
	var s scan.Scanner
	s.InitFromString("", v)

	needsQuotes := false
	if !IsValuePrefix(s.This, s.Next) {
		needsQuotes = true
	} else {
		scan.Until(&s, scan.IsSpace, s.Discard)
		if s.HasMore() {
			needsQuotes = true
		}
	}

	if !needsQuotes {
		return v
	}

	s.InitFromString("", v)
	s.Val.WriteRune('\'')
	for s.HasMore() {
		if s.This == '\'' {
			s.Val.WriteString("\\'")
			s.Skip()
		} else {
			s.Keep()
		}
	}
	s.Val.WriteRune('\'')
	return s.Emit().Val
}

func RemoveTrailingZeros(s string) string {
	chs := []rune(s)
	end := len(chs)
	zerosDone := false
	for i := len(chs) - 1; i >= 0; i-- {
		ch := chs[i]

		// Done when a decimal separator is found
		if ch == '.' {
			return string(chs[:end])
		}
		// If we have already found all trailing zeros and just waiting
		// for the decimal separator
		if zerosDone {
			continue
		}

		// Are we still seeing zeros?
		if ch != '0' {
			zerosDone = true
			continue
		}
		end--

		// If we do see a zero, check to see if the next is a decimal point.
		// Gobble that up too and we are now done.
		next := rune(0)
		if i-1 > 0 {
			next = chs[i-1]
		}
		if next == '.' {
			end--
			return string(chs[:end])
		}
	}
	return s
}
