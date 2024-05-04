package zc

import (
	"fmt"
	"strings"
)

func FormatStackValues(vals ...any) string {
	var strs []string
	for _, val := range vals {
		str := fmt.Sprintf("%v", val)
		strs = append(strs, str)
	}
	return strings.Join(strs, " | ")
}

func FormatStack(items []Item) string {
	var strs []string
	for _, item := range items {
		str := fmt.Sprintf("%v", item.Value)
		strs = append(strs, str)
	}
	return strings.Join(strs, " | ")
}
