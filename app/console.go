package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/internal/ansi"
	"github.com/peterh/liner"
)

var errQuit = errors.New("quit")

type Console struct {
	MaxUndo     int
	calc        *zc.Calc
	cli         *liner.State
	homeDir     string
	localDir    string
	historyFile string
	undoStack   []*zc.Stack
	redoStack   []*zc.Stack
}

func NewConsole(calc *zc.Calc) *Console {
	return &Console{
		calc: calc,
	}
}

func (c *Console) Init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("unable to determine home directory: %v", err)
		dir = ""
	}
	c.homeDir = dir
	c.localDir = path.Join(c.homeDir, ".local", "share", "zc")
	c.historyFile = path.Join(c.localDir, "history")

	consoleInit()

	c.cli = liner.NewLiner()
	c.cli.SetCtrlCAborts(true)
	c.cli.SetTabCompletionStyle(liner.TabPrints)
	c.cli.SetWordCompleter(c.calc.WordCompleter)

	c.loadHistory()

	ansi.Write(ansi.ClearScreen)
	ansi.Write(ansi.MoveToBottom)
}

func (c *Console) Close() {
	c.saveHistory()
	c.cli.Close()
}

func (c *Console) ReadLine() (string, error) {
	text, err := c.cli.Prompt(c.getPrompt())
	if err != nil {
		log.Printf("error: %v", err)
		return "", err
	}
	return text, nil
}

func (c *Console) Eval(text string) bool {
	var execError error

	prev := c.calc.Env.Main.Copy()

	cmd := strings.TrimRight(text, " ")
	switch cmd {
	case "":
		execError = c.pop()
	case "quit":
		execError = c.quit()
	case "redo", "r":
		execError = c.redo()
	case "undo", "u":
		execError = c.undo()
	default:
		execError = c.eval(cmd)
	}
	if execError == errQuit {
		return false
	}

	// Print out previous stack in dark gray
	fmt.Print(ansi.DarkGray)
	if execError == nil {
		for _, val := range prev.Items() {
			fmt.Println(raw(val))
		}
		fmt.Println()
		c.undoStack = append([]*zc.Stack{prev}, c.undoStack...)
		c.redoStack = nil
	} else {
		c.calc.Env.SetMain(prev)
	}
	fmt.Print(ansi.Reset)

	for i, val := range c.calc.Env.Main.Items() {
		color := ansi.LightBlue
		if i == c.calc.Env.Stack.Len()-1 {
			color = ansi.Bold
		}
		fmt.Print(colorize(color, val))
		fmt.Println()
	}
	if execError != nil {
		calcError, ok := execError.(zc.CalcError)
		if ok {
			for _, f := range calcError.Frames {
				fmt.Printf("%v @ %v:%v\n", f.Pos.File, f.Pos.Line, f.Pos.Column)
				if f.Func != "" {
					fmt.Printf("\t%v\n", f.Func)
				}
			}
		}
		ansi.Write(ansi.BrightYellow)
		fmt.Printf("(!) %v\n", execError)
		ansi.Write(ansi.Reset)
	} else if c.calc.Info != "" {
		ansi.Write(ansi.LightGreen)
		fmt.Println(c.calc.Info)
		ansi.Write(ansi.Reset)
	} else {
		fmt.Println()
	}
	if strings.TrimSpace(text) != "" {
		if c.cli != nil {
			c.cli.AppendHistory(text)
		}
	}
	return true
}

func (c *Console) eval(line string) error {
	return c.calc.EvalString("<cli>", line)
}

func (c *Console) pop() error {
	c.calc.Env.Stack.Pop()
	return nil
}

func (c *Console) redo() error {
	if len(c.redoStack) == 0 {
		return fmt.Errorf("redo stack is empty")
	}
	c.undoStack = append([]*zc.Stack{c.calc.Env.Main}, c.undoStack...)
	c.calc.Env.SetMain(c.redoStack[0])
	c.redoStack = c.redoStack[1:]
	return nil
}

func (c *Console) quit() error {
	return errQuit
}

func (c *Console) undo() error {
	if len(c.undoStack) == 0 {
		return fmt.Errorf("undo stack is empty")
	}
	c.redoStack = append([]*zc.Stack{c.calc.Env.Main.Copy()}, c.redoStack...)
	c.calc.Env.SetMain(c.undoStack[0])
	c.undoStack = c.undoStack[1:]
	return nil
}

func (c *Console) loadHistory() {
	file, err := os.Open(c.historyFile)
	if err != nil {
		log.Printf("unable to load history: %v", err)
		return
	}
	defer file.Close()
	if _, err := c.cli.ReadHistory(file); err != nil {
		log.Printf("unable to load history: %v", err)
	}
}

func (c *Console) saveHistory() {
	if err := os.MkdirAll(c.localDir, 0o700); err != nil {
		log.Printf("unable to create local directory: %v", err)
		return
	}
	file, err := os.OpenFile(c.historyFile, os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		log.Printf("unable to save history: %v", err)
		return
	}
	_, err = c.cli.WriteHistory(file)
	if err != nil {
		log.Printf("unable to save history: %v", err)
	}
}

func (c *Console) getPrompt() string {
	prompt := "zc"
	if c.calc.Mode != "" {
		prompt += ":" + c.calc.Mode
	}
	return prompt + " > "
}

func colorize(color string, text string) string {
	if strings.HasPrefix(text, "#raw:") {
		return text[5:] + ansi.Reset
	}
	return color + text + ansi.Reset
}

func raw(text string) string {
	var s strings.Builder
	for _, ch := range text {
		if ch == '\033' {
			s.WriteString("\\033")
		} else {
			s.WriteRune(ch)
		}
	}
	return s.String()
}
