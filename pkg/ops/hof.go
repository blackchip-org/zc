package ops

import "github.com/blackchip-org/zc/pkg/zc"

/*
oper	eval
func	Eval Val* expr:Str -- Val*
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
func	Repeat Val* expr:Str n:Int -- Val*
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
