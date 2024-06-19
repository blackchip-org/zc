package repl

import (
	"testing"

	"github.com/blackchip-org/zc/v6/pkg/ansi"
)

func TestUndo(t *testing.T) {
	ansi.Enabled = false
	repl := NewReplTester(t)

	repl.Eval("1")
	repl.Eval("2")
	repl.Eval("3")
	repl.AssertStack(1, 2, 3)

	repl.Eval("undo")
	repl.AssertStack(1, 2)

	repl.Eval("undo")
	repl.AssertStack(1)

	repl.Eval("undo")
	repl.AssertStack()

	repl.Eval("undo")
	repl.AssertError("undo stack is empty")

	repl.Eval("redo")
	// 1
	repl.AssertStack("1")

	repl.Eval("redo")
	// 1 2
	repl.Eval("redo")
	// 1 2 3
	repl.AssertStack("1", "2", "3")
	repl.Eval("redo")
	repl.AssertError("redo stack is empty")
}

func TestQuote(t *testing.T) {
	ansi.Enabled = false
	repl := NewReplTester(t)

	repl.Eval("quote EOF")
	repl.Eval("1 2 add")
	repl.Eval("2 3 sub")
	repl.Eval("4")
	repl.Eval("EOF")
	repl.Eval("2 pow")

	repl.AssertStack("1 2 add", "2 3 sub", "16")
}

func TestQuoteBlanks(t *testing.T) {
	ansi.Enabled = false
	repl := NewReplTester(t)

	repl.Eval("quote EOF")
	repl.Eval("1 2 add")
	repl.Eval("2 3 sub")
	repl.Eval("")
	repl.Eval("")
	repl.Eval("EOF")

	repl.AssertStack("1 2 add", "2 3 sub", "", "")
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
