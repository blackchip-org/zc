package lexer

import (
	"fmt"
	"testing"

	"github.com/blackchip-org/zc/lang/token"
	"github.com/blackchip-org/zc/scanner"
	"github.com/stretchr/testify/assert"
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

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			s := New("", []byte(test.src))
			assert.Equal(t, test.want, s.Next())
		})
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
			assert.Equal(t, test.want, have)
		})
	}
}
