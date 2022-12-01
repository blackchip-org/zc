package lang

import (
	"embed"
	"strings"
	"testing"
)

//go:embed parser_tests/*
var parserTestData embed.FS

func TestParser(t *testing.T) {
	tests := []string{
		"expr2",
		"fn2",
		"while",
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			in, _ := parserTestData.ReadFile("parser_tests/" + test + ".zc")
			out, _ := parserTestData.ReadFile("parser_tests/" + test + ".json")

			ast, err := Parse("", in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			have := strings.TrimSpace(ast.String())
			want := strings.TrimSpace(string(out))
			if have != want {
				t.Errorf("\n have \n%v\n want \n%v", have, want)
			}
		})
	}
}
