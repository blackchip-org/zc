package lang

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type Scanner struct {
	src        []byte
	ch         rune     // current rune being scanned
	w          int      // width, in bytes, of the current rune
	idx        int      // index into src of where ch is located
	pos        Position // file, line, and column where ch is located
	start      Position // position of the scanner when Next() was called
	indents    int      // count of current indentation level, one for each tab
	nextTokens []Token  // pending dedent tokens to emit
}

// When the src stream is exhausted, ch is set to this value
const end rune = -1

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
	// If there are dedent tokens buffered up, emit those first
	// before continuing the scan.
	if len(s.nextTokens) > 0 {
		var tok Token
		tok, s.nextTokens = s.nextTokens[0], s.nextTokens[1:]
		return tok
	}

	// When at the start of the line, check to see what the current
	// indentation level is and emit indent and dedent tokens as needed
	s.start = s.pos
	if s.pos.Column == 1 {
		if tok, yes := s.scanIndent(); yes {
			return tok
		}
	}

	s.skipSpace()

	s.start = s.pos
	next := s.lookahead()

	var tok Token
	switch {
	case s.ch == end:
		tok = Token{EndToken, "", s.pos}
	case s.ch == '\n':
		tok = s.scanOp(NewlineToken)
	case s.ch == '/':
		tok = s.scanSlash()
	case s.ch == '"':
		tok = s.scanQuotedValue('"')
	case s.ch == '\'':
		tok = s.scanQuotedValue('\'')
	case unicode.IsDigit(s.ch):
		tok = s.scanValue()
	case (s.ch == '-' || s.ch == '+') && unicode.IsDigit(next):
		tok = s.scanValue()
	default:
		tok = s.scanId()
	}
	return tok
}

func (s *Scanner) scanId() Token {
	startL := s.idx
	for s.ch != end && !unicode.IsSpace(s.ch) {
		s.scan()
	}
	lit := string(s.src[startL:s.idx])
	// If the identifier is a keyword, use the keyword specific token type,
	// otherwise use IdentToken
	return Token{LookupKeyword(lit), lit, s.start}
}

// Returns true if there is an indent/dedent token to emit
func (s *Scanner) scanIndent() (Token, bool) {
	startL := s.idx

	for s.ch != end && s.ch == '\t' {
		s.scan()
	}
	lit := string(s.src[startL:s.idx])

	// Count the number of tabs to determine the indentation level. If this
	// is the same indentation level of the previous line, do not emit
	// an indent/dedent token
	n := len(lit)
	diff := n - s.indents
	var token Token
	if diff == 0 {
		return token, false
	}

	if diff > 0 {
		token = Token{IndentToken, lit, s.start}
	} else {
		token = Token{DedentToken, lit, s.start}
	}
	if diff < 0 {
		diff = -diff
	}

	// If multiple dedent tokens need to be emitted, emit one now and
	// put the remaining ones in nextTokens
	for i := 1; i < diff; i++ {
		s.nextTokens = append([]Token{token}, s.nextTokens...)
	}
	s.indents = n

	return token, true
}

func (s *Scanner) scanOp(name TokenType) Token {
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
		return Token{DoubleSlashToken, "//", s.start}
	}
	if s.ch == end || unicode.IsSpace(s.ch) {
		return Token{IdToken, "/", s.start}
	}
	return Token{SlashToken, "/", s.start}
}

func (s *Scanner) scan() {
	if s.ch == end {
		return
	}

	if s.ch == '\n' {
		s.pos.Line++
		s.pos.Column = 1
	} else if s.ch == '\t' {
		s.pos.Column += 4
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

func (s *Scanner) skipSpace() {
	// Newlines have their own tokens so do not include here
	for s.ch != end && s.ch != '\n' && unicode.IsSpace(s.ch) {
		s.scan()
	}
}

func (s *Scanner) lookahead() rune {
	next := s.idx + s.w
	if next >= len(s.src) {
		return end
	}
	ch, _ := utf8.DecodeRune(s.src[next:])
	return ch
}
