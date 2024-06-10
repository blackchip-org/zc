package zc

import (
	"testing"
)

func TestScanner(t *testing.T) {
	tests := []struct {
		src  string
		n    int
		tok  Token
		lit  string
		name string
	}{
		{"a 0", 0, TokName, "a", "one name"},
		{"a 0", 1, TokValue, "0", "one value"},
		{"abc def", 0, TokName, "abc", "name"},
		{"abc def", 1, TokName, "def", "name, name"},
		{"/abc def", 0, TokValue, "abc", "slash val"},
		{"/abc def", 1, TokName, "def", "slash val, name"},
		{"\"abc\" def", 0, TokValue, "abc", "double quote val"},
		{"'abc' def", 0, TokValue, "abc", "single quote val"},
		{"[abc] def", 0, TokValue, "abc", "bracket quote val"},
		{"[abc def", 0, TokValue, "abc def", "bracket quote val dangling"},
	}

	for _, test := range tests {
		var s Scanner
		t.Run(test.name, func(t *testing.T) {
			s.InitFromString(test.src)
			for i := 0; i < test.n; i++ {
				s.Next()
			}
			tok, lit := s.Next()
			if tok != test.tok || lit != test.lit {
				t.Errorf("\n have %v: %v \n want %v: %v", tok, lit, test.tok, test.lit)
			}
		})
	}
}
