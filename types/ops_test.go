package types

import (
	"reflect"
	"testing"
)

type opTest struct {
	name   string
	args   []string
	result []string
	err    string
	exact  bool
}

var opTests = []opTest{
	{OpAbs, []string{"-6"}, []string{"6"}, "", false},
	{OpAbs, []string{"-2.2"}, []string{"2.2"}, "", false},
	{OpAbs, []string{"-1e30"}, []string{"1e+30"}, "", false},
	{OpAbs, []string{"3+4i"}, []string{"5"}, "", false},

	{OpAdd, []string{"6", "2"}, []string{"8"}, "", true},
	{OpAdd, []string{"0xa", "0x2"}, []string{"12"}, "", true},
	{OpAdd, []string{"2.2", "1.1"}, []string{"3.3"}, "", true},
	{OpAdd, []string{"2.2f", "1.1f"}, []string{"3.3"}, "", true},
	{OpAdd, []string{"1e3", "1e2"}, []string{"1100"}, "", true},
	{OpAdd, []string{"1e30", "1e29"}, []string{"1.1e+30"}, "", true},
	{OpAdd, []string{"6+2i", "2+6i"}, []string{"8+8i"}, "", true},
	{OpAdd, []string{"1/2", "1/4"}, []string{"3/4"}, "", true},
	{OpAdd, []string{"3-3/4", "2-1/2"}, []string{"6 1/4"}, "", true},

	{OpCeil, []string{"6"}, []string{"6"}, "", true},
	{OpCeil, []string{"6.2"}, []string{"7"}, "", true},
	{OpCeil, []string{"6.2f"}, []string{"7"}, "", true},

	{OpDiv, []string{"6", "2"}, []string{"3"}, "", true},
	{OpDiv, []string{"2.2", "1.1"}, []string{"2"}, "", true},
	{OpDiv, []string{"1e3", "1e2"}, []string{"10"}, "", true},
	{OpDiv, []string{"6+8i", "2+2i"}, []string{"3.5+0.5i"}, "", true},
	{OpDiv, []string{"1/4", "3/4"}, []string{"1/3"}, "", true},

	{OpDiv, []string{"6", "0"}, []string{}, ErrDivisionByZero.Error(), false},
	{OpDiv, []string{"2.2", "0"}, []string{}, ErrDivisionByZero.Error(), false},
	{OpDiv, []string{"-1.1e0", "0"}, []string{}, ErrDivisionByZero.Error(), false},
	{OpDiv, []string{"6+8i", "0"}, []string{"3.5+0.5i"}, ErrDivisionByZero.Error(), false},
	{OpDiv, []string{"1/4", "0"}, []string{}, ErrDivisionByZero.Error(), false},

	{OpFloor, []string{"6"}, []string{"6"}, "", true},
	{OpFloor, []string{"6.2"}, []string{"6"}, "", true},
	{OpFloor, []string{"6.2f"}, []string{"6"}, "", true},

	{OpMul, []string{"6", "2"}, []string{"12"}, "", true},
	{OpMul, []string{"6.6", "2.2"}, []string{"14.52"}, "", true},
	{OpMul, []string{"1e3", "1e2"}, []string{"100000"}, "", true},
	{OpMul, []string{"6+8i", "2+2i"}, []string{"-4+28i"}, "", true},
	{OpMul, []string{"1/2", "1/4"}, []string{"1/8"}, "", true},

	{OpMod, []string{"7", "2"}, []string{"1"}, "", true},
	{OpMod, []string{"5.75", "0.5"}, []string{"0.25"}, "", true},
	{OpMod, []string{"5.75f", "0.5f"}, []string{"0.25"}, "", true},

	{OpMod, []string{"7", "0"}, []string{}, ErrDivisionByZero.Error(), true},
	{OpMod, []string{"5.75", "0"}, []string{}, ErrDivisionByZero.Error(), false},
	{OpMod, []string{"5.75f", "0f"}, []string{}, ErrDivisionByZero.Error(), true},

	{OpNeg, []string{"6"}, []string{"-6"}, "", false},
	{OpNeg, []string{"2.2"}, []string{"-2.2"}, "", false},
	{OpNeg, []string{"1e30"}, []string{"-1e+30"}, "", false},
	{OpNeg, []string{"3/4"}, []string{"-3/4"}, "", false},

	{OpRem, []string{"7", "2"}, []string{"1"}, "", true},
	{OpRem, []string{"5.75f", "0.5f"}, []string{"-0.25"}, "", true},

	{OpPow, []string{"6", "2"}, []string{"36"}, "", true},
	{OpPow, []string{"3.3", "2.2"}, []string{"13.82708611804415"}, "", false},
	{OpPow, []string{"1e3", "1e2"}, []string{"1e+300"}, "", true},
	{OpPow, []string{"6+2i", "2+6i"}, []string{"3.802646420739731-4.383464777273336i"}, "", true},

	{OpSub, []string{"6", "2"}, []string{"4"}, "", true},
	{OpSub, []string{"3.3", "2.2"}, []string{"1.1"}, "", true},
	{OpSub, []string{"1e3", "1e2"}, []string{"900"}, "", true},
	{OpSub, []string{"6+2i", "2+6i"}, []string{"4-4i"}, "", true},
	{OpSub, []string{"3/4", "1/2"}, []string{"1/4"}, "", true},

	{OpSign, []string{"-6"}, []string{"-1"}, "", false},
	{OpSign, []string{"-2.2"}, []string{"-1"}, "", false},
	{OpSign, []string{"-1e30"}, []string{"-1"}, "", false},
	{OpSign, []string{"-3/4"}, []string{"-1"}, "", false},
}

func doEval(t *testing.T, test opTest, exact bool) {
	rs, err := eval(test.name, MustParseNumbers(test.args), exact)
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
	result := FormatGenerics(rs)
	if !reflect.DeepEqual(result, test.result) {
		t.Fatalf("\n have: %v \n want: %v", result, test.result)
	}
}

func TestOps(t *testing.T) {
	for _, test := range opTests {
		t.Run(test.name+"/exact", func(t *testing.T) {
			if test.exact {
				doEval(t, test, test.exact)
			}
		})
		t.Run(test.name, func(t *testing.T) {
			doEval(t, test, false)
		})
	}
}
