package token

import (
	"fmt"
	"strings"
	"unicode"
)

type Pos struct {
	File   string
	Line   int
	Column int
}

func NewPos(file string, line int, column int) Pos {
	return Pos{file, line, column}
}

func (p Pos) IsValid() bool {
	return p.Line > 0
}

func (p Pos) String() string {
	if p.File != "" {
		return fmt.Sprintf("%v:%v:%v", p.File, p.Line, p.Column)
	}
	return fmt.Sprintf("%v:%v", p.Line, p.Column)
}

type Type int

const (
	Invalid Type = iota
	Dedent
	DoubleSlash
	Elif
	Else
	End
	For
	Func
	Id
	Include
	Import
	Indent
	If
	Loop
	Macro
	Newline
	Return
	Semicolon
	Slash
	String
	Try
	Use
	Value
	While
)

var tokStr = map[Type]string{
	Invalid:     "invalid",
	Dedent:      "dedent",
	DoubleSlash: "//",
	Elif:        "elif",
	Else:        "else",
	End:         "end",
	For:         "for",
	Func:        "func",
	Id:          "id",
	Include:     "include",
	Import:      "import",
	Indent:      "indent",
	If:          "if",
	Loop:        "loop",
	Macro:       "macro",
	Newline:     "newline",
	Return:      "return",
	Semicolon:   ";",
	Slash:       "/",
	String:      "string",
	Try:         "try",
	Use:         "use",
	Value:       "value",
	While:       "while",
}

var keywords = map[string]Type{
	"elif":    Elif,
	"else":    Else,
	"for":     For,
	"func":    Func,
	"include": Include,
	"import":  Import,
	"if":      If,
	"loop":    Loop,
	"macro":   Macro,
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
	Pos     Pos
}

func New(t Type, lit string, pos Pos) Token {
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
	}
	return true
}
