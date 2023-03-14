package fscan

import (
	"bufio"
	"io"
	"strings"
)

const End = rune(-1)

type Scanner struct {
	Error  error
	Behind rune
	This   rune
	Ahead  rune
	Out    strings.Builder
	Line   int
	Column int
	src    *bufio.Reader
}

func NewForReader(src io.Reader) *Scanner {
	s := &Scanner{}
	s.SetReader(src)
	return s
}

func NewForString(src string) *Scanner {
	s := &Scanner{}
	s.SetString(src)
	return s
}

func (s *Scanner) SetReader(src io.Reader) {
	s.src = bufio.NewReader(src)
	s.init()
}

func (s *Scanner) SetString(src string) {
	s.SetReader(strings.NewReader(src))
}

func (s *Scanner) init() {
	s.Error = nil
	s.Behind, s.This, s.Ahead = 0, 0, 0
	s.Next()
	s.Next()
	s.Behind = End
	s.Line = 1
	s.Column = 1
}

func (s *Scanner) NextToken(fn Func) string {
	s.Out.Reset()
	fn(s)
	token := s.Out.String()
	s.Out.Reset()
	return token
}

func (s *Scanner) Next() {
	if s.This == End || s.src == nil {
		return
	}

	s.Behind = s.This
	s.This = s.Ahead

	r, _, err := s.src.ReadRune()
	if err != nil {
		s.Ahead = End
		if err != io.EOF {
			s.Error = err
		}
	} else {
		s.Ahead = r
	}

	if s.This == '\n' {
		s.Line++
		s.Column = 0
	} else {
		s.Column++
	}
}

func (s *Scanner) Keep() {
	s.Out.WriteRune(s.This)
	s.Next()
}

func (s *Scanner) IsEnd() bool {
	return s.This == End
}

func (s *Scanner) ScanWhitespace() string {
	return s.NextToken(Whitespace)
}

func (s *Scanner) ScanUDec() string {
	return s.NextToken(UDec)
}
