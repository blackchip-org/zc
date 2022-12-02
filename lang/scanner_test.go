package lang

import (
	"reflect"
	"testing"
)

func TestPosition(t *testing.T) {
	src := []byte("one two\nthree four\n\n")
	pos := []Position{
		{"", 1, 1},
		{"", 1, 5},
		{"", 1, 8},
		{"", 2, 1},
		{"", 2, 7},
		{"", 2, 11},
		{"", 3, 1},
		{"", 4, 1},
	}

	s := NewScanner("", src)
	for i, want := range pos {
		have := s.Next()
		if have.At != want {
			t.Fatalf("\n %v \n have %v \n want %v", i, have.At, want)
		}
	}
}

func TestToken(t *testing.T) {
	start := Position{File: "", Line: 1, Column: 1}
	tests := []struct {
		src  string
		want Token
	}{
		{"foo", Token{IdToken, "foo", start}},
		{"123", Token{ValueToken, "123", start}},
		{"-123", Token{ValueToken, "-123", start}},
		{"+123", Token{ValueToken, "+123", start}},
		{"\"foo bar\"", Token{ValueToken, "foo bar", start}},
		{"'foo bar'", Token{ValueToken, "foo bar", start}},
		{"'\\'foo bar\\''", Token{ValueToken, "'foo bar'", start}},
		{"", Token{EndToken, "", start}},
		{"\n", Token{NewlineToken, "\n", start}},
		{"/", Token{IdToken, "/", start}},
		{"/foo;", Token{SlashToken, "/", start}},
		{"//", Token{IdToken, "//", start}},
		{"//foo;", Token{DoubleSlashToken, "//", start}},
	}

	for i, test := range tests {
		s := NewScanner("", []byte(test.src))
		have := s.Next()
		if have != test.want {
			t.Errorf("\n %v \n have %+v\n want %+v", i, have, test.want)
		}
	}
}

func TestIndent(t *testing.T) {
	tests := []struct {
		src  string
		want []TokenType
	}{
		{"\tfoo", []TokenType{IndentToken, IdToken}},
		{"foo\n\tbar", []TokenType{IdToken, NewlineToken, IndentToken, IdToken}},
		{"\tfoo\nbar", []TokenType{IndentToken, IdToken, NewlineToken, DedentToken, IdToken}},
		{"\tfoo\n\tbar", []TokenType{IndentToken, IdToken, NewlineToken, IdToken}},
	}

	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			s := NewScanner("", []byte(test.src))
			var have []TokenType
			for tok := s.Next(); tok.Type != EndToken; tok = s.Next() {
				have = append(have, tok.Type)
			}
			if !reflect.DeepEqual(have, test.want) {
				t.Errorf("\n have %v \n want %v", have, test.want)
			}
		})
	}
}
