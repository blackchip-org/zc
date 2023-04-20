package ops

import (
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/pkg/scanner"
)

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
