package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/blackchip-org/zc/lang/scanner"
	"github.com/blackchip-org/zc/lang/token"
)

type Lexer struct {
	src        []byte
	ch         rune          // current rune being scanned
	w          int           // width, in bytes, of the current rune
	idx        int           // index into src of where ch is located
	pos        scanner.Pos   // file, line, and column where ch is located
	start      scanner.Pos   // position of the scanner when Next() was called
	indents    int           // count of current indentation level, one for each tab
	nextTokens []token.Token // pending dedent tokens to emit
	inBlock    bool
}

// When the src stream is exhausted, ch is set to this value
const end rune = -1

func New(file string, src []byte) *Lexer {
	s := &Lexer{
		src: src,
		pos: scanner.Pos{
			Name: file,
			Line: 1,
		},
	}
	s.scan()
	return s
}

func (l *Lexer) Next() token.Token {
	// If there are dedent tokens buffered up, emit those first
	// before continuing the scan.
	if len(l.nextTokens) > 0 {
		var tok token.Token
		tok, l.nextTokens = l.nextTokens[0], l.nextTokens[1:]
		return tok
	}

	if l.ch == '-' && l.lookahead() == '-' {
		l.skipComment()
		// If we have consumed a comment and we are in a block, we need to
		// consume any indentation on the next line.
		if l.inBlock {
			l.skipSpace()
		}
	}

	// When at the start of the line, check to see what the current
	// indentation level is and emit indent and dedent tokens as needed
	l.start = l.pos
	if l.pos.Column == 1 && !l.inBlock {
		if tok, yes := l.scanIndent(); yes {
			return tok
		}
	}

	l.skipSpace()

	l.start = l.pos
	next := l.lookahead()

	if l.inBlock {
		l.skipSpaceAndNewlines()
	}
	if l.ch == '[' {
		l.inBlock = true
		l.scan()
		l.skipSpaceAndNewlines()
	}
	if l.ch == ']' {
		l.inBlock = false
		l.scan()
		l.skipSpace()
	}

	switch {
	case l.ch == end:
		return token.New(token.End, "", l.pos)
	case l.ch == '\n':
		return l.scanOp(token.Newline)
	case l.ch == ';':
		return l.scanOp(token.Semicolon)
	case l.ch == '/':
		return l.scanSlash()
	case l.ch == '"':
		return l.scanString('"')
	case l.ch == '\'':
		return l.scanString('\'')
	case isValue(l.ch, next):
		return l.scanValue()
	}
	return l.scanId()
}

func (l *Lexer) scanId() token.Token {
	startL := l.idx
	for l.ch != end && token.IsIdRune(l.ch) {
		l.scan()
	}
	lit := string(l.src[startL:l.idx])
	if len(lit) == 0 {
		panic("id literal is zero in length")
	}
	// If the identifier is a keyword, use the keyword specific token type,
	// otherwise use IdentToken
	return token.New(token.LookupKeyword(lit), lit, l.start)
}

// Returns true if there is an indent/dedent token to emit
func (l *Lexer) scanIndent() (token.Token, bool) {
	var indent strings.Builder

	spaces := 0
	for l.ch != end && (l.ch == '\t' || l.ch == ' ') {
		if l.ch == ' ' {
			spaces++
			if spaces == 4 {
				spaces = 0
				indent.WriteRune('\t')
			}
		} else if l.ch == '\t' {
			spaces = 0
			indent.WriteRune('\t')
		}
		l.scan()
	}

	lit := indent.String()

	// If the entire line is blank, ignore it
	if (l.ch == end || l.ch == '\n') && strings.TrimSpace(lit) == "" {
		return token.Token{}, false
	}

	// Count the number of tabs to determine the indentation level. If this
	// is the same indentation level of the previous line, do not emit
	// an indent/dedent token
	n := len(lit)
	diff := n - l.indents
	var tok token.Token
	if diff == 0 {
		return tok, false
	}

	if diff > 0 {
		tok = token.New(token.Indent, lit, l.start)
	} else {
		tok = token.New(token.Dedent, lit, l.start)
	}
	if diff < 0 {
		diff = -diff
	}

	// If multiple dedent tokens need to be emitted, emit one now and
	// put the remaining ones in nextTokens
	for i := 1; i < diff; i++ {
		l.nextTokens = append([]token.Token{tok}, l.nextTokens...)
	}
	l.indents = n

	return tok, true
}

func (l *Lexer) scanOp(name token.Type) token.Token {
	lit := string(l.ch)
	l.scan()
	return token.New(name, lit, l.start)
}

func (l *Lexer) scanString(term rune) token.Token {
	l.scan()
	var lit strings.Builder
	escaping := false
	for l.ch != end && l.ch != '\n' {
		if !escaping && l.ch == term {
			break
		}
		if l.ch == '\\' {
			escaping = true
		} else {
			if escaping {
				switch l.ch {
				case 'n':
					lit.WriteRune('\n')
				default:
					lit.WriteRune(l.ch)
				}
			} else {
				lit.WriteRune(l.ch)
			}
			escaping = false
		}
		l.scan()
	}
	l.scan()
	t := token.String
	if term == '"' {
		t = token.StringPlain
	}
	return token.New(t, lit.String(), l.start)
}

func (l *Lexer) scanValue() token.Token {
	startL := l.idx
	for l.ch != end && !unicode.IsSpace(l.ch) {
		l.scan()
	}
	lit := string(l.src[startL:l.idx])
	t := token.Value
	return token.New(t, lit, l.start)
}

func (l *Lexer) scanSlash() token.Token {
	l.scan()
	if l.ch == '/' {
		l.scan()
		if l.ch == end || unicode.IsSpace(l.ch) {
			return token.New(token.Id, "//", l.start)
		}
		return token.New(token.DoubleSlash, "//", l.start)
	}
	if l.ch == '-' {
		l.scan()
		return token.New(token.SlashDash, "/-", l.start)
	}
	if l.ch == end || unicode.IsSpace(l.ch) {
		return token.New(token.Id, "/", l.start)
	}
	return token.New(token.Slash, "/", l.start)
}

func (l *Lexer) scan() {
	if l.ch == end {
		return
	}

	if l.ch == '\n' {
		l.pos.Line++
		l.pos.Column = 1
	} else if l.ch == '\t' {
		l.pos.Column += 4
	} else {
		l.pos.Column++
	}
	l.idx += l.w
	if l.idx >= len(l.src) {
		l.ch = end
		l.idx = len(l.src)
		return
	}
	l.ch, l.w = utf8.DecodeRune(l.src[l.idx:])
}

func (l *Lexer) skipSpace() {
	// Newlines have their own tokens so do not include here
	for l.ch != end && l.ch != '\n' && unicode.IsSpace(l.ch) {
		l.scan()
	}
}

func (l *Lexer) skipSpaceAndNewlines() {
	for l.ch != end && unicode.IsSpace(l.ch) {
		l.scan()
	}
}

func (l *Lexer) skipComment() {
	l.scan()
	l.scan()
	if l.ch == '-' {
		// Block comment
		l.scan()
		for l.ch != end && !(l.ch == '-' && l.lookahead() == '-') {
			l.scan()
		}
		l.scan()
		l.scan()
		l.scan()
	} else {
		// Line comment
		l.scan()
		for l.ch != end && l.ch != '\n' {
			l.scan()
		}
		l.scan()
	}
}

func (l *Lexer) lookahead() rune {
	next := l.idx + l.w
	if next >= len(l.src) {
		return end
	}
	ch, _ := utf8.DecodeRune(l.src[next:])
	return ch
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
