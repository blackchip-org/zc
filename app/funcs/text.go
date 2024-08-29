package funcs

import (
	"unicode/utf8"

	"github.com/blackchip-org/zc/v6"
)

func Len(c zc.Calc) {
	x := zc.String.Pop(c)
	zc.Int.Push(c, utf8.RuneCountInString(x))
}
