package lang

import (
	"fmt"
)

type parser struct {
	s      *Scanner
	tok    Token
	next   Token
	errors Errors
}

func Parse(file string, src []byte) (NodeAST, error) {
	p := &parser{
		s: NewScanner(file, src),
	}
	p.scan()
	p.scan()

	nodes := p.parseFile()
	root := &FileNode{Name: file, Nodes: nodes}
	if len(p.errors) > 0 {
		return root, p.errors
	}
	return root, nil
}

func (p *parser) parseBody() []NodeAST {
	var body []NodeAST
	if p.tok.Type != IndentToken {
		p.err("expecting %v but got %v", IndentToken, p.tok)
	}
	p.scan()

	for p.tok.Type != EndToken && p.tok.Type != DedentToken {
		stmt := p.parseStatement()
		body = append(body, stmt)
	}
	p.scan()
	return body
}

func (p *parser) parseFile() []NodeAST {
	var nodes []NodeAST

	done := false
	for !done {
		switch p.tok.Type {
		case EndToken:
			done = true
		case NewlineToken:
			p.scan()
		default:
			nodes = append(nodes, p.parseStatement())
		}
	}
	return nodes
}

func (p *parser) parseFn() NodeAST {
	var fn FnNode
	p.scan()
	if p.tok.Type != IdToken {
		p.err("expecting %v but got %v", IdToken, p.tok)
	}
	fn.Name = p.tok.Literal

	p.scan()
	for p.tok.Type == AllRefToken || p.tok.Type == TopRefToken {
		ref := p.parseRef()
		fn.Params = append(fn.Params, ref)
	}

	if p.tok.Type != NewlineToken {
		p.err("expecting %v but got %v", NewlineToken, p.tok)
	}
	p.scan()
	fn.Body = p.parseBody()

	return fn
}

func (p *parser) parseExpr() NodeAST {
	line := &ExprNode{}
	for {
		var node NodeAST
		switch p.tok.Type {
		case AllRefToken, TopRefToken:
			node = p.parseRef()
		case IdToken:
			node = &InvokeNode{Token: p.tok, Name: p.tok.Literal}
			p.scan()
		case ValueToken:
			node = &ValueNode{Token: p.tok, Value: p.tok.Literal}
			p.scan()
		}
		if node != nil {
			line.Nodes = append(line.Nodes, node)
		} else {
			break
		}
	}
	if p.tok.Type != NewlineToken && p.tok.Type != EndToken {
		return p.err("unexpected %v", p.tok)
	}
	p.scan()
	if len(line.Nodes) == 0 {
		return nil
	}
	return line
}

func (p *parser) parseRef() NodeAST {
	var ref RefNode
	switch p.tok.Type {
	case AllRefToken:
		ref.Type = RefAll
	case TopRefToken:
		ref.Type = RefTop
	default:
		return p.err("expected %v or %v but got %v", AllRefToken, TopRefToken, p.tok)
	}
	p.scan()

	if p.tok.Type != IdToken {
		return p.err("expected %v but got %v", IdToken, p.tok)
	}
	ref.Name = p.tok.Literal
	p.scan()

	return ref
}

func (p *parser) parseStatement() NodeAST {
	switch p.tok.Type {
	case AllRefToken:
		return p.parseExpr()
	case FnToken:
		return p.parseFn()
	case IdToken:
		return p.parseExpr()
	case TopRefToken:
		return p.parseExpr()
	case ValueToken:
		return p.parseExpr()
	case WhileToken:
		return p.parseWhile()
	}
	return p.err("unexpected: %v", p.tok)
}

func (p *parser) parseWhile() NodeAST {
	while := WhileNode{Pos: p.tok.At}
	p.scan()

	while.Expr = p.parseExpr()
	while.Body = p.parseBody()
	return while
}

func (p *parser) scan() {
	p.tok = p.next
	p.next = p.s.Next()
	//fmt.Printf("[%v] %v %v\n", p.tok.At, p.tok.Type, p.tok.Literal)
}

func (p *parser) err(spec string, a ...interface{}) *BadNode {
	msg := fmt.Sprintf(spec, a...)
	//panic(msg)
	e := fmt.Errorf("%v: %v", p.tok.At, msg)
	p.errors = append(p.errors, e)
	return &BadNode{Tokens: p.recover()}
}

func (p *parser) recover() []Token {
	var tokens []Token
	for p.tok.Type != NewlineToken && p.tok.Type != EndToken {
		tokens = append(tokens, p.tok)
		p.scan()
	}
	return tokens
}
