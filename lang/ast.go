package lang

import (
	"encoding/json"
)

type NodeAST interface {
	At() Position
	String() string
}

type BadNode struct {
	Tokens []Token
}

func (n BadNode) At() Position   { return n.Tokens[0].At }
func (n BadNode) String() string { return nodeStringJSON(n) }

type ExprNode struct {
	NodeAST `json:"-"`
	Target  *RefNode `json:",omitempty"`
	Nodes   []NodeAST
}

func (n ExprNode) At() Position {
	if n.Target == nil {
		return n.Nodes[0].At()
	}
	return n.Target.At()
}

func (n ExprNode) String() string { return nodeStringJSON(n) }

type FileNode struct {
	NodeAST `json:"-"`
	Name    string `json:",omitempty"`
	Nodes   []NodeAST
}

func (n FileNode) At() Position   { return n.Nodes[0].At() }
func (n FileNode) String() string { return nodeStringJSON(n) }

type FnNode struct {
	NodeAST `json:"-"`
	Token   Token `json:"-"`
	Name    string
	Params  []NodeAST
	Body    []NodeAST
}

func (n FnNode) At() Position   { return n.Token.At }
func (n FnNode) String() string { return nodeStringJSON(n) }

type InvokeNode struct {
	NodeAST `json:"-"`
	Token   Token `json:"-"`
	Name    string
}

func (n InvokeNode) At() Position   { return n.Token.At }
func (n InvokeNode) String() string { return nodeStringJSON(n) }

type RefNode struct {
	NodeAST `json:"-"`
	First   Token `json:"-"`
	Name    string
	Type    RefType
}

func (n RefNode) At() Position   { return n.First.At }
func (n RefNode) String() string { return nodeStringJSON(n) }

type RefType int

const (
	RefInvalid RefType = iota
	RefTop
	RefAll
)

func (r RefType) String() string {
	switch r {
	case RefTop:
		return "top"
	case RefAll:
		return "all"
	}
	return "???"
}

type ValueNode struct {
	Token Token `json:"-"`
	Value string
}

func (n ValueNode) At() Position   { return n.Token.At }
func (n ValueNode) String() string { return nodeStringJSON(n) }

func nodeStringJSON(n NodeAST) string {
	b, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(b)
}

type WhileNode struct {
	Pos  Position `json:"-"`
	Expr NodeAST
	Body []NodeAST
}

func (n WhileNode) At() Position   { return n.Pos }
func (n WhileNode) String() string { return nodeStringJSON(n) }

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

func (n FnNode) MarshalJSON() ([]byte, error) {
	type Alias FnNode
	return json.Marshal(&struct {
		Node string
		Alias
	}{
		Node:  "Fn",
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
