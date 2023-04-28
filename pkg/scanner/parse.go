package scanner

import (
	"fmt"
	"strconv"
)

func ParseFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(fmt.Sprintf("not a float64: %v", s))
	}
	return f
}

func ParseInt(s string) int {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		panic(fmt.Sprintf("not an int: %v", s))
	}
	return int(i)
}
