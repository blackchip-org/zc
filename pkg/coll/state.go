package coll

type State interface {
	Var(string) (any, bool)
	NewVar(string, any)
}

var NullState = nullState{}

type nullState struct{}

func (s nullState) Var(string) (any, bool) {
	return nil, false
}

func (s nullState) NewVar(string, any) {}
