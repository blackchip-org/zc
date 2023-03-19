package scanner

import "unicode"

func OptionalClass(r RuneClass) RuneClass {
	if r == nil {
		return Never
	}
	return r
}

type NumberDef struct {
	Digit    RuneClass
	DecSep   RuneClass
	Sign     RuneClass
	Exponent RuneClass
}

type QuotedDef struct {
	Escape    RuneClass
	EscapeMap map[rune]rune
	AltEnd    RuneClass
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
	StringDef = QuotedDef{
		Escape: Backslash,
		AltEnd: Never,
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
	String     = QuotedFunc(StringDef)
	Whitespace = WhileFunc(unicode.IsSpace)
	UInt       = NumberFunc(UIntDef)
	UDec       = NumberFunc(UDecRef)
	Word       = UntilFunc(unicode.IsSpace)
)
