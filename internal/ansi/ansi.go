package ansi

import "fmt"

var Enabled bool = true

const (
	ClearScreen  = "\033[2J"
	MoveToBottom = "\033[200;0H" // go to line 200, column 0
	Reset        = "\033[0m"
	Bold         = "\033[1m"
	LightGreen   = "\033[1;32m"
	LightBlue    = "\033[1;36m"
	BrightYellow = "\033[1;93m"
)

func Write(v string) {
	if !Enabled {
		return
	}
	fmt.Print(v)
}
