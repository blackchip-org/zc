package ops

import (
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

/*
oper	anno
func	Anno p0:Val anno:Str -- p0:Val
alias	annotate
title	Annotate value

desc
Annotate the value of *p0* with *anno*.
end

example
42 -- 42
'the answer' anno -- 42 # the answer
end
*/
func Anno(c zc.Calc) {
	anno := zc.PopString(c)
	p0 := zc.PopString(c)
	zc.PushString(c, p0)
	zc.Annotate(c, anno)
}

/*
oper	no-anno
func	NoAnno p0:Val -- p0:Val
alias	noa
alias	no-annotation
title 	Remove annotation

desc
Remove the annotation, if any, from *p0*.
end

example
42 -- 42
'the answer' anno -- 42 # the answer
noa -- 42
end
*/
func NoAnno(c zc.Calc) {
	p0 := zc.PopString(c)
	zc.PushString(c, p0)
}
