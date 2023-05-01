package ops

import (
	"bytes"
	"math/big"
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/zc/pkg/zc"
)

/*
oper	char-codepoint
func	CharToCodePoint p0:Char -- Int32
alias	char-cp
title	Character to code point

desc
Converts the character *p0* into an integer code point.
end

example
[°] -- °
char-cp -- 176
hex -- 0xb0
end
*/
func CharToCodePoint(c zc.Calc) {
	a0 := zc.PopRune(c)
	r0 := int32(a0)
	zc.PushInt32(c, r0)
}

/*
oper	codepoint-char
func	CodePointToChar p0:Int32 -- Char
alias	cp-char
title	Code point to character

desc
Coverts the code point *p0* to a character.
end

example
0xb0 -- 0xb0
cp-char -- °
end
*/
func CodePointToChar(c zc.Calc) {
	a0 := zc.PopInt32(c)
	r0 := rune(a0)
	zc.PushRune(c, r0)
}

/*
oper	join
func	Join Val* sep:Str -- Str
title	Join stack elements

desc
Join all stack elements into a single string separated by *sep*.
end

example
128 8 74 2 -- 128 | 8 | 74 | 2
'.' join -- 128.8.74.2
end
*/
func Join(c zc.Calc) {
	sep := zc.PopString(c)
	r0 := strings.Join(c.Stack(), sep)
	c.SetStack([]string{r0})
}

/*
oper	left
func	Left p0:Str n:Int -- Str
title 	Left substring

desc
Substring of *s* from the left.

If *n* is positive, *m* characters are taken from the left. If *n* is negative,
characters are taken from the left until there are *n* characters remaining. If
*n* is zero, *s* is returned without change.

If the absolute value of *n* is greater then then length of *s*, an
'illegal arguments' error is raised.
end

example
'abcdef -- abcdef
4 left -- abcd
-1 left -- abc
end
*/
func Left(c zc.Calc) {
	i := zc.PopInt(c)
	s := zc.PopString(c)

	if i > len(s) || i < -len(s) {
		zc.ErrInvalidArgs(c, "index out of range")
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

/*
oper	len
func	Len p0:Str -- Int
title 	Length of string

desc
Length of string in bytes.
end

example
'abcd -- abcd
len -- 4
end
*/
func Len(c zc.Calc) {
	a0 := zc.PopString(c)
	r0 := len(a0)
	zc.PushInt(c, r0)
}

/*
oper	lower
func	Lower p0:Str -- Str
title	Lowercase

desc
Converts the string *p0* to lowercase.
end

example
'AbCd -- AbCd
lower -- abcd
end
*/
func Lower(c zc.Calc) {
	a0 := zc.PopString(c)
	r0 := strings.ToLower(a0)
	zc.PushString(c, r0)
}

/*
oper	right
func 	Right p0:Str n:Int -- Str
title	Right substring

desc
Substring of *p0* from the right.

If *n* is positive, *n* characters are taken from the right. If *n* is
negative, characters are taken from the right until there are *n* characters
remaining. If *n* is zero, *s* is returned without change.

If the absolute value of *n* is greater then then length of *s*, an
'illegal arguments' error is raised.
end

example
'abcdef -- abcdef
4 right -- cdef
-1 right -- def
end
*/
func Right(c zc.Calc) {
	i := zc.PopInt(c)
	s := zc.PopString(c)

	if i > len(s) || i < -len(s) {
		zc.ErrInvalidArgs(c, "index out of range")
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

/*
oper	split
func	Split p0:Str sep:Str -- Str*
title	Split string

desc
Split *p0* into multiple strings that are separated by *sep*.
end

example
128.8.74.2 -- 128.8.74.2
'.' split -- 128 | 8 | 74 | 2
end
*/
func Split(c zc.Calc) {
	sep := zc.PopString(c)
	s := zc.PopString(c)

	rs := strings.Split(s, sep)
	for _, r := range rs {
		zc.PushString(c, r)
	}
}

/*
oper	upper
func	Upper p0:Str -- Str
title	Uppercase

desc
Convert string *p0* to uppercase.
end

example
'AbCd -- AbCd
upper -- ABCD
end
*/
func Upper(c zc.Calc) {
	a0 := zc.PopString(c)
	r0 := strings.ToUpper(a0)
	zc.PushString(c, r0)
}

/*
oper	utf8-decode
func	UTF8Decode p0:BigInt -- Str
alias	u8de
title	Decode UTF-8 bytes

desc
Decode UTF-8 bytes in *p0* into to a string.
end

example
0x3534c2b0 -- 0x3534c2b0
utf8-decode -- 54°
end
*/
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

/*
oper	utf8-encode
func	UTF8Encode p0:Str -- BigInt
alias	u8en
title	Encode UTF-8 bytes

desc
Encode a string into UTF-8 bytes.
end

example
54° -- 54°
utf8-encode hex -- 0x3534c2b0
end
*/
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
