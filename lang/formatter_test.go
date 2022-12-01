package lang

// import (
// 	"fmt"
// 	"testing"
// )

// func TestFormatRoundTrip(t *testing.T) {
// 	tests := []string{
// 		"1234 0xabcd\n",
// 		"1234\n0xabcd\n",
// 	}

// 	for i, want := range tests {
// 		t.Run(fmt.Sprintf("test %v", i), func(t *testing.T) {
// 			ast, errors := Parse("", []byte(want))
// 			if errors != nil {
// 				t.Fatalf("unexpected error: %v", errors)
// 			}
// 			have := Format(ast)
// 			if have != want {
// 				t.Fatalf("\n have \n%v\n want \n%v", have, want)
// 			}
// 		})
// 	}
// }
