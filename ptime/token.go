package ptime

import "fmt"

type TokenType int

const (
	End TokenType = iota
	Number
	Text
	Indicator
)

func (t TokenType) String() string {
	switch t {
	case Number:
		return "number"
	case Text:
		return "text"
	case Indicator:
		return "indicator"
	}
	return "invalid"
}

type Token struct {
	Type TokenType
	Val  string
	Pos  int
}

func (t Token) String() string {
	return fmt.Sprintf("%v '%v' (c%v)", t.Type, t.Val, t.Pos)
}
