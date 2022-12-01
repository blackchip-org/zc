package lang

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

const end rune = -1

type Scanner struct {
	src        []byte
	ch         rune
	w          int
	idx        int
	pos        Position
	start      Position
	lastIndent string
	// Errors       Errors
	// ErrorHandler ErrorHandler
}

func NewScanner(file string, src []byte) *Scanner {
	s := &Scanner{
		src: src,
		pos: Position{
			File: file,
			Line: 1,
		},
	}
	s.scan()
	return s
}

func (s *Scanner) Next() Token {
	if s.pos.Column == 1 {
		if tok, yes := s.scanIndent(); yes {
			return tok
		}
	} else {
		s.skipSpace()
	}
	s.start = s.pos
	next := s.lookahead()
	switch {
	case s.ch == end:
		return Token{EndToken, "", s.pos}
	case s.ch == '\n':
		return s.op(NewlineToken)
	case s.ch == '/':
		return s.scanSlash()
	case s.ch == '"':
		return s.scanQuotedValue('"')
	case s.ch == '\'':
		return s.scanQuotedValue('\'')
	case unicode.IsDigit(s.ch):
		return s.scanValue()
	case (s.ch == '-' || s.ch == '+') && unicode.IsDigit(next):
		return s.scanValue()

	}
	return s.scanIdent()
}

func (s *Scanner) scanIndent() (Token, bool) {
	var lit strings.Builder
	for s.ch != end && (s.ch == ' ' || s.ch == '\t') {
		if s.ch == '\t' {
			lit.WriteRune(' ')
			for lit.Len()%8 != 0 {
				lit.WriteRune(' ')
			}
		} else {
			lit.WriteRune(s.ch)
		}
		s.scan()
	}
	thisIndent := lit.String()
	if len(thisIndent) > len(s.lastIndent) {
		s.lastIndent = thisIndent
		return Token{IndentToken, thisIndent, s.start}, true
	} else if len(thisIndent) < len(s.lastIndent) {
		s.lastIndent = thisIndent
		return Token{DedentToken, thisIndent, s.start}, true
	}
	return Token{}, false
}

func (s *Scanner) skipSpace() {
	for s.ch != end && s.ch != '\n' && unicode.IsSpace(s.ch) {
		s.scan()
	}
}

func (s *Scanner) op(name TokenType) Token {
	lit := string(s.ch)
	s.scan()
	return Token{name, lit, s.start}
}

func (s *Scanner) scanQuotedValue(term rune) Token {
	s.scan()
	var lit strings.Builder
	escaping := false
	for s.ch != end && s.ch != '\n' {
		if !escaping && s.ch == term {
			break
		}
		if s.ch == '\\' {
			escaping = true
		} else {
			lit.WriteRune(s.ch)
			escaping = false
		}
		s.scan()
	}
	s.scan()
	return Token{ValueToken, lit.String(), s.start}
}

func (s *Scanner) scanIdent() Token {
	startL := s.idx
	for s.ch != end && !unicode.IsSpace(s.ch) {
		s.scan()
	}
	lit := string(s.src[startL:s.idx])
	return Token{LookupKeyword(lit), lit, s.start}
}

func (s *Scanner) scanValue() Token {
	startL := s.idx
	for s.ch != end && !unicode.IsSpace(s.ch) {
		s.scan()
	}
	lit := string(s.src[startL:s.idx])
	return Token{ValueToken, lit, s.start}
}

func (s *Scanner) scanSlash() Token {
	s.scan()
	if s.ch == '/' {
		s.scan()
		if s.ch == end || unicode.IsSpace(s.ch) {
			return Token{IdToken, "//", s.start}
		}
		return Token{AllRefToken, "//", s.start}
	}
	if s.ch == end || unicode.IsSpace(s.ch) {
		return Token{IdToken, "/", s.start}
	}
	return Token{TopRefToken, "/", s.start}
}

func (s *Scanner) scan() {
	if s.ch == end {
		return
	}

	if s.ch == '\n' {
		s.pos.Line++
		s.pos.Column = 1
	} else {
		s.pos.Column++
	}
	s.idx += s.w
	if s.idx >= len(s.src) {
		s.ch = end
		s.idx = len(s.src)
		return
	}
	s.ch, s.w = utf8.DecodeRune(s.src[s.idx:])
}

func (s *Scanner) lookahead() rune {
	next := s.idx + s.w
	if next >= len(s.src) {
		return end
	}
	ch, _ := utf8.DecodeRune(s.src[next:])
	return ch
}

// func (s *Scanner) err(format string, a ...any) {
// 	msg := fmt.Sprintf(format, a...)
// 	e := fmt.Errorf("%v: %v", l.pos, msg)
// 	s.Errors = append(s.Errors, e)
// 	if s.ErrorHandler != nil {
// 		s.ErrorHandler(e)
// 	}
// }
