package lexer

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/blackchip-org/zc/lang/scanner"
	"github.com/blackchip-org/zc/lang/token"
)

func TestToken(t *testing.T) {
	start := scanner.Pos{Name: "", Line: 1, Column: 1}
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
		{"'foo bar\nbaz", token.New(token.String, "foo bar", start)},
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

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d_%s", i, test.src), func(t *testing.T) {
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
