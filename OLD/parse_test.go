package zc

import (
	"testing"

	"github.com/blackchip-org/scan"
)

func TestScanner(t *testing.T) {
	tests := []scan.Test{
		scan.NewTest("abc", "abc", 1, 1, TokenName),
		scan.NewTest("abc def", "abc", 1, 1, TokenName).
			And("def", 1, 5, TokenName),
		scan.NewTest("/abc", "abc", 1, 1, TokenValue),
		scan.NewTest("/abc abc", "abc", 1, 1, TokenValue).
			And("abc", 1, 6, TokenName),
		scan.NewTest(`"abc"`, "abc", 1, 1, TokenValue),
		scan.NewTest(`'abc'`, "abc", 1, 1, TokenValue),
		scan.NewTest(`[abc]`, "abc", 1, 1, TokenValue),
		scan.NewTest(`"abc`, "abc", 1, 1, TokenValue),
		scan.NewTest(`'abc`, "abc", 1, 1, TokenValue),
		scan.NewTest(`[abc`, "abc", 1, 1, TokenValue),
	}
	scan.RunTests(t, rules, tests)
}

func TestPreParseNumber(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{"1234", "1234"},
		{"1,234", "1234"},
		{"1_234", "1234"},
		{"1 234", "1234"},
		{"$1234", "1234"},
		{"1234×1045", "1234e45"},
		{"1234x1045", "1234e45"},
	}

	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out := PreParseNumber(test.in)
			if out != test.out {
				t.Fatalf("\n have: %v \n want: %v", out, test.out)
			}
		})
	}
}
