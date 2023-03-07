package scanner

import (
	"math/big"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/blackchip-org/zc/lang/token"
)

type Scanner struct {
	src        []byte
	ch         rune          // current rune being scanned
	w          int           // width, in bytes, of the current rune
	idx        int           // index into src of where ch is located
	pos        token.Pos     // file, line, and column where ch is located
	start      token.Pos     // position of the scanner when Next() was called
	indents    int           // count of current indentation level, one for each tab
	nextTokens []token.Token // pending dedent tokens to emit
	inBlock    bool
	scanType   ScanType
}

type ScanType int

const (
	Compiler ScanType = iota
	Runtime
)

// When the src stream is exhausted, ch is set to this value
const end rune = -1

func New(file string, src []byte, scanType ScanType) *Scanner {
	s := &Scanner{
		src: src,
		pos: token.Pos{
			File: file,
			Line: 1,
		},
		scanType: scanType,
	}
	s.scan()
	return s
}

func (s *Scanner) Next() token.Token {
	// If there are dedent tokens buffered up, emit those first
	// before continuing the scan.
	if len(s.nextTokens) > 0 {
		var tok token.Token
		tok, s.nextTokens = s.nextTokens[0], s.nextTokens[1:]
		return tok
	}

	if s.ch == '-' && s.lookahead() == '-' {
		s.skipComment()
		// If we have consumed a comment and we are in a block, we need to
		// consume any indentation on the next line.
		if s.inBlock {
			s.skipSpace()
		}
	}

	// When at the start of the line, check to see what the current
	// indentation level is and emit indent and dedent tokens as needed
	s.start = s.pos
	if s.pos.Column == 1 && !s.inBlock {
		if tok, yes := s.scanIndent(); yes {
			return tok
		}
	}

	s.skipSpace()

	s.start = s.pos
	next := s.lookahead()

	if s.inBlock {
		s.skipSpaceAndNewlines()
	}
	if s.ch == '[' {
		s.inBlock = true
		s.scan()
		s.skipSpaceAndNewlines()
	}
	if s.ch == ']' {
		s.inBlock = false
		s.scan()
		s.skipSpace()
	}

	switch {
	case s.ch == end:
		return token.New(token.End, "", s.pos)
	case s.ch == '\n':
		return s.scanOp(token.Newline)
	case s.ch == ';':
		return s.scanOp(token.Semicolon)
	case s.ch == '/':
		return s.scanSlash()
	case s.ch == '"':
		return s.scanString('"')
	case s.ch == '\'':
		return s.scanString('\'')
	case isValue(s.ch, next):
		return s.scanValue()
	}
	return s.scanId()
}

func (s *Scanner) scanId() token.Token {
	startL := s.idx
	for s.ch != end && token.IsIdRune(s.ch) {
		s.scan()
	}
	lit := string(s.src[startL:s.idx])
	if len(lit) == 0 {
		panic("id literal is zero in length")
	}
	// If the identifier is a keyword, use the keyword specific token type,
	// otherwise use IdentToken
	return token.New(token.LookupKeyword(lit), lit, s.start)
}

// Returns true if there is an indent/dedent token to emit
func (s *Scanner) scanIndent() (token.Token, bool) {
	var indent strings.Builder

	spaces := 0
	for s.ch != end && (s.ch == '\t' || s.ch == ' ') {
		if s.ch == ' ' {
			spaces++
			if spaces == 4 {
				spaces = 0
				indent.WriteRune('\t')
			}
		} else if s.ch == '\t' {
			spaces = 0
			indent.WriteRune('\t')
		}
		s.scan()
	}

	lit := indent.String()

	// If the entire line is blank, ignore it
	if (s.ch == end || s.ch == '\n') && strings.TrimSpace(lit) == "" {
		return token.Token{}, false
	}

	// Count the number of tabs to determine the indentation level. If this
	// is the same indentation level of the previous line, do not emit
	// an indent/dedent token
	n := len(lit)
	diff := n - s.indents
	var tok token.Token
	if diff == 0 {
		return tok, false
	}

	if diff > 0 {
		tok = token.New(token.Indent, lit, s.start)
	} else {
		tok = token.New(token.Dedent, lit, s.start)
	}
	if diff < 0 {
		diff = -diff
	}

	// If multiple dedent tokens need to be emitted, emit one now and
	// put the remaining ones in nextTokens
	for i := 1; i < diff; i++ {
		s.nextTokens = append([]token.Token{tok}, s.nextTokens...)
	}
	s.indents = n

	return tok, true
}

