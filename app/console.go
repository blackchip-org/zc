package app

import (
	"fmt"
	"log"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/ansi"
	"github.com/peterh/liner"
)

func RunConsole(calc *zc.Calc) {
	log.SetFlags(0)

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)

	ansi.Write(ansi.ClearScreen)
	ansi.Write(ansi.MoveToBottom)

	prompt := "zc > "
	text, err := line.Prompt(prompt)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// homedir, err := os.UserHomeDir()
	// if err != nil {
	// 	log.Fatalf("unable to determine home directory: %v", err)
	// }

	// historyFile := path.Join(homedir, ".zc-history")
	// f, err := os.OpenFile(historyFile, os.O_CREATE, 0o640)
	// if err != nil {
	// 	log.Printf("unable to save history: %v", err)
	// 	return
	// }

	// defer func() {
	// 	fmt.Println("****** SAVING HISTORY")
	// 	_, err := line.WriteHistory(f)
	// 	if err != nil {
	// 		log.Printf("unable to save history: %v", err)
	// 	}
	// 	f.Close()
	// }()

	for ; err == nil; text, err = line.Prompt(prompt) {
		var err error
		if strings.TrimSpace(text) == "" {
			if calc.Stack.Len() > 0 {
				_, err = calc.Stack.Pop()
			}
		} else {
			err = calc.EvalString(text)
		}

		ansi.Write(ansi.ClearScreen)
		fmt.Println()

		for i, val := range calc.Stack.Items() {
			color := ansi.LightBlue
			if i == calc.Stack.Len()-1 {
				color = ansi.Bold
			}
			ansi.Write(color)
			fmt.Println(val)
			ansi.Write(ansi.Reset)
		}
		if err != nil {
			ansi.Write(ansi.BrightYellow)
			fmt.Printf("(!) %v\n", err)
			ansi.Write(ansi.Reset)
		} else {
			fmt.Println()
		}
		if strings.TrimSpace(text) != "" {
			line.AppendHistory(text)
		}
	}
}
