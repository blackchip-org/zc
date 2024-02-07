package ops

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"strings"

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
title	MD5 hash function

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
title	SHA-1 hash function

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
oper	sha1hmac
func	Sha1Hmac p0:Str key:Str -- Str
title	SHA-1 keyed hash function

desc
Hash *p0* with *key* using the SHA-1 function
end

example
'Behind the tree!' /swordfish -- Behind the tree! | swordfish
sha1hmac                      -- 0x75859001097e8ad6907a879d340170ef571c8e5c
end
*/
func Sha1Hmac(c zc.Calc) {
	key := zc.PopString(c)
	evalHash(c, hmac.New(sha1.New, []byte(key)))
}

/*
oper	sha224
func	Sha224 p0:Str -- Str
title	SHA-224 hash function

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
oper	sha224hmac
func	Sha224Hmac p0:Str key:Str -- Str
title	SHA-224 keyed hash function

desc
Hash *p0* with *key* using the SHA-224 function
end

example
'Behind the tree!' /swordfish -- Behind the tree! | swordfish
sha224hmac                    -- 0x30ecbacd9aedf890b796eb79032105382ce323f18835f9bbb3867c4a
end
*/
func Sha224Hmac(c zc.Calc) {
	key := zc.PopString(c)
	evalHash(c, hmac.New(sha256.New224, []byte(key)))
}

/*
oper	sha256
func	Sha256 p0:Str -- Str
title	SHA-256 hash function

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
oper	sha256hmac
func	Sha256Hmac p0:Str key:Str -- Str
title	SHA-256 keyed hash function

desc
Hash *p0* with *key* using the SHA-256 function
end

example
'Behind the tree!' /swordfish -- Behind the tree! | swordfish
sha256hmac                    -- 0x0155d56d7485e9db843792ff48b97a1a96ae0655a92ad7c2324d71ed0410b907
end
*/
func Sha256Hmac(c zc.Calc) {
	key := zc.PopString(c)
	evalHash(c, hmac.New(sha256.New, []byte(key)))
}

/*
oper	sha384
func	Sha384 p0:Str -- Str
title	SHA-384 hash function

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
oper	sha384hmac
func	Sha384Hmac p0:Str key:Str -- Str
title	SHA-384 keyed hash function

desc
Hash *p0* with *key* using the SHA-384 function
end

example
'Behind the tree!' /swordfish -- Behind the tree! | swordfish
sha384hmac                    -- 0x401d0de7dd5cdf1ed82002cc1a696ec8d3d636b0388b89cc6f33ec4b237382a16b976f4773c3b0a3979392289a0ceaf9
end
*/
func Sha384Hmac(c zc.Calc) {
	key := zc.PopString(c)
	evalHash(c, hmac.New(sha512.New384, []byte(key)))
}

/*
oper	sha512
func	Sha512 p0:Str -- Str
title   SHA-512 hash function

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
oper	sha512hmac
func	Sha512Hmac p0:Str key:Str -- Str
title	SHA-512 keyed hash function

desc
Hash *p0* with *key* using the SHA-512 function
end

example
'Behind the tree!' /swordfish -- Behind the tree! | swordfish
sha512hmac                    -- 0x8e4074cf371bc7e83cd508e9a3d4c0ecd9e014aa808f3234b05de26ddad9895ad6d375651b3ac8231806ccd0b9cd477f3df024c1ecd75032a0e827e7f2f98fcf
end
*/
func Sha512Hmac(c zc.Calc) {
	key := zc.PopString(c)
	evalHash(c, hmac.New(sha512.New, []byte(key)))
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
		if ch >= 'a' && ch <= 'z' {
			lower, upper = 'a', 'z'
		}
		if ch >= 'A' && ch <= 'Z' {
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
