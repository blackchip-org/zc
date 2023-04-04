package scanner

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

const EndCh = rune(-1)

type Scanner struct {
	Error     error
	Ch        rune
	Lookahead rune
	Text      strings.Builder
	TokenPos  Pos
	ChPos     Pos
	src       *bufio.Reader
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
	s.ChPos.Name = name
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
	s.Ch, s.Lookahead = 0, 0
	s.Next()
	s.Next()
	s.ChPos.Line = 1
	s.ChPos.Column = 1
}

func (s *Scanner) Scan(fn Func) string {
	s.Start()
	fn(s)
	return s.Text.String()
}

func (s *Scanner) ScanUntil(c RuneClass) string {
	return s.Scan(UntilFunc(c))
}

func (s *Scanner) ScanWhile(c RuneClass) string {
	return s.Scan(WhileFunc(c))
}

func (s *Scanner) Next() {
	if s.Ch == EndCh || s.src == nil {
		return
	}

	s.Ch = s.Lookahead

	r, _, err := s.src.ReadRune()
	if err != nil {
		s.Lookahead = EndCh
		if err != io.EOF {
			s.Error = err
		}
	} else {
		s.Lookahead = r
	}

	if s.Ch == '\n' {
		s.ChPos.Line++
		s.ChPos.Column = 0
	} else {
		s.ChPos.Column++
	}
}

func (s *Scanner) Start() {
	s.Error = nil
	s.TokenPos = s.ChPos
	s.Text.Reset()
}

func (s *Scanner) Token() string {
	return s.Text.String()
}

func (s *Scanner) Keep() {
	s.Text.WriteRune(s.Ch)
	s.Next()
}

func (s *Scanner) End() bool {
	return s.src == nil || s.Ch == EndCh
}

func (s *Scanner) Ok() bool {
	return !s.End() && s.Error == nil
}