func (s *Scanner) scanOp(name token.Type) token.Token {
	lit := string(s.ch)
	s.scan()
	return token.New(name, lit, s.start)
}

func (s *Scanner) scanString(term rune) token.Token {
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
			if escaping {
				switch s.ch {
				case 'n':
					lit.WriteRune('\n')
				default:
					lit.WriteRune(s.ch)
				}
			} else {
				lit.WriteRune(s.ch)
			}
			escaping = false
		}
		s.scan()
	}
	s.scan()
	t := token.String
	if term == '"' {
		t = token.StringPlain
	}
	return token.New(t, lit.String(), s.start)
}

func (s *Scanner) scanValue() token.Token {
	startL := s.idx
	for s.ch != end && !unicode.IsSpace(s.ch) {
		s.scan()
	}
	lit := string(s.src[startL:s.idx])
	t := token.Value
	if s.scanType == Compiler && isNumber(lit) {
		t = token.Number
	}
	return token.New(t, lit, s.start)
}

func (s *Scanner) scanSlash() token.Token {
	s.scan()
	if s.ch == '/' {
		s.scan()
		if s.ch == end || unicode.IsSpace(s.ch) {
			return token.New(token.Id, "//", s.start)
		}
		return token.New(token.DoubleSlash, "//", s.start)
	}
	if s.ch == '-' {
		s.scan()
		return token.New(token.SlashDash, "/-", s.start)
	}
	if s.ch == end || unicode.IsSpace(s.ch) {
		return token.New(token.Id, "/", s.start)
	}
	return token.New(token.Slash, "/", s.start)
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

func (s *Scanner) skipSpaceAndNewlines() {
	for s.ch != end && unicode.IsSpace(s.ch) {
		s.scan()
	}
}

func (s *Scanner) skipComment() {
	s.scan()
	s.scan()
	if s.ch == '-' {
		// Block comment
		s.scan()
		for s.ch != end && !(s.ch == '-' && s.lookahead() == '-') {
			s.scan()
		}
		s.scan()
		s.scan()
		s.scan()
	} else {
		// Line comment
		s.scan()
		for s.ch != end && s.ch != '\n' {
			s.scan()
		}
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

func isNumber(lit string) bool {
	var bi big.Int
	if _, ok := bi.SetString(lit, 0); ok {
		return true
	}
	if _, err := strconv.ParseFloat(lit, 64); err == nil {
		return true
	}
	return false
}

func isValue(ch rune, next rune) bool {
	switch {
	case unicode.IsDigit(ch), unicode.Is(unicode.Sc, ch):
		return true
	case (ch == '-' || ch == '+' || ch == '.') && unicode.IsDigit(next):
		return true
	}
	return false
}

func Quote(v string) string {
	required := false
	runes := []rune(v)
	for i, ch := range runes {
		if i == 0 {
			var next rune
			if len(runes) > 1 {
				next = runes[i+1]
			}
			if !isValue(ch, next) {
				required = true
				break
			}
		}
		if unicode.IsSpace(ch) {
			required = true
			break
		}
	}
	if !required {
		return v
	}
	var ret strings.Builder
	ret.WriteRune('\'')
	for _, ch := range runes {
		if ch == '\'' {
			ret.WriteString("\\'")
		} else {
			ret.WriteRune(ch)
		}
	}
	ret.WriteRune('\'')
	return ret.String()
}
