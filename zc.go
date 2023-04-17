package zc

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/zc/coll"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/scanner"
	"github.com/blackchip-org/zc/types"
)

const (
	ProgName  = "zc"
	MainStack = "main"
)

type FuncDecl struct {
	Name   string
	Params []types.Type
	Func   CalcFunc
}

type ModuleDef struct {
	Name    string
	Script  string
	Natives map[string]FuncDecl
	GenOps  map[string]FuncDecl
	Init    func(Env)
}

func NewModuleDef(name string, script string) *ModuleDef {
	return &ModuleDef{Name: name, Script: script}
}

func (d *ModuleDef) Native(name string, fn CalcFunc, params ...types.Type) *ModuleDef {
	d.Natives[name] = FuncDecl{
		Name:   name,
		Func:   fn,
		Params: params,
	}
	return d
}

type Calc interface {
	Eval(string, []byte) error
	Trace() bool
	SetTrace(bool)
	Info() string
	SetInfo(string, ...string)
	Stack() coll.Deque[string]
	SetStack([]string)
}

func EvalString(c Calc, name string, source string) error {
	return c.Eval(name, []byte(source))
}

type Env interface {
	Calc() Calc
	Eval(string, []byte) error
	Stack() coll.Deque[string]
	UseStack(string) bool
}

type CalcFunc func(Env, []types.Value) ([]types.Value, error)

type Library interface {
	Load(string) ([]byte, error)
	Define(ModuleDef) error
	Module(name string) (ModuleDef, bool)
	Parse(name string) (*ast.File, bool)
}

type Frame interface {
	Pos() scanner.Pos
	FuncName() string
}

type CalcError struct {
	Message string
	Frames  []Frame
}

func (c CalcError) Error() string {
	return c.Message
}

func ErrorWithStack(err error) string {
	if calcErr, ok := err.(CalcError); ok {
		var buf strings.Builder
		for _, f := range calcErr.Frames {
			buf.WriteString(fmt.Sprintf("[%v] %v", f.Pos(), f.FuncName()))
			buf.WriteRune('\n')
		}
		buf.WriteString(calcErr.Error())
		buf.WriteRune('\n')
		return buf.String()
	}
	return err.Error()
}

func GenOpsDecl(params ...types.Type) FuncDecl {
	return FuncDecl{Params: params}
}
