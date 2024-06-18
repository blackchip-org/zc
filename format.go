package zc

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func Format(a any) string {
	switch v := a.(type) {
	case *big.Int:
		return v.String()
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
