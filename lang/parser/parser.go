package parser

import (
	"errors"
	"fmt"
	"path"

	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/lexer"
	"github.com/blackchip-org/zc/lang/token"
)

type parser struct {
	l    *lexer.Lexer
	tok  token.Token
	next token.Token
}

func Parse(file string, src []byte) (*ast.File, error) {
	return parse(file, src)
}

func ParseRuntime(file string, src []byte) (*ast.File, error) {
	return parse(file, src)
}

func parse(file string, src []byte) (*ast.File, error) {
	p := &parser{
		l: lexer.New(file, src),
	}

	// Current and next tokens are tracked for lookahead. First scan puts the
	// first token in 'next'. Next scan puts 'next' into 'tok'
	p.scan()
	p.scan()

	// If there is an error, return as much of the AST as possible
	stmts, err := p.parseFile()
	root := &ast.File{Name: file, Stmts: stmts}
	return root, err
}

func (p *parser) parseAliasStmt() (*ast.AliasStmt, error) {
	node := &ast.AliasStmt{Token: p.tok}

	p.scan()
	if p.tok.Type != token.Id {
		return node, p.err("expecting %v but got %v", token.Id, p.tok)
	}
	node.To = p.tok.Literal

	p.scan()
	if p.tok.Type != token.Id {
		return node, p.err("expecting %v but got %v", token.Id, p.tok)
	}
	node.From = p.tok.Literal

	p.scan()
	if p.tok.Type != token.Newline && p.tok.Type != token.End {
		return node, p.err("expecting %v but got %v", token.Newline, p.tok)
	}

	p.scan()
	return node, nil
}

func (p *parser) parseExpr() (*ast.Expr, error) {
	expr := &ast.Expr{Token: p.tok}
	done := false
	for !done {
		var atom ast.Atom
		var err error
		switch p.tok.Type {
		case token.DoubleSlash, token.Slash, token.SlashDash:
			atom, err = p.parseRefAtom()
		case token.Id:
			if p.next.Type == token.Semicolon {
				atom, err = p.parseSelectAtom()
			} else {
				atom = &ast.InvokeAtom{Token: p.tok, Name: p.tok.Literal}
				p.scan()
			}
		case token.Newline, token.End:
			done = true
			p.scan()
		case token.String:
			atom = &ast.ValueAtom{Token: p.tok, Value: p.tok.Literal, IsString: true}
			p.scan()
		case token.StringPlain:
			atom = &ast.ValueAtom{Token: p.tok, Value: p.tok.Literal, IsString: true, IsPlain: true}
			p.scan()
		case token.Value:
			atom = &ast.ValueAtom{Token: p.tok, Value: p.tok.Literal}
			p.scan()
		default:
			return expr, p.err("unexpected %v", p.tok)
		}
		if err != nil {
			return expr, err
		}
		if atom != nil {
			expr.Atoms = append(expr.Atoms, atom)
		}
	}

	return expr, nil
}

func (p *parser) parseExprStmt() (ast.Stmt, error) {
	stmt := &ast.ExprStmt{Token: p.tok}
	expr, err := p.parseExpr()
	stmt.Expr = expr
	return stmt, err
}

func (p *parser) parseFile() ([]ast.Stmt, error) {
	var stmts []ast.Stmt

	done := false
	for !done {
		switch p.tok.Type {
		case token.End:
			done = true
		case token.Newline:
			p.scan()
		default:
			stmt, err := p.parseStmtTop()
			stmts = append(stmts, stmt)
			if err != nil {
				return stmts, err
			}
		}
	}
	return stmts, nil
}

func (p *parser) parseForStmt() (*ast.ForStmt, error) {
	node := &ast.ForStmt{Token: p.tok}

	p.scan()
	stack, err := p.parseSelectAtom()
	node.Stack = stack
	if err != nil {
		return node, err
	}

	expr, err := p.parseExpr()
	node.Expr = expr
	if err != nil {
		return node, err
	}

	stmts, err := p.parseStmts()
	node.Stmts = stmts
	if err != nil {
		return node, err
	}
	return node, nil
}

