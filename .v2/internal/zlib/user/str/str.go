package str

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/zc"
)

func Join(env *zc.Env) error {
	sep, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	items := env.Stack.Items()
	env.Stack.Clear()
	env.Stack.Push(strings.Join(items, sep))
	return nil
}

func Left(env *zc.Env) error {
	i, err := env.Stack.PopInt()
	if err != nil {
		return err
	}
	s, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	if i > len(s) || i < -len(s) {
		return fmt.Errorf("invalid length: %v", i)
	}
	switch {
	case i > 0:
		env.Stack.Push(s[:i])
	case i == 0:
		env.Stack.Push(s)
	case i < 0:
		env.Stack.Push(s[:len(s)+i])
	}
	return nil
}

func Len(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	r := len(a)
	env.Stack.PushInt(r)
	return nil
}

func LowerStr(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	z := strings.ToLower(a)
	env.Stack.Push(z)
	return nil
}

func Right(env *zc.Env) error {
	i, err := env.Stack.PopInt()
	if err != nil {
		return err
	}
	s, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	if i > len(s) || i < -len(s) {
		return fmt.Errorf("invalid length: %v", i)
	}
	switch {
	case i > 0:
		env.Stack.Push(s[len(s)-i:])
	case i == 0:
		env.Stack.Push(s)
	case i < 0:
		env.Stack.Push(s[-i:])
	}
	return nil
}

func Split(env *zc.Env) error {
	sep, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	s, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	xs := strings.Split(s, sep)
	for _, x := range xs {
		env.Stack.Push(x)
	}
	return nil
}

func StartsWith(env *zc.Env) error {
	prefix, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	str, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	r := strings.HasPrefix(str, prefix)
	env.Stack.PushBool(r)
	return nil
}

func UpperStr(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	z := strings.ToUpper(a)
	env.Stack.Push(z)
	return nil
}
