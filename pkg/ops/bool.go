package ops

import "github.com/blackchip-org/zc/pkg/zc"

/*
oper	and
func	AndBool p0:Bool p1:Bool -- Bool
alias 	a
alias 	+
title 	Logical conjunction

desc
The logical conjunction of *p0* and *p1*.
end

example
c true  true   and -- true
c true  false  and -- false
c false false  and -- false
end
*/
func AndBool(c zc.Calc) {
	a1 := zc.PopBool(c)
	a0 := zc.PopBool(c)
	r0 := a0 && a1
	zc.PushBool(c, r0)
}

/*
oper	false
func	- -- Bool
macro	[false]
title	False

desc
Places `false` on the stack
end

example
false -- false
end
*/

/*
oper	not
func	NotBool Bool -- Bool
title 	Negation

desc
If `true`, returns `false`, otherwise returns `true`.
end

example
true not -- false
not -- true
end
*/
func NotBool(c zc.Calc) {
	a0 := zc.PopBool(c)
	r0 := !a0
	zc.PushBool(c, r0)
}

/*
oper	or
func	OrBool p0:Bool p1:Bool -- Bool
title	Logical disjunction

desc
The logical disjunction of *p0* and *p1*.
end

example
c true  true  or -- true
c true  false or -- true
c false false or -- false
end
*/
func OrBool(c zc.Calc) {
	a1 := zc.PopBool(c)
	a0 := zc.PopBool(c)
	r0 := a0 || a1
	zc.PushBool(c, r0)
}

/*
oper	true
func	- -- Bool
macro	[true]
title	False

desc
Places `true` on the stack
end

example
true -- true
end
*/
