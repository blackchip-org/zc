package repl

import (
	"testing"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/calc"
	"github.com/blackchip-org/zc/v6/pkg/ansi"
)

func TestUndo(t *testing.T) {
	ansi.Enabled = false
	c := calc.New()
	repl := New(c)

	repl.Eval("1")
	repl.Eval("2")
	repl.Eval("3")
	zc.AssertStack(t, c, 1, 2, 3)

	repl.Eval("undo")
	zc.AssertStack(t, c, 1, 2)

	repl.Eval("undo")
	zc.AssertStack(t, c, 1)

	repl.Eval("undo")
	zc.AssertStack(t, c)

	repl.Eval("undo")
	if repl.Error() == nil {
		t.Fatalf("expected error")
	}
	repl.Eval("redo")
	// 1
	zc.AssertStack(t, c, "1")

	repl.Eval("redo")
	// 1 2
	repl.Eval("redo")
	// 1 2 3
	zc.AssertStack(t, c, "1", "2", "3")
	repl.Eval("redo")
	if repl.Error() == nil {
		t.Fatalf("expected error")
	}
}

func TestQuote(t *testing.T) {
	ansi.Enabled = false
	c := calc.New()
	repl := New(c)

	repl.Eval("quote EOF")
	repl.Eval("1 2 add")
	repl.Eval("2 3 sub")
	repl.Eval("4")
	repl.Eval("EOF")
	repl.Eval("2 pow")

	zc.AssertStack(t, c, "1 2 add", "2 3 sub", "16")
}

func TestQuoteBlanks(t *testing.T) {
	ansi.Enabled = false
	c := calc.New()
	repl := New(c)

	repl.Eval("quote EOF")
	repl.Eval("1 2 add")
	repl.Eval("2 3 sub")
	repl.Eval("")
	repl.Eval("")
	repl.Eval("EOF")

	zc.AssertStack(t, c, "1 2 add", "2 3 sub", "", "")
}

func TestCommonPrefix(t *testing.T) {
	tests := []struct {
		common string
		vals   []string
	}{
		{"abc", []string{"abc", "abc", "abc"}},
		{"a", []string{"abc", "ab", "a"}},
		{"a", []string{"a", "ab", "abc"}},
		{"", []string{"a", "b", "c"}},
		{"abc", []string{"abcde", "abcfg", "abch"}},
		{"char-c", []string{"char-codepoint", "char-cp"}},
	}

	for _, test := range tests {
		t.Run(test.common, func(t *testing.T) {
			common := CommonPrefix(test.vals)
			if common != test.common {
				t.Errorf("\n have: %v \n want: %v", common, test.common)
			}
		})
	}
}
