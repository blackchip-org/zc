package zc

import (
	"log"
	"math/big"

	"github.com/shopspring/decimal"
	"golang.org/x/text/message"
)

type decRoundFunc func(decimal.Decimal, int32) decimal.Decimal

var (
	printer    = message.NewPrinter(message.MatchLanguage("en"))
	roundModes = map[string]decRoundFunc{
		"ceil":      func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundCeil(places) },
		"down":      func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundDown(places) },
		"floor":     func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundFloor(places) },
		"half-up":   func(d decimal.Decimal, places int32) decimal.Decimal { return d.Round(places) },
		"half-even": func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundBank(places) },
		"up":        func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundUp(places) },
	}
)

func FormatBigInt(v *big.Int) string {
	if v.IsUint64() {
		return printer.Sprint(v.Uint64())
	} else if v.IsInt64() {
		return printer.Sprint(v.Int64())
	}
	return printer.Sprintf("%d", v)
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
	return printer.Sprint(i)
}

func FormatValue(v string) string {
	if IsInt(v) {
		return FormatInt(MustParseInt(v))
	}
	return v
}
