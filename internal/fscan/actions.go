package fscan

func Discard(s *Scanner) {
	s.Next()
}

func Keep(s *Scanner) {
	s.Out.WriteRune(s.This)
	s.Next()
}

func None(s *Scanner) {}

func ReplaceRune(r rune) Func {
	return func(s *Scanner) {
		s.Out.WriteRune(r)
		s.Next()
	}
}

var (
	ToPeriod = ReplaceRune('.')
)
