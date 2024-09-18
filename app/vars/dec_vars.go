package vars

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6/msg"
	"github.com/blackchip-org/zc/v6/pkg/coll"
	"github.com/cockroachdb/apd/v3"
)

const DecID = "dec"

const (
	DefaultDecPrec         = 28
	DefaultDecRoundingMode = big.ToNearestEven
)

const (
	RoundingModeCeil     = "ceil"
	RoundingModeDown     = "down"
	RoundingModeFloor    = "floor"
	RoundingModeHalfUp   = "half.up"
	RoundingModeHalfEven = "half.even"
	RoundingModeUp       = "up"
)

type Dec struct {
	Math         *apd.Context
	RoundingMode big.RoundingMode
}

func (d *Dec) GetRoundingMode() string {
	switch d.RoundingMode {
	case big.ToPositiveInf:
		return RoundingModeCeil
	case big.ToZero:
		return RoundingModeDown
	case big.ToNegativeInf:
		return RoundingModeFloor
	case big.ToNearestAway:
		return RoundingModeHalfUp
	case big.ToNearestEven:
		return RoundingModeHalfEven
	case big.AwayFromZero:
		return RoundingModeUp
	default:
		panic("unreachable")
	}
}

func (d *Dec) SetRoundingMode(rm string) error {
	switch rm {
	case RoundingModeCeil:
		d.RoundingMode = big.ToPositiveInf
	case RoundingModeDown:
		d.RoundingMode = big.ToZero
	case RoundingModeFloor:
		d.RoundingMode = big.ToNegativeInf
	case RoundingModeHalfUp:
		d.RoundingMode = big.ToNearestAway
	case RoundingModeHalfEven:
		d.RoundingMode = big.ToNearestEven
	case RoundingModeUp:
		d.RoundingMode = big.AwayFromZero
	default:
		return msg.ErrInvalidRoundingMode(rm)
	}
	d.Math.Rounding = DecRounder(d.RoundingMode)
	return nil
}

func ForDec(state coll.State) *Dec {
	dv, ok := state.Var(DecID)
	if !ok {
		context := apd.BaseContext.WithPrecision(DefaultDecPrec)
		dec := &Dec{
			Math:         context,
			RoundingMode: DefaultDecRoundingMode,
		}
		dec.Math.Rounding = DecRounder(DefaultDecRoundingMode)
		state.NewVar(DecID, dec)
		dv = dec
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
