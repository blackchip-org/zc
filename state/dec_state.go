package state

import (
	"fmt"
	"math/big"

	"github.com/cockroachdb/apd/v3"
)

const DecID = "dec"

type Dec struct {
	Context *apd.Context
}

func ForDec(state State) *Dec {
	conf := ForConf(state)
	s := &Dec{
		Context: apd.BaseContext.WithPrecision(conf.DecPrec),
	}
	s.Context.Rounding = DecRounder(conf.RoundingMode)
	return s
}

func DecRounder(rm big.RoundingMode) apd.Rounder {
	switch rm {
	case big.ToPositiveInf:
		return apd.RoundCeiling
	case big.ToZero:
		return apd.RoundDown
	case big.ToNegativeInf:
		return apd.RoundFloor
	case big.ToNearestAway:
		return apd.RoundHalfUp
	case big.ToNearestEven:
		return apd.RoundHalfEven
	case big.AwayFromZero:
		return apd.RoundUp
	default:
		panic(fmt.Errorf("unexpected rounding mode: %v", rm))
	}
}
