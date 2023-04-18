package calc

import (
	"fmt"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/coll"
	"github.com/blackchip-org/zc/scanner"
	"github.com/blackchip-org/zc/types"
)

type frame struct {
	pos      scanner.Pos
	funcName string
	env      zc.Env
}

func (f frame) Pos() scanner.Pos {
	return f.pos
}

func (f frame) FuncName() string {
	return f.funcName
}

func funcDecl(name string, params []types.Type) string {
	paramTypes := coll.Map(params, func(t types.Type) string { return t.String() })
	return fmt.Sprintf("%v(%v)", name, paramTypes)
}

type calc struct {
	lib     zc.Library
	env     *env
	modules map[string]*env
	genOps  map[string]zc.CalcFunc
	frames  coll.Deque[frame]
	trace   bool
	info    string
}

func New() (zc.Calc, error) {
	c := &calc{
		lib:    zlib,
		genOps: make(map[string]zc.CalcFunc),
		frames: coll.NewDequeSlice[frame](),
	}
	c.env = newEnv(c, zc.ProgName)
	return c, nil
}

func MustNew() zc.Calc {
	c, err := New()
	if err != nil {
		panic(err)
	}
	return c
}

func (c *calc) RegisterGenOp(name string, fn zc.CalcFunc, params ...types.Type) {
	c.genOps[funcDecl(name, params)] = fn
}

func (c *calc) GenOp(name string, args []types.Value) (zc.CalcFunc, error) {
	argTypes := coll.Map(args, func(v types.Value) types.Type { return v.Type() })
	decl := funcDecl(name, argTypes)
	fn, ok := c.genOps[decl]
	if !ok {
		return nil, fmt.Errorf("no operation for %v", decl)
	}
	return fn, nil
}

func (c *calc) Eval(name string, src []byte) error {
	return c.env.Eval(name, src)
}

func (c *calc) Trace() bool {
	return c.trace
}

func (c *calc) SetTrace(t bool) {
	c.trace = t
}

func (c *calc) Info() string {
	return c.info
}

func (c *calc) SetInfo(format string, args ...string) {
	c.info = fmt.Sprintf(format, args)
}

func (c *calc) Stack() coll.Deque[string] {
	return c.env.stack
}

func (c *calc) SetStack(items []string) {
	c.env.stack.Clear()
	coll.Push(c.env.stack, items...)
}
