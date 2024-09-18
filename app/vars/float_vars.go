package vars

import "github.com/blackchip-org/zc/v6/pkg/coll"

const (
	FloatID          = "float"
	DefaultFloatPrec = 113
)

type Float struct {
	Prec uint
}

func ForFloat(state coll.State) *Float {
	v, ok := state.Var(FloatID)
	if !ok {
		v = &Float{Prec: DefaultFloatPrec}
		state.NewVar(FloatID, v)
	}
	return v.(*Float)
}
