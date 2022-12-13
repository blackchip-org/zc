package zc

import (
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/shopspring/decimal"
)

type decRoundFunc func(decimal.Decimal, int32) decimal.Decimal

var (
	roundModes = map[string]decRoundFunc{
		"ceil":      func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundCeil(places) },
		"down":      func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundDown(places) },
		"floor":     func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundFloor(places) },
		"half-up":   func(d decimal.Decimal, places int32) decimal.Decimal { return d.Round(places) },
		"half-even": func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundBank(places) },
		"up":        func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundUp(places) },
	}
)

func (c *Calc) parseDigits(sep rune, v string) ([]rune, []rune) {
	var intDigits, fracDigits []rune
	inInt := true
	for _, ch := range v {
		if ch == sep {
			if !inInt {
				fracDigits = append(fracDigits, ch)
			}
			inInt = false
		} else if inInt {
			intDigits = append(intDigits, ch)
		} else {
			fracDigits = append(fracDigits, ch)
		}
	}
	return intDigits, fracDigits
}

func (c *Calc) FormatNumberString(v string, opts NumberFormatOptions) string {
	var digits strings.Builder
	intDigits, fracDigits := c.parseDigits('.', v)

	if opts.IntPat == "" {
		digits.WriteString(string(intDigits))
	} else {
		var intResult []rune
		intPat := []rune(opts.IntPat)

		idxPat := len(opts.IntPat) - 1
		idxDig := len(intDigits) - 1
		for idxDig >= 0 {
			if intDigits[idxDig] == '-' {
				intResult = append([]rune{intDigits[idxDig]}, intResult...)
				idxDig--
			} else if intPat[idxPat] == '0' {
				intResult = append([]rune{intDigits[idxDig]}, intResult...)
				idxDig--
				idxPat--
			} else {
				intResult = append([]rune{intPat[idxPat]}, intResult...)
				idxPat--
			}
			if idxPat < 0 {
				idxPat = len(intPat) - 1
			}
		}
		digits.WriteString(string(intResult))
	}

	if len(fracDigits) == 0 {
		return digits.String()
	}

	point := opts.Point
	if point == 0 {
		point = '.'
	}
	digits.WriteRune(point)

	if opts.FracPat == "" {
		digits.WriteString(string(fracDigits))
	} else {
		var fracResult []rune
		fracPat := []rune(opts.FracPat)

		idxPat := 0
		idxDig := 0
		for idxDig < len(fracDigits) {
			if fracPat[idxPat] == '0' {
				fracResult = append(fracDigits, fracResult[idxDig])
				idxDig++
				idxPat++
			} else {
				fracResult = append(fracDigits, fracPat[idxPat])
				idxPat++
			}
			if idxPat >= len(fracDigits) {
				idxPat = 0
			}
		}
		digits.WriteString(string(fracResult))
	}

	return digits.String()

}

func (c *Calc) FormatBigInt(v *big.Int) string {
	return fmt.Sprintf("%d", v)
}

func (c *Calc) FormatBigIntBase(v *big.Int, radix int) string {
	switch radix {
	case 16:
		return fmt.Sprintf("0x%x", v)
	case 8:
		return fmt.Sprintf("0o%o", v)
	case 2:
		return fmt.Sprintf("0b%b", v)
	}
	return c.FormatBigInt(v)
}

func (c *Calc) FormatBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func (c *Calc) FormatDecimal(v decimal.Decimal) string {
	fn, ok := roundModes[RoundMode]
	if !ok {
		log.Panicf("invalid rounding mode: %v", RoundMode)
	}
	return fn(v, Places).String()
}

func (c *Calc) FormatInt(i int) string {
	return fmt.Sprintf("%v", i)
}

func (c *Calc) FormatValue(v string) string {
	r := ParseRadix(v)
	switch {
	case r != 10:
		return v
	case c.IsBigInt(v):
		v := c.FormatBigIntBase(c.MustParseBigInt(v), r)
		return c.FormatNumberString(v, NumberFormat)
	case c.IsDecimal(v):
		v := c.FormatDecimal(c.MustParseDecimal(v))
		return c.FormatNumberString(v, NumberFormat)
	}
	return v
}
