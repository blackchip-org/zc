package app

import (
	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/pkg/coll"
)

type Calc struct {
	Catalog  *zc.Catalog
	Notice   string
	Err      error
	Listener zc.Listener
	stack    coll.Stack[zc.Item]
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
	c.stack.Push(item)
	if c.Listener != nil {
		c.Listener(zc.NewStackEvent(c, "push"))
	}
}

func (c *Calc) Pop() zc.Item {
	item := c.stack.Pop()
	if c.Listener != nil {
		c.Listener(zc.NewStackEvent(c, "pop"))
	}
	return item
}

func (c *Calc) Stack() []zc.Item {
	return c.stack.Items()
}

func (c *Calc) SetStack(items []zc.Item) {
	c.stack.SetItems(items)
}

func (c *Calc) Len() int {
	return c.stack.Len()
}

func (c *Calc) String() string {
	return zc.FormatList(zc.FormatItems(c.Stack()))
}

func (c *Calc) Var(name string) (any, bool) {
	s, ok := c.state[name]
	return s, ok
}

func (c *Calc) NewVar(name string, a any) {
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

func (c *Calc) Label() string {
	if c.stack.Len() == 0 {
		panic(zc.ErrStackEmpty)
	}
	return c.stack.Get(0).Label
}

func (c *Calc) Unit() string {
	if c.stack.Len() == 0 {
		panic(zc.ErrStackEmpty)
	}
	return c.stack.Get(0).Unit
}

func (c *Calc) SetLabel(label string) {
	if c.stack.Len() == 0 {
		panic(zc.ErrStackEmpty)
	}
	item := c.stack.Get(0)
	item.Label = label
	c.stack.Set(0, item)
}

func (c *Calc) SetUnit(unit string) {
	if c.stack.Len() == 0 {
		panic(zc.ErrStackEmpty)
	}
	item := c.stack.Get(0)
	item.Unit = unit
	c.stack.Set(0, item)
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
	fn, ok, err := c.ResolveOp(op)
	switch {
	case err != nil:
		c.Raise(zc.ErrOp(name, err))
		return
	case !ok:
		c.Raise(zc.ErrOp(name, c.errTypeMismatch(op)))
		return
	}

	if c.Listener != nil {
		c.Listener(zc.NewOpEvent(name))
	}
	fn.Eval(c)

	if c.Err == nil {
		ok, err := c.isFuncMatch(fn.Returns, fn.VarReturn)
		switch {
		case err != nil:
			c.Raise(zc.ErrOp(name, err))
			return
		case !ok:
			panic("return mismatch: " + name)
		}
	} else {
		c.Err = zc.ErrOp(name, c.Err)
	}
}

func (c *Calc) ResolveOp(op zc.Op) (zc.Func, bool, error) {
	for _, fn := range op.Funcs {
		ok, err := c.isFuncMatch(fn.Params, fn.VarParam)
		switch {
		case err != nil:
			return zc.Func{}, false, err
		case ok:
			c.convert(fn.Params, fn.VarParam)
			return fn, true, nil
		}
	}
	return zc.Func{}, false, nil
}

func (c *Calc) isFuncMatch(params []zc.Type, varParam zc.Type) (bool, error) {
	if c.Len() < len(params) {
		return false, nil
	}

	for i, param := range params {
		ok, err := c.isParamMatch(i, param)
		switch {
		case err != nil:
			return false, err
		case !ok:
			return false, nil
		}
	}
	if varParam != nil {
		for i := len(params); i < c.stack.Len(); i++ {
			ok, err := c.isParamMatch(i, varParam)
			switch {
			case err != nil:
				return false, err
			case !ok:
				return false, nil
			}
		}
	}
	return true, nil
}

func (c *Calc) isParamMatch(index int, param zc.Type) (bool, error) {
	arg := c.stack.Get(index)
	if arg.Type == param {
		return true, nil
	}
	_, ok, err := param.Parse(c, arg.Val())
	return ok, err
}

func (c *Calc) convertArg(index int, param zc.Type) {
	arg := c.stack.Get(index)
	if arg.Type == param {
		return
	}
	conv, ok, err := param.Parse(c, arg.Val())
	if !ok || err != nil {
		panic("unexpected: should be correct type")
	}
	c.stack.Set(index, zc.Item{TypeVal: conv, Type: param})
}

func (c *Calc) convert(params []zc.Type, varParam zc.Type) {
	for i, param := range params {
		c.convertArg(i, param)
	}
	if varParam != nil {
		for i := len(params); i < c.stack.Len(); i++ {
			c.convertArg(i, varParam)
		}
	}
}

func (c *Calc) errTypeMismatch(op zc.Op) error {
	for i := 0; i < c.stack.Len(); i++ {
		arg := c.stack.Get(i)
		good := false
		for _, fn := range op.Funcs {
			var param zc.Type
			if i >= len(fn.Params) {
				param = fn.VarParam
			} else {
				param = fn.Params[i]
			}
			if param == nil {
				continue
			}
			if ok, _ := c.isParamMatch(i, param); ok {
				good = true
				break
			}
		}
		if !good {
			return zc.ErrUnexpectedType(arg.Val())
		}
	}
	panic("unexpected")
}
