package state

import (
	"math"
	"math/big"
)

const ConfID = "conf"
const SettingDecPrec = "dec.prec"

type Conf struct {
	FloatPrec      uint
	DecPrec        uint32
	DecMaxExponent int32
	DecMinExponent int32
	RoundingMode   big.RoundingMode
}

func (c Conf) DecPrecAsInt32() int32 {
	if c.DecPrec > math.MaxInt32 {
		return int32(math.MaxInt32)
	}
	return int32(c.DecPrec)
}

func ForConf(state State) *Conf {
	s, ok := state[ConfID]
	if !ok {
		s = &Conf{
			FloatPrec:      53,
			DecPrec:        16,
			DecMaxExponent: 100_000,
			DecMinExponent: -100_000,
			RoundingMode:   big.ToNearestEven,
		}
		state[ConfID] = s
	}
	return s.(*Conf)
}
