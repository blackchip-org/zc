package ast

import (
	"encoding/json"

	"github.com/blackchip-org/zc/lang/token"
)

type Node interface {
	Pos() token.Pos
	String() string
}

type Atom interface {
	Node
	atomNode()
}

type Stmt interface {
	Node
	stmtNode()
}

type ModuleRef struct {
	Zlib  bool
	Name  string
	Alias string
}

type RefType int

const (
	InvalidRef RefType = iota
	TopRef
	AllRef
	PopRef
)

func (r RefType) String() string {
	switch r {
	case TopRef:
		return "/"
	case AllRef:
		return "//"
	case PopRef:
		return "/-"
	}
	return "???"
}

type AliasStmt struct {
	Token token.Token
	From  string
	To    string
}

func (n AliasStmt) Pos() token.Pos { return n.Token.Pos }
func (n AliasStmt) String() string { return nodeStringJSON(n) }
func (n AliasStmt) stmtNode()      {}

type BadStmt struct {
	Token token.Token
}

func (n BadStmt) Pos() token.Pos { return n.Token.Pos }
func (n BadStmt) String() string { return nodeStringJSON(n) }
func (n BadStmt) stmtNode()      {}

type Expr struct {
	Token  token.Token `json:"-"`
	Target *RefAtom    `json:",omitempty"`
	Atoms  []Atom
}

func (n Expr) Pos() token.Pos { return n.Token.Pos }
func (n Expr) String() string { return nodeStringJSON(n) }

type ExprStmt struct {
	Token token.Token `json:"-"`
	Expr  *Expr
}

func (n ExprStmt) Pos() token.Pos { return n.Token.Pos }
func (n ExprStmt) String() string { return nodeStringJSON(n) }
func (n ExprStmt) stmtNode()      {}

type File struct {
	Token token.Token `json:"-"`
	Name  string      `json:",omitempty"`
	Stmts []Stmt
}

func (n File) Pos() token.Pos { return n.Token.Pos }
func (n File) String() string { return nodeStringJSON(n) }

type ForStmt struct {
	Token token.Token `json:"-"`
	Stack *SelectAtom
	Expr  *Expr
	Stmts []Stmt
}

func (n ForStmt) Pos() token.Pos { return n.Token.Pos }
func (n ForStmt) String() string { return nodeStringJSON(n) }
func (n ForStmt) stmtNode()      {}

type FuncStmt struct {
	Token  token.Token `json:"-"`
	Name   string
	Params []*RefAtom
	Stmts  []Stmt
}

func (n FuncStmt) Pos() token.Pos { return n.Token.Pos }
func (n FuncStmt) String() string { return nodeStringJSON(n) }
func (n FuncStmt) stmtNode()      {}

type IfStmt struct {
	Token token.Token `json:"-"`
	Cases []*IfCaseNode
}

func (n IfStmt) Pos() token.Pos { return n.Token.Pos }
func (n IfStmt) String() string { return nodeStringJSON(n) }
func (n IfStmt) stmtNode()      {}

type IfCaseNode struct {
	Token token.Token `json:"-"`
	Cond  *Expr       `json:",omitempty"` // for the final "else", this is nil
	Stmts []Stmt
}

func (n IfCaseNode) Pos() token.Pos { return n.Token.Pos }
func (n IfCaseNode) String() string { return nodeStringJSON(n) }

type ImportStmt struct {
	Token   token.Token `json:"-"`
	Modules []ModuleRef
}

func (n ImportStmt) Pos() token.Pos { return n.Token.Pos }
func (n ImportStmt) String() string { return nodeStringJSON(n) }
func (n ImportStmt) stmtNode()      {}

type InvokeAtom struct {
	Token token.Token `json:"-"`
	Name  string
}

func (n InvokeAtom) Pos() token.Pos { return n.Token.Pos }
func (n InvokeAtom) String() string { return nodeStringJSON(n) }
func (n InvokeAtom) atomNode()      {}

type MacroStmt struct {
	Token token.Token `json:"-"`
	Name  string
	Expr  *Expr
}

func (n MacroStmt) Pos() token.Pos { return n.Token.Pos }
func (n MacroStmt) String() string { return nodeStringJSON(n) }
func (n MacroStmt) stmtNode()      {}

type NativeStmt struct {
	Token  token.Token `json:"-"`
	Name   string
	Export string `json:",omitempty"`
}

func (n NativeStmt) Pos() token.Pos { return n.Token.Pos }
func (n NativeStmt) String() string { return nodeStringJSON(n) }
func (n NativeStmt) stmtNode()      {}

type NumberAtom struct {
	Token token.Token `json:"-"`
	Value string
}

func (n NumberAtom) Pos() token.Pos { return n.Token.Pos }
func (n NumberAtom) String() string { return nodeStringJSON(n) }
func (n NumberAtom) atomNode()      {}

type RefAtom struct {
	Token token.Token `json:"-"`
	Name  string
	Type  RefType
}

func (n RefAtom) Pos() token.Pos { return n.Token.Pos }
func (n RefAtom) String() string { return nodeStringJSON(n) }
func (n RefAtom) atomNode()      {}

type ReturnStmt struct {
	Token token.Token `json:"-"`
}

