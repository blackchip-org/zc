package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		t.Run(test.src, func(t *testing.T) {
			assert.Equal(t, test.want, Quote(test.src))
		})
	}
}
