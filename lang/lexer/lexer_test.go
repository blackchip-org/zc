package lexer

import (
	"reflect"
	"testing"

	"github.com/blackchip-org/zc/lang/scanner"
	"github.com/blackchip-org/zc/lang/token"
)

func TestPosition(t *testing.T) {
	src := []byte("one two\nthree four\n\n")
	pos := []scanner.Pos{
		scanner.NewPos("", 1, 1),
		scanner.NewPos("", 1, 5),
		scanner.NewPos("", 1, 8),
		scanner.NewPos("", 2, 1),
		scanner.NewPos("", 2, 7),
		scanner.NewPos("", 2, 11),
		scanner.NewPos("", 3, 1),
		scanner.NewPos("", 4, 1),
	}

	s := New("", src)
	for i, want := range pos {
		have := s.Next()
		if have.Pos != want {
			t.Fatalf("\n %v \n have %v \n want %v", i, have.Pos, want)
		}
	}
}

func TestToken(t *testing.T) {
	start := scanner.Pos{File: "", Line: 1, Column: 1}
	tests := []struct {
		src  string
		want token.Token
	}{
		{"foo", token.New(token.Id, "foo", start)},
		{"123", token.New(token.Value, "123", start)},
		{"-123", token.New(token.Value, "-123", start)},
		{"+123", token.New(token.Value, "+123", start)},
		{".123", token.New(token.Value, ".123", start)},
		{"\"foo bar\"", token.New(token.StringPlain, "foo bar", start)},
		{"'foo bar'", token.New(token.String, "foo bar", start)},
		{"'\\'foo bar\\''", token.New(token.String, "'foo bar'", start)},
		{"", token.New(token.End, "", start)},
		{"\n", token.New(token.Newline, "\n", start)},
		{"/", token.New(token.Id, "/", start)},
		{"/foo;", token.New(token.Slash, "/", start)},
		{"//", token.New(token.Id, "//", start)},
		{"//foo;", token.New(token.DoubleSlash, "//", start)},
		{"/-foo;", token.New(token.SlashDash, "/-", start)},
	}

	for i, test := range tests {
		s := New("", []byte(test.src))
		have := s.Next()
		if have != test.want {
			t.Errorf("\n %v \n have %+v\n want %+v", i, have, test.want)
		}
	}
}

func TestIndent(t *testing.T) {
	tests := []struct {
		src  string
		want []token.Type
	}{
		{"\tfoo", []token.Type{token.Indent, token.Id}},
		{"foo\n\tbar", []token.Type{token.Id, token.Newline, token.Indent, token.Id}},
		{"\tfoo\nbar", []token.Type{token.Indent, token.Id, token.Newline, token.Dedent, token.Id}},
		{"\tfoo\n\tbar", []token.Type{token.Indent, token.Id, token.Newline, token.Id}},
		{"[foo\nbar]\n", []token.Type{token.Id, token.Id, token.Newline}},
		{"a; [\n1\n2\n]\n", []token.Type{token.Id, token.Semicolon, token.Value, token.Value, token.Newline}},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			s := New("", []byte(test.src))
			var have []token.Type
			for tok := s.Next(); tok.Type != token.End; tok = s.Next() {
				have = append(have, tok.Type)
			}
			if !reflect.DeepEqual(have, test.want) {
				t.Errorf("\n have %v \n want %v", have, test.want)
			}
		})
	}
}

func TestQuote(t *testing.T) {
	tests := []struct {
		src  string
		want string
	}{
		{"1234", "1234"},
		{"abcd", "'abcd'"},
		{"12 34", "'12 34'"},
		{"a", "'a'"},
	}

	for _, test := range tests {
		have := Quote(test.src)
		if have != test.want {
			t.Errorf("\n have: %v \n want: %v", have, test.want)
		}
	}
}
