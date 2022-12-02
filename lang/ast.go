package lang

import (
	"encoding/json"
)

type NodeAST interface {
	At() Position
	String() string
}

type BadNode struct {
	Token Token
}

func (n BadNode) At() Position   { return n.Token.At }
func (n BadNode) String() string { return nodeStringJSON(n) }

type ExprNode struct {
	Pos    Position `json:"-"`
	Target *RefNode `json:",omitempty"`
	Nodes  []NodeAST
}

func (n ExprNode) At() Position   { return n.Pos }
func (n ExprNode) String() string { return nodeStringJSON(n) }

type FileNode struct {
	Pos   Position `json:"-"`
	Name  string   `json:",omitempty"`
	Nodes []NodeAST
}

func (n FileNode) At() Position   { return n.Pos }
func (n FileNode) String() string { return nodeStringJSON(n) }

type FuncNode struct {
	Pos    Position `json:"-"`
	Name   string
	Params []*RefNode
	Body   []NodeAST
}

func (n FuncNode) At() Position   { return n.Pos }
func (n FuncNode) String() string { return nodeStringJSON(n) }

type IncludeNode struct {
	Pos   Position `json:"-"`
	Names []string
}

func (n IncludeNode) At() Position   { return n.Pos }
func (n IncludeNode) String() string { return nodeStringJSON(n) }

type InvokeNode struct {
	Pos  Position `json:"-"`
	Name string
}

func (n InvokeNode) At() Position   { return n.Pos }
func (n InvokeNode) String() string { return nodeStringJSON(n) }

type RefNode struct {
	Pos  Position `json:"-"`
	Name string
	Type RefType
}

func (n RefNode) At() Position   { return n.Pos }
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
	Pos  Position `json:"-"`
	Name string
}

func (n StackNode) At() Position   { return n.Pos }
func (n StackNode) String() string { return nodeStringJSON(n) }

type ValueNode struct {
	Pos   Position `json:"-"`
	Value string
}

func (n ValueNode) At() Position   { return n.Pos }
func (n ValueNode) String() string { return nodeStringJSON(n) }

type WhileNode struct {
	Pos  Position `json:"-"`
	Expr NodeAST
	Body []NodeAST
}

func (n WhileNode) At() Position   { return n.Pos }
func (n WhileNode) String() string { return nodeStringJSON(n) }

func nodeStringJSON(n NodeAST) string {
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
