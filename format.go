package zc

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/msg"
)

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
				s.Val.WriteString(msg.Escape(s.This))
				s.Skip()
			}
		}
		strs = append(strs, s.Emit().Val)
	}
	return strings.Join(strs, " | ")
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
