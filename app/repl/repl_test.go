package repl

import (
	"testing"
)

func TestUndo(t *testing.T) {
	r := NewReplTester(t)

	r.Eval("1")
	r.Eval("2")
	r.Eval("3")
	r.AssertStack(1, 2, 3)

	r.Eval("undo")
	r.AssertStack(1, 2)

	r.Eval("undo")
	r.AssertStack(1)

	r.Eval("undo")
	r.AssertStack()

	r.Eval("undo")
	r.AssertError("undo stack is empty")

	r.Eval("redo")
	// 1
	r.AssertStack("1")

	r.Eval("redo")
	// 1 2
	r.Eval("redo")
	// 1 2 3
	r.AssertStack("1", "2", "3")
	r.Eval("redo")
	r.AssertError("redo stack is empty")
}

func TestQuote(t *testing.T) {
	r := NewReplTester(t)

	r.Eval("quote EOF")
	r.Eval("1 2 add")
	r.Eval("2 3 sub")
	r.Eval("4")
	r.Eval("EOF")
	r.Eval("2 pow")

	r.AssertStack("1 2 add", "2 3 sub", "16")
}

func TestQuoteBlanks(t *testing.T) {
	r := NewReplTester(t)

	r.Eval("quote EOF")
	r.Eval("1 2 add")
	r.Eval("2 3 sub")
	r.Eval("")
	r.Eval("")
	r.Eval("EOF")

	r.AssertStack("1 2 add", "2 3 sub", "", "")
}

func TestDrop(t *testing.T) {
	r := NewReplTester(t)
	r.Eval("1 2 3")
	r.Eval("")
	r.AssertStack(1, 2)
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
