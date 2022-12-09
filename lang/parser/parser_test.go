package parser

import (
	"embed"
	"strings"
	"testing"
)

//go:embed tests/*
var parserTestData embed.FS

func TestParser(t *testing.T) {
	tests := []string{
		"blank_lines",
		"dedent2",
		"expr2",
		"for",
		"func2",
		"if-elif-else",
		"if-else",
		"if",
		"import",
		"include",
		"indent_spaces",
		"indent_tabs",
		"stack",
		"try",
		"while",
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			in, err := parserTestData.ReadFile("tests/" + test + ".zc")
			if err != nil {
				t.Fatal(err)
			}

			out, err := parserTestData.ReadFile("tests/" + test + ".json")
			if err != nil {
				t.Fatal(err)
			}

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
