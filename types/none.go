package types

type gNone struct{}

func (g gNone) Type() Type     { return None }
func (g gNone) Format() string { return "" }
func (g gNone) String() string { return "None" }
func (g gNone) Value() any     { return nil }

type noneType struct{}

func (t noneType) String() string { return "None" }

func (t noneType) ParseGeneric(s string) (Generic, bool) {
	return gNone{}, false
}
