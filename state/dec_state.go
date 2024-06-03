package state

import (
	"github.com/cockroachdb/apd/v3"
)

const DecID = "dec"

type Dec struct {
	Context *apd.Context
}

func ForDec(state State) *Dec {
	conf := ForConf(state)
	s, ok := state[DecID]
	if !ok {
		s = &Dec{
			Context: apd.BaseContext.WithPrecision(conf.DecPrec),
		}
		state[DecID] = s
	}
	d := s.(*Dec)
	d.Context.Precision = conf.DecPrec
	return d
}
