package zc

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blackchip-org/zc/internal"
	"github.com/blackchip-org/zc/lang/ast"
	"github.com/blackchip-org/zc/lang/parser"
	"github.com/blackchip-org/zc/lang/token"
)

type Settings struct {
	Places       int32
	RoundMode    string
	NumberFormat NumberFormatOptions
}

func DefaultSettings() *Settings {
	return &Settings{
		Places:       16,
		RoundMode:    "half-up",
		NumberFormat: DefaultNumberFormatOptions(),
	}
}

type NumberFormatOptions struct {
	IntPat  string
	Point   rune
	FracPat string
}

func (n NumberFormatOptions) Separators() map[rune]struct{} {
	seps := make(map[rune]struct{})
	for _, pat := range []string{n.IntPat, n.FracPat} {
		for _, ch := range pat {
			if ch != '0' {
				seps[ch] = struct{}{}
			}
		}
	}
	return seps
}

func DefaultNumberFormatOptions() NumberFormatOptions {
	return NumberFormatOptions{
		IntPat: ",000",
		Point:  '.',
	}
}

var (
	//Places       int32  = 16
	RoundMode    string = "half-up"
	NumberFormat        = DefaultNumberFormatOptions()
)

type Config struct {
	ModuleDefs []ModuleDef
	PreludeCLI []string
	PreludeDev []string
	Trace      bool
}

type ModuleDef struct {
	Name       string
	Include    bool
	ScriptPath string
	Natives    map[string]CalcFunc
}

type Frame struct {
	Pos  token.Pos
	Func string
}

func (f Frame) String() string {
	return fmt.Sprintf("[%v] %v", f.Pos, f.Func)
}

type CalcError struct {
	Message string
	Frames  []Frame
}

func (c CalcError) Error() string {
	return c.Message
}

type CalcFunc func(*Calc) error

type Calc struct {
	Out      *strings.Builder
	Info     string
	Stack    *Stack
	name     string
	config   Config
	main     *Stack
	global   map[string]*Stack
	local    map[string]*Stack
	Funcs    map[string]CalcFunc
	Exports  map[string]CalcFunc
	Natives  map[string]CalcFunc
	defs     map[string]ModuleDef
	Modules  map[string]*Calc
	Settings *Settings
}

func NewCalc() *Calc {
	c, err := NewCalcWithConfig(Config{})
	if err != nil {
		panic(err)
	}
	return c
}

func NewCalcWithConfig(config Config) (*Calc, error) {
	c := &Calc{
		Out:      &strings.Builder{},
		name:     "<cli>",
		config:   config,
		main:     NewStack("main"),
		global:   make(map[string]*Stack),
		defs:     make(map[string]ModuleDef),
		Modules:  make(map[string]*Calc),
		Funcs:    make(map[string]CalcFunc),
		Exports:  make(map[string]CalcFunc),
		Natives:  make(map[string]CalcFunc),
		Settings: DefaultSettings(),
	}
	c.Stack = c.main
	c.global["main"] = c.Stack
	c.local = c.global

	for _, def := range config.ModuleDefs {
		c.Install(def)
	}
	for name, fn := range builtin {
		c.Funcs[name] = fn
	}
	c.Funcs["eval"] = eval

	for _, prelude := range config.PreludeCLI {
		if err := c.Include(prelude); err != nil {
			return nil, err
		}
	}

	c.Define("conf-places").Set("16")

	return c, nil
}

func (c *Calc) Eval(name string, src []byte) (err error) {
	c.Out.Reset()
	c.Info = ""
	ast, err := parser.Parse(name, src)
	if err != nil {
		return
	}
	err = c.evalNode(ast)

	return
}

func (c *Calc) EvalString(name string, src string) error {
	return c.Eval(name, []byte(src))
}

func (c *Calc) Define(name string) *Stack {
	stack, ok := c.local[name]
	if !ok {
		stack, ok = c.global[name]
		if !ok {
			stack = NewStack(name)
			c.local[name] = stack
		}
	}
	return stack
}

func (c *Calc) import_(def ModuleDef, prefix string) error {
	dc, err := c.load(def)
	if err != nil {
		return err
	}
	for funcName, fn := range dc.Exports {
		var qName string
		if prefix != "" {
			qName = prefix + "." + funcName
		} else {
			qName = funcName
		}
		c.Funcs[qName] = fn
	}
	return nil
}

