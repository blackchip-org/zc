package zc

import (
	"log"
	"math/big"
	"strconv"

	"github.com/shopspring/decimal"
)

type decRoundFunc func(decimal.Decimal, int32) decimal.Decimal

var roundModes = map[string]decRoundFunc{
	"ceil":      func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundCeil(places) },
	"down":      func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundDown(places) },
	"floor":     func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundFloor(places) },
	"half-up":   func(d decimal.Decimal, places int32) decimal.Decimal { return d.Round(places) },
	"half-even": func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundBank(places) },
	"up":        func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundUp(places) },
}

func FormatBigInt(v *big.Int) string {
	return v.String()
}

func FormatBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func FormatDecimal(v decimal.Decimal) string {
	fn, ok := roundModes[RoundMode]
	if !ok {
		log.Panicf("invalid rounding mode: %v", RoundMode)
	}
	return fn(v, Places).String()
}

func FormatInt(i int) string {
	return strconv.Itoa(i)
}
