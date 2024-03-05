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

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v5/pkg/ansi"
	"github.com/blackchip-org/zc/v5/pkg/zc"
	"github.com/peterh/liner"
)

var errQuit = errors.New("quit")

type REPL struct {
	MaxUndo     int
	Calc        zc.Calc
	Out         io.Writer
	EndQuote    string
	cli         *liner.State
	homeDir     string
	localDir    string
	historyFile string
	undoStack   [][]string
	redoStack   [][]string
	ops         map[string]struct{}
	macros      map[string]string
	info        string
	err         error
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
	r.cli.SetWordCompleter(r.WordCompleter)

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

func (r *REPL) evalLine(line string) error {
	var s scan.Scanner
	var out []string
	s.InitFromString("", line)

	for s.HasMore() {
		word := scan.Word(&s)
		if mac, ok := r.macros[word]; ok {
			out = append(out, mac)
		} else {
			out = append(out, word)
		}
		scan.Space(&s)
	}
	outLine := strings.Join(out, " ")
	return r.Calc.Eval(outLine)
}

func (r *REPL) Eval(line string) error {
	var s scan.Scanner

	r.info = ""
	r.err = nil
	prev := r.Calc.Stack()

	line = strings.TrimSpace(line)
	if r.EndQuote != "" {
		if line == r.EndQuote {
			r.EndQuote = ""
		} else {
			r.Calc.Push(line)
		}
		return nil
	}

	s.InitFromString("", line)

	cmdName := scan.Word(&s)
	cmd, ok := cmds[cmdName]

	var err error
	if ok {
		err = cmd(r, &s)
	} else {
		err = r.evalLine(line)
	}

	if err == nil && cmdName != "undo" && cmdName != "u" && cmdName != "redo" {
		r.undoStack = append([][]string{prev}, r.undoStack...)
		r.redoStack = nil
	}
	r.err = err
	if r.Calc.Info() != "" && r.info == "" {
		r.info = r.Calc.Info()
	}
	return err
}

func (r *REPL) Info() string {
	return r.info
}

func (r *REPL) Error() error {
	return r.err
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
	if r.EndQuote != "" {
		return fmt.Sprintf("… %s> ", r.EndQuote)
	}
	return zc.ProgName + " > "
}

func (r *REPL) WordCompleter(line string, pos int) (string, []string, string) {
	endPos := pos
	for endPos < len(line) {
		if line[endPos] == ' ' {
			break
		}
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
	return prefix, candidates, suffix
}

func CommonPrefix(vals []string) string {
	if len(vals) == 0 {
		return ""
	}
	var result []rune
	for i, sval := range vals {
		val := []rune(sval)
		if i == 0 {
			result = val
			continue
		}
		if len(result) == 0 {
			return ""
		}
		for j, a := range result {
			if j >= len(val) {
				result = result[:j]
				break
			}
			b := val[j]
			if a != b {
				result = result[:j]
				break
			}
		}
	}
	return string(result)
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
	r := New(c)
	r.Init()
	defer r.Close()
	for {
		line, err := r.ReadLine()
		if err != nil {
			if err.Error() != "prompt aborted" {
				log.Println(err)
			}
			return
		}
		ansi.Write(ansi.ClearScreen)

		prev := c.Stack()
		err = r.Eval(line)
		if err == errQuit {
			break
		}

		// Print out previous stack in dark gray
		ansi.Write(ansi.DarkGray)
		if err == nil {
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
		if err != nil {
			ansi.Write(ansi.BrightYellow)
			fmt.Fprintf(r.Out, "(!) %v\n", err)
			ansi.Fprint(r.Out, ansi.Reset)
		} else if r.Info() != "" {
			ansi.Fprint(r.Out, ansi.LightGreen)
			fmt.Fprintln(r.Out, r.Info())
			ansi.Write(ansi.Reset)
		} else {
			fmt.Fprintln(r.Out)
		}
		if strings.TrimSpace(line) != "" {
			if r.cli != nil {
				r.cli.AppendHistory(line)
			}
		}
	}
	for _, item := range c.Stack() {
		fmt.Println(item)
	}
	fmt.Println()
}
