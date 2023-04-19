package repl

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/pkg/ansi"
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
	undoStack   [][]string
	redoStack   [][]string
}

func New(calc zc.Calc) *REPL {
	return &REPL{calc: calc}
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

	consoleInit(r)

	r.cli = liner.NewLiner()
	r.cli.SetCtrlCAborts(true)
	r.cli.SetTabCompletionStyle(liner.TabPrints)
	// r.cli.SetWordCompleter(r.calc.WordCompleter)

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
	return text, err
}

func (r *REPL) Eval(text string) bool {
	var execError error

	prev := r.calc.Stack()

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
	ansi.Write(ansi.DarkGray)
	if execError == nil {
		for _, val := range prev {
			fmt.Println(raw(val))
		}
		fmt.Println()
		r.undoStack = append([][]string{prev}, r.undoStack...)
		r.redoStack = nil
	} else {
		r.calc.SetStack(prev)
	}
	ansi.Write(ansi.Reset)

	for i, val := range r.calc.Stack() {
		color := ansi.LightBlue
		if i == len(r.calc.Stack())-1 {
			color = ansi.Bold
		}
		fmt.Print(colorize(color, val))
		fmt.Println()
	}
	if execError != nil {
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
	return r.calc.Eval(line)
}

func (r *REPL) pop() error {
	r.calc.Pop()
	return nil
}

func (r *REPL) redo() error {
	if len(r.redoStack) == 0 {
		return fmt.Errorf("redo stack is empty")
	}
	r.undoStack = append([][]string{r.calc.Stack()}, r.undoStack...)
	r.calc.SetStack(r.redoStack[0])
	r.redoStack = r.redoStack[1:]
	return nil
}

func (r *REPL) quit() error {
	return errQuit
}

func (r *REPL) undo() error {
	if len(r.undoStack) == 0 {
		return fmt.Errorf("undo stack is empty")
	}
	r.redoStack = append([][]string{r.calc.Stack()}, r.redoStack...)
	r.calc.SetStack(r.undoStack[0])
	r.undoStack = r.undoStack[1:]
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
	return zc.ProgName + "> "
}

func colorize(color string, text string) string {
	if !ansi.Enabled {
		return text
	}
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

func Run(c zc.Calc) {
	repl := New(c)
	repl.Init()
	defer repl.Close()
	for {
		line, err := repl.ReadLine()
		if err != nil {
			if err.Error() != "prompt aborted" {
				log.Println(err)
			}
			return
		}
		ansi.Write(ansi.ClearScreen)
		if ok := repl.Eval(line); !ok {
			break
		}
	}
	for _, item := range c.Stack() {
		fmt.Println(item)
	}
	fmt.Println()
}
