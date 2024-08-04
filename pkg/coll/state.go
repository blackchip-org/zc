package coll

type State interface {
	Var(string) (any, bool)
	NewVar(string, any)
}
