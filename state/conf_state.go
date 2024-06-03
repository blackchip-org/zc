package state

import "math/big"

const ConfID = "conf"

type Conf struct {
	FloatPrec    uint
	DecPrec      uint32
	RoundingMode big.RoundingMode
}

func ForConf(state State) *Conf {
	s, ok := state[ConfID]
	if !ok {
		s = &Conf{
			FloatPrec:    64,
			DecPrec:      16,
			RoundingMode: big.ToNearestEven,
		}
		state[ConfID] = s
	}
	return s.(*Conf)
}
