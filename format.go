package zc

import "math/big"

func FormatBigInt(v *big.Int) string {
	return v.String()
}

func FormatBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}
