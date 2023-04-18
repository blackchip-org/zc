package lexer

import (
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/lang/token"
	"github.com/blackchip-org/zc/scanner"
)

var isSpaceOrTab = scanner.Rune2(' ', '\t')

type Lexer struct {
	s          *scanner.Scanner
	indents    int           // count of current indentation level, one for each tab
	nextTokens []token.Token // pending dedent tokens to emit
	inBlock    bool
}

func New(file string, src []byte) *Lexer {
	s := scanner.NewBytes(src)
	s.SetName(file)
	return &Lexer{s: s}
}

func (l *Lexer) Next() token.Token {
	// If there are dedent tokens buffered up, emit those first
	// before continuing the scan.
	if len(l.nextTokens) > 0 {
		var tok token.Token
		tok, l.nextTokens = l.nextTokens[0], l.nextTokens[1:]
		return tok
	}

	if l.s.Ch == '-' && l.s.Lookahead == '-' {
		l.skipComment()
		// If we have consumed a comment and we are in a block, we need to
		// consume any indentation on the next line.
		if l.inBlock {
			l.s.ScanWhile(isSpaceOrTab)
		}
	}

	l.s.Start()
	if l.s.ChPos.Column == 1 && !l.inBlock {
		if tok, yes := l.scanIndent(); yes {
			return tok
		}
	}
	l.s.ScanWhile(isSpaceOrTab)

	l.s.Start()
	if l.inBlock {
		l.s.ScanWhile(unicode.IsSpace)
	}
	if l.s.Ch == '[' {
		l.inBlock = true
		l.s.Next()
		l.s.ScanWhile(unicode.IsSpace)
	}
	if l.s.Ch == ']' {
		l.inBlock = false
		l.s.Next()
		l.s.ScanWhile(isSpaceOrTab)
	}

	switch {
	case l.s.End():
		return token.New(token.End, "", l.s.ChPos)
	case l.s.Ch == '\n':
		return l.scanOp(token.Newline)
	case l.s.Ch == ';':
		return l.scanOp(token.Semicolon)
	case l.s.Ch == '/':
		return l.scanSlash()
	case l.s.Ch == '"', l.s.Ch == '\'':
		return l.scanString()
	case isValue(l.s.Ch, l.s.Lookahead):
		return l.scanValue()
	}
	return l.scanId()
}

func (l *Lexer) scanIndent() (token.Token, bool) {
	spaces := 0
	for l.s.Ok() && isSpaceOrTab(l.s.Ch) {
		if l.s.Ch == ' ' {
			spaces++
			if spaces == 4 {
				spaces = 0
				l.s.Text.WriteRune('\t')
			}
		} else if l.s.Ch == '\t' {
			spaces = 0
			l.s.Text.WriteRune('\t')
		}
		l.s.Next()
	}

	text := l.s.Token()

	// If the entire line is blank, ignore it
	if (l.s.End() || l.s.Ch == '\n') && strings.TrimSpace(text) == "" {
		return token.Token{}, false
	}

	// Count the number of tabs to determine the indentation level. If this
	// is the same indentation level of the previous line, do not emit
	// an indent/dedent token
	n := len(text)
	diff := n - l.indents
	var tok token.Token
	if diff == 0 {
		return tok, false
	}

	if diff > 0 {
		tok = token.New(token.Indent, text, l.s.TokenPos)
	} else {
		tok = token.New(token.Dedent, text, l.s.TokenPos)
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

func (l *Lexer) scanOp(t token.Type) token.Token {
	l.s.Keep()
	return token.New(t, l.s.Token(), l.s.TokenPos)
}

func (l *Lexer) scanId() token.Token {
	text := l.s.Scan(scanner.WhileFunc(token.IsIdRune))
	if len(text) == 0 {
		panic("id literal is zero in length")
	}
	// If the identifier is a keyword, use the keyword specific token type,
	// otherwise use IdentToken
	return token.New(token.LookupKeyword(text), text, l.s.TokenPos)
}

func (l *Lexer) scanValue() token.Token {
	text := l.s.Scan(scanner.UntilFunc(unicode.IsSpace))
	return token.New(token.Value, text, l.s.TokenPos)
}

var quoteFunc = scanner.QuotedFunc(scanner.QuotedDef{
	Escape: scanner.Rune('\\'),
	AltEnd: scanner.Rune2('\n', scanner.EndCh),
	EscapeMap: map[rune]rune{
		'n': '\n',
	},
})

func (l *Lexer) scanString() token.Token {
	t := token.String
	if l.s.Ch == '"' {
		t = token.StringPlain
	}
	text := l.s.Scan(quoteFunc)
	return token.New(t, text, l.s.TokenPos)
}

func (l *Lexer) scanSlash() token.Token {
	l.s.Next()
	if l.s.Ch == '/' {
		l.s.Next()
		if l.s.End() || unicode.IsSpace(l.s.Ch) {
			return token.New(token.Id, "//", l.s.TokenPos)
		}
		return token.New(token.DoubleSlash, "//", l.s.TokenPos)
	}
	if l.s.Ch == '-' {
		l.s.Next()
		return token.New(token.SlashDash, "/-", l.s.TokenPos)
	}
	if l.s.End() || unicode.IsSpace(l.s.Ch) {
		return token.New(token.Id, "/", l.s.TokenPos)
	}
	return token.New(token.Slash, "/", l.s.TokenPos)
}

func (l *Lexer) skipComment() {
	l.s.Next()
	l.s.Next()
	if l.s.Ch == '-' {
		// Block comment
		l.s.Next()
		l.s.Scan(scanner.UntilRepeatsFunc(scanner.Rune('-'), 3))
	} else {
		// Line comment
		l.s.ScanUntil(scanner.Rune('\n'))
	}
}
