package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/cockroachdb/apd/v3"
)

func Format(a any) string {
	switch v := a.(type) {
	case *big.Int:
		return v.String()
	case *apd.Decimal:
		f := RemoveTrailingZeros(v.String())
		f = FormatExponent(f)
		return f
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", a)
	}
}

func FormatList(vals ...any) string {
	var strs []string
	for _, val := range vals {
		strs = append(strs, Format(val))
	}
	return strings.Join(strs, " | ")
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

func RemoveTrailingZeros(v string) string {
	s := scan.NewScannerFromString("", v)

	// Keep everything up to the decimal separator
	scan.Until(s, scan.Rune('.'), s.Keep)
	if s.This == '.' {
		s.Skip()
	}
	if !s.HasMore() {
		return s.Emit().Val
	}

	// If the value is something like 100.00, we want to remove the decimal
	// separator along with the zeros. The write of the separator is delayed
	// until a digit is written
	pendingDecSep := true

	inZeros := false
	zerosSeen := 0
	for s.HasMore() {
		// Count up the number of zeros seen until there is a non-zero
		// rune. Then emit those zeros when not trailing.
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
