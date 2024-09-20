package msg

import (
	"fmt"
	"math/big"
)

var (
	Copied           = "copied"
	Inexact          = "inexact"
	LocalTimeZoneSet = func(z string) string {
		return fmt.Sprintf("local time zone is now %v", z)
	}
	NowSet = func(dt string) string {
		return fmt.Sprintf("now set to %v", Quote(dt))
	}
	PrecisionSet = func(p uint) string {
		return fmt.Sprintf("precision set to %v", p)
	}
	RoundingModeSet = func(m string) string {
		return fmt.Sprintf("rounding mode set to %v", m)
	}
	Reset   = "reset"
	SeedSet = func(s *big.Int) string {
		return fmt.Sprintf("seed set to %v", s)
	}
	Stored = "stored"
)
