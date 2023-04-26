package ops

import "github.com/blackchip-org/zc/pkg/zc"

/*
oper	clear
func	Clear ... --
alias	c
title	Clear the stack

desc
Remove all items from the stack.
end

example
1 -- 1
1 2 -- 1 | 2
clear --
end
*/
func Clear(c zc.Calc) {
	c.SetStack([]string{})
}

/*
oper	down
func	Down ... -- ...
alias	dn
title	Rotate stack downward

desc
Rotate items on the stack by moving downward.

In the interactive calculator, the top of the stack is towards the bottom of
the terminal so downward means seeing all items moves toward the bottom. The
top of the stack wraps around to be the bottom of the stack.
end

example
1 2 3 -- 1 | 2 | 3
down -- 3 | 1 | 2
down -- 2 | 3 | 1
end
*/
func Down(c zc.Calc) {
	a0, ok := c.Pop()
	if ok {
		c.SetStack(append([]string{a0}, c.Stack()...))
	}
}

/*
oper	drop
func	Drop Val --
title	Drop top item from stack

desc
Remove the top item from the stack.
end

example
1 -- 1
2 -- 1 | 2
drop -- 1
end
*/
func Drop(c zc.Calc) {
	c.Pop()
}

/*
oper	dup
func	Dup p0:Val -- p0:Val p0:Val
title	Duplicate top stack item

desc
Duplicate the top item on the stack.
end

example
1 -- 1
dup -- 1 | 1
end
*/
func Dup(c zc.Calc) {
	a0 := zc.PopString(c)
	zc.PushString(c, a0)
	zc.PushString(c, a0)
}

/*
oper	n
func	N -- Int
title	Number of stack items

desc
Number of items on the stack.
end

example
1 1 1 1 -- 1 | 1 | 1 | 1
n -- 1 | 1 | 1 | 1 | 4
end
*/
func N(c zc.Calc) {
	r0 := len(c.Stack())
	zc.PushInt(c, r0)
}

/*
oper	reverse
func	Reverse ... -- ...
alias	rev
title	Reverse stack

desc
Reverses the elements on the stack.
end

example
1 2 3 4 5 -- 1 | 2 | 3 | 4 | 5
reverse -- 5 | 4 | 3 | 2 | 1
end
*/
func Reverse(c zc.Calc) {
	var rs []string
	for c.StackLen() > 0 {
		a := c.MustPop()
		rs = append(rs, a)
	}
	c.SetStack(rs)
}

/*
oper	swap
func	Swap p0:Val p1:Val -- p1:Val p0:Val
title	Swap top two items on the stack

desc
Swap the first two items on the stack.
end

example
1 2 -- 1 | 2
swap -- 2 | 1
swap -- 1 | 2
end
*/
func Swap(c zc.Calc) {
	a1 := zc.PopString(c)
	a0 := zc.PopString(c)
	zc.PushString(c, a1)
	zc.PushString(c, a0)
}

/*
oper	take
func	Take ... n:Int -- ....
title	Take elements from the stack

desc
Take the top *n* elements from the stack and discard the rest.
end

example
1 2 3 4 5 -- 1 | 2 | 3 | 4 | 5
2 take -- 4 | 5
end
*/
func Take(c zc.Calc) {
	var rs []string
	n := zc.PopInt(c)
	for i := 0; i < n; i++ {
		a, ok := c.Pop()
		if !ok {
			break
		}
		rs = append([]string{a}, rs...)
	}
	c.SetStack(rs)
}

/*
oper	top
func	Top ... -- Val
macro	1 take
title	Take top element from the stack

desc
Keep the top of the stack and discard the rest.
end

example
1 2 3 4 5 -- 1 | 2 | 3 | 4 | 5
top -- 5
end
*/

/*
oper	up
func	Up ... -- ...
title 	Rotate items upward

desc
Rotate items on the stack by moving upward.

In the interactive calculator, the top of the stack is towards the bottom of
the terminal so upwards means seeing all items move toward the top. The
bottom of the stack wraps around to be the top of the stack.
end

example
1 2 3 -- 1 | 2 | 3
up -- 2 | 3 | 1
up -- 3 | 1 | 2
end
*/
func Up(c zc.Calc) {
	s := c.Stack()
	c.SetStack(append(s[1:], s[0]))
}
