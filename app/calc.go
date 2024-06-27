package app

import (
	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app/state"
	"github.com/blackchip-org/zc/v6/app/vols"
)

var mainCatalog *zc.Catalog

func init() {
	mainCatalog = zc.NewCatalog()

	// Order here is important. BasicInt adds overloads for basic before
	// BasicDec does. Do not try to sort this list.
	mainCatalog.AddVolume(
		vols.Anno,
		vols.Basic,
		vols.BasicInt,
		vols.BasicFloat,
		vols.BasicRat,
		vols.BasicComplex,
		vols.BasicIntU,
		vols.BasicIntU8,
		vols.BasicIntU16,
		vols.BasicIntU32,
		vols.BasicIntU64,
		vols.Conf,
		vols.Format,
		vols.Prog,
		vols.Sci,
		vols.SciComplex,
		vols.SciFloat64,
		vols.Stack,
		vols.Types,
		vols.TypesIntU,
		vols.TypesIntU8,
		vols.TypesIntU16,
		vols.TypesIntU32,
		vols.TypesIntU64,
	)
}

type Calc struct {
	zc.Stack[zc.Item]
	Catalog  *zc.Catalog
	state    state.State
	Err      error
	Info     string
	Listener zc.Listener
}

func NewCalc() *Calc {
	return &Calc{
		Catalog: mainCatalog,
		state:   state.New(),
	}
}

func (c *Calc) Push(item zc.Item) {
	c.Stack.Push(item)
	if c.Listener != nil {
		c.Listener(zc.NewStackEvent("push", c.Stack))
	}
}

func (c *Calc) PushVal(vals ...any) {
	for _, val := range vals {
		c.Push(zc.Item{Val: val})
	}
}

func (c *Calc) Pop() zc.Item {
	item := c.Stack.Pop()
	if c.Listener != nil {
		c.Listener(zc.NewStackEvent("pop", c.Stack))
	}
	return item
}

func (c *Calc) Eval(line string) error {
	toks := zc.ScanWords(line)
	for _, tok := range toks {
		if c.Err != nil {
			return c.Err
		}
		c.EvalToken(tok)
	}
	return c.Err
}

func (c *Calc) EvalToken(toks ...scan.Token) {
	for _, tok := range toks {
		if c.Err != nil {
			return
		}
		switch tok.Type {
		case zc.TokenName:
			c.evalName(tok.Val)
		case zc.TokenValue:
			c.evalValue(tok.Val)
		default:
			panic("unknown token type: " + tok.Type)
		}
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
		c.Err = zc.ErrArgMismatch(name)
		return
	}
	c.Do(op)
}

func (c *Calc) LookupOp(name string) ([]zc.Op, bool) {
	ops, ok := c.Catalog.OpFor(name)
	return ops, ok
}

func (c *Calc) ResolveOp(ops []zc.Op) (zc.Op, bool) {
	var op zc.Op
	for _, op = range ops {
		if c.isTypeMatch(op.Params) {
			return op, true
		}
	}
	return op, false
}

func (c *Calc) isTypeMatch(def []zc.Type) bool {
	if c.Len() < len(def) {
		return false
	}
	for i, param := range def {
		arg := c.Get(len(def) - i - 1)
		_, _, ok := param.From(c.state, arg.Val)
		if !ok {
			return false
		}
	}
	return true
}

func (c *Calc) Do(op zc.Op) {
	if len(op.Macro) > 0 {
		for _, tok := range op.Macro {
			if c.Err != nil {
				return
			}
			c.EvalToken(tok)
		}
		return
	}

	nParams := len(op.Params)
	for i, param := range op.Params {
		arg := c.Get(nParams - i - 1)
		convVal, argType, ok := param.From(c.state, arg.Val)
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
		State: c.state,
	}
	if c.Listener != nil {
		c.Listener(zc.NewOpEvent(&env))
	}
	op.Func(&env)

	c.Err = env.Err
	c.Info = env.Info

	if c.Err == nil && !c.isTypeMatch(op.Returns) {
		c.Err = zc.ErrReturnMismatch(op.Name)
	}
}
