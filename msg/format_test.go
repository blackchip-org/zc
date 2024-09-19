package msg

import "testing"

func TestEscapeString(t *testing.T) {
	tests := []struct {
		str string
		esc string
	}{
		{"1", "1"},
		{"\a", "\\a"},
		{"\b", "\\b"},
		{"\f", "\\f"},
		{"\n", "\\n"},
		{"\r", "\\r"},
		{"\t", "\\t"},
		{"\v", "\\v"},
		{"\x01", "\\x01"},
		{"\ufffe", "\\ufffe"},
		{"\U00100000", "\\U00100000"},
		{"1\a\b\f2", "1\\a\\b\\f2"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			esc := EscapeString(test.str)
			if esc != test.esc {
				t.Fatalf("\n have: %v \n want: %v", esc, test.esc)
			}
		})
	}
}

func TestQuote(t *testing.T) {
	tests := []struct {
		str    string
		quoted string
	}{
		{`123`, `123`},
		{`12 34`, `'12 34'`},
		{"12\n34", `12\n34`},
	}

	for _, test := range tests {
		t.Run(test.quoted, func(t *testing.T) {
			quoted := Quote(test.str)
			if quoted != test.quoted {
				t.Fatalf("\n have: %v \n want: %v", quoted, test.quoted)
			}
		})
	}
}