func (c *Calc) Import(modName string, alias string) error {
	def, ok := c.defs[modName]
	if !ok {
		return fmt.Errorf("no such module: %v", modName)
	}

	prefix := modName
	if alias != "" {
		prefix = alias
	}
	return c.import_(def, prefix)
}

func (c *Calc) ImportFile(file string, name string) error {
	def := ModuleDef{Name: name, ScriptPath: file}
	return c.import_(def, name)
}

func (c *Calc) Include(modName string) error {
	def, ok := c.defs[modName]
	if !ok {
		return fmt.Errorf("no such module: %v", modName)
	}
	dc, err := c.load(def)
	if err != nil {
		return err
	}
	for funcName, fn := range dc.Exports {
		c.Funcs[funcName] = fn
	}
	return nil
}

func (c *Calc) IncludeFile(file string) error {
	src, err := c.LoadFile(file)
	if err != nil {
		return err
	}

	ast, err := parser.Parse(file, src)
	if err != nil {
		return nil
	}

	return c.evalNode(ast)
}

func (c *Calc) Install(def ModuleDef) {
	if def.Name == "" {
		panic(fmt.Sprintf("unable to install a module with no name: %+v", def))
	}
	c.defs[def.Name] = def
}

func (c *Calc) Interpolate(v string) (string, error) {
	var result, name strings.Builder

	inQuote := false
	inEscape := false

	for _, ch := range v {
		if ch == '`' && !inQuote && !inEscape {
			inQuote = true
		} else if ch == '`' && !inEscape {
			inQuote = false
			stack, err := c.StackFor(name.String())
			if err != nil {
				return "", fmt.Errorf("no such stack: %v", name.String())
			}
			for i, item := range stack.Items() {
				if i != 0 {
					result.WriteString("  ")
				}
				result.WriteString(item)
			}
			name.Reset()
		} else if ch == '\\' {
			inEscape = true
		} else {
			inEscape = false
			if inQuote {
				name.WriteRune(ch)
			} else {
				result.WriteRune(ch)
			}
		}
	}
	if name.Len() > 0 {
		return "", fmt.Errorf("expected`")
	}
	return result.String(), nil
}

func (c *Calc) Printf(format string, a ...any) {
	c.Out.WriteString(fmt.Sprintf(format, a...))
}

func (c *Calc) Print(a any) {
	c.Out.WriteString(fmt.Sprint(a))
}

func (c *Calc) StackFor(name string) (*Stack, error) {
	stack, ok := c.local[name]
	if ok {
		return stack, nil
	}
	stack, ok = c.global[name]
	if ok {
		return stack, nil
	}
	return nil, fmt.Errorf("no such stack: %v", name)
}

func (c *Calc) LoadFile(p string) ([]byte, error) {
	if strings.HasPrefix(p, "zc:") {
		p = p[3:]
		return internal.Files.ReadFile(p)
	}
	return os.ReadFile(p)
}

func (c *Calc) chain(node *ast.InvokeNode, err error) error {
	frame := Frame{Pos: node.Pos()}

	errCalc, ok := err.(CalcError)
	if ok {
		if len(errCalc.Frames) > 0 {
			errCalc.Frames[len(errCalc.Frames)-1].Func = node.Name
		}
		errCalc.Frames = append(errCalc.Frames, frame)
		return errCalc
	}
	return CalcError{
		Message: err.Error(),
		Frames:  []Frame{frame},
	}
}

func (c *Calc) err(node ast.Node, err error) error {
	errCalc, ok := err.(CalcError)
	if ok {
		return errCalc
	}
	frame := Frame{Pos: node.Pos()}
	return CalcError{
		Message: err.Error(),
		Frames:  []Frame{frame},
	}
}

func (c *Calc) trace(node ast.Node, format string, a ...any) {
	if c.config.Trace {
		msg := fmt.Sprintf(format, a...)
		if c.Stack.Len() > 0 {
			log.Printf("eval: %v(%v)", c.Stack.Name, c.Stack)
		}
		log.Printf("eval:     %v @ %v", msg, node.Pos())
		//log.Println()
	}
}
