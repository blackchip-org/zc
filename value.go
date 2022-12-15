package zc

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/shopspring/decimal"
)

type RoundingMode int

const (
	RoundingModeHalfUp RoundingMode = iota
	RoundingModeCeil
	RoundingModeDown
	RoundingModeFloor
	RoundingModeHalfEven
	RoundingModeUp
)

func (r RoundingMode) String() string {
	switch r {
	case RoundingModeHalfUp:
		return "half-up"
	case RoundingModeCeil:
		return "ceil"
	case RoundingModeDown:
		return "down"
	case RoundingModeFloor:
		return "floor"
	case RoundingModeHalfEven:
		return "half-even"
	case RoundingModeUp:
		return "up"
	}
	panic("unknown rounding mode")
}

func ParseRoundingMode(v string) (RoundingMode, bool) {
	switch strings.ToLower(v) {
	case "half-up":
		return RoundingModeHalfUp, true
	case "ceil":
		return RoundingModeCeil, true
	case "down":
		return RoundingModeDown, true
	case "floor":
		return RoundingModeFloor, true
	case "half-even":
		return RoundingModeHalfEven, true
	case "up":
		return RoundingModeUp, true
	}
	return 0, false
}

type RoundingFuncFix func(decimal.Decimal, int32) decimal.Decimal

var (
	RoundingFuncsFix = map[RoundingMode]RoundingFuncFix{
		RoundingModeCeil:     func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundCeil(places) },
		RoundingModeDown:     func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundDown(places) },
		RoundingModeFloor:    func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundFloor(places) },
		RoundingModeHalfUp:   func(d decimal.Decimal, places int32) decimal.Decimal { return d.Round(places) },
		RoundingModeHalfEven: func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundBank(places) },
		RoundingModeUp:       func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundUp(places) },
	}
)

type ValueOps struct {
	Places       int32
	RoundingMode RoundingMode
	IntPat       string
	Point        rune
	FracPat      string
}

func DefaultValueOps() ValueOps {
	return ValueOps{
		Places:       16,
		RoundingMode: RoundingModeHalfUp,
		IntPat:       ",000",
		Point:        '.',
	}
}

func (o ValueOps) parseDigits(sep rune, v string) ([]rune, []rune) {
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

func (o ValueOps) FormatNumberString(v string) string {
	var digits strings.Builder
	intDigits, fracDigits := o.parseDigits('.', v)

	if o.IntPat == "" {
		digits.WriteString(string(intDigits))
	} else {
		var intResult []rune
		intPat := []rune(o.IntPat)

		idxPat := len(o.IntPat) - 1
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

	point := o.Point
	if point == 0 {
		point = '.'
	}
	digits.WriteRune(point)

	if o.FracPat == "" {
		digits.WriteString(string(fracDigits))
	} else {
		var fracResult []rune
		fracPat := []rune(o.FracPat)

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

func (o ValueOps) FormatBigInt(v *big.Int) string {
	return fmt.Sprintf("%d", v)
}

func (o ValueOps) FormatBigIntBase(v *big.Int, radix int) string {
	switch radix {
	case 16:
		return fmt.Sprintf("0x%x", v)
	case 8:
		return fmt.Sprintf("0o%o", v)
	case 2:
		return fmt.Sprintf("0b%b", v)
	}
	return o.FormatBigInt(v)
}

func (o ValueOps) FormatBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func (o ValueOps) FormatFixed(v decimal.Decimal) string {
	fn, ok := RoundingFuncsFix[o.RoundingMode]
	if !ok {
		log.Panicf("invalid rounding mode: %v", o.RoundingMode)
	}

	return fn(v, o.Places).String()
}

func (o ValueOps) FormatInt(i int) string {
	return fmt.Sprintf("%v", i)
}

func (o ValueOps) FormatValue(v string) string {
	r := ParseRadix(v)
	switch {
	case r != 10:
		return v
	case o.IsBigInt(v):
		v := o.FormatBigIntBase(o.MustParseBigInt(v), r)
		return o.FormatNumberString(v)
	case o.IsFixed(v):
		v := o.FormatFixed(o.MustParseFixed(v))
		return o.FormatNumberString(v)
	}
	return v
}

func (o ValueOps) cleanNumString(v string) string {
	var sb strings.Builder
	// FIXME
	// seps := c.Settings.NumberFormat.Separators()
	for _, ch := range v {
		// if _, ok := seps[ch]; ok {
		// 	continue
		// }
		if ch == ',' {
			continue
		}
		if unicode.Is(unicode.Sc, ch) {
			continue
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func (o ValueOps) ParseBigInt(v string) (*big.Int, error) {
	i := new(big.Int)
	_, ok := i.SetString(o.cleanNumString(v), 0)
	if !ok {
		return i, fmt.Errorf("expecting integer but got %v", v)
	}
	return i, nil
}

func (o ValueOps) IsBigInt(v string) bool {
	_, err := o.ParseBigInt(v)
	return err == nil
}

func (o ValueOps) MustParseBigInt(v string) *big.Int {
	i, err := o.ParseBigInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

func (o ValueOps) ParseBool(v string) (bool, error) {
	vl := strings.ToLower(v)
	switch vl {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, fmt.Errorf("expecting boolean but got %v", v)
}

func (o ValueOps) IsBool(v string) bool {
	_, err := o.ParseBool(v)
	return err == nil
}

func (o ValueOps) MustParseBool(v string) bool {
	b, err := o.ParseBool(v)
	if err != nil {
		panic(err)
	}
	return b
}

func (o ValueOps) ParseFixed(v string) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(o.cleanNumString(v))
	if err != nil {
		return decimal.Zero, fmt.Errorf("expecting fixed-point but got %v", v)
	}
	return d, nil
}

func (o ValueOps) IsFixed(v string) bool {
	_, err := o.ParseFixed(v)
	return err == nil
}

func (o ValueOps) MustParseFixed(v string) decimal.Decimal {
	d, err := o.ParseFixed(v)
	if err != nil {
		panic(err)
	}
	return d
}

func (o ValueOps) ParseInt(v string) (int, error) {
	i, err := strconv.ParseInt(o.cleanNumString(v), 0, 64)
	if err != nil {
		return 0, fmt.Errorf("expecting integer but got %v", v)
	}
	return int(i), nil
}

func (o ValueOps) IsInt(v string) bool {
	_, err := o.ParseInt(v)
	return err == nil
}

func (o ValueOps) MustParseInt(v string) int {
	i, err := o.ParseInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

func (o ValueOps) ParseInt32(v string) (int32, error) {
	i, err := strconv.ParseInt(o.cleanNumString(v), 0, 32)
	if err != nil {
		return 0, fmt.Errorf("expecting int32 but got %v", v)
	}
	return int32(i), nil
}

func ParseRadix(v string) int {
	if len(v) < 2 {
		return 10
	}
	prefix := strings.ToLower(v[:2])
	switch prefix {
	case "0b":
		return 2
	case "0o":
		return 8
	case "0x":
		return 16
	}
	return 10
}
