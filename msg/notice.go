package msg

import "fmt"

func PrecisionSet(p uint) string {
	return fmt.Sprintf("precision set to %v", p)
}

func Inexact() string {
	return "inexact"
}
