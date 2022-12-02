package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func ParseBigInt(v string) (*big.Int, error) {
	i := new(big.Int)
	_, ok := i.SetString(v, 0)
	if !ok {
		return i, fmt.Errorf("expecting integer but got %v", v)
	}
	return i, nil
}

func IsBigInt(v string) bool {
	_, err := ParseBigInt(v)
	return err == nil
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

func ParseDecimal(v string) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(v)
	if err != nil {
		return decimal.Zero, fmt.Errorf("expecting decimal but got %v", v)
	}
	return d, nil
}

func IsDecimal(v string) bool {
	_, ok := ParseDecimal(v)
	return ok == nil
}

func ParseInt(v string) (int, error) {
	i, err := strconv.ParseInt(v, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("expecting integer but got %v", v)
	}
	return int(i), nil
}

func ParseInt32(v string) (int32, error) {
	i, err := strconv.ParseInt(v, 0, 32)
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
