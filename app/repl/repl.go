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
	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app"
	"github.com/blackchip-org/zc/v6/pkg/ansi"
	"github.com/peterh/liner"
)

var errQuit = errors.New("quit")

type Repl struct {
	MaxUndo     int
	Calc        *app.Calc
	Out         io.Writer
	EndQuote    string
	cli         *liner.State
	homeDir     string
	localDir    string
	historyFile string
	undoStack   []zc.Stack[zc.Item]
	redoStack   []zc.Stack[zc.Item]
	ops         map[string]struct{}
	macros      map[string][]scan.Token
	info        string
	err         error
}

func New(calc *app.Calc) *Repl {
	r := &Repl{
		Calc:   calc,
		Out:    os.Stdout,
		ops:    make(map[string]struct{}),
		macros: make(map[string][]scan.Token),
	}
	for _, n := range calc.Catalog.OpNames() {
		r.ops[n] = struct{}{}
	}
	return r
}

func (r *Repl) Init() {
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

func (r *Repl) Close() {
	r.saveHistory()
	r.cli.Close()
}

func (r *Repl) ReadLine() (string, error) {
	text, err := r.cli.Prompt(r.getPrompt())
	return text, err
}

func (r *Repl) evalLine(toks []scan.Token) error {
	r.Calc.Err = nil
	for _, tok := range toks {
		var mac []scan.Token
		if tok.Type == zc.TokenName {
			mac = r.macros[tok.Val]
		}
		if len(mac) > 0 {
			r.Calc.EvalToken(mac...)
		} else {
			r.Calc.EvalToken(tok)
		}
	}
	return r.Calc.Err
}

func (r *Repl) Eval(line string) error {
	r.info = ""
	r.err = nil
	prev := r.Calc.Stack.Clone()

	if r.EndQuote != "" {
		if strings.TrimSpace(line) == r.EndQuote {
			r.EndQuote = ""
		} else {
			r.Calc.PushVal(line)
		}
		return nil
	}

	var cmdName string
	var err error
	toks := zc.ScanWords(line)

	switch {
	case len(toks) == 0:
		if r.EndQuote == "" && r.Calc.Stack.Len() > 0 {
			r.Calc.Stack.Pop()
		}
	case toks[0].Type == zc.TokenName:
		cmdName = toks[0].Val
		cmd, ok := cmds[cmdName]
		if ok {
			err = cmd(r, toks[1:])
		} else {
			r.evalLine(toks)
		}
	default:
		r.evalLine(toks)
	}

	if err == nil && cmdName != "undo" && cmdName != "u" && cmdName != "redo" {
		r.undoStack = append([]zc.Stack[zc.Item]{prev}, r.undoStack...)
		r.redoStack = nil
	}
	r.err = err
	if r.Calc.Info != "" && r.info == "" {
		r.info = r.Calc.Info
	}
	return err
}

func (r *Repl) Info() string {
	return r.info
}

func (r *Repl) Error() error {
	return r.err
}

func (r *Repl) loadHistory() {
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

func (r *Repl) saveHistory() {
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

func (r *Repl) getPrompt() string {
	if r.EndQuote != "" {
		return fmt.Sprintf("… %s> ", r.EndQuote)
	}
	return zc.ProgName + " > "
}

func (r *Repl) WordCompleter(line string, pos int) (string, []string, string) {
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

func colorize(color string, item zc.Item) string {
	// FIXME: Should be done another way now
	// if strings.HasPrefix(text, "#raw:") {
	// 	return text[5:] + ansi.Reset
	// }

	var b strings.Builder

	if item.Label != "" {
		ansi.Fprint(&b, ansi.DarkGray)
		fmt.Fprintf(&b, "%v: ", item.Label)
	}

	ansi.Fprint(&b, ansi.Reset)
	ansi.Fprint(&b, color)
	b.WriteString(zc.Format(item.Val))

	if item.Unit != "" {
		ansi.Fprint(&b, ansi.DarkGray)
		fmt.Fprintf(&b, item.Unit)
	}

	// if item.Label != "" {
	// 	ansi.Fprint(&b, ansi.DarkGray)
	// 	fmt.Fprintf(&b, " (%v)", item.Label)
	// }

	ansi.Fprint(&b, ansi.Reset)
	return b.String()
}

func raw(item zc.Item) string {
	var s strings.Builder
	for _, ch := range item.String() {
		if ch == '\033' {
			s.WriteString("\\033")
		} else {
			s.WriteRune(ch)
		}
	}
	return s.String()
}

func Run(c *app.Calc) {
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

		prev := c.Stack.Clone()
		err = r.Eval(line)
		if err == errQuit {
			break
		}

		// Print out previous stack in dark gray
		ansi.Write(ansi.DarkGray)
		if err == nil {
			if ansi.Enabled {
				for _, val := range prev.Items() {
					fmt.Fprintln(r.Out, raw(val))
				}
				fmt.Fprintln(r.Out)
			}
		} else {
			r.Calc.Stack = prev
		}
		ansi.Write(ansi.Reset)

		for i, val := range r.Calc.Stack.Items() {
			color := ansi.LightBlue
			if i == r.Calc.Stack.Len()-1 {
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
	for _, item := range c.Stack.Items() {
		fmt.Println(item)
	}
	fmt.Println()
}
