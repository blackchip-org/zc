package ops

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/blackchip-org/zc/types"
)

type opTest struct {
	name   string
	args   []string
	result []string
	err    string
	exact  bool
}

var opTests = []opTest{
	{Abs, []string{"-6"}, []string{"6"}, "", false},
	{Abs, []string{"-2.2"}, []string{"2.2"}, "", false},
	{Abs, []string{"-1e30"}, []string{"1e+30"}, "", false},
	{Abs, []string{"3+4i"}, []string{"5"}, "", false},

	{Add, []string{"6", "2"}, []string{"8"}, "", true},
	{Add, []string{"0xa", "0x2"}, []string{"12"}, "", true},
	{Add, []string{"2.2", "1.1"}, []string{"3.3"}, "", true},
	{Add, []string{"2.2f", "1.1f"}, []string{"3.3"}, "", true},
	{Add, []string{"1e3", "1e2"}, []string{"1100"}, "", true},
	{Add, []string{"1e30", "1e29"}, []string{"1.1e+30"}, "", true},
	{Add, []string{"6+2i", "2+6i"}, []string{"8+8i"}, "", true},
	{Add, []string{"1/2", "1/4"}, []string{"3/4"}, "", true},
	{Add, []string{"3-3/4", "2-1/2"}, []string{"6 1/4"}, "", true},

	{And, []string{"true", "false"}, []string{"false"}, "", true},

	{Ceil, []string{"6"}, []string{"6"}, "", true},
	{Ceil, []string{"6.2"}, []string{"7"}, "", true},
	{Ceil, []string{"6.2f"}, []string{"7"}, "", true},

	{Div, []string{"6", "2"}, []string{"3"}, "", true},
	{Div, []string{"1", "2"}, []string{"0.5"}, "", false},
	{Div, []string{"2", "3"}, []string{"0.6666666666666666"}, "", false},
	{Div, []string{"2.2", "1.1"}, []string{"2"}, "", true},
	{Div, []string{"1e3", "1e2"}, []string{"10"}, "", true},
	{Div, []string{"6+8i", "2+2i"}, []string{"3.5+0.5i"}, "", true},
	{Div, []string{"1/4", "3/4"}, []string{"1/3"}, "", true},

	{Div, []string{"6", "0"}, []string{}, ErrDivisionByZero.Error(), false},
	{Div, []string{"2.2", "0"}, []string{}, ErrDivisionByZero.Error(), false},
	{Div, []string{"-1.1e0", "0"}, []string{}, ErrDivisionByZero.Error(), false},
	{Div, []string{"6+8i", "0"}, []string{"3.5+0.5i"}, ErrDivisionByZero.Error(), false},
	{Div, []string{"1/4", "0"}, []string{}, ErrDivisionByZero.Error(), false},

	{Eq, []string{"6", "6.0"}, []string{"true"}, "", false},
	{Eq, []string{"6.6", "6.6"}, []string{"true"}, "", false},
	{Eq, []string{"123e-1", "12.3"}, []string{"true"}, "", false},
	{Eq, []string{"1/2", "0.5"}, []string{"true"}, "", false},
	{Eq, []string{"true", "true"}, []string{"true"}, "", false},
	{Eq, []string{"foo", "foo"}, []string{"true"}, "", false},

	{Floor, []string{"6"}, []string{"6"}, "", true},
	{Floor, []string{"6.2"}, []string{"6"}, "", true},
	{Floor, []string{"6.2f"}, []string{"6"}, "", true},

	{Gt, []string{"6", "5"}, []string{"true"}, "", false},
	{Gt, []string{"6.6", "5.5"}, []string{"true"}, "", false},
	{Gt, []string{"123e-1", "1.23"}, []string{"true"}, "", false},
	{Gt, []string{"1/2", "0.25"}, []string{"true"}, "", false},
	{Gt, []string{"foo", "bar"}, []string{"true"}, "", false},

	{Gte, []string{"6", "5"}, []string{"true"}, "", false},
	{Gte, []string{"6.6", "5.5"}, []string{"true"}, "", false},
	{Gte, []string{"123e-1", "1.23"}, []string{"true"}, "", false},
	{Gte, []string{"1/2", "0.25"}, []string{"true"}, "", false},
	{Gte, []string{"foo", "bar"}, []string{"true"}, "", false},

	{Lt, []string{"6", "5"}, []string{"false"}, "", false},
	{Lt, []string{"6.6", "5.5"}, []string{"false"}, "", false},
	{Lt, []string{"123e-1", "1.23"}, []string{"false"}, "", false},
	{Lt, []string{"1/2", "0.25"}, []string{"false"}, "", false},
	{Lt, []string{"foo", "bar"}, []string{"false"}, "", false},

	{Lte, []string{"6", "5"}, []string{"false"}, "", false},
	{Lte, []string{"6.6", "5.5"}, []string{"false"}, "", false},
	{Lte, []string{"123e-1", "1.23"}, []string{"false"}, "", false},
	{Lte, []string{"1/2", "0.25"}, []string{"false"}, "", false},
	{Lte, []string{"foo", "bar"}, []string{"false"}, "", false},

	{Mul, []string{"6", "2"}, []string{"12"}, "", true},
	{Mul, []string{"6.6", "2.2"}, []string{"14.52"}, "", true},
	{Mul, []string{"1e3", "1e2"}, []string{"100000"}, "", true},
	{Mul, []string{"6+8i", "2+2i"}, []string{"-4+28i"}, "", true},
	{Mul, []string{"1/2", "1/4"}, []string{"1/8"}, "", true},

	{Mod, []string{"7", "2"}, []string{"1"}, "", true},
	{Mod, []string{"5.75", "0.5"}, []string{"0.25"}, "", true},
	{Mod, []string{"5.75f", "0.5f"}, []string{"0.25"}, "", true},

	{Mod, []string{"7", "0"}, []string{}, ErrDivisionByZero.Error(), true},
	{Mod, []string{"5.75", "0"}, []string{}, ErrDivisionByZero.Error(), false},
	{Mod, []string{"5.75f", "0f"}, []string{}, ErrDivisionByZero.Error(), true},

	{Neg, []string{"6"}, []string{"-6"}, "", false},
	{Neg, []string{"2.2"}, []string{"-2.2"}, "", false},
	{Neg, []string{"1e30"}, []string{"-1e+30"}, "", false},
	{Neg, []string{"3/4"}, []string{"-3/4"}, "", false},

	{Neq, []string{"6", "6.0"}, []string{"false"}, "", false},
	{Neq, []string{"6.6", "6.6"}, []string{"false"}, "", false},
	{Neq, []string{"123e-1", "12.3"}, []string{"false"}, "", false},
	{Neq, []string{"1/2", "0.5"}, []string{"false"}, "", false},
	{Neq, []string{"true", "true"}, []string{"false"}, "", false},
	{Neq, []string{"foo", "foo"}, []string{"false"}, "", false},

	{Not, []string{"true"}, []string{"false"}, "", false},

	{Or, []string{"true", "false"}, []string{"true"}, "", true},

	{Rem, []string{"7", "2"}, []string{"1"}, "", true},
	{Rem, []string{"5.75f", "0.5f"}, []string{"-0.25"}, "", true},

	{Pow, []string{"6", "2"}, []string{"36"}, "", true},
	{Pow, []string{"3.3", "2.2"}, []string{"13.82708611804415"}, "", false},
	{Pow, []string{"1e3", "1e2"}, []string{"1e+300"}, "", true},
	{Pow, []string{"6+2i", "2+6i"}, []string{"3.802646420739731-4.383464777273336i"}, "", true},

	{Sqrt, []string{"256"}, []string{"16"}, "", false},

	{Sub, []string{"6", "2"}, []string{"4"}, "", true},
	{Sub, []string{"3.3", "2.2"}, []string{"1.1"}, "", true},
	{Sub, []string{"1e3", "1e2"}, []string{"900"}, "", true},
	{Sub, []string{"6+2i", "2+6i"}, []string{"4-4i"}, "", true},
	{Sub, []string{"3/4", "1/2"}, []string{"1/4"}, "", true},

	{Sign, []string{"-6"}, []string{"-1"}, "", false},
	{Sign, []string{"-2.2"}, []string{"-1"}, "", false},
	{Sign, []string{"-1e30"}, []string{"-1"}, "", false},
	{Sign, []string{"-3/4"}, []string{"-1"}, "", false},
}

func doEval(t *testing.T, test opTest, exact bool) {
	rs, err := eval(test.name, types.ParseN(test.args), exact)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	if errMsg != test.err {
		t.Fatalf("\n have: %v \n want: %v", err, test.err)
	}
	if err != nil {
		return
	}
	result := types.FormatN(rs)
	if !reflect.DeepEqual(result, test.result) {
		t.Fatalf("\n have: %v \n want: %v", result, test.result)
	}
}

func TestOps(t *testing.T) {
	for _, test := range opTests {
		name := fmt.Sprintf("%v_%v", test.name, test.args)
		t.Run(name+"/exact", func(t *testing.T) {
			if test.exact {
				doEval(t, test, test.exact)
			}
		})
		t.Run(name, func(t *testing.T) {
			doEval(t, test, false)
		})
	}
}

/*
func TestSingle(t *testing.T) {
	args := []types.Generic{
		types.Parse("2"),
		types.Parse("3"),
	}
	result, err := Eval(Div, args)
	fmt.Printf("%v %v\n", result, err)
	t.Fail()
}
*/
