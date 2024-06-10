package zc

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/scan"
)

func FormatList(vals ...any) string {
	var strs []string
	for _, val := range vals {
		str := fmt.Sprintf("%v", val)
		strs = append(strs, str)
	}
	return strings.Join(strs, " | ")
}

func FormatStack(s Stack) string {
	var strs []string
	for _, item := range s.Items {
		strs = append(strs, item.String())
	}
	return strings.Join(strs, " | ")
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

func ToString(v any) string {
	return fmt.Sprintf("%v", v)
}
