package ops

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"strings"

	"github.com/blackchip-org/zc/v5/pkg/scanner"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

func evalHash(c zc.Calc, h hash.Hash) {
	var r0 strings.Builder

	a0 := zc.PopString(c)
	h.Write([]byte(a0))
	h0 := h.Sum(nil)
	r0.WriteString("0x")
	for _, b := range h0 {
		r0.WriteString(fmt.Sprintf("%02x", b))
	}
	zc.PushString(c, r0.String())
}

/*
oper	md5
func	Md5 p0:Str -- Str
title	Hash *p0* with the MD5 function

desc
Hash *p0* with the MD5 function.
end

example
'Behind the tree! -- Behind the tree!
md5               -- 0xbbce0fb98a6a1b308e589d1408968ac2
end
*/
func Md5(c zc.Calc) {
	evalHash(c, md5.New())
}

/*
oper	sha1
func	Sha1 p0:Str -- Str
title	Hash *p0* with the SHA-1 function

desc
Hash *p0* with the SHA-1 function.
end

example
'Behind the tree! -- Behind the tree!
sha1              -- 0xda6e7f530a9b42a679944f0c9fc0f86ae5534450
end
*/
func Sha1(c zc.Calc) {
	evalHash(c, sha1.New())
}

/*
oper	sha224
func	Sha224 p0:Str -- Str
title	Hash *p0* with the SHA-224 function

desc
Hash *p0* with the SHA-224 function.
end

example
'Behind the tree! -- Behind the tree!
sha224            -- 0x15eac2f886e0e09a44ce08da58f3386b707885150d6142d1a6e7c608
end
*/
func Sha224(c zc.Calc) {
	evalHash(c, sha256.New224())
}

/*
oper	sha256
func	Sha256 p0:Str -- Str
title	Hash *p0* with the SHA-256 function

desc
Hash *p0* with the SHA-256 function.
end

example
'Behind the tree! -- Behind the tree!
sha256            -- 0x5e19fc5f8ec2aabccef7970385bb9151a421f398d048ced2b2c86757aafebfc3
end
*/
func Sha256(c zc.Calc) {
	evalHash(c, sha256.New())
}

/*
oper	sha384
func	Sha384 p0:Str -- Str
title	Hash *p0* with the SHA-384 function

desc
Hash *p0* with the SHA-384 function.
end

example
'Behind the tree! -- Behind the tree!
sha384            -- 0x54489c547782d201bb0c8c2c81e77e034695067c98087bd949d13de752dd3843323c7244c1d15776ad52093598420dca
end
*/
func Sha384(c zc.Calc) {
	evalHash(c, sha512.New384())
}

/*
oper	sha512
func	Sha512 p0:Str -- Str
title	Hash *p0* with the SHA-512 function

desc
Hash *p0* with the SHA-512 function.
end

example
'Behind the tree! -- Behind the tree!
sha512            -- 0x431777a80ed22c45b4fe0dc8c7e3a07d8d20df3b796a39068f2fc6f57cd69b6c60f4a6e3151189b97a1ad2fe5888c255e93f28c1e6c9b6f0241b10c34f8f9e86
end
*/
func Sha512(c zc.Calc) {
	evalHash(c, sha512.New())
}

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
