package ops

import (
	"errors"

	"github.com/blackchip-org/zc/pkg/zc"
)

type memory struct {
	stack [][]string
	vars  map[string][]string
}

func getMemoryState(c zc.Calc) *memory {
	mem, ok := c.State("memory")
	if !ok {
		mem = &memory{
			stack: make([][]string, 0),
			vars:  make(map[string][]string),
		}
		c.NewState("memory", mem)
	}
	return mem.(*memory)
}

/*
oper	clear
func	Clear Val* --
alias	c
title	Clear the stack

desc
Remove all items from the stack.
end

example
1 -- 1
2 -- 1 | 2
clear --
end
*/
func Clear(c zc.Calc) {
	c.SetStack([]string{})
}

/*
oper	clear-all
func	ClearAll Val* --
alias	ca
title 	Clear stack and memory

desc
Remove all items from the stack and any stacks stored in memory.
end

example
1 -- 1
2 -- 1 | 2
store -- *stored*
clear-all -- *cleared*
recall -- memory empty
end
*/
func ClearAll(c zc.Calc) {
	c.SetStack([]string{})
	s := getMemoryState(c)
	s.stack = [][]string{}
	clear(s.vars)
	c.SetInfo("cleared")
}

/*
oper	down
func	Down Val* -- Val*
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
oper	get
func	Get Val* name:Str -- Val* Val*
title	Get a named stack from memory
*/
func Get(c zc.Calc) {
	s := getMemoryState(c)
	name := zc.PopString(c)
	v, ok := s.vars[name]
	if !ok || len(v) == 0 {
		c.SetInfo("empty")
		return
	}
	v = append(c.Stack(), v...)
	c.SetStack(v)
}

/*
oper	size
func	Size Val* -- Val* n:Int
title	Size of the current stack
alias	n

desc
Number of items on the stack.
end

example
1 1 1 1 -- 1 | 1 | 1 | 1
n -- 1 | 1 | 1 | 1 | 4 # size
end
*/
func Size(c zc.Calc) {
	r0 := len(c.Stack())
	zc.PushInt(c, r0)
	zc.Annotate(c, "size")
}

/*
oper 	recall
func	Recall Val* -- Val* Val*
alias	re
title	Recall stack from memory

desc
Recall a stack from memory. The recalled stack is placed before an existing
items on the current stack. Memory is also a stack so that multiple stacks can
be stored at one time.
end

example
1 2 3 4 store average -- 2.5
recall -- 1 | 2 | 3 | 4 | 2.5
end
*/
func Recall(c zc.Calc) {
	s := getMemoryState(c)
	if len(s.stack) == 0 {
		c.SetError(errors.New("memory empty"))
		return
	}
	var v []string
	v, s.stack = s.stack[0], s.stack[1:]
	v = append(v, c.Stack()...)
	c.SetStack(v)
}

/*
oper	reverse
func	Reverse Val* -- Val*
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
oper	set
func	Set Val* name:Str --
title	Place a named stack to memory
*/
func Set(c zc.Calc) {
	s := getMemoryState(c)
	name := zc.PopString(c)
	s.vars[name] = c.Stack()
	c.SetInfo("set")
}

/*
oper 	store
func	Store Val* -- Val*
alias	st
title	Store stack to memory

desc
Store a copy of the current stack to memory for later recall. Memory is also
a stack so that multiple stacks can be stored at one time.
end

example
1 2 3 4 store average -- 2.5
recall -- 1 | 2 | 3 | 4 | 2.5
end
*/
func Store(c zc.Calc) {
	if c.StackLen() == 0 {
		c.SetError(errors.New("empty"))
		return
	}
	s := getMemoryState(c)
	s.stack = append([][]string{c.Stack()}, s.stack...)
	c.SetInfo("stored")
}

/*
oper	swap
func	Swap p0:Val p1:Val -- p1:Val p0:Val
alias	sw
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
func	Take Val* n:Int -- Val*
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
func	Top Val* -- Val
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
func	Up Val* -- Val*
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
