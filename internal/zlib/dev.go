package zlib

import (
	"errors"
	"fmt"
	"os"

	"github.com/blackchip-org/zc"
)

func Abort(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return errors.New("aborted")
	}
	return errors.New(a)
}

func Eval(env *zc.Env) error {
	fmt.Println(env.Funcs)
	node, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	return zc.Eval(env, "<eval>", []byte(node))
}

func Exit(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	code, err := env.Calc.ParseInt(a)
	if err != nil {
		return err
	}
	os.Exit(code)
	return nil
}

func Nothing(env *zc.Env) error {
	return nil
}

func Trace(env *zc.Env) error {
	env.Calc.Trace = true
	return nil
}

func TraceOff(env *zc.Env) error {
	env.Calc.Trace = false
	return nil
}

// FIXME: This should be handled better. Maybe a statement?
func Undef(env *zc.Env) error {
	return errors.New("not implemented")
	// target, err := env.Stack.Pop()
	// if err != nil {
	// 	return err
	// }

	// var n = 0
	// for name := range env.funcs {
	// 	parts := strings.Split(name, ".")
	// 	if parts[0] == target {
	// 		delete(env.Funcs, name)
	// 		n++
	// 	}
	// }
	// env.Printf("%v undefined", n)
	// return nil
}
