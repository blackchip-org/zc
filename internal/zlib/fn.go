package zlib

import (
	"github.com/blackchip-org/zc"
)

func Eval(env *zc.Env) error {
	context := env.Calc.Frames[len(env.Calc.Frames)-2].Env
	node, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	return zc.Eval(context, "<eval>", []byte(node))
}

func Filter(env *zc.Env) error {
	context := env.Calc.Frames[len(env.Calc.Frames)-2].Env
	fn, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	items := env.Stack.Items()
	env.Stack.Clear()
	for _, item := range items {
		env.Stack.Push(item)
		if err := zc.Eval(context, "<map>", []byte(fn)); err != nil {
			return err
		}
		v, err := env.Stack.PopBool()
		if err != nil {
			return err
		}
		if v {
			env.Stack.Push(item)
		}
	}
	return nil
}

func Map(env *zc.Env) error {
	context := env.Calc.Frames[len(env.Calc.Frames)-2].Env
	fn, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	items := env.Stack.Items()
	env.Stack.Clear()
	for _, item := range items {
		env.Stack.Push(item)
		if err := zc.Eval(context, "<map>", []byte(fn)); err != nil {
			return err
		}
	}
	return nil
}

func Fold(env *zc.Env) error {
	context := env.Calc.Frames[len(env.Calc.Frames)-2].Env
	fn, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	if env.Stack.Len() < 2 {
		return nil
	}
	items := env.Stack.Items()
	env.Stack.Clear()
	env.Stack.Push(items[0])
	for _, item := range items[1:] {
		env.Stack.Push(item)
		if err := zc.Eval(context, "<fold>", []byte(fn)); err != nil {
			return err
		}
	}
	return nil
}

func Repeat(env *zc.Env) error {
	context := env.Calc.Frames[len(env.Calc.Frames)-2].Env
	n, err := env.Stack.PopInt()
	if err != nil {
		return err
	}
	fn, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	for i := 0; i < n; i++ {
		if err := zc.Eval(context, "<repeat>", []byte(fn)); err != nil {
			return err
		}
	}
	return nil
}
