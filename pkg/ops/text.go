package ops

import (
	"bytes"
	"math/big"
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/zc/pkg/zc"
)

func CharToCodePoint(c zc.Calc) {
	a0 := zc.PopRune(c)
	r0 := int32(a0)
	zc.PushInt32(c, r0)
}

func CodePointToChar(c zc.Calc) {
	a0 := zc.PopInt32(c)
	r0 := rune(a0)
	zc.PushRune(c, r0)
}

func Is(c zc.Calc) {
	a1 := zc.PopString(c)
	a0 := zc.PopString(c)
	r0 := a0 == a1
	zc.PushBool(c, r0)
}

func Join(c zc.Calc) {
	sep := zc.PopString(c)
	r0 := strings.Join(c.Stack(), sep)
	c.SetStack([]string{r0})
}

func Left(c zc.Calc) {
	i := zc.PopInt(c)
	s := zc.PopString(c)

	if i > len(s) || i < -len(s) {
		zc.ErrInvalidArgs(c)
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
	zc.PushString(c, r0)
}

func Len(c zc.Calc) {
	a0 := zc.PopString(c)
	r0 := len(a0)
	zc.PushInt(c, r0)
}

func Lower(c zc.Calc) {
	a0 := zc.PopString(c)
	r0 := strings.ToLower(a0)
	zc.PushString(c, r0)
}

func Right(c zc.Calc) {
	i := zc.PopInt(c)
	s := zc.PopString(c)

	if i > len(s) || i < -len(s) {
		zc.ErrInvalidArgs(c)
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
	zc.PushString(c, r0)
}

func Split(c zc.Calc) {
	sep := zc.PopString(c)
	s := zc.PopString(c)

	rs := strings.Split(s, sep)
	for _, r := range rs {
		zc.PushString(c, r)
	}
}

func Upper(c zc.Calc) {
	a0 := zc.PopString(c)
	r0 := strings.ToUpper(a0)
	zc.PushString(c, r0)
}

func UTF8Decode(c zc.Calc) {
	a0 := zc.PopBigInt(c)

	var r0 strings.Builder
	bytes := a0.Bytes()
	pos := 0
	for pos < len(bytes) {
		r, size := utf8.DecodeRune(bytes[pos:])
		r0.WriteRune(r)
		pos += size
	}

	zc.PushString(c, r0.String())
}

func UTF8Encode(c zc.Calc) {
	a0 := zc.PopString(c)

	var t0 bytes.Buffer
	codePoint := make([]byte, 4)
	for _, r := range a0 {
		size := utf8.EncodeRune(codePoint, r)
		for i := 0; i < size; i++ {
			t0.WriteByte(codePoint[i])
		}
	}
	var r0 big.Int
	r0.SetBytes(t0.Bytes())
	zc.PushBigInt(c, &r0)
}