func (p *parser) parseFuncStmt() (*ast.FuncStmt, error) {
	fn := &ast.FuncStmt{Token: p.tok}

	p.scan()
	if p.tok.Type != token.Id {
		return nil, p.err("expecting %v but got %v", token.Id, p.tok)
	}
	fn.Name = p.tok.Literal

	p.scan()
	for p.tok.Type == token.DoubleSlash || p.tok.Type == token.Slash {
		ref, err := p.parseRefAtom()
		if err != nil {
			return nil, err
		}
		fn.Params = append([]*ast.RefAtom{ref}, fn.Params...)
	}
	if p.tok.Type != token.Newline {
		return nil, p.err("expecting %v but got %v", token.Newline, p.tok)
	}

	p.scan()
	stmts, err := p.parseStmts()
	fn.Stmts = stmts
	return fn, err
}

func (p *parser) parseIfStmt() (*ast.IfStmt, error) {
	ifNode := &ast.IfStmt{Token: p.tok}

	caseToken := p.tok
	p.scan()
	expr, err := p.parseExpr()
	if err != nil {
		return ifNode, err
	}
	stmts, err := p.parseStmts()
	caseNode := &ast.IfCaseNode{Token: caseToken, Cond: expr, Stmts: stmts}
	ifNode.Cases = append(ifNode.Cases, caseNode)
	if err != nil {
		return ifNode, err
	}

	for {
		var expr *ast.Expr
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
		stmts, err := p.parseStmts()

		caseNode := &ast.IfCaseNode{Token: caseToken, Cond: expr, Stmts: stmts}
		ifNode.Cases = append(ifNode.Cases, caseNode)
		if err != nil {
			return ifNode, err
		}
	}
	return ifNode, nil
}

func (p *parser) parseImportStmt() (*ast.ImportStmt, error) {
	node := &ast.ImportStmt{Token: p.tok}
	var refs []ast.ModuleRef

	p.scan()
	for p.tok.Type != token.Newline && p.tok.Type != token.End {
		var ref ast.ModuleRef
		switch p.tok.Type {
		case token.Id:
			ref.Name = p.tok.Literal
			ref.Zlib = true
		case token.String, token.Value:
			ref.Name = p.tok.Literal
			base := path.Base(ref.Name)
			ext := path.Ext(base)
			ref.Alias = base[:len(base)-len(ext)]
		default:
			return node, p.err("expecting %v or %v but got %v", token.Id, token.Value, p.tok)
		}
		refs = append(refs, ref)
		p.scan()
	}

	node.Modules = refs
	p.scan()
	return node, nil
}

