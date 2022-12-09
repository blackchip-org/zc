package ast

import (
	"encoding/json"

	"github.com/blackchip-org/zc/lang/token"
)

type Node interface {
	Pos() token.Pos
	String() string
}

type BadNode struct {
	Token token.Token
}

func (n BadNode) Pos() token.Pos { return n.Token.Pos }
func (n BadNode) String() string { return nodeStringJSON(n) }

type ExprNode struct {
	Token  token.Token `json:"-"`
	Target *RefNode    `json:",omitempty"`
	Expr   []Node
}

func (n ExprNode) Pos() token.Pos { return n.Token.Pos }
func (n ExprNode) String() string { return nodeStringJSON(n) }

type FileNode struct {
	Token token.Token `json:"-"`
	Name  string      `json:",omitempty"`
	Block []Node
}

func (n FileNode) Pos() token.Pos { return n.Token.Pos }
func (n FileNode) String() string { return nodeStringJSON(n) }

type ForNode struct {
	Token token.Token `json:"-"`
	Stack *StackNode
	Expr  *ExprNode
	Block []Node
}

func (n ForNode) Pos() token.Pos { return n.Token.Pos }
func (n ForNode) String() string { return nodeStringJSON(n) }

type FuncNode struct {
	Token  token.Token `json:"-"`
	Name   string
	Params []*RefNode
	Block  []Node
}

func (n FuncNode) Pos() token.Pos { return n.Token.Pos }
func (n FuncNode) String() string { return nodeStringJSON(n) }

type IfNode struct {
	Token token.Token `json:"-"`
	Cases []*IfCaseNode
}

func (n IfNode) Pos() token.Pos { return n.Token.Pos }
func (n IfNode) String() string { return nodeStringJSON(n) }

type IfCaseNode struct {
	Token token.Token `json:"-"`
	Cond  *ExprNode   `json:",omitempty"` // for the final "else", this is nil
	Block []Node
}

func (n IfCaseNode) Pos() token.Pos { return n.Token.Pos }
func (n IfCaseNode) String() string { return nodeStringJSON(n) }

type ImportNode struct {
	Token   token.Token `json:"-"`
	Modules []ModuleRef
}

func (n ImportNode) Pos() token.Pos { return n.Token.Pos }
func (n ImportNode) String() string { return nodeStringJSON(n) }

type ModuleRef struct {
	Zlib  bool
	Name  string
	Alias string
}

type IncludeNode struct {
	Token token.Token `json:"-"`
	Names []string
}

func (n IncludeNode) Pos() token.Pos { return n.Token.Pos }
func (n IncludeNode) String() string { return nodeStringJSON(n) }

type InvokeNode struct {
	Token token.Token `json:"-"`
	Name  string
}

func (n InvokeNode) Pos() token.Pos { return n.Token.Pos }
func (n InvokeNode) String() string { return nodeStringJSON(n) }

type MacroNode struct {
	Token token.Token `json:"-"`
	Name  string
	Expr  *ExprNode
}

func (n MacroNode) Pos() token.Pos { return n.Token.Pos }
func (n MacroNode) String() string { return nodeStringJSON(n) }

type RefNode struct {
	Token token.Token `json:"-"`
	Name  string
	Type  RefType
}

func (n RefNode) Pos() token.Pos { return n.Token.Pos }
func (n RefNode) String() string { return nodeStringJSON(n) }

type RefType int

const (
	InvalidRef RefType = iota
	TopRef
	AllRef
)

func (r RefType) String() string {
	switch r {
	case TopRef:
		return "/"
	case AllRef:
		return "//"
	}
	return "???"
}

type StackNode struct {
	Token token.Token `json:"-"`
	Name  string
}

func (n StackNode) Pos() token.Pos { return n.Token.Pos }
func (n StackNode) String() string { return nodeStringJSON(n) }

type TryNode struct {
	Token token.Token `json:"-"`
	Expr  *ExprNode
}

func (n TryNode) Pos() token.Pos { return n.Token.Pos }
func (n TryNode) String() string { return nodeStringJSON(n) }

type ValueNode struct {
	Token token.Token `json:"-"`
	Value string
}

func (n ValueNode) Pos() token.Pos { return n.Token.Pos }
func (n ValueNode) String() string { return nodeStringJSON(n) }

type WhileNode struct {
	Token token.Token `json:"-"`
	Cond  *ExprNode
	Block []Node
}

func (n WhileNode) Pos() token.Pos { return n.Token.Pos }
func (n WhileNode) String() string { return nodeStringJSON(n) }

func nodeStringJSON(n Node) string {
	b, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

// Add synthetic field
// http://choly.ca/post/go-json-marshalling/

func (n BadNode) MarshalJSON() ([]byte, error) {
	type Alias BadNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Bad",
		Alias: (Alias)(n),
	})
}

func (n ExprNode) MarshalJSON() ([]byte, error) {
	type Alias ExprNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Expr",
		Alias: (Alias)(n),
	})
}

func (n FileNode) MarshalJSON() ([]byte, error) {
	type Alias FileNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "File",
		Alias: (Alias)(n),
	})
}

func (n ForNode) MarshalJSON() ([]byte, error) {
	type Alias ForNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "For",
		Alias: (Alias)(n),
	})
}

func (n FuncNode) MarshalJSON() ([]byte, error) {
	type Alias FuncNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Func",
		Alias: (Alias)(n),
	})
}

func (n IfNode) MarshalJSON() ([]byte, error) {
	type Alias IfNode
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

func (n ImportNode) MarshalJSON() ([]byte, error) {
	type Alias ImportNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Import",
		Alias: (Alias)(n),
	})
}

func (n IncludeNode) MarshalJSON() ([]byte, error) {
	type Alias IncludeNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Include",
		Alias: (Alias)(n),
	})
}

func (n InvokeNode) MarshalJSON() ([]byte, error) {
	type Alias InvokeNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Invoke",
		Alias: (Alias)(n),
	})
}

func (n MacroNode) MarshalJSON() ([]byte, error) {
	type Alias MacroNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Macro",
		Alias: (Alias)(n),
	})
}

func (n RefNode) MarshalJSON() ([]byte, error) {
	type Alias RefNode
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

func (n StackNode) MarshalJSON() ([]byte, error) {
	type Alias StackNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Stack",
		Alias: (Alias)(n),
	})
}

func (n TryNode) MarshalJSON() ([]byte, error) {
	type Alias TryNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Try",
		Alias: (Alias)(n),
	})
}

func (n ValueNode) MarshalJSON() ([]byte, error) {
	type Alias ValueNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Value",
		Alias: (Alias)(n),
	})
}

func (n WhileNode) MarshalJSON() ([]byte, error) {
	type Alias WhileNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "While",
		Alias: (Alias)(n),
	})
}
