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
