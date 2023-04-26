package scanner

import (
	"bufio"
	"bytes"
	"fmt"
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
	Debug     bool
	src       *bufio.Reader
}

func New(src io.Reader) *Scanner {
	s := &Scanner{}
	s.SetReader(src)
	return s
}

func NewString(src string) *Scanner {
	s := &Scanner{}
	s.SetString(src)
	return s
}

func NewBytes(src []byte) *Scanner {
	s := &Scanner{}
	s.SetBytes(src)
	return s
}

func (s *Scanner) SetReader(src io.Reader) {
	s.src = bufio.NewReader(src)
	s.init()
}

func (s *Scanner) SetString(src string) {
	s.SetReader(strings.NewReader(src))
}

func (s *Scanner) SetBytes(src []byte) {
	s.SetReader(bytes.NewReader(src))
}

func (s *Scanner) SetName(name string) {
	s.ChPos.Name = name
}

func (s *Scanner) init() {
	s.Error = nil
	s.Ch, s.Lookahead = 0, 0
	s.Text.Reset()
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
		if s.Debug {
			fmt.Println("scan: end")
		}
		return
	}

	s.Ch = s.Lookahead
	if s.Debug {
		fmt.Printf("scan: %v(%c)\n", s.Ch, s.Ch)
	}

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

func (s *Scanner) SkipIf(c RuneClass) {
	if c(s.Ch) {
		s.Next()
	} else {
		s.Keep()
	}
}

func (s *Scanner) KeepIf(c RuneClass) {
	if c(s.Ch) {
		s.Keep()
	} else {
		s.Next()
	}
}

func (s *Scanner) End() bool {
	return s.src == nil || s.Ch == EndCh
}

func (s *Scanner) Ok() bool {
	return !s.End() && s.Error == nil
}

func (s *Scanner) setErr(err error) {
	s.Ch = EndCh
	if err != io.EOF {
		s.Error = err
	}
}
