package app

import (
	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/vols"
)

var mainCatalog *zc.Catalog

func init() {
	mainCatalog = zc.NewCatalog()
	mainCatalog.AddVolume(vols.BasicInt, vols.Stack)
}

type Calc struct {
	zc.Stack[zc.Item]
	cat  *zc.Catalog
	Err  error
	Info string
}

func NewCalc() *Calc {
	return &Calc{cat: mainCatalog}
}

func (c *Calc) PushVal(vals ...any) {
	for _, val := range vals {
		c.Push(zc.Item{Val: val})
	}
}

func (c *Calc) Eval(line string) {
	toks := zc.ScanWords(line)
	for _, tok := range toks {
		if c.Err != nil {
			return
		}
		c.EvalToken(tok)
	}
}

func (c *Calc) EvalToken(tok scan.Token) {
	switch tok.Type {
	case zc.TokenName:
		c.evalName(tok.Val)
	case zc.TokenValue:
		c.evalValue(tok.Val)
	default:
		panic("unknown token type: " + tok.Type)
	}
}

func (c *Calc) evalValue(val string) {
	c.PushVal(val)
}

func (c *Calc) evalName(name string) {
	ops, ok := c.LookupOp(name)
	if !ok {
		c.Err = zc.ErrNoSuchOp(name)
		return
	}
	op, ok := c.ResolveOp(ops)
	if !ok {
		c.Err = zc.ErrNoMatchForOp(name)
		return
	}
	c.Do(op)
}

func (c *Calc) LookupOp(name string) ([]zc.Op, bool) {
	ops, ok := c.cat.OpFor(name)
	return ops, ok
}

func (c *Calc) ResolveOp(ops []zc.Op) (zc.Op, bool) {
	var op zc.Op
	for _, op = range ops {
		if c.isParamMatch(op) {
			return op, true
		}
	}
	return op, false
}

func (c *Calc) isParamMatch(op zc.Op) bool {
	if c.Len() < len(op.Params) {
		return false
	}
	for i, param := range op.Params {
		arg := c.Get(len(op.Params) - i - 1)
		_, _, ok := param.From(arg.Val)
		if !ok {
			return false
		}
	}
	return true
}

func (c *Calc) Do(op zc.Op) {
	nParams := len(op.Params)
	for i, param := range op.Params {
		arg := c.Get(nParams - i - 1)
		convVal, argType, ok := param.From(arg.Val)
		if !ok {
			panic("cannot convert")
		}
		if argType != param {
			argType.Recycle(arg.Val)
			arg.Val = convVal
			c.Set(nParams-i-1, arg)
		}
	}

	env := zc.OpEnv{
		Op:    op,
		Stack: &c.Stack,
	}
	op.Func(&env)
	c.Err = env.Err
	c.Info = env.Info
}