func (n ReturnStmt) Pos() token.Pos { return n.Token.Pos }
func (n ReturnStmt) String() string { return nodeStringJSON(n) }
func (n ReturnStmt) stmtNode()      {}

type SelectAtom struct {
	Token token.Token `json:"-"`
	Name  string
}

func (n SelectAtom) Pos() token.Pos { return n.Token.Pos }
func (n SelectAtom) String() string { return nodeStringJSON(n) }
func (n SelectAtom) atomNode()      {}

type TryStmt struct {
	Token token.Token `json:"-"`
	Expr  *Expr
}

func (n TryStmt) Pos() token.Pos { return n.Token.Pos }
func (n TryStmt) String() string { return nodeStringJSON(n) }
func (n TryStmt) stmtNode()      {}

type UseStmt struct {
	Token   token.Token `json:"-"`
	Modules []ModuleRef
}

func (n UseStmt) Pos() token.Pos { return n.Token.Pos }
func (n UseStmt) String() string { return nodeStringJSON(n) }
func (n UseStmt) stmtNode()      {}

type ValueAtom struct {
	Token    token.Token `json:"-"`
	IsString bool
	Value    string
}

func (n ValueAtom) Pos() token.Pos { return n.Token.Pos }
func (n ValueAtom) String() string { return nodeStringJSON(n) }
func (n ValueAtom) atomNode()      {}

type WhileStmt struct {
	Token token.Token `json:"-"`
	Cond  *Expr
	Stmts []Stmt
}

func (n WhileStmt) Pos() token.Pos { return n.Token.Pos }
func (n WhileStmt) String() string { return nodeStringJSON(n) }
func (n WhileStmt) stmtNode()      {}

func nodeStringJSON(n Node) string {
	b, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Add synthetic field
// http://choly.ca/post/go-json-marshalling/

func (n AliasStmt) MarshalJSON() ([]byte, error) {
	type Alias AliasStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Alias",
		Alias: (Alias)(n),
	})
}

func (n BadStmt) MarshalJSON() ([]byte, error) {
	type Alias BadStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Bad",
		Alias: (Alias)(n),
	})
}

func (n Expr) MarshalJSON() ([]byte, error) {
	type Alias Expr
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Expr",
		Alias: (Alias)(n),
	})
}

func (n ExprStmt) MarshalJSON() ([]byte, error) {
	type Alias ExprStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "ExprStmt",
		Alias: (Alias)(n),
	})
}

func (n File) MarshalJSON() ([]byte, error) {
	type Alias File
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "File",
		Alias: (Alias)(n),
	})
}

func (n ForStmt) MarshalJSON() ([]byte, error) {
	type Alias ForStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "For",
		Alias: (Alias)(n),
	})
}

func (n FuncStmt) MarshalJSON() ([]byte, error) {
	type Alias FuncStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Func",
		Alias: (Alias)(n),
	})
}

func (n IfStmt) MarshalJSON() ([]byte, error) {
	type Alias IfStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "If",
		Alias: (Alias)(n),
	})
}

func (n IfCaseNode) MarshalJSON() ([]byte, error) {
	type Alias IfCaseNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "IfCase",
		Alias: (Alias)(n),
	})
}

func (n ImportStmt) MarshalJSON() ([]byte, error) {
	type Alias ImportStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Import",
		Alias: (Alias)(n),
	})
}

func (n InvokeAtom) MarshalJSON() ([]byte, error) {
	type Alias InvokeAtom
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Invoke",
		Alias: (Alias)(n),
	})
}

func (n MacroStmt) MarshalJSON() ([]byte, error) {
	type Alias MacroStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Macro",
		Alias: (Alias)(n),
	})
}

func (n NativeStmt) MarshalJSON() ([]byte, error) {
	type Alias NativeStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Native",
		Alias: (Alias)(n),
	})
}

func (n NumberAtom) MarshalJSON() ([]byte, error) {
	type Alias NumberAtom
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Number",
		Alias: (Alias)(n),
	})
}

func (n RefAtom) MarshalJSON() ([]byte, error) {
	type Alias RefAtom
	return json.Marshal(&struct {
		Node string
		Type string
		Alias
	}{
		Node:  "Ref",
		Type:  n.Type.String(),
		Alias: (Alias)(n),
	})
}

func (n ReturnStmt) MarshalJSON() ([]byte, error) {
	type Alias ReturnStmt
	return json.Marshal(&struct {
		Node string
		Type string
		Alias
	}{
		Node:  "Return",
		Alias: (Alias)(n),
	})
}

func (n SelectAtom) MarshalJSON() ([]byte, error) {
	type Alias SelectAtom
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Select",
		Alias: (Alias)(n),
	})
}

func (n TryStmt) MarshalJSON() ([]byte, error) {
	type Alias TryStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Try",
		Alias: (Alias)(n),
	})
}

func (n UseStmt) MarshalJSON() ([]byte, error) {
	type Alias UseStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Use",
		Alias: (Alias)(n),
	})
}

func (n ValueAtom) MarshalJSON() ([]byte, error) {
	type Alias ValueAtom
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Value",
		Alias: (Alias)(n),
	})
}

func (n WhileStmt) MarshalJSON() ([]byte, error) {
	type Alias WhileStmt
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "While",
		Alias: (Alias)(n),
	})
}
