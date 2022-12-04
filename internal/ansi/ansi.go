package ansi

import "fmt"

var Enabled bool = true

const (
	ClearScreen   = "\033[2J"
	MoveToBottom  = "\033[200;0H" // go to line 200, column 0
	Reset         = "\033[0m"
	Bold          = "\033[1m"
	Cyan          = "\033[36m"
	LightGreen    = "\033[1;32m"
	LightBlue     = "\033[1;36m"
	BrightRed     = "\033[1;31m"
	BrightYellow  = "\033[1;93m"
	BrightMagenta = "\033[1;95m"
)

func FgColor8(v uint8) string {
	return fmt.Sprintf("\033[38;5;%vm", v)
}

func Write(v string) {
	if !Enabled {
		return
	}
	fmt.Print(v)
}
