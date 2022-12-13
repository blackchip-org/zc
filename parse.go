package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"github.com/shopspring/decimal"
)

func (c *Calc) cleanNumString(v string) string {
	var sb strings.Builder
	seps := c.Settings.NumberFormat.Separators()
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

func (c *Calc) ParseBigInt(v string) (*big.Int, error) {
	i := new(big.Int)
	_, ok := i.SetString(c.cleanNumString(v), 0)
	if !ok {
		return i, fmt.Errorf("expecting integer but got %v", v)
	}
	return i, nil
}

func (c *Calc) IsBigInt(v string) bool {
	_, err := c.ParseBigInt(v)
	return err == nil
}

func (c *Calc) MustParseBigInt(v string) *big.Int {
	i, err := c.ParseBigInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

func (c *Calc) ParseBool(v string) (bool, error) {
	vl := strings.ToLower(v)
	switch vl {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, fmt.Errorf("expecting boolean but got %v", v)
}

func (c *Calc) IsBool(v string) bool {
	_, err := c.ParseBool(v)
	return err == nil
}

func (c *Calc) MustParseBool(v string) bool {
	b, err := c.ParseBool(v)
	if err != nil {
		panic(err)
	}
	return b
}

func (c *Calc) ParseDecimal(v string) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(c.cleanNumString(v))
	if err != nil {
		return decimal.Zero, fmt.Errorf("expecting decimal but got %v", v)
	}
	return d, nil
}

func (c *Calc) IsDecimal(v string) bool {
	_, err := c.ParseDecimal(v)
	return err == nil
}

func (c *Calc) MustParseDecimal(v string) decimal.Decimal {
	d, err := c.ParseDecimal(v)
	if err != nil {
		panic(err)
	}
	return d
}

func (c *Calc) ParseInt(v string) (int, error) {
	i, err := strconv.ParseInt(c.cleanNumString(v), 0, 64)
	if err != nil {
		return 0, fmt.Errorf("expecting integer but got %v", v)
	}
	return int(i), nil
}

func (c *Calc) IsInt(v string) bool {
	_, err := c.ParseInt(v)
	return err == nil
}

func (c *Calc) MustParseInt(v string) int {
	i, err := c.ParseInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

func (c *Calc) ParseInt32(v string) (int32, error) {
	i, err := strconv.ParseInt(c.cleanNumString(v), 0, 32)
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
