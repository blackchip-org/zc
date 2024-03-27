package ops

import (
	"bytes"
	"math/big"
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/pkg/zc"
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
title 	Length of text in characters

desc
Length of text in characters.
end

example
c 'abcd' len -- 4
c '🥇🥈🥉👏' len -- 4
end
*/
func Len(c zc.Calc) {
	a0 := zc.PopString(c)
	r0 := utf8.RuneCountInString(a0)
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
oper 	tone
func	Tone p0:Str tone:Int -- Str
title	Apply a skin tone to an emoji

desc
Apply a skin *tone*, an integer between 1 and 5 inclusive, to the first
character in *p0* which should be an emoji that supports a skin tone.
end

example
:waving-hand: -- 👋
2 tone -- 👋🏽
end
*/
func Tone(c zc.Calc) {
	tone := zc.PopInt(c)
	p0 := zc.PopString(c)

	if tone < 1 || tone > 5 {
		zc.ErrInvalidArgs(c, "for tone use [1-5]")
		return
	}

	var out bytes.Buffer
	runes := []rune(p0)
	out.WriteRune(runes[0])
	out.WriteRune(rune(0x1f3fb + tone))
	for i := 1; i < len(runes); i++ {
		out.WriteRune(runes[i])
	}

	r0 := out.String()
	zc.PushString(c, r0)
}

/*
oper	unescape
func	Unescape p0:Str -- Str
alias	unesc
title 	Unescape

desc
Unescapes characters in a string value that are prefixed by a backslash. The
escape sequences as defined for Go strings are used. A value such as
`\n` is converted to a new line. Values such as `\x7f`, `\u007f`, '\U0000007f`,
and `\077` are converted to the characters of that code point.
end

example
"\u65e5\u672c\u8a9e" -- \u65e5\u672c\u8a9e
unescape -- 日本語
c "\U000065e5\U0000672c\U00008a9e" -- \U000065e5\U0000672c\U00008a9e
unescape -- 日本語
c "\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e" -- \xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e
unescape -- 日本語
end
*/
func Unescape(c zc.Calc) {
	a0 := zc.PopString(c)
	s := scan.NewScannerFromString("", a0)
	rule := scan.NewRuleSet(
		scan.NewCharEncRule(
			scan.AlertEnc,
			scan.BackspaceEnc,
			scan.FormFeedEnc,
			scan.LineFeedEnc,
			scan.CarriageReturnEnc,
			scan.HorizontalTabEnc,
			scan.VerticalTabEnc,
		),
		scan.Hex2EncRule.AsByte(true),
		scan.Hex4EncRule,
		scan.Hex8EncRule,
		scan.OctEnc)
	for s.HasMore() {
		if s.This == '\\' {
			s.Skip()
			if !rule.Eval(s) {
				s.Keep()
			}
		} else {
			s.Keep()
		}
	}
	tok := s.Emit()
	if tok.Type == scan.IllegalType {
		c.SetError(tok.Errs[0])
	} else {
		zc.PushString(c, tok.Val)
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
