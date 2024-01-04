package ops

import (
	"slices"
	"strings"

	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper    apply
func    Apply args:Val* fn:Str nargs:Int -- Val*
title 	Apply a function using arguments on stack

desc
Evaluates the expression in *fn* by first popping *nargs* off the stack and
pushing them back as a single argument. This is useful for higher order
functions, like map, where some of the arguments are from existing results
found on the stack.
end

example
1 2 3 4 -- 1 | 2 | 3 | 4
n -- 1 | 2 | 3 | 4 | 4 # size
[swap sub] [map] 2 apply -- 3 | 2 | 1 | 0
end
*/
func Apply(c zc.Calc) {
	nArgs := zc.PopInt(c)
	fnName := zc.PopString(c)
	if c.StackLen() < nArgs {
		zc.ErrNotEnoughArgs(c, "apply", nArgs)
		return
	}
	var args []string
	for i := 0; i < nArgs; i++ {
		args = slices.Insert(args, 0, zc.PopString(c))
	}
	c.Push(strings.Join(args, " "))
	c.Eval(fnName)
}

/*
oper	eval
func	Eval expr:Str -- Val*
title 	Evaluate top of stack

desc
Evaluate *expr* as if it was input to the calculator.
end

example
[1 2 add -- 1 2 add
eval -- 3
end
*/
func Eval(c zc.Calc) {
	fn := zc.PopString(c)
	c.Eval(fn)
}

/*
oper	filter
func	Filter Val* expr:Str -- Val*
title	Filter items in the stack

desc
Filter the stack by keeping items that are true when evaluated by
expression *expr*.
end

example
1 2 3 4 5 6 -- 1 | 2 | 3 | 4 | 5 | 6
[2 mod 0 eq] filter -- 2 | 4 | 6
end
*/
func Filter(c zc.Calc) {
	var rs []string
	fn := zc.PopString(c)
	for _, v := range c.Stack() {
		dc := c.Derive()
		dc.Push(v)
		dc.Eval(fn)
		out, ok := dc.Pop()
		if !ok {
			zc.ErrInvalidFunc(c, fn, "no results")
			return
		}
		r, ok := zc.Bool.Parse(out)
		if !ok {
			zc.ErrExpectedType(c, zc.Bool, out)
			return
		}
		if r {
			rs = append(rs, v)
		}
	}
	c.SetStack(rs)
}

/*
oper	fold
func	Fold Val* expr:Str -- Val
alias	reduce
title	Reduce items to a single value

desc
Reduce the stack to a single value using the expression *expr*. An
'invalid function' error is raised if *expr* does not reduce.
end

example
1 2 3 4 5 -- 1 | 2 | 3 | 4 | 5
[add] fold -- 15
end
*/
func Fold(c zc.Calc) {
	fn := zc.PopString(c)
	for c.StackLen() > 1 {
		before := c.StackLen()
		c.Eval(fn)
		if c.Error() != nil {
			return
		}
		if c.StackLen() >= before {
			zc.ErrInvalidFunc(c, fn, "does not reduce")
			return
		}
	}
}

/*
oper	map
func	Map Val* expr:Str -- Val*
title	Apply a function to each item on the stack

desc
Apply expression *expr* to each value in the stack.
end

example
1 2 3 4 5 -- 1 | 2 | 3 | 4 | 5
[2 mul] map -- 2 | 4 | 6 | 8 | 10
end
*/
func Map(c zc.Calc) {
	var rs []string
	fn := zc.PopString(c)
	for _, a := range c.Stack() {
		dc := c.Derive()
		dc.Push(a)
		if err := dc.Eval(fn); err != nil {
			c.SetError(err)
			return
		}
		if r0, ok := dc.Pop(); ok {
			rs = append(rs, r0)
		}
	}
	c.SetStack(rs)
}

/*
oper	repeat
func	Repeat expr:Val n:Int -- Val*
title	Repeat the execution of a function

desc
Repeat execution of expression *expr* for *n* times.
end

example
1 -- 1
[2 mul] 8 repeat -- 256
end
*/
func Repeat(c zc.Calc) {
	n := zc.PopInt(c)
	fn := zc.PopString(c)
	for i := 0; i < n; i++ {
		c.Eval(fn)
	}
}
