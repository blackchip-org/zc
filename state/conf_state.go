package state

import "github.com/blackchip-org/zc/v6"

const ConfID = "conf"

type Conf struct {
	FloatPrec uint
	DecPrec   uint32
}

func ForConf(e *zc.OpEnv) *Conf {
	s, ok := e.State[ConfID]
	if !ok {
		s = &Conf{
			FloatPrec: 53,
			DecPrec:   16,
		}
		e.State[ConfID] = s
	}
	return s.(*Conf)
}
