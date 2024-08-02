package funcs

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
)

const (
	RoundingModeCeil     = "ceil"
	RoundingModeDown     = "down"
	RoundingModeFloor    = "floor"
	RoundingModeHalfUp   = "half.up"
	RoundingModeHalfEven = "half.even"
	RoundingModeUp       = "up"
)

func RoundingModeSet(e *zc.OpEnv) {
	s := state.ForConf(e.State)
	m := zc.String.Pop(e)
	switch m {
	case RoundingModeCeil:
		s.RoundingMode = big.ToPositiveInf
	case RoundingModeDown:
		s.RoundingMode = big.ToZero
	case RoundingModeFloor:
		s.RoundingMode = big.ToNegativeInf
	case RoundingModeHalfUp:
		s.RoundingMode = big.ToNearestAway
	case RoundingModeHalfEven:
		s.RoundingMode = big.ToNearestEven
	case RoundingModeUp:
		s.RoundingMode = big.AwayFromZero
	default:
		e.Err = zc.ErrInvalidArg(e, "invalid rounding mode: %v", m)
		return
	}
	e.Info = fmt.Sprintf("rounding mode set to %v", m)
}

func RoundingModeGet(e *zc.OpEnv) {
	s := state.ForConf(e.State)
	var m string
	switch s.RoundingMode {
	case big.ToPositiveInf:
		m = RoundingModeCeil
	case big.ToZero:
		m = RoundingModeDown
	case big.ToNegativeInf:
		m = RoundingModeFloor
	case big.ToNearestAway:
		m = RoundingModeHalfUp
	case big.ToNearestEven:
		m = RoundingModeHalfEven
	case big.AwayFromZero:
		m = RoundingModeUp
	default:
		panic(fmt.Errorf("invalid rounding mode: %v", s.RoundingMode))
	}
	zc.String.Push(e, m)
}
