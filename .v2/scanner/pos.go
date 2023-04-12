package scanner

import "fmt"

type Pos struct {
	Name   string
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
	if p.Name != "" {
		return fmt.Sprintf("%v:%v:%v", p.Name, p.Line, p.Column)
	}
	return fmt.Sprintf("%v:%v", p.Line, p.Column)
}
