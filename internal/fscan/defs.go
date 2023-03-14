package fscan

import "unicode"

type Rule struct {
	Is     RuneClass
	Action Func
}

func NewRule(class RuneClass, action Func) Rule {
	return Rule{class, action}
}

var NoRule = NewRule(Never, None)

type NumberDef struct {
	Digit    Rule
	DecSep   Rule
	Sign     Rule
	Exponent Rule
}

type NumberSepDef struct {
	Left  Rule
	Right Rule
}

var (
	NoSepDef = NumberSepDef{
		Left:  NewRule(Comma, Discard),
		Right: NoRule,
	}
	BinDef = NumberDef{
		Digit:    NewRule(Digit01, Keep),
		DecSep:   NoRule,
		Sign:     NoRule,
		Exponent: NoRule,
	}
	DecDef = NumberDef{
		Digit:    NewRule(Digit09, Keep),
		DecSep:   NewRule(Period, Keep),
		Sign:     NewRule(PlusMinus, Keep),
		Exponent: NoRule,
	}
	FloatDef = NumberDef{
		Digit:    NewRule(Digit09, Keep),
		DecSep:   NewRule(Period, Keep),
		Sign:     NewRule(PlusMinus, Keep),
		Exponent: NewRule(ExponentE, Keep),
	}
	HexDef = NumberDef{
		Digit:    NewRule(Digit0F, Keep),
		DecSep:   NoRule,
		Sign:     NoRule,
		Exponent: NoRule,
	}
	IntDef = NumberDef{
		Digit:    NewRule(Digit09, Keep),
		DecSep:   NoRule,
		Sign:     NewRule(PlusMinus, Keep),
		Exponent: NoRule,
	}
	OctDef = NumberDef{
		Digit:    NewRule(Digit07, Keep),
		DecSep:   NoRule,
		Sign:     NoRule,
		Exponent: NoRule,
	}
	UIntDef = NumberDef{
		Digit:    NewRule(Digit09, Keep),
		DecSep:   NoRule,
		Sign:     NoRule,
		Exponent: NoRule,
	}
	UDecRef = NumberDef{
		Digit:    NewRule(Digit09, Keep),
		DecSep:   NewRule(Period, Keep),
		Sign:     NoRule,
		Exponent: NoRule,
	}
)

var (
	Bin        = NumberFunc(BinDef)
	Dec        = NumberFunc(DecDef)
	Float      = NumberFunc(FloatDef)
	Hex        = NumberFunc(HexDef)
	Int        = NumberFunc(IntDef)
	Oct        = NumberFunc(OctDef)
	Remaining  = WhileFunc(Always)
	Whitespace = WhileFunc(unicode.IsSpace)
	UInt       = NumberFunc(UIntDef)
	UDec       = NumberFunc(UDecRef)
	Word       = UntilFunc(unicode.IsSpace)
)
