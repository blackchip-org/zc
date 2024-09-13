package funcs

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"strings"

	"github.com/blackchip-org/zc/v6"
)

func evalHash(c zc.Calc, h hash.Hash) {
	d := zc.Data.Pop(c)
	h.Write(d.Bytes())
	d.Reset()
	d.Write(h.Sum(nil))
	zc.Data.Push(c, d)
}

func Md5(c zc.Calc) {
	evalHash(c, md5.New())
}

func Sha1(c zc.Calc) {
	evalHash(c, sha1.New())
}

func Sha1Hmac(c zc.Calc) {
	key := zc.Data.Pop(c)
	evalHash(c, hmac.New(sha1.New, key.Bytes()))
	zc.Data.Recycle(key)
}

func Sha224(c zc.Calc) {
	evalHash(c, sha256.New224())
}

func Sha224Hmac(c zc.Calc) {
	key := zc.Data.Pop(c)
	evalHash(c, hmac.New(sha256.New224, key.Bytes()))
	zc.Data.Recycle(key)
}

func Sha256(c zc.Calc) {
	evalHash(c, sha256.New())
}

func Sha256Hmac(c zc.Calc) {
	key := zc.Data.Pop(c)
	evalHash(c, hmac.New(sha256.New, key.Bytes()))
	zc.Data.Recycle(key)
}

func Sha384(c zc.Calc) {
	evalHash(c, sha512.New384())
}

func Sha384Hmac(c zc.Calc) {
	key := zc.Data.Pop(c)
	evalHash(c, hmac.New(sha512.New384, key.Bytes()))
	zc.Data.Recycle(key)
}

func Sha512(c zc.Calc) {
	evalHash(c, sha512.New())
}

func Sha512Hmac(c zc.Calc) {
	key := zc.Data.Pop(c)
	evalHash(c, hmac.New(sha512.New, key.Bytes()))
	zc.Data.Recycle(key)
}

func Rotate13(c zc.Calc) {
	a0 := zc.String.Pop(c)
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
	zc.String.Push(c, r0.String())
}
