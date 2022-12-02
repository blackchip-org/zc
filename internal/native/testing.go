package native

// import (
// 	"errors"

// 	"github.com/blackchip-org/zc"
// )

// const (
// 	gTestingName   = "__testing.name"
// 	gTestingFailed = "__testing.failed"
// )

// func TestSuite(calc zc.Env) error {
// 	path, err := calc.Stack().Pop()
// 	if err != nil {
// 		return err
// 	}
// 	text, _ := zc.Static.ReadFile(path)
// 	return calc.Eval(text)
// }

// func Test(calc zc.Env) error {
// 	name, err := calc.Stack().Pop()
// 	if err != nil {
// 		return err
// 	}
// 	calc.Global(gTestingName).Set(name)
// 	return nil
// }

// func Assert(calc zc.Env) error {
// 	name, err := calc.Global(gTestingName).Get()
// 	if err != nil {
// 		return errors.New("assert called outside of test")
// 	}

// 	a, err := calc.Stack().Pop()
// 	if err != nil {
// 		return err
// 	}

// 	ab, err := zc.ParseBool(a)
// 	if err != nil {
// 		return err
// 	}

// 	if !ab {
// 		calc.Printf("FAIL: %v\n", name)
// 		calc.Global(gTestingFailed).Push("true")
// 	} else {
// 		calc.Printf("pass: %v\n", name)
// 	}
// 	return nil
// }
