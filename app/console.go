package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/ansi"
	"github.com/blackchip-org/zc/internal/modules"
	"github.com/peterh/liner"
)

func RunConsole() {
	log.SetFlags(0)

	line := liner.NewLiner()
	defer line.Close()

	calc := zc.NewCalc(modules.Prelude)
	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)

	fmt.Print(ansi.ClearScreen)
	fmt.Print(ansi.MoveToBottom)

	prompt := fmt.Sprintf("%vzc%v> ", ansi.LightGreen, ansi.Reset)
	fmt.Print(prompt)
	text, err := line.Prompt("")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for ; err == nil; text, err = line.Prompt("") {
		err := calc.EvalString(text)
		fmt.Print(ansi.ClearScreen)
		for i, val := range calc.Stack().Items() {
			color := ansi.LightBlue
			if i == calc.Stack().Len()-1 {
				color = ansi.Bold
			}
			fmt.Printf("%v%v%v\n", color, val, ansi.Reset)
		}
		if err != nil {
			fmt.Printf("%v(!) %v%v\n", ansi.BrightYellow, err, ansi.Reset)
		}
		if strings.TrimSpace(text) != "" {
			line.AppendHistory(text)
		}
		fmt.Print(prompt)
	}

}
