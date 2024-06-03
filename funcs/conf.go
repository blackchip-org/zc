package funcs

import (
	"fmt"
	"math/big"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/state"
)

func FloatPrecSet(e *zc.OpEnv) {
	s := state.ForConf(e.State)
	p := e.Args[0].(uint)
	s.FloatPrec = p
	e.Info = fmt.Sprintf("precision set to %v", p)
}

func FloatPrecGet(e *zc.OpEnv) {
	s := state.ForConf(e.State)
	e.Returns = []any{s.FloatPrec}
	e.Annos = []string{"precision"}
}

func RoundingModeSet(e *zc.OpEnv) {
	s := state.ForConf(e.State)
	m := e.Args[0].(string)
	switch m {
	case "ceil":
		s.RoundingMode = big.ToPositiveInf
	case "down":
		s.RoundingMode = big.ToZero
	case "floor":
		s.RoundingMode = big.ToNegativeInf
	case "half.up":
		s.RoundingMode = big.ToNearestAway
	case "half.even":
		s.RoundingMode = big.ToNearestEven
	case "up":
		s.RoundingMode = big.AwayFromZero
	default:
		e.Err = zc.ErrInvalidArg(e, "invalid rounding mode: %v", m)
		return
	}
	e.Info = fmt.Sprintf("rounding mode set to %v", m)
}
