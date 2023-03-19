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
		Digit: IsDigit01,
	}
	DecDef = NumberDef{
		Digit:  IsDigit09,
		DecSep: Rune('.'),
		Sign:   Rune2('+', '-'),
	}
	FloatDef = NumberDef{
		Digit:    IsDigit09,
		DecSep:   Rune('.'),
		Sign:     Rune2('+', '-'),
		Exponent: Rune2('e', 'E'),
	}
	HexDef = NumberDef{
		Digit: IsDigit0F,
	}
	IntDef = NumberDef{
		Digit: IsDigit09,
		Sign:  Rune2('+', '-'),
	}
	OctDef = NumberDef{
		Digit: IsDigit07,
	}
	StringDef = QuotedDef{
		Escape: Rune('\\'),
	}
	UIntDef = NumberDef{
		Digit: IsDigit09,
	}
	UDecRef = NumberDef{
		Digit:  IsDigit09,
		DecSep: Rune('.'),
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
