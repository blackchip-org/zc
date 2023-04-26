package ops

import (
	"strings"

	"github.com/blackchip-org/zc/pkg/scanner"
	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper	rotate-13
func	Rot13 p0:Str -- Str
alias	rot13
title	Rotate characters by 13

desc
Rotate all characters in string *p0* by 13.
end

example
'Behind the tree! -- Behind the tree!
rot13             -- Oruvaq gur gerr!
rot13             -- Behind the tree!
end
*/
func Rot13(c zc.Calc) {
	a0 := zc.PopString(c)
	var r0 strings.Builder
	for _, ch := range a0 {
		var lower, upper rune
		if scanner.IsLowerCharAZ(ch) {
			lower, upper = 'a', 'z'
		}
		if scanner.IsUpperCharAZ(ch) {
			lower, upper = 'A', 'Z'
		}
		if lower != 0 {
			ch += 13
			if ch > upper {
				ch = lower + (ch - upper) - 1
			}
		}
		r0.WriteRune(ch)
	}
	zc.PushString(c, r0.String())
}
