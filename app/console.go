package app

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/ansi"
	"github.com/peterh/liner"
)

func RunConsole(calc *zc.Calc) {
	log.SetFlags(0)

	line := liner.NewLiner()
	defer line.Close()

	loadHistory(line)
	defer saveHistory(line)

	line.SetCtrlCAborts(true)
	line.SetTabCompletionStyle(liner.TabPrints)
	line.SetWordCompleter(calc.WordCompleter)

	ansi.Write(ansi.ClearScreen)
	ansi.Write(ansi.MoveToBottom)

	text, err := line.Prompt(getPrompt(calc))
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	for ; err == nil; text, err = line.Prompt(getPrompt(calc)) {
		var err error
		ansi.Write(ansi.ClearScreen)
		if strings.TrimSpace(text) == "" {
			if calc.Env.Stack.Len() > 0 {
				_, err = calc.Env.Stack.Pop()
			}
		} else {
			err = calc.EvalString("<cli>", text)
		}

		fmt.Println()

		for i, val := range calc.Env.Stack.Items() {
			color := ansi.LightBlue
			if i == calc.Env.Stack.Len()-1 {
				color = ansi.Bold
			}
			ansi.Write(color)
			fmt.Print(val)
			ansi.Write(ansi.Reset)
			fmt.Println()
		}
		if err != nil {
			calcError, ok := err.(zc.CalcError)
			if ok {
				for _, f := range calcError.Frames {
					fmt.Printf("%v @ %v:%v\n", f.Pos.File, f.Pos.Line, f.Pos.Column)
					if f.Func != "" {
						fmt.Printf("\t%v\n", f.Func)
					}
				}
			}
			ansi.Write(ansi.BrightYellow)
			fmt.Printf("(!) %v\n", err)
			ansi.Write(ansi.Reset)
		} else if calc.Info != "" {
			ansi.Write(ansi.LightGreen)
			fmt.Println(calc.Info)
			ansi.Write(ansi.Reset)
		} else {
			fmt.Println()
		}
		if strings.TrimSpace(text) != "" {
			line.AppendHistory(text)
		}
	}
}

func loadHistory(line *liner.State) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("unable to determine home directory: %v", err)
		return
	}

	historyFile := path.Join(homedir, ".local", "share", "zc", "history")
	file, err := os.Open(historyFile)
	if err != nil {
		log.Printf("unable to load history: %v", err)
		return
	}
	defer file.Close()
	if _, err := line.ReadHistory(file); err != nil {
		log.Printf("unable to load history: %v", err)
	}
}

func saveHistory(line *liner.State) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("unable to determine home directory: %v", err)
		return
	}

	localDir := path.Join(homedir, ".local", "share", "zc")
	if err := os.MkdirAll(localDir, 0o700); err != nil {
		log.Printf("unable to create local directory: %v", err)
		return
	}

	historyFile := path.Join(localDir, "history")
	file, err := os.OpenFile(historyFile, os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		log.Printf("unable to save history: %v", err)
		return
	}

	_, err = line.WriteHistory(file)
	if err != nil {
		log.Printf("unable to save history: %v", err)
	}
}

func getPrompt(calc *zc.Calc) string {
	prompt := "zc"
	if calc.Mode != "" {
		prompt += ":" + calc.Mode
	}
	return prompt + " > "
}
