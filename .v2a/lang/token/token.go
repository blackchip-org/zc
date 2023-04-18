package token

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/scanner"
)

type Type int

const (
	Invalid Type = iota
	Alias
	Dedent
	DoubleSlash
	Elif
	Else
	End
	For
	Func
	Id
	If
	Import
	Include
	Indent
	Loop
	Macro
	Native
	Newline
	Return
	Semicolon
	Slash
	SlashDash
	String
	StringPlain
	Try
	Use
	Value
	While
)

var tokStr = map[Type]string{
	Invalid:     "invalid",
	Alias:       "alias",
	Dedent:      "dedent",
	DoubleSlash: "//",
	Elif:        "elif",
	Else:        "else",
	End:         "end",
	For:         "for",
	Func:        "func",
	Id:          "id",
	If:          "if",
	Import:      "import",
	Include:     "include",
	Indent:      "indent",
	Loop:        "loop",
	Native:      "native",
	Macro:       "macro",
	Newline:     "newline",
	Return:      "return",
	Semicolon:   ";",
	Slash:       "/",
	SlashDash:   "/-",
	String:      "string",
	StringPlain: "plain",
	Try:         "try",
	Use:         "use",
	Value:       "value",
	While:       "while",
}

var keywords = map[string]Type{
	"alias":   Alias,
	"elif":    Elif,
	"else":    Else,
	"for":     For,
	"func":    Func,
	"if":      If,
	"import":  Import,
	"include": Include,
	"loop":    Loop,
	"macro":   Macro,
	"native":  Native,
	"return":  Return,
	"try":     Try,
	"use":     Use,
	"while":   While,
}

// If id is a keyword, returns the specific keyword token type, otherwise
// returns IdToken
func LookupKeyword(id string) Type {
	tok, ok := keywords[id]
	if ok {
		return tok
	}
	return Id
}

func (t Type) String() string {
	str, ok := tokStr[t]
	if ok {
		return str
	}
	return "unknown"
}

type Token struct {
	Type    Type
	Literal string
	Pos     scanner.Pos
}

func New(t Type, lit string, pos scanner.Pos) Token {
	return Token{t, lit, pos}
}

func (t Token) String() string {
	var quoted strings.Builder
	for _, ch := range t.Literal {
		switch ch {
		case '\n':
			quoted.WriteString("\\n")
		case '\t':
			quoted.WriteString("\\t")
		default:
			quoted.WriteRune(ch)
		}
	}

	if t.Type.String() != t.Literal {
		return fmt.Sprintf("%v(%v)", t.Type, quoted.String())
	}
	return t.Type.String()
}

func IsIdRune(ch rune) bool {
	switch {
	case unicode.IsSpace(ch):
		return false
	case unicode.Is(unicode.Sc, ch):
		return false
	case ch == ';', ch == ',':
		return false
	case ch == '[', ch == ']':
		return false
	}
	return true
}
