package zc

// import (
// 	"math/big"
// 	"testing"
// )

// func TestParseBigInt(t *testing.T) {
// 	tests := []struct {
// 		s    string
// 		want *big.Int
// 	}{
// 		{"1234", new(big.Int).SetInt64(1234)},
// 		{"0xffd2", new(big.Int).SetInt64(65490)},
// 	}

// 	value := DefaultValueOps()
// 	for _, test := range tests {
// 		t.Run(test.s, func(t *testing.T) {
// 			have, err := value.ParseBigInt(test.s)
// 			if err != nil {
// 				t.Fatalf("unexpected error: %v", err)
// 			}
// 			if have.Cmp(test.want) != 0 {
// 				t.Errorf("\n have: %v \n want: %v", have, test.want)
// 			}
// 		})
// 	}
// }

// func TestParseBigIntInvalid(t *testing.T) {
// 	tests := []string{
// 		"abcd",
// 	}

// 	value := DefaultValueOps()
// 	for _, test := range tests {
// 		t.Run(test, func(t *testing.T) {
// 			_, err := value.ParseBigInt(test)
// 			if err == nil {
// 				t.Errorf("expecting error")
// 			}
// 		})
// 	}
// }

// func TestParseRadix(t *testing.T) {
// 	tests := []struct {
// 		s     string
// 		radix int
// 	}{
// 		{"1234", 10},
// 		{"abcd", 10},
// 		{"0xabcd", 16},
// 		{"0XABCD", 16},
// 		{"0b0101", 2},
// 		{"0o755", 8},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.s, func(t *testing.T) {
// 			radix := ParseRadix(test.s)
// 			if radix != test.radix {
// 				t.Errorf("\n have: %v \n want: %v", radix, test.radix)
// 			}
// 		})
// 	}
// }

// func TestFormatNumberString(t *testing.T) {
// 	opsDefault := DefaultValueOps()
// 	opsFR := ValueOps{
// 		IntPat: ".000",
// 		Point:  ',',
// 	}
// 	opsBin := ValueOps{
// 		IntPat: "__0000_0000",
// 	}
// 	opsEmpty := ValueOps{}

// 	tests := []struct {
// 		in   string
// 		want string
// 		ops  ValueOps
// 	}{
// 		{"1", "1", opsDefault},
// 		{"123", "123", opsDefault},
// 		{"1234", "1,234", opsDefault},
// 		{"123456", "123,456", opsDefault},
// 		{"1234567", "1,234,567", opsDefault},
// 		{"-123", "-123", opsDefault},
// 		{".123", ".123", opsDefault},
// 		{".12345", ".12345", opsDefault},
// 		{"1234567.8901", "1,234,567.8901", opsDefault},
// 		{"1234567.8901", "1234567.8901", opsEmpty},
// 		{"1234567.8901", "1.234.567,8901", opsFR},
// 		{"11110000", "1111_0000", opsBin},
// 		{"1111000011110000", "1111_0000__1111_0000", opsBin},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.in, func(t *testing.T) {
// 			have := test.ops.FormatNumberString(test.in)
// 			if have != test.want {
// 				t.Errorf("\n have: %v \n want: %v", have, test.want)
// 			}
// 		})
// 	}
// }
