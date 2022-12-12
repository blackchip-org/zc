package parser

import (
	"errors"
	"fmt"
	"path"

	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/scanner"
	"github.com/blackchip-org/zc/lang/token"
)

type parser struct {
	s    *scanner.Scanner
	tok  token.Token
	next token.Token
}

func Parse(file string, src []byte) (ast.Node, error) {
	p := &parser{
		s: scanner.New(file, src),
	}

	// Current and next tokens are tracked for lookahead. First scan puts the
	// first token in 'next'. Next scan puts 'next' into 'tok'
	p.scan()
	p.scan()

	// If there is an error, return as much of the AST as possible
	nodes, err := p.parseFile()
	root := &ast.FileNode{Name: file, Block: nodes}
	return root, err
}

func (p *parser) parseBlock() ([]ast.Node, error) {
	var body []ast.Node
	if p.tok.Type != token.Indent {
		return body, p.err("expecting %v but got %v", token.Indent, p.tok)
	}

	p.scan()
	for p.tok.Type != token.End && p.tok.Type != token.Dedent {
		if p.tok.Type == token.Newline {
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

func (p *parser) parseExpr() (*ast.ExprNode, error) {
	expr := &ast.ExprNode{Token: p.tok}
	done := false
	for !done {
		var node ast.Node
		var err error
		switch p.tok.Type {
		case token.DoubleSlash, token.Slash:
			node, err = p.parseRef()
		case token.Id:
			if p.next.Type == token.Semicolon {
				node, err = p.parseStack()
			} else {
				node = &ast.InvokeNode{Token: p.tok, Name: p.tok.Literal}
				p.scan()
			}
		case token.Newline, token.End:
			done = true
			p.scan()
		case token.String:
			node = &ast.ValueNode{Token: p.tok, Value: p.tok.Literal, IsString: true}
			p.scan()
		case token.Value:
			node = &ast.ValueNode{Token: p.tok, Value: p.tok.Literal}
			p.scan()
		default:
			return expr, p.err("unexpected %v", p.tok)
		}
		if err != nil {
			return expr, err
		}
		if node != nil {
			expr.Expr = append(expr.Expr, node)
		}
	}

	return expr, nil
}

func (p *parser) parseFile() ([]ast.Node, error) {
	var nodes []ast.Node

	done := false
	for !done {
		switch p.tok.Type {
		case token.End:
			done = true
		case token.Newline:
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

func (p *parser) parseFor() (*ast.ForNode, error) {
	node := &ast.ForNode{Token: p.tok}

	p.scan()
	stack, err := p.parseStack()
	node.Stack = stack
	if err != nil {
		return node, err
	}

	expr, err := p.parseExpr()
	node.Expr = expr
	if err != nil {
		return node, err
	}

	block, err := p.parseBlock()
	node.Block = block
	if err != nil {
		return node, err
	}
	return node, nil
}

func (p *parser) parseFunc() (*ast.FuncNode, error) {
	fn := &ast.FuncNode{Token: p.tok}

	p.scan()
	if p.tok.Type != token.Id {
		return nil, p.err("expecting %v but got %v", token.Id, p.tok)
	}
	fn.Name = p.tok.Literal

	p.scan()
	for p.tok.Type == token.DoubleSlash || p.tok.Type == token.Slash {
		ref, err := p.parseRef()
		if err != nil {
			return nil, err
		}
		fn.Params = append(fn.Params, ref)
	}
	if p.tok.Type != token.Newline {
		return nil, p.err("expecting %v but got %v", token.Newline, p.tok)
	}

	p.scan()
	block, err := p.parseBlock()
	fn.Block = block
	return fn, err
}

func (p *parser) parseIf() (*ast.IfNode, error) {
	ifNode := &ast.IfNode{Token: p.tok}

	caseToken := p.tok
	p.scan()
	expr, err := p.parseExpr()
	if err != nil {
		return ifNode, err
	}
	block, err := p.parseBlock()
	caseNode := &ast.IfCaseNode{Token: caseToken, Cond: expr, Block: block}
	ifNode.Cases = append(ifNode.Cases, caseNode)
	if err != nil {
		return ifNode, err
	}

	for {
		var expr *ast.ExprNode
		caseToken := p.tok
		if p.tok.Type == token.Elif {
			p.scan()
			expr, err = p.parseExpr()
			if err != nil {
				return ifNode, err
			}
		} else if p.tok.Type == token.Else {
			p.scan()
			if p.tok.Type != token.Newline {
				return ifNode, p.err("expected %v but got %v", token.Newline, p.tok)
			}
			p.scan()
		} else {
			break
		}
		block, err := p.parseBlock()

		caseNode := &ast.IfCaseNode{Token: caseToken, Cond: expr, Block: block}
		ifNode.Cases = append(ifNode.Cases, caseNode)
		if err != nil {
			return ifNode, err
		}
	}
	return ifNode, nil
}

func (p *parser) parseInclude() (*ast.IncludeNode, error) {
	node := &ast.IncludeNode{Token: p.tok}

	p.scan()
	switch p.tok.Type {
	case token.Id:
		node.Module.Name = p.tok.Literal
		node.Module.Zlib = true
	case token.String, token.Value:
		node.Module.Name = p.tok.Literal
	default:
		return node, p.err("expecting %v or %v but got %v", token.Id, token.Value, p.tok)
	}

	p.scan()
	if p.tok.Type != token.Newline && p.tok.Type != token.End {
		return node, p.err("expecting %v but got %v", token.Newline, p.tok)
	}

	p.scan()
	return node, nil
}

func (p *parser) parseImport() (*ast.ImportNode, error) {
	node := &ast.ImportNode{Token: p.tok}

	p.scan()
	switch p.tok.Type {
	case token.Id:
		node.Module.Name = p.tok.Literal
		node.Module.Zlib = true
	case token.String, token.Value:
		node.Module.Name = p.tok.Literal
		base := path.Base(node.Module.Name)
		ext := path.Ext(base)
		node.Module.Alias = base[:len(base)-len(ext)]
	default:
		return node, p.err("expecting %v or %v but got %v", token.Id, token.Value, p.tok)
	}

	p.scan()
	if p.tok.Type == token.Id {
		node.Module.Alias = p.tok.Literal
		p.scan()
	}

	if p.tok.Type != token.Newline && p.tok.Type != token.End {
		return node, p.err("expecting %v but got %v", token.Newline, p.tok)
	}

	p.scan()
	return node, nil
}

func (p *parser) parseMacro() (*ast.MacroNode, error) {
	macro := &ast.MacroNode{Token: p.tok}

	p.scan()
	if p.tok.Type != token.Id {
		return macro, p.err("expecting %v but got %v", token.Id, p.tok)
	}
	macro.Name = p.tok.Literal

	p.scan()
	expr, err := p.parseExpr()
	if err != nil {
		return macro, err
	}
	macro.Expr = expr

	return macro, nil
}

func (p *parser) parseRef() (*ast.RefNode, error) {
	ref := &ast.RefNode{Token: p.tok}

	switch p.tok.Type {
	case token.DoubleSlash:
		ref.Type = ast.AllRef
	case token.Slash:
		ref.Type = ast.TopRef
	default:
		return ref, p.err("expected %v or %v but got %v", token.DoubleSlash, token.Slash, p.tok)
	}

	p.scan()
	if p.tok.Type != token.Id {
		return ref, p.err("expected %v but got %v", token.Id, p.tok)
	}
	ref.Name = p.tok.Literal

	p.scan()
	return ref, nil
}

func (p *parser) parseStack() (*ast.StackNode, error) {
	stack := &ast.StackNode{Token: p.tok, Name: p.tok.Literal}

	p.scan()
	if p.tok.Type != token.Semicolon {
		return stack, p.err("expecting %v but got %v", token.Semicolon, p.tok)
	}

	p.scan()
	return stack, nil
}

func (p *parser) parseStatement() (ast.Node, error) {
	switch p.tok.Type {
	case token.DoubleSlash:
		return p.parseExpr()
	case token.For:
		return p.parseFor()
	case token.Func:
		return p.parseFunc()
	case token.If:
		return p.parseIf()
	case token.Id:
		return p.parseExpr()
	case token.Import:
		return p.parseImport()
	case token.Include:
		return p.parseInclude()
	case token.Macro:
		return p.parseMacro()
	case token.Slash:
		return p.parseExpr()
	case token.String:
		return p.parseExpr()
	case token.Try:
		return p.parseTry()
	case token.Use:
		return p.parseUse()
	case token.Value:
		return p.parseExpr()
	case token.While:
		return p.parseWhile()
	}
	return &ast.BadNode{Token: p.tok}, p.err("unexpected: %v", p.tok)
}

func (p *parser) parseTry() (*ast.TryNode, error) {
	try := &ast.TryNode{Token: p.tok}
	p.scan()

	expr, err := p.parseExpr()
	try.Expr = expr
	if err != nil {
		return try, err
	}

	return try, err
}

func (p *parser) parseUse() (*ast.UseNode, error) {
	node := &ast.UseNode{Token: p.tok}

	p.scan()
	if p.tok.Type != token.Id {
		return node, p.err("expecting %v but got %v", token.Id, p.tok)
	}
	node.Name = p.tok.Literal

	p.scan()
	return node, nil
}

func (p *parser) parseWhile() (*ast.WhileNode, error) {
	while := &ast.WhileNode{Token: p.tok}
	p.scan()

	expr, err := p.parseExpr()
	while.Cond = expr
	if err != nil {
		return while, err
	}

	block, err := p.parseBlock()
	while.Block = block
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
	msg := "[" + p.tok.Pos.String() + "] " + fmt.Sprintf(format, a...)
	return errors.New(msg)
}
