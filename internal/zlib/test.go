package zlib

import (
	"fmt"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/parser"
)

func TestFile(env *zc.Env) error {
	file, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	testError := func(err error) {
		fmt.Printf("%v: error: %v", file, err)
		env.SetInt("errors", env.GetInt("errors")+1)
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
			env.SetInt("failed", env.GetInt("failed")+1)
		} else {
			if env.GetBool("verbose") {
				fmt.Printf("PASS: %v(%v)\n", file, fn.Name)
			}
			env.SetInt("passed", env.GetInt("passed")+1)
		}
	}
	return nil
}
