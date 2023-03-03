package zlib

import (
	"bytes"
	"math/big"
	"strings"
	"unicode/utf8"

	"github.com/blackchip-org/zc"
)

func Decode(env *zc.Env) error {
	var z strings.Builder
	for _, item := range env.Stack.Items() {
		i, err := env.Calc.ParseInt32(item)
		if err != nil {
			return err
		}
		z.WriteRune(rune(i))
	}
	env.Stack.Clear()
	env.Stack.Push(z.String())
	return nil
}

func Encode(env *zc.Env) error {
	s, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	for _, r := range s {
		env.Stack.PushInt(int(r))
	}
	return nil
}

func UTF8Decode(env *zc.Env) error {
	i, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}

	var result strings.Builder
	bytes := i.Bytes()
	pos := 0
	for pos < len(bytes) {
		r, size := utf8.DecodeRune(bytes[pos:])
		result.WriteRune(r)
		pos += size
	}

	env.Stack.Push(result.String())
	return nil
}

func UTF8Encode(env *zc.Env) error {
	s, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	var result bytes.Buffer
	codePoint := make([]byte, 4)
	for _, r := range s {
		size := utf8.EncodeRune(codePoint, r)
		for i := 0; i < size; i++ {
			result.WriteByte(codePoint[i])
		}
	}
	var i big.Int
	i.SetBytes(result.Bytes())

	env.Stack.PushBigInt(&i)
	return nil
}
