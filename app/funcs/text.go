package funcs

import (
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/msg"
)

func Concat(c zc.Calc) {
	y := zc.String.Pop(c)
	x := zc.String.Pop(c)
	zc.String.Push(c, x+y)
}

func CodePointToText(c zc.Calc) {
	x := zc.Int32.Pop(c)
	zc.String.Push(c, string(x))
}

func Join(c zc.Calc) {
	sep := zc.String.Pop(c)
	y := zc.String.Pop(c)
	x := zc.String.Pop(c)
	zc.String.Push(c, x+sep+y)
}

func Left(c zc.Calc) {
	i := zc.Int.Pop(c)
	s := zc.String.Pop(c)

	if i > len(s) || i < -len(s) {
		c.Raise(msg.ErrIndexOutOfRange(i))
		return
	}

	var r0 string
	switch {
	case i > 0:
		r0 = s[:i]
	case i < 0:
		r0 = s[:len(s)+i]
	default:
		r0 = s
	}
	zc.String.Push(c, r0)
}

func Len(c zc.Calc) {
	x := zc.String.Pop(c)
	zc.Int.Push(c, utf8.RuneCountInString(x))
}

func Lower(c zc.Calc) {
	x := zc.String.Pop(c)
	zc.String.Push(c, strings.ToLower(x))
}

func Right(c zc.Calc) {
	i := zc.Int.Pop(c)
	s := zc.String.Pop(c)

	if i > len(s) || i < -len(s) {
		c.Raise(msg.ErrIndexOutOfRange(i))
		return
	}

	var r0 string
	switch {
	case i > 0:
		r0 = s[len(s)-i:]
	case i < 0:
		r0 = s[-i:]
	default:
		r0 = s
	}
	zc.String.Push(c, r0)
}

func Split(c zc.Calc) {
	sep := zc.String.Pop(c)
	x := zc.String.Pop(c)
	rs := strings.Split(x, sep)
	for _, r := range rs {
		zc.String.Push(c, r)
	}
}

func TextToCodePoints(c zc.Calc) {
	i := 0
	x := zc.String.Pop(c)
	for i < len(x) {
		cp, w := utf8.DecodeRuneInString(x[i:])
		zc.Int32.Push(c, int32(cp))
		i += w
	}
}

func TextToUtf8(c zc.Calc) {
	x := zc.String.Pop(c)

	d := zc.Data.New()
	cp := make([]byte, 4)
	for _, r := range x {
		w := utf8.EncodeRune(cp, r)
		for i := 0; i < w; i++ {
			d.WriteByte(cp[i])
		}
	}
	zc.Data.Push(c, d)
}

func Upper(c zc.Calc) {
	x := zc.String.Pop(c)
	zc.String.Push(c, strings.ToUpper(x))
}

func Utf8ToText(c zc.Calc) {
	x := zc.Data.Pop(c)
	var r strings.Builder
	i := 0
	bytes := x.Bytes()

	for i < len(bytes) {
		cp, w := utf8.DecodeRune(bytes[i:])
		r.WriteRune(cp)
		i += w
	}
	zc.String.Push(c, r.String())
	zc.Data.Recycle(x)
}
