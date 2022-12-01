package lang

import "fmt"

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
	AllRefToken
	DedentToken
	ElifToken
	ElseToken
	EndToken
	FnToken
	IdToken
	IndentToken
	IfToken
	LoopToken
	NewlineToken
	ReturnToken
	TopRefToken
	ValueToken
	WhileToken
)

var tokStr = map[TokenType]string{
	InvalidToken: "invalid",
	AllRefToken:  "ref",
	DedentToken:  "dedent",
	ElifToken:    "elif",
	ElseToken:    "else",
	EndToken:     "end",
	FnToken:      "fn",
	IdToken:      "id",
	IndentToken:  "indent",
	IfToken:      "if",
	LoopToken:    "loop",
	NewlineToken: "newline",
	ReturnToken:  "return",
	TopRefToken:  "ref",
	ValueToken:   "value",
	WhileToken:   "while",
}

var keywords = map[string]TokenType{
	"elif":  ElifToken,
	"else":  ElseToken,
	"fn":    FnToken,
	"if":    IfToken,
	"loop":  LoopToken,
	"while": WhileToken,
}

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
	switch t.Type {
	case IdToken, ValueToken, AllRefToken, TopRefToken:
		return fmt.Sprintf("%v(%v)", t.Type, t.Literal)
	}
	return t.Type.String()
}
