package zc

import (
	"strings"
)

type Token uint8

const (
	TokEnd Token = iota
	TokName
	TokValue
)

type Scanner struct {
	src []byte
	pos int
	ch  byte
	la  byte
	tok strings.Builder
}

func (s *Scanner) InitFromString(src string) {
	s.src = []byte(src)
	s.pos = 0
	s.ch = 0xff
	s.la = 0xff
	s.next()
	s.next()
}

func (s *Scanner) Next() (Token, string) {
	if s.ch == 0 {
		return TokEnd, ""
	}

	s.tok.Reset()
	s.skipSpace()

	switch {
	case s.ch == '/':
		return s.scanSlashValue()
	case s.ch == '"':
		return s.scanQuoted('"', false)
	case s.ch == '\'':
		return s.scanQuoted('\'', false)
	case s.ch == '[':
		return s.scanQuoted(']', true)
	case s.ch >= '0' && s.ch <= '9':
		return s.scanWord(TokValue)
	case (s.ch == '-' || s.ch == '+' || s.ch == '.') && (s.la >= '0' && s.la <= '9'):
		return s.scanWord(TokValue)
	default:
		return s.scanWord(TokName)
	}
}

func (s *Scanner) skipSpace() {
	for s.isSpace() {
		s.next()
	}
}

func (s *Scanner) scanSlashValue() (Token, string) {
	s.next()
	for !s.isSpace() && s.ch != 0 {
		s.tok.WriteByte(s.ch)
		s.next()
	}
	return TokValue, s.tok.String()
}

func (s *Scanner) scanQuoted(end byte, raw bool) (Token, string) {
	s.next()
	for {
		switch s.ch {
		case 0:
			return TokValue, s.tok.String()
		case end:
			s.next()
			return TokValue, s.tok.String()
		case '\\':
			if raw {
				s.tok.WriteByte(s.ch)
			} else {
				s.next()
				switch s.ch {
				case end:
					s.tok.WriteByte(s.ch)
				case 'n':
					s.tok.WriteByte('\n')
				case 't':
					s.tok.WriteByte('\t')
				case '\\':
					s.tok.WriteByte('\\')
				default:
					s.tok.WriteByte('\\')
					s.tok.WriteByte(s.ch)
				}
				s.next()
			}
		default:
			s.tok.WriteByte(s.ch)
		}
		s.next()
	}
}

func (s *Scanner) scanWord(tok Token) (Token, string) {
	s.tok.WriteByte(s.ch)
	s.next()
	for !s.isSpace() && s.ch != 0 {
		s.tok.WriteByte(s.ch)
		s.next()
	}
	return tok, s.tok.String()
}

func (s *Scanner) next() {
	if s.ch == 0 {
		return
	}
	s.ch = s.la
	if s.pos >= len(s.src) {
		s.la = 0
	} else {
		s.la = s.src[s.pos]
		s.pos++
	}
}

func (s *Scanner) isSpace() bool {
	return s.ch == ' ' || s.ch == '\t'
}