func (p *parser) parseMacroStmt() (*ast.MacroStmt, error) {
	macro := &ast.MacroStmt{Token: p.tok}

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

func (p *parser) parseNativeStmt() (*ast.NativeStmt, error) {
	node := &ast.NativeStmt{Token: p.tok}

	p.scan()
	if p.tok.Type != token.Id {
		return node, p.err("expecting %v but got %v", token.Id, p.tok)
	}
	node.Name = p.tok.Literal

	p.scan()
	if p.tok.Type == token.Id {
		node.Export = p.tok.Literal
		p.scan()
	}

	if p.tok.Type != token.Newline && p.tok.Type != token.End {
		return node, p.err("expecting %v but got %v", token.Newline, p.tok)
	}

	p.scan()
	return node, nil
}

func (p *parser) parseRefAtom() (*ast.RefAtom, error) {
	ref := &ast.RefAtom{Token: p.tok}

	switch p.tok.Type {
	case token.DoubleSlash:
		ref.Type = ast.AllRef
	case token.Slash:
		ref.Type = ast.TopRef
	case token.SlashDash:
		ref.Type = ast.PopRef
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

func (p *parser) parseReturnStmt() (*ast.ReturnStmt, error) {
	p.scan()
	return &ast.ReturnStmt{Token: p.tok}, nil
}

func (p *parser) parseSelectAtom() (*ast.SelectAtom, error) {
	stack := &ast.SelectAtom{Token: p.tok, Name: p.tok.Literal}

	p.scan()
	if p.tok.Type != token.Semicolon {
		return stack, p.err("expecting %v but got %v", token.Semicolon, p.tok)
	}

	p.scan()
	return stack, nil
}

func (p *parser) parseStmtTop() (ast.Stmt, error) {
	switch p.tok.Type {
	case token.Alias:
		return p.parseAliasStmt()
	case token.Func:
		return p.parseFuncStmt()
	case token.Import:
		return p.parseImportStmt()
	case token.Macro:
		return p.parseMacroStmt()
	case token.Native:
		return p.parseNativeStmt()
	case token.Use:
		return p.parseUseStmt()
	}
	return p.parseStmtNested()
}

func (p *parser) parseStmtNested() (ast.Stmt, error) {
	switch p.tok.Type {
	case token.DoubleSlash:
		return p.parseExprStmt()
	case token.For:
		return p.parseForStmt()
	case token.If:
		return p.parseIfStmt()
	case token.Id:
		return p.parseExprStmt()
	case token.Return:
		return p.parseReturnStmt()
	case token.Slash:
		return p.parseExprStmt()
	case token.SlashDash:
		return p.parseExprStmt()
	case token.String:
		return p.parseExprStmt()
	case token.StringPlain:
		return p.parseExprStmt()
	case token.Try:
		return p.parseTryStmt()
	case token.Value:
		return p.parseExprStmt()
	case token.While:
		return p.parseWhileStmt()
	}
	return &ast.BadStmt{Token: p.tok}, p.err("unexpected: %v", p.tok)
}

func (p *parser) parseStmts() ([]ast.Stmt, error) {
	var stmts []ast.Stmt
	if p.tok.Type != token.Indent {
		return stmts, p.err("expecting %v but got %v", token.Indent, p.tok)
	}

	p.scan()
	for p.tok.Type != token.End && p.tok.Type != token.Dedent {
		if p.tok.Type == token.Newline {
			p.scan()
			continue
		}
		stmt, err := p.parseStmtNested()
		stmts = append(stmts, stmt)
		if err != nil {
			return stmts, err
		}
	}

	p.scan()
	return stmts, nil
}

func (p *parser) parseTryStmt() (*ast.TryStmt, error) {
	try := &ast.TryStmt{Token: p.tok}
	p.scan()

	expr, err := p.parseExpr()
	try.Expr = expr
	if err != nil {
		return try, err
	}

	return try, err
}

func (p *parser) parseUseStmt() (*ast.UseStmt, error) {
	node := &ast.UseStmt{Token: p.tok}

	p.scan()
	var refs []ast.ModuleRef

	for p.tok.Type != token.Newline && p.tok.Type != token.End {
		var ref ast.ModuleRef
		switch p.tok.Type {
		case token.Id:
			ref.Name = p.tok.Literal
			ref.Zlib = true
		case token.String, token.Value:
			ref.Name = p.tok.Literal
		default:
			return node, p.err("expecting %v or %v but got %v", token.Id, token.Value, p.tok)
		}
		refs = append(refs, ref)
		p.scan()
	}
	node.Modules = refs

	return node, nil
}

func (p *parser) parseWhileStmt() (*ast.WhileStmt, error) {
	while := &ast.WhileStmt{Token: p.tok}
	p.scan()

	expr, err := p.parseExpr()
	while.Cond = expr
	if err != nil {
		return while, err
	}

	stmts, err := p.parseStmts()
	while.Stmts = stmts
	if err != nil {
		return while, err
	}

	return while, nil
}

func (p *parser) scan() {
	p.tok = p.next
	p.next = p.l.Next()
}

func (p *parser) err(format string, a ...any) error {
	msg := "[" + p.tok.Pos.String() + "] " + fmt.Sprintf(format, a...)
	return errors.New(msg)
}
