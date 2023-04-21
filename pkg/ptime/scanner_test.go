package ptime

import (
	"reflect"
	"testing"
)

func TestScanner(t *testing.T) {
	tests := []struct {
		text   string
		tokens []Token
	}{
		{"1234", []Token{{Number, "1234", 1}}},
		{"abcd", []Token{{Text, "abcd", 1}}},
		{"+", []Token{{Indicator, "+", 1}}},
		{"+0700", []Token{
			{Indicator, "+", 1},
			{Number, "0700", 2},
		}},
		{"1/2/2003", []Token{
			{Number, "1", 1},
			{Indicator, "/", 2},
			{Number, "2", 3},
			{Indicator, "/", 4},
			{Number, "2003", 5},
		}},
		{"2-Jan-06", []Token{
			{Number, "2", 1},
			{Indicator, "-", 2},
			{Text, "Jan", 3},
			{Indicator, "-", 6},
			{Number, "06", 7},
		}},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			tokens := Scan(test.text)
			if !reflect.DeepEqual(tokens, test.tokens) {
				t.Errorf("\n have: %v \n want: %v", tokens, test.tokens)
			}
		})
	}
}
