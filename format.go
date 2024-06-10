package zc

import (
	"fmt"
	"math/big"
	"strings"
)

func FormatList(vals ...any) string {
	var strs []string
	for _, val := range vals {
		strs = append(strs, String(val))
	}
	return strings.Join(strs, " | ")
}

func String(a any) string {
	if a == nil {
		return "{!:nil}"
	}
	switch v := a.(type) {
	case big.Int:
		return v.String()
	}
	return fmt.Sprintf("%v", a)
}
