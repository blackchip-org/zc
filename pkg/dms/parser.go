package dms

import "github.com/blackchip-org/scan"

type Angle struct {
	Sign int
	Deg  float64
	Min  float64
	Sec  float64
	Hemi string
}

type Parser struct {
	ctx     *Context
	scanner scan.Scanner
}

func NewParser(ctx *Context) *Parser {
	return &Parser{ctx: ctx}
}

func (p *Parser) Parse(v string) (Angle, error) {
	var a Angle
	p.scanner.InitFromString("", v)
	r := scan.NewRunner(&p.scanner, p.ctx.RuleSet)

	var deg, min, sec scan.Token

	tok := r.Next()
	switch tok.Type {
	case "+":
		a.Sign = 1
		tok = r.Next()
	case "-":
		a.Sign = -1
		tok = r.Next()
	}
}

func (p *Parser) parseDeg() (scan.Token, error) {

}
