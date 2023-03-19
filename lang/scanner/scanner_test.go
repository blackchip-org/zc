package scanner

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
			s.SetString("", test.data)
			tok := s.Scan(test.fn)
			if tok != test.tok {
				t.Errorf("\n tok have: %v \n tok want: %v", tok, test.tok)
			}
			rest := s.Scan(Remaining)
			if rest != test.rest {
				t.Errorf("\n rest have: %v \n rest want: %v", rest, test.rest)
			}
		})
	}
}

func TestThisPos(t *testing.T) {
	data := "123 567 9\n12 456"
	wants := []Pos{
		NewPos("", 1, 5),
		NewPos("", 1, 9),
		NewPos("", 2, 1),
		NewPos("", 2, 4),
		NewPos("", 2, 7),
	}
	s := NewString("", data)
	for _, want := range wants {
		s.Scan(Word)
		s.ScanWhitespace()
		if want != s.ThisPos {
			t.Fatalf("\n want: %v \n have: %v", want, s.ThisPos)
		}
	}
}

func TestOutPos(t *testing.T) {
	data := "123 567 9\n12 456"
	wants := []Pos{
		NewPos("", 1, 1),
		NewPos("", 1, 5),
		NewPos("", 1, 9),
		NewPos("", 2, 1),
		NewPos("", 2, 4),
	}
	s := NewString("", data)
	for _, want := range wants {
		s.Scan(Word)
		if want != s.OutPos {
			t.Fatalf("\n want: %v \n have: %v", want, s.OutPos)
		}
		s.ScanWhitespace()
	}
}

func TestManual(t *testing.T) {
	s := NewString("", "1234")
	s.Start()
	s.Keep()
	s.Keep()
	want := "12"
	have := s.Emit()
	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}

	s.Start()
	s.Keep()
	s.Keep()
	want = "34"
	have = s.Emit()
	if have != want {
		t.Errorf("\n have: %v \n want: %v", have, want)
	}
}
