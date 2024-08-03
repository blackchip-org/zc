package app

import (
	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6"
)

type Calc struct {
	Catalog  *zc.Catalog
	Notice   string
	Err      error
	Listener zc.Listener
	items    []zc.Item
	pos      int
	state    map[string]any
}

func NewCalc() *Calc {
	c := &Calc{
		state:   make(map[string]any),
		Catalog: mainCatalog,
	}
	return c
}

func (c *Calc) Push(item zc.Item) {
	if c.pos < len(c.items) {
		c.items[c.pos] = item
	} else {
		c.items = append(c.items, item)
	}
	if c.Listener != nil {
		c.Listener(zc.NewStackEvent(c, "push"))
	}
	c.pos++
}

func (c *Calc) Pop() zc.Item {
	if c.pos == 0 {
		panic(zc.ErrStackEmpty)
	}
	c.pos--
	if c.Listener != nil {
		c.Listener(zc.NewStackEvent(c, "pop"))
	}
	return c.items[c.pos]
}

func (c *Calc) Items() []zc.Item {
	return c.items[:c.pos]
}

func (c *Calc) SetItems(items []zc.Item) {
	c.items = items
	c.pos = len(c.items)
}

func (c *Calc) Len() int {
	return c.pos
}

func (c *Calc) String() string {
	return zc.FormatList(zc.FormatItems(c.Items()))
}

func (c *Calc) State(name string) (any, bool) {
	s, ok := c.state[name]
	return s, ok
}

func (c *Calc) NewState(name string, a any) {
	c.state[name] = a
}

func (c *Calc) Notify(msg string) {
	c.Notice = msg
}

func (c *Calc) Raise(err error) {
	if c.Err == nil {
		c.Err = err
	}
}

func (c *Calc) Label(label string) {
	item := c.Pop()
	item.Label = label
	c.Push(item)
}

func (c *Calc) Unit(unit string) {
	item := c.Pop()
	item.Unit = unit
	c.Push(item)
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
	c.Push(zc.Item{TypeVal: val, Type: zc.String})
}

func (c *Calc) evalName(name string) {
	op, ok := c.Catalog.OpFor(name)
	if !ok {
		c.Raise(zc.ErrNoSuchOp(name))
		return
	}
	if len(op.Macro) > 0 {
		c.EvalToken(op.Macro...)
		return
	}
	fn, ok := c.ResolveOp(op)
	if !ok {
		c.Raise(zc.ErrArgMismatch(name))
		return
	}

	if c.Listener != nil {
		c.Listener(zc.NewOpEvent(name))
	}
	fn.Eval(c)

	if c.Err == nil {
		if !c.isTypeMatch(fn.Returns, fn.VarReturn) {
			panic("return mismatch: " + name)
		}
	} else {
		c.Err = zc.ErrOp(name, c.Err)
	}
}

func (c *Calc) ResolveOp(op zc.Op) (zc.Func, bool) {
	for _, fn := range op.Funcs {
		if c.isTypeMatch(fn.Params, fn.VarParam) {
			c.convert(fn.Params, fn.VarParam)
			return fn, true
		}
	}
	return zc.Func{}, false
}

func (c *Calc) isTypeMatch(params []zc.Type, varParam zc.Type) bool {
	if c.Len() < len(params) {
		return false
	}

	for i, param := range params {
		if !c.isParamMatch(i, param) {
			return false
		}
	}
	if varParam != nil {
		for i := len(params); i < c.pos; i++ {
			if !c.isParamMatch(i, varParam) {
				return false
			}
		}
	}
	return true
}

func (c *Calc) isParamMatch(index int, param zc.Type) bool {
	arg := c.items[c.pos-index-1]
	if arg.Type == param {
		return true
	}
	_, ok := param.Parse(arg.Val())
	return ok
}

func (c *Calc) convertArg(index int, param zc.Type) {
	arg := c.items[c.pos-index-1]
	if arg.Type == param {
		return
	}
	conv, ok := param.Parse(arg.Val())
	if !ok {
		panic("unexpected: should be correct type")
	}
	c.items[c.pos-index-1] = zc.Item{TypeVal: conv, Type: param}
}

func (c *Calc) convert(params []zc.Type, varParam zc.Type) {
	for i, param := range params {
		c.convertArg(i, param)
	}
	if varParam != nil {
		for i := len(params); i < c.pos; i++ {
			c.convertArg(i, varParam)
		}
	}
}
