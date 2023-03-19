package scanner

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

const End = rune(-1)

type Scanner struct {
	Error   error
	Behind  rune
	This    rune
	Ahead   rune
	Out     strings.Builder
	OutPos  Pos
	ThisPos Pos
	src     *bufio.Reader
}

func New(name string, src io.Reader) *Scanner {
	s := &Scanner{}
	s.SetReader(name, src)
	return s
}

func NewString(name string, src string) *Scanner {
	s := &Scanner{}
	s.SetString(name, src)
	return s
}

func NewBytes(name string, src []byte) *Scanner {
	s := &Scanner{}
	s.SetBytes(name, src)
	return s
}

func (s *Scanner) SetReader(name string, src io.Reader) {
	s.src = bufio.NewReader(src)
	s.ThisPos.Name = name
	s.init()
}

func (s *Scanner) SetString(name string, src string) {
	s.SetReader(name, strings.NewReader(src))
}

func (s *Scanner) SetBytes(name string, src []byte) {
	s.SetReader(name, bytes.NewReader(src))
}

func (s *Scanner) init() {
	s.Error = nil
	s.Behind, s.This, s.Ahead = 0, 0, 0
	s.Next()
	s.Next()
	s.Behind = End
	s.ThisPos.Line = 1
	s.ThisPos.Column = 1
}

func (s *Scanner) Scan(fn Func) string {
	s.Out.Reset()
	s.OutPos = s.ThisPos
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
		s.ThisPos.Line++
		s.ThisPos.Column = 0
	} else {
		s.ThisPos.Column++
	}
}

func (s *Scanner) Start() {
	s.OutPos = s.ThisPos
	s.Out.Reset()
}

func (s *Scanner) Emit() string {
	return s.Out.String()
}

func (s *Scanner) Keep() {
	s.Out.WriteRune(s.This)
	s.Next()
}

func (s *Scanner) IsEnd() bool {
	return s.This == End
}

func (s *Scanner) ScanWhitespace() string {
	return s.Scan(Whitespace)
}

func (s *Scanner) ScanUDec() string {
	return s.Scan(UDec)
}
