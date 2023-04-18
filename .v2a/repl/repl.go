package repl

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/coll"
	"github.com/blackchip-org/zc/internal/ansi"
	"github.com/peterh/liner"
)

var errQuit = errors.New("quit")

type REPL struct {
	MaxUndo     int
	calc        zc.Calc
	cli         *liner.State
	homeDir     string
	localDir    string
	historyFile string
	undoStack   coll.Deque[[]string]
	redoStack   coll.Deque[[]string]
}

func New(calc zc.Calc) *REPL {
	return &REPL{
		calc:      calc,
		undoStack: coll.NewDequeSlice[[]string](),
		redoStack: coll.NewDequeSlice[[]string](),
	}
}

func (r *REPL) Init() {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("unable to determine home directory: %v", err)
		dir = ""
	}
	r.homeDir = dir
	r.localDir = path.Join(r.homeDir, ".local", "share", "zc")
	r.historyFile = path.Join(r.localDir, "history")

	consoleInit()

	r.cli = liner.NewLiner()
	r.cli.SetCtrlCAborts(true)
	r.cli.SetTabCompletionStyle(liner.TabPrints)
	//r.cli.SetWordCompleter(r.calc.WordCompleter)

	r.loadHistory()

	ansi.Write(ansi.ClearScreen)
	ansi.Write(ansi.MoveToBottom)
}

func (r *REPL) Close() {
	r.saveHistory()
	r.cli.Close()
}

func (r *REPL) ReadLine() (string, error) {
	text, err := r.cli.Prompt(r.getPrompt())
	if err != nil {
		log.Printf("error: %v", err)
		return "", err
	}
	return text, nil
}

func (r *REPL) Eval(text string) bool {
	var execError error

	prev := r.calc.Stack().Items()

	cmd := strings.TrimRight(text, " ")
	switch cmd {
	case "":
		execError = r.pop()
	case "quit":
		execError = r.quit()
	case "redo", "r":
		execError = r.redo()
	case "undo", "u":
		execError = r.undo()
	default:
		execError = r.eval(cmd)
	}
	if execError == errQuit {
		return false
	}

	// Print out previous stack in dark gray
	fmt.Print(ansi.DarkGray)
	if execError == nil {
		for _, val := range prev {
			fmt.Println(raw(val))
		}
		fmt.Println()
		coll.Push(r.undoStack, prev)
		r.redoStack.Clear()
	} else {
		r.calc.SetStack(prev)
	}
	fmt.Print(ansi.Reset)

	for i, val := range r.calc.Stack().Items() {
		color := ansi.LightBlue
		if i == r.calc.Stack().Len()-1 {
			color = ansi.Bold
		}
		fmt.Print(colorize(color, val))
		fmt.Println()
	}
	if execError != nil {
		calcError, ok := execError.(zc.CalcError)
		if ok {
			for _, f := range calcError.Frames {
				fmt.Printf("%v @ %v:%v\n", f.Pos().Name, f.Pos().Line, f.Pos().Column)
				if f.FuncName() != "" {
					fmt.Printf("\t%v\n", f.FuncName())
				}
			}
		}
		ansi.Write(ansi.BrightYellow)
		fmt.Printf("(!) %v\n", execError)
		ansi.Write(ansi.Reset)
	} else if r.calc.Info() != "" {
		ansi.Write(ansi.LightGreen)
		fmt.Println(r.calc.Info())
		ansi.Write(ansi.Reset)
	} else {
		fmt.Println()
	}
	if strings.TrimSpace(text) != "" {
		if r.cli != nil {
			r.cli.AppendHistory(text)
		}
	}
	return true
}

func (r *REPL) eval(line string) error {
	return zc.EvalString(r.calc, "<cli>", line)
}

func (r *REPL) pop() error {
	coll.Pop(r.calc.Stack())
	return nil
}

func (r *REPL) redo() error {
	redo, ok := coll.Pop(r.redoStack)
	if !ok {
		return fmt.Errorf("redo stack is empty")
	}
	coll.Push(r.undoStack, r.calc.Stack().Items())
	r.calc.SetStack(redo)
	return nil
}

func (r *REPL) quit() error {
	return errQuit
}

func (r *REPL) undo() error {
	undo, ok := coll.Pop(r.undoStack)
	if !ok {
		return fmt.Errorf("undo stack is empty")
	}
	coll.Push(r.redoStack, r.calc.Stack().Items())
	r.calc.SetStack(undo)
	return nil
}

func (r *REPL) loadHistory() {
	file, err := os.Open(r.historyFile)
	if err != nil {
		log.Printf("unable to load history: %v", err)
		return
	}
	defer file.Close()
	if _, err := r.cli.ReadHistory(file); err != nil {
		log.Printf("unable to load history: %v", err)
	}
}

func (r *REPL) saveHistory() {
	if err := os.MkdirAll(r.localDir, 0o700); err != nil {
		log.Printf("unable to create local directory: %v", err)
		return
	}
	file, err := os.OpenFile(r.historyFile, os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		log.Printf("unable to save history: %v", err)
		return
	}
	_, err = r.cli.WriteHistory(file)
	if err != nil {
		log.Printf("unable to save history: %v", err)
	}
}

func (r *REPL) getPrompt() string {
	prompt := "zc"
	// if r.calc.Mode != "" {
	// 	prompt += ":" + r.calc.Mode
	// }
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
