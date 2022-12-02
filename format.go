package zc

import (
	"math/big"

	"github.com/shopspring/decimal"
)

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
	return v.String()
}
