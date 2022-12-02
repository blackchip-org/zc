package lang

import (
	"errors"
	"fmt"
)

type parser struct {
	s    *Scanner
	tok  Token
	next Token
}

func Parse(file string, src []byte) (NodeAST, error) {
	p := &parser{
		s: NewScanner(file, src),
	}

	// Current and next tokens are tracked for lookahead. First scan puts the
	// first token in 'next'. Next scan puts 'next' into 'tok'
	p.scan()
	p.scan()

	// If there is an error, return as much of the AST as possible
	nodes, err := p.parseFile()
	root := &FileNode{Name: file, Nodes: nodes}
	return root, err
}

func (p *parser) parseBody() ([]NodeAST, error) {
	var body []NodeAST
	if p.tok.Type != IndentToken {
		return body, p.err("expecting %v but got %v", IndentToken, p.tok)
	}

	p.scan()
	for p.tok.Type != EndToken && p.tok.Type != DedentToken {
		if p.tok.Type == NewlineToken {
			p.scan()
			continue
		}
		stmt, err := p.parseStatement()
		body = append(body, stmt)
		if err != nil {
			return body, err
		}
	}

	p.scan()
	return body, nil
}

func (p *parser) parseExpr() (*ExprNode, error) {
	expr := &ExprNode{Pos: p.tok.At}
	done := false
	for !done {
		var node NodeAST
		var err error
		switch p.tok.Type {
		case DoubleSlashToken, SlashToken:
			node, err = p.parseRef()
		case IdToken:
			if p.next.Type == SemicolonToken {
				node, err = p.parseStack()
			} else {
				node = &InvokeNode{Pos: p.tok.At, Name: p.tok.Literal}
				p.scan()
			}
		case NewlineToken, EndToken:
			done = true
			p.scan()
		case ValueToken:
			node = &ValueNode{Pos: p.tok.At, Value: p.tok.Literal}
			p.scan()
		default:
			return expr, p.err("unexpected %v", p.tok)
		}
		if err != nil {
			return expr, err
		}
		if node != nil {
			expr.Nodes = append(expr.Nodes, node)
		}
	}

	return expr, nil
}

func (p *parser) parseFile() ([]NodeAST, error) {
	var nodes []NodeAST

	done := false
	for !done {
		switch p.tok.Type {
		case EndToken:
			done = true
		case NewlineToken:
			p.scan()
		default:
			stmt, err := p.parseStatement()
			nodes = append(nodes, stmt)
			if err != nil {
				return nodes, err
			}
		}
	}
	return nodes, nil
}

func (p *parser) parseFunc() (*FuncNode, error) {
	fn := &FuncNode{Pos: p.tok.At}

	p.scan()
	if p.tok.Type != IdToken {
		return nil, p.err("expecting %v but got %v", IdToken, p.tok)
	}
	fn.Name = p.tok.Literal

	p.scan()
	for p.tok.Type == DoubleSlashToken || p.tok.Type == SlashToken {
		ref, err := p.parseRef()
		if err != nil {
			return nil, err
		}
		fn.Params = append(fn.Params, ref)
	}
	if p.tok.Type != NewlineToken {
		return nil, p.err("expecting %v but got %v", NewlineToken, p.tok)
	}

	p.scan()
	body, err := p.parseBody()
	fn.Body = body
	return fn, err
}

func (p *parser) parseInclude() (*IncludeNode, error) {
	include := &IncludeNode{Pos: p.tok.At}

	p.scan()
	for p.tok.Type == IdToken {
		include.Names = append(include.Names, p.tok.Literal)
		p.scan()
	}
	if p.tok.Type != NewlineToken && p.tok.Type != EndToken {
		return include, p.err("expecting %v but got %v", NewlineToken, p.tok)
	}

	p.scan()
	return include, nil
}

func (p *parser) parseRef() (*RefNode, error) {
	ref := &RefNode{Pos: p.tok.At}

	switch p.tok.Type {
	case DoubleSlashToken:
		ref.Type = AllRef
	case SlashToken:
		ref.Type = TopRef
	default:
		return ref, p.err("expected %v or %v but got %v", DoubleSlashToken, SlashToken, p.tok)
	}

	p.scan()
	if p.tok.Type != IdToken {
		return ref, p.err("expected %v but got %v", IdToken, p.tok)
	}
	ref.Name = p.tok.Literal

	p.scan()
	return ref, nil
}

func (p *parser) parseStack() (*StackNode, error) {
	stack := &StackNode{Pos: p.tok.At, Name: p.tok.Literal}

	p.scan()
	if p.tok.Type != SemicolonToken {
		return stack, p.err("expecting %v but got %v", SemicolonToken, p.tok)
	}

	p.scan()
	return stack, nil
}

func (p *parser) parseStatement() (NodeAST, error) {
	switch p.tok.Type {
	case DoubleSlashToken:
		return p.parseExpr()
	case FuncToken:
		return p.parseFunc()
	case IdToken:
		return p.parseExpr()
	case IncludeToken:
		return p.parseInclude()
	case SlashToken:
		return p.parseExpr()
	case ValueToken:
		return p.parseExpr()
	case WhileToken:
		return p.parseWhile()
	}
	return &BadNode{Token: p.tok}, p.err("unexpected: %v", p.tok)
}

func (p *parser) parseWhile() (*WhileNode, error) {
	while := &WhileNode{Pos: p.tok.At}
	p.scan()

	expr, err := p.parseExpr()
	while.Expr = expr
	if err != nil {
		return while, err
	}

	body, err := p.parseBody()
	while.Body = body
	if err != nil {
		return while, err
	}

	return while, nil
}

func (p *parser) scan() {
	p.tok = p.next
	p.next = p.s.Next()
}

func (p *parser) err(format string, a ...any) error {
	msg := "[" + p.tok.At.String() + "] " + fmt.Sprintf(format, a...)
	return errors.New(msg)
}
