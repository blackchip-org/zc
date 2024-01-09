package repl

import (
	"reflect"

	"github.com/blackchip-org/zc/v5/pkg/scanner"
	"github.com/blackchip-org/zc/v5/pkg/zc"
)

func Test(r *REPL, test string) bool {
	var s scanner.Scanner
	s.SetString(test)

	var input string
	errorTest := false
	for s.Ok() {
		if (s.Ch == '-' || s.Ch == '!') && s.Lookahead == '-' {
			if s.Ch == '!' {
				errorTest = true
			}
			input = s.Token()
			s.Next()
			s.Next()
			break
		}
		s.Keep()
	}
	output := s.Scan(scanner.LineTrimSpace)
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
