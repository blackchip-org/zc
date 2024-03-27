package repl

import (
	"reflect"
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/pkg/zc"
)

func Test(r *REPL, test string) bool {
	var s scan.Scanner
	s.InitFromString("", test)

	var input string
	errorTest := false
	for s.HasMore() {
		if (s.This == '-' || s.This == '!') && s.Next == '-' {
			if s.This == '!' {
				errorTest = true
			}
			input = s.Emit().Val
			s.Discard()
			s.Discard()
			break
		}
		s.Keep()
	}

	output := strings.TrimSpace(scan.Line(&s))
	r.Eval(input)
	err := r.Error()
	if errorTest {
		if err == nil {
			return false
		}
		return err.Error() == output
	}
	return reflect.DeepEqual(zc.StackString(r.Calc), output)
}
