package types

import (
	"reflect"
	"testing"
)

func TestOps(t *testing.T) {
	tests := []struct {
		name   string
		args   []string
		result []string
		err    string
		exact  bool
	}{
		{"add", []string{"6", "2"}, []string{"8"}, "", true},
		{"add", []string{"0xa", "0x2"}, []string{"12"}, "", true},
		{"add", []string{"2.2", "1.1"}, []string{"3.3"}, "", true},
		{"add", []string{"1e3", "1e2"}, []string{"1100"}, "", true},
		{"add", []string{"1e3", "1e2"}, []string{"1100"}, "", true},
		{"add", []string{"6+2i", "2+6i"}, []string{"8+8i"}, "", true},

		{"div", []string{"6", "2"}, []string{"3"}, "", true},
		{"div", []string{"2.2", "1.1"}, []string{"2"}, "", true},
		{"div", []string{"1e3", "1e2"}, []string{"10"}, "", true},
		{"div", []string{"6+8i", "2+2i"}, []string{"3.5+0.5i"}, "", true},

		{"div", []string{"6", "0"}, []string{}, ErrDivisionByZero.Error(), false},
		{"div", []string{"2.2", "0"}, []string{}, ErrDivisionByZero.Error(), false},
		{"div", []string{"-1.1e0", "0"}, []string{}, ErrDivisionByZero.Error(), false},
		{"div", []string{"6+8i", "0"}, []string{"3.5+0.5i"}, "+Inf+Infi", false},

		{"mul", []string{"6", "2"}, []string{"12"}, "", true},
		{"mul", []string{"6.6", "2.2"}, []string{"14.52"}, "", true},
		{"mul", []string{"1e3", "1e2"}, []string{"100000"}, "", true},
		{"mul", []string{"6+8i", "2+2i"}, []string{"-4+28i"}, "", true},

		{"sub", []string{"6", "2"}, []string{"4"}, "", true},
		{"sub", []string{"3.3", "2.2"}, []string{"1.1"}, "", true},
		{"sub", []string{"1e3", "1e2"}, []string{"900"}, "", true},
		{"sub", []string{"6+2i", "2+6i"}, []string{"4-4i"}, "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rs, err := eval(test.name, MustParseNumbers(test.args), test.exact)
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
			result := FormatValues(rs)
			if !reflect.DeepEqual(result, test.result) {
				t.Fatalf("\n have: %v \n want: %v", result, test.result)
			}
		})
	}
}
