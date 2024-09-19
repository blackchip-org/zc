package msg

import (
	"fmt"
	"math/big"
)

func Copied() string {
	return "copied"
}

func Inexact() string {
	return "inexact"
}

func LocalTimeZoneSet(z string) string {
	return fmt.Sprintf("local time zone is now %v", z)
}

func NowSet(dt string) string {
	return fmt.Sprintf("now set to %v", Quote(dt))
}

func PrecisionSet(p uint) string {
	return fmt.Sprintf("precision set to %v", p)
}

func RoundingModeSet(m string) string {
	return fmt.Sprintf("rounding mode set to %v", m)
}

func Reset() string {
	return "reset"
}

func SeedSet(s *big.Int) string {
	return fmt.Sprintf("seed set to %v", s)
}

func Stored() string {
	return "stored"
}
