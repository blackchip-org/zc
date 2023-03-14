package fscan

import (
	"fmt"
	"testing"
)

func TestScanners(t *testing.T) {
	tests := []struct {
		fn   Func
		data string
		tok  string
		rest string
	}{
		{Bin, "010123", "0101", "23"},
		{Dec, "12.345 678", "12.345", " 678"},
		{Dec, "-12.345 678", "-12.345", " 678"},
		{Dec, "-12.34-5 678", "-12.34", "-5 678"},
		{Dec, "-12.34.5 678", "-12.34", ".5 678"},
		{Dec, "-123.45e10 abc", "-123.45", "e10 abc"},
		{Float, "-123.45e10 abc", "-123.45e10", " abc"},
		{Float, "-123.45E-10 abc", "-123.45E-10", " abc"},
		{Hex, "cdefghi", "cdef", "ghi"},
		{Int, "1234 567", "1234", " 567"},
		{Int, "-1234 567", "-1234", " 567"},
		{Int, "0x1234 567", "0", "x1234 567"},
		{Oct, "56789", "567", "89"},
		{Remaining, "123 456", "123 456", ""},
		{UDec, "12.345 678", "12.345", " 678"},
		{UDec, "-12.345 678", "", "-12.345 678"},
		{UInt, "1234 567", "1234", " 567"},
		{UInt, "-1234 567", "", "-1234 567"},
		{Whitespace, "     1234", "     ", "1234"},
		{Word, "foo bar", "foo", " bar"},
	}

	var s Scanner
	for i, test := range tests {
		t.Run(fmt.Sprintf("[%v] %v", i, test.data), func(t *testing.T) {
			s.SetString(test.data)
			tok := s.NextToken(test.fn)
			if tok != test.tok {
				t.Errorf("\n tok have: %v \n tok want: %v", tok, test.tok)
			}
			rest := s.NextToken(Remaining)
			if rest != test.rest {
				t.Errorf("\n rest have: %v \n rest want: %v", rest, test.rest)
			}
		})
	}
}

func TestSepScanners(t *testing.T) {
	commaUnder := NumberSepDef{
		Left:  NewRule(Comma, Discard),
		Right: NewRule(Underscore, Discard),
	}

	tests := []struct {
		numDef NumberDef
		sepDef NumberSepDef
		data   string
		tok    string
		rest   string
	}{
		{DecDef, commaUnder, "1,234.567_890 abc", "1234.567890", " abc"},
		{DecDef, commaUnder, "1,23,. abc", "123", ",. abc"},
		{DecDef, commaUnder, ",123", "", ",123"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%v] %v", i, test.data), func(t *testing.T) {
			l := NewForString(test.data)
			tok := l.NextToken(SepNumberFunc(test.numDef, test.sepDef))
			if tok != test.tok {
				t.Errorf("\n tok have: %v \n tok want: %v", tok, test.tok)
			}
			rest := l.NextToken(Remaining)
			if rest != test.rest {
				t.Errorf("\n rest have: %v \n rest want: %v", rest, test.rest)
			}
		})
	}
}

func TestPos(t *testing.T) {
	data := "123 567 9\n12 456"
	want := []int{
		1, 5,
		1, 9,
		2, 1,
		2, 4,
		2, 7,
	}
	s := NewForString(data)
	for i := 0; i < len(want); i += 2 {
		s.NextToken(Word)
		s.ScanWhitespace()
		if want[i] != s.Line || want[i+1] != s.Column {
			t.Fatalf("\n want: %v.%v \n have: %v.%v", want[i], want[i+1], s.Line, s.Column)
		}
	}
}
