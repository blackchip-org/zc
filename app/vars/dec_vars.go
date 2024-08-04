package vars

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6/pkg/coll"
	"github.com/cockroachdb/apd/v3"
)

const DecID = "dec"

type Dec struct {
	Context *apd.Context
}

func ForDec(state coll.State) *Dec {
	dv, ok := state.Var(DecID)
	if !ok {
		context := apd.BaseContext.WithPrecision(16)
		context.MaxExponent = 100_000_000
		context.MinExponent = -100_000_000
		dv = &Dec{
			Context: context,
		}
		state.NewVar(DecID, dv)
	}
	return dv.(*Dec)
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
