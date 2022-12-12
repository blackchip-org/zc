package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/shopspring/decimal"
)

func cleanNumString(v string) string {
	var sb strings.Builder
	seps := NumberFormat.Separators()
	for _, ch := range v {
		if _, ok := seps[ch]; ok {
			continue
		}
		if unicode.Is(unicode.Sc, ch) {
			continue
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func ParseBigInt(v string) (*big.Int, error) {
	i := new(big.Int)
	_, ok := i.SetString(cleanNumString(v), 0)
	if !ok {
		return i, fmt.Errorf("expecting integer but got %v", v)
	}
	return i, nil
}

func IsBigInt(v string) bool {
	_, err := ParseBigInt(v)
	return err == nil
}

func MustParseBigInt(v string) *big.Int {
	i, err := ParseBigInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseBool(v string) (bool, error) {
	vl := strings.ToLower(v)
	switch vl {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, fmt.Errorf("expecting boolean but got %v", v)
}

func IsBool(v string) bool {
	_, err := ParseBool(v)
	return err == nil
}

func MustParseBool(v string) bool {
	b, err := ParseBool(v)
	if err != nil {
		panic(err)
	}
	return b
}

func ParseDecimal(v string) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(cleanNumString(v))
	if err != nil {
		return decimal.Zero, fmt.Errorf("expecting decimal but got %v", v)
	}
	return d, nil
}

func IsDecimal(v string) bool {
	_, err := ParseDecimal(v)
	return err == nil
}

func MustParseDecimal(v string) decimal.Decimal {
	d, err := ParseDecimal(v)
	if err != nil {
		panic(err)
	}
	return d
}

func ParseInt(v string) (int, error) {
	i, err := strconv.ParseInt(cleanNumString(v), 0, 64)
	if err != nil {
		return 0, fmt.Errorf("expecting integer but got %v", v)
	}
	return int(i), nil
}

func IsInt(v string) bool {
	_, err := ParseInt(v)
	return err == nil
}

func MustParseInt(v string) int {
	i, err := ParseInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

func ParseInt32(v string) (int32, error) {
	i, err := strconv.ParseInt(cleanNumString(v), 0, 32)
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
