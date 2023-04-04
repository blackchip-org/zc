package ptime

import (
	"unicode"
	"unicode/utf8"
)

// When the src stream is exhausted, ch is set to this value
const end rune = -1

type scanner struct {
	src    string
	n      int
	ch     rune
	w      int
	idx    int
	inWord bool
}

func Scan(text string) []Token {
	s := scanner{
		src: text,
		n:   len(text),
		idx: -1,
	}
	s.scan()
	var tokens []Token
	for t := s.next(); t.Type != End; t = s.next() {
		tokens = append(tokens, t)
	}
	return tokens
}

func (s *scanner) next() Token {
	if s.ch == end {
		return Token{End, "", s.idx}
	}
	s.skipWhitespace()
	switch {
	case unicode.IsDigit(s.ch):
		return s.scanNumber()
	case unicode.IsLetter(s.ch):
		return s.scanText()
	}
	return s.scanIndicator()
}

func (s *scanner) scanNumber() Token {
	start := s.idx
	for unicode.IsDigit(s.ch) {
		s.scan()
	}
	return Token{Number, s.src[start:s.idx], start + 1}
}

func (s *scanner) scanText() Token {
	start := s.idx
	for unicode.IsLetter(s.ch) {
		s.scan()
	}
	return Token{Text, s.src[start:s.idx], start + 1}
}

func (s *scanner) scanIndicator() Token {
	start := s.idx
	s.scan()
	for s.ch != end && !unicode.IsSpace(s.ch) && !unicode.IsDigit(s.ch) && !unicode.IsLetter(s.ch) {
		s.scan()
	}
	return Token{Indicator, s.src[start:s.idx], start + 1}
}

func (s *scanner) skipWhitespace() {
	if s.idx != 0 {
		s.inWord = true
	}
	for s.ch != end && unicode.IsSpace(s.ch) {
		s.scan()
		s.inWord = false
	}
}

func (s *scanner) scan() {
	if s.ch == end {
		return
	}
	s.idx++
	if s.idx >= len(s.src) {
		s.ch = end
		s.idx = s.n
		return
	}
	s.ch, s.w = utf8.DecodeRuneInString(s.src[s.idx:])
}
