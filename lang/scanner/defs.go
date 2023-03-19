package scanner

import "unicode"

type NumberDef struct {
	Digit    RuneClass
	DecSep   RuneClass
	Sign     RuneClass
	Exponent RuneClass
}

var (
	BinDef = NumberDef{
		Digit:    Digit01,
		DecSep:   Never,
		Sign:     Never,
		Exponent: Never,
	}
	DecDef = NumberDef{
		Digit:    Digit09,
		DecSep:   Period,
		Sign:     PlusMinus,
		Exponent: Never,
	}
	FloatDef = NumberDef{
		Digit:    Digit09,
		DecSep:   Period,
		Sign:     PlusMinus,
		Exponent: ExponentE,
	}
	HexDef = NumberDef{
		Digit:    Digit0F,
		DecSep:   Never,
		Sign:     Never,
		Exponent: Never,
	}
	IntDef = NumberDef{
		Digit:    Digit09,
		DecSep:   Never,
		Sign:     PlusMinus,
		Exponent: Never,
	}
	OctDef = NumberDef{
		Digit:    Digit07,
		DecSep:   Never,
		Sign:     Never,
		Exponent: Never,
	}
	UIntDef = NumberDef{
		Digit:    Digit09,
		DecSep:   Never,
		Sign:     Never,
		Exponent: Never,
	}
	UDecRef = NumberDef{
		Digit:    Digit09,
		DecSep:   Period,
		Sign:     Never,
		Exponent: Never,
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
