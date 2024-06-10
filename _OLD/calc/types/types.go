package types

import (
	"reflect"
	"strings"

	"github.com/blackchip-org/scan"
)

func GoName(v any) string {
	var name strings.Builder
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
		name.WriteRune('*')
	}
	name.WriteString(t.Name())
	return name.String()
}

func FormatExponent(str string) string {
	s := scan.NewScannerFromString("", str)

	// Keep everything before the exponent
	scan.Until(s, scan.Rune('E', 'e'), s.Keep)

	// If no more, we didn't see the exponent
	if s.HasMore() {
		// We did see the exponent. Always write this out in lower case.
		s.Skip()
		s.Val.WriteRune('e')

		// Omit positive signs but keep the negative ones
		if s.This == '+' {
			s.Skip()
		}

		// Remove all leading zeros and long.
		if s.This == '0' && s.Next != scan.EndOfText {
			scan.While(s, scan.Rune('0'), s.Skip)
		}

		// Actual digits of the exponent
		scan.While(s, scan.IsAny, s.Keep)
	}
	return s.Emit().Val
}

func RemoveTrailingZeros(v string) string {
	s := scan.NewScannerFromString("", v)
	scan.Until(s, scan.Rune('.'), s.Keep)
	pendingDecSep := true
	s.Skip()

	inZeros := false
	zerosSeen := 0
	for s.HasMore() {
		if s.This == '0' {
			if inZeros {
				zerosSeen++
			} else {
				inZeros = true
				zerosSeen = 1
			}
			s.Skip()
		} else {
			if pendingDecSep {
				pendingDecSep = false
				s.Val.WriteRune('.')
			}
			if inZeros && scan.IsDigit09(s.This) {
				inZeros = false
				for i := 0; i < zerosSeen; i++ {
					s.Val.WriteRune('0')
				}
			}
			if !scan.IsDigit09(s.This) {
				scan.While(s, scan.IsAny, s.Keep)
				break
			}
			s.Keep()
		}
	}
	return s.Emit().Val
}
