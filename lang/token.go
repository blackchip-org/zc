package lang

import (
	"fmt"
	"strings"
	"unicode"
)

type Position struct {
	File   string
	Line   int
	Column int
}

func (p Position) String() string {
	if p.File != "" {
		return fmt.Sprintf("%v:%v:%v", p.File, p.Line, p.Column)
	}
	return fmt.Sprintf("%v:%v", p.Line, p.Column)
}

type TokenType int

const (
	InvalidToken TokenType = iota
	DedentToken
	DoubleSlashToken
	ElifToken
	ElseToken
	EndToken
	FuncToken
	IdToken
	IncludeToken
	IndentToken
	IfToken
	LoopToken
	NewlineToken
	ReturnToken
	SemicolonToken
	SlashToken
	ValueToken
	WhileToken
)

var tokStr = map[TokenType]string{
	InvalidToken:     "invalid",
	DedentToken:      "dedent",
	DoubleSlashToken: "//",
	ElifToken:        "elif",
	ElseToken:        "else",
	EndToken:         "end",
	FuncToken:        "func",
	IdToken:          "id",
	IncludeToken:     "include",
	IndentToken:      "indent",
	IfToken:          "if",
	LoopToken:        "loop",
	NewlineToken:     "newline",
	ReturnToken:      "return",
	SlashToken:       "/",
	SemicolonToken:   ";",
	ValueToken:       "value",
	WhileToken:       "while",
}

var keywords = map[string]TokenType{
	"elif":    ElifToken,
	"else":    ElseToken,
	"func":    FuncToken,
	"include": IncludeToken,
	"if":      IfToken,
	"loop":    LoopToken,
	"while":   WhileToken,
}

// If id is a keyword, returns the specific keyword token type, otherwise
// returns IdToken
func LookupKeyword(id string) TokenType {
	tok, ok := keywords[id]
	if ok {
		return tok
	}
	return IdToken
}

func (t TokenType) String() string {
	str, ok := tokStr[t]
	if ok {
		return str
	}
	return "unknown"
}

type Token struct {
	Type    TokenType
	Literal string
	At      Position
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
	if unicode.IsSpace(ch) {
		return false
	}
	switch ch {
	case end, ';':
		return false
	}
	return true
}
