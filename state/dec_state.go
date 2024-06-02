package state

import (
	"github.com/blackchip-org/zc/v6"
	"github.com/cockroachdb/apd/v3"
)

const DecID = "dec"

type Dec struct {
	Context *apd.Context
}

func ForDec(e *zc.OpEnv) *Dec {
	conf := ForConf(e)
	s, ok := e.State[DecID]
	if !ok {
		s = &Dec{
			Context: apd.BaseContext.WithPrecision(conf.DecPrec),
		}
		e.State[DecID] = s
	}
	d := s.(*Dec)
	d.Context.Precision = conf.DecPrec
	return d
}
