package scanner

import (
	"errors"
	"strings"
	"unicode"
)

type Func func(*Scanner)

func WhileFunc(while RuneClass) Func {
	return func(s *Scanner) {
		for while(s.Ch) && s.Ch != EndCh {
			s.Keep()
		}
	}
}

func UntilFunc(until RuneClass) Func {
	return func(s *Scanner) {
		for !until(s.Ch) && s.Ch != EndCh {
			s.Keep()
		}
	}
}

func NumberFunc(def NumberDef) Func {
	def.DecSep = OptionalClass(def.DecSep)
	def.Sign = OptionalClass(def.Sign)
	def.Exponent = OptionalClass(def.Exponent)

	return func(s *Scanner) {
		if def.Sign(s.Ch) {
			s.Keep()
		}
		seenDecSep := false
		exponent := false
		for {
			if def.DecSep(s.Ch) {
				if seenDecSep {
					break
				}
				seenDecSep = true
				s.Keep()
			} else if def.Exponent(s.Ch) {
				exponent = true
				s.Keep()
				break
			} else if def.Digit(s.Ch) {
				s.Keep()
			} else {
				break
			}
		}
		if !exponent {
			return
		}
		if def.Sign(s.Ch) {
			s.Keep()
		}
		for def.Digit(s.Ch) {
			s.Keep()
		}
	}
}

var ErrNotTerminated = errors.New("not terminated")

func QuotedFunc(def QuotedDef) Func {
	def.AltEnd = OptionalClass(def.AltEnd)
	return func(s *Scanner) {
		endQuote := s.Ch
		s.Next()
		for s.Ch != endQuote && !def.AltEnd(s.Ch) && !s.End() {
			if !def.Escape(s.Ch) {
				s.Keep()
				continue
			}
			s.Next()
			if def.EscapeMap == nil {
				s.Keep()
				continue
			}
			mapped, ok := def.EscapeMap[s.Ch]
			if ok {
				s.Text.WriteRune(mapped)
				s.Next()
			} else {
				s.Keep()
			}
		}
		if s.Ch != endQuote && !def.AltEnd(s.Ch) {
			s.Error = ErrNotTerminated
		}
		s.Next()
	}
}

func UntilRepeatsFunc(is RuneClass, n int) Func {
	return func(s *Scanner) {
		count := 0
		for s.Ok() && count < n {
			if is(s.Ch) {
				count++
			} else {
				count = 0
			}
			s.Keep()
		}
		if count == n {
			text := s.Text.String()
			s.Text.Reset()
			s.Text.WriteString(text[0 : len(text)-count])
		}
	}
}

func Word(s *Scanner) {
	s.ScanWhile(unicode.IsSpace)
	s.ScanUntil(unicode.IsSpace)
}

func Line(s *Scanner) {
	s.ScanUntil(IsNewline)
	s.Next()
}

func LineTrimSpace(s *Scanner) {
	Line(s)
	l := s.Token()
	s.Text.Reset()
	s.Text.WriteString(strings.TrimSpace(l))
}
