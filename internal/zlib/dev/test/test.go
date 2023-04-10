package test

import (
	"errors"
	"fmt"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/parser"
)

func testError(env *zc.Env, err error) {
	fmt.Printf("ERROR %v: %v\n", env.Get("name"), err)
	if calcErr, ok := err.(zc.CalcError); ok {
		for _, frame := range calcErr.Frames {
			fmt.Println(frame)
		}
	}
	env.SetInt("errors", env.GetInt("errors")+1)
}

func TestFile(env *zc.Env) error {
	file, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	env.Set("name", file)

	text, err := zc.LoadFile(file)
	if err != nil {
		testError(env, err)
		return nil
	}
	root, err := parser.Parse(file, text)
	if err != nil {
		testError(env, err)
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

		c, err := zc.NewCalc(env.Calc.Config())
		if err != nil {
			panic(fmt.Sprintf("unable to create calc: %v", err))
		}

		cmd := []string{
			"import test",
			fmt.Sprintf("%v test.verbose", env.GetBool("verbose")),
			fmt.Sprintf("'%v/%v' test.name", file, fn.Name),
			fmt.Sprintf("use '%v'", file),
			fn.Name,
			"test.passed",
			"test.failed",
		}
		if err := c.EvalLines(file, cmd); err != nil {
			testError(env, err)
			return nil
		} else {
			failed, _ := c.Env.Stack.PopInt()
			passed, _ := c.Env.Stack.PopInt()
			if passed == 0 && failed == 0 {
				testError(c.Env, errors.New("no assertions"))
			}
			if env.GetBool("verbose") {
				fmt.Printf("PASS %v/%v\n", file, fn.Name)
			}
			env.SetInt("passed", env.GetInt("passed")+passed)
			env.SetInt("failed", env.GetInt("failed")+failed)
		}
	}
	return nil
}
