package state

const ConfID = "conf"

type Conf struct {
	FloatPrec uint
	DecPrec   uint32
}

func ForConf(state State) *Conf {
	s, ok := state[ConfID]
	if !ok {
		s = &Conf{
			FloatPrec: 64,
			DecPrec:   16,
		}
		state[ConfID] = s
	}
	return s.(*Conf)
}
