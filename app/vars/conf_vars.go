package vars

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6/pkg/coll"
	"github.com/cockroachdb/apd/v3"
)

const ConfID = "conf"

const (
	DefaultPrec         = 16
	DefaultRoundingMode = big.ToNearestEven
)

const (
	RoundingModeCeil     = "ceil"
	RoundingModeDown     = "down"
	RoundingModeFloor    = "floor"
	RoundingModeHalfUp   = "half.up"
	RoundingModeHalfEven = "half.even"
	RoundingModeUp       = "up"
)

type Conf struct {
	DecMath      *apd.Context
	RoundingMode big.RoundingMode
}

func (d *Conf) GetRoundingMode() string {
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
		panic(fmt.Errorf("invalid rounding mode: %v", d.RoundingMode))
	}
}

func (d *Conf) SetRoundingMode(rm string) error {
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
		return errors.New("invalid rounding mode")
	}
	d.DecMath.Rounding = DecRounder(d.RoundingMode)
	return nil
}

func ForConf(state coll.State) *Conf {
	dv, ok := state.Var(ConfID)
	if !ok {
		context := apd.BaseContext.WithPrecision(DefaultPrec)
		dec := &Conf{
			DecMath:      context,
			RoundingMode: DefaultRoundingMode,
		}
		dec.DecMath.Rounding = DecRounder(DefaultRoundingMode)
		state.NewVar(ConfID, dec)
		dv = dec
	}
	return dv.(*Conf)
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
