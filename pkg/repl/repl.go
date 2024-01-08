package repl

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sort"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/v5/pkg/ansi"
	"github.com/blackchip-org/zc/v5/pkg/scanner"
	"github.com/blackchip-org/zc/v5/pkg/zc"
	"github.com/peterh/liner"
)

var errQuit = errors.New("quit")

type REPL struct {
	MaxUndo     int
	Calc        zc.Calc
	Out         io.Writer
	cli         *liner.State
	homeDir     string
	localDir    string
	historyFile string
	undoStack   [][]string
	redoStack   [][]string
	ops         map[string]struct{}
	macros      map[string]string
	quoteEnd    string
}

func New(calc zc.Calc) *REPL {
	r := &REPL{
		Calc:   calc,
		Out:    os.Stdout,
		ops:    make(map[string]struct{}),
		macros: make(map[string]string),
	}
	for _, o := range calc.OpNames() {
		r.ops[o] = struct{}{}
	}
	return r
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
	r.cli.SetWordCompleter(r.wordCompleter)

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
	var s scanner.Scanner

	prev := r.Calc.Stack()

	s.SetString(text)
	s.ScanWhile(unicode.IsSpace)
	cmdName := s.Scan(scanner.Word)
	cmd, ok := cmds[cmdName]
	// FIXME: should these errors always be set in the calculator?
	if ok {
		execError = cmd(r, &s)
		r.Calc.SetError(execError)
	} else {
		execError = r.eval(text)
		if execError == nil {
			r.undoStack = append([][]string{prev}, r.undoStack...)
			r.redoStack = nil
		}
	}
	if execError == errQuit {
		return false
	}

	// Print out previous stack in dark gray
	ansi.Write(ansi.DarkGray)
	if execError == nil {
		if ansi.Enabled {
			for _, val := range prev {
				fmt.Fprintln(r.Out, raw(val))
			}
			fmt.Fprintln(r.Out)
		}
	} else {
		r.Calc.SetStack(prev)
	}
	ansi.Write(ansi.Reset)

	for i, val := range r.Calc.Stack() {
		color := ansi.LightBlue
		if i == len(r.Calc.Stack())-1 {
			color = ansi.Bold
		}
		fmt.Fprint(r.Out, colorize(color, val))
		fmt.Fprintln(r.Out)
	}
	if execError != nil {
		ansi.Write(ansi.BrightYellow)
		fmt.Fprintf(r.Out, "(!) %v\n", execError)
		ansi.Fprint(r.Out, ansi.Reset)
	} else if r.Calc.Info() != "" {
		ansi.Fprint(r.Out, ansi.LightGreen)
		fmt.Fprintln(r.Out, r.Calc.Info())
		ansi.Write(ansi.Reset)
	} else {
		fmt.Fprintln(r.Out)
	}
	if strings.TrimSpace(text) != "" {
		if r.cli != nil {
			r.cli.AppendHistory(text)
		}
	}
	return true
}

func (r *REPL) eval(line string) error {
	if r.quoteEnd != "" {
		if line == r.quoteEnd {
			r.quoteEnd = ""
		} else {
			r.Calc.Push(line)
		}
		return nil
	}

	var out []string
	var s scanner.Scanner
	s.SetString(line)

	for s.Ok() {
		word := s.Scan(scanner.Word)
		if mac, ok := r.macros[word]; ok {
			out = append(out, mac)
		} else {
			out = append(out, word)
		}
	}
	outLine := strings.Join(out, " ")
	return r.Calc.Eval(outLine)
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
	if r.quoteEnd != "" {
		return "...>"
	}
	return zc.ProgName + " > "
}

func (r *REPL) wordCompleter(line string, pos int) (string, []string, string) {
	endPos := pos
	for endPos < len(line) {
		if line[endPos] == ' ' {
			break
		}
		endPos++
	}
	startPos := pos - 1
	if startPos < 0 {
		startPos = 0
	}
	if startPos >= len(line) && len(line) > 0 {
		startPos = len(line) - 1
	}
	for startPos > 0 {
		if line[startPos] == ' ' {
			startPos++
			break
		}
		startPos--
	}
	prefix := line[:startPos]
	word := line[startPos:endPos]
	suffix := line[endPos:]

	var candidates []string
	for name := range r.ops {
		if strings.HasPrefix(name, word) {
			candidates = append(candidates, name)
		}
	}
	sort.Strings(candidates)
	//fmt.Printf("\n[%v] (%v)[%v] [%v]\n", prefix, word, candidates, suffix)
	return prefix, candidates, suffix
}

func colorize(color string, text string) string {
	if !ansi.Enabled {
		return zc.FormatStackItem(text)
	}
	if strings.HasPrefix(text, "#raw:") {
		return text[5:] + ansi.Reset
	}

	var display string
	parts := strings.SplitN(text, zc.AnnotationMarker, 2)
	display = parts[0]
	if len(parts) > 1 {
		display += ansi.DarkGray + "#" + parts[1]
	}
	return color + display + ansi.Reset
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
	return zc.FormatStackItem(s.String())
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
