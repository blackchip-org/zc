package scanner

import "fmt"

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
