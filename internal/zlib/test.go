package zlib

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/parser"
)

type testState struct {
	pass   int
	fail   int
	errors int
}

func getTestState(env *zc.Env) *testState {
	return env.Calc.States["test"].(*testState)
}

func InitTest(env *zc.Env) error {
	env.Calc.States["test"] = &testState{}
	return nil
}

func TestBegin(env *zc.Env) error {
	s := getTestState(env)
	s.pass = 0
	s.fail = 0
	s.errors = 0
	return nil
}

func TestEnd(env *zc.Env) error {
	s := getTestState(env)
	fmt.Printf("\n%v passed, %v failed, %v error(s)\n", s.pass, s.fail, s.errors)
	return nil
}

func TestFile(env *zc.Env) error {
	s := getTestState(env)

	file, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	testError := func(err error) {
		fmt.Printf("%v: error: %v", file, err)
		s.errors++
	}

	text, err := zc.LoadFile(file)
	if err != nil {
		testError(err)
		return nil
	}

	root, err := parser.Parse(file, text)
	if err != nil {
		testError(err)
		return nil
	}

	for _, stmt := range root.Stmts {
		fn, ok := stmt.(*ast.FuncStmt)
		if !ok {
			continue
		}
		if !strings.HasPrefix(fn.Name, "test-") {
			continue
		}
		c, err := zc.NewCalc(env.Calc.Config)
		if err != nil {
			panic(fmt.Sprintf("unable to create calc: %v", err))
		}
		cmd := fmt.Sprintf("use '%v'\n%v", file, fn.Name)
		if err := c.EvalString(file, cmd); err != nil {
			fmt.Printf("FAIL: %v(%v): %v\n", file, fn.Name, err)
			s.fail++
		} else {
			fmt.Printf("PASS: %v(%v)\n", file, fn.Name)
			s.pass++
		}
	}
	return nil
}

func TestOk(env *zc.Env) error {
	s := getTestState(env)
	env.Stack.PushBool(s.fail == 0 && s.errors == 0)
	return nil
}
