package native

import (
	"errors"
	"fmt"

	"github.com/blackchip-org/zc"
)

const (
	gTestingName   = "__testing.name"
	gTestingFailed = "__testing.failed"
)

func TestSuite(calc *zc.Calc) error {
	path, err := calc.Stack().Pop()
	if err != nil {
		return err
	}
	text, _ := zc.Static.ReadFile(path)
	return calc.Eval(text)
}

func Test(calc *zc.Calc) error {
	name, err := calc.Stack().Pop()
	if err != nil {
		return err
	}
	calc.Global(gTestingName).Set(name)
	return nil
}

func Assert(calc *zc.Calc) error {
	name, err := calc.Global(gTestingName).Get()
	if err != nil {
		return errors.New("assert called outside of test")
	}

	a, err := calc.Stack().Pop()
	if err != nil {
		return err
	}

	ab, err := zc.ParseBool(a)
	if err != nil {
		return err
	}

	if !ab {
		fmt.Fprintf(calc.Out, "FAIL: %v\n", name)
		calc.Global(gTestingFailed).Push("true")
	} else {
		fmt.Fprintf(calc.Out, "pass: %v\n", name)
	}
	return nil
}
