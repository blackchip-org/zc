package zc

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/blackchip-org/zc/internal"
	"github.com/blackchip-org/zc/lang/lexer"
	"github.com/blackchip-org/zc/lang/parser"
	"github.com/blackchip-org/zc/scanner"
	"github.com/shopspring/decimal"
)

const (
	ValidSeparators = ",. _"
	ValidPoints     = ",."
)

const (
	DefaultMaxHistory = 10
)

var ErrFunctionNotAvailable = errors.New("feature not available")

type Config struct {
	ModuleDefs   []ModuleDef
	Preload      []string
	PreludeCLI   []string
	PreludeDev   []string
	Trace        bool
	RoundingMode RoundingMode
	MaxHistory   int
}

type ModuleDef struct {
	Name       string
	Include    bool
	ScriptPath string
	Natives    map[string]CalcFunc
	Init       CalcFunc
}

type CalcFunc func(*Env) error

type Frame struct {
	Pos  scanner.Pos
	Func string
	Env  *Env
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

type Calc struct {
	Config
	Mode       string
	Out        *strings.Builder
	Info       string
	Env        *Env
	ModuleDefs map[string]ModuleDef
	Modules    map[string]*Env
	Natives    map[string]CalcFunc
	States     map[string]any
	Frames     []Frame
}

func NewCalc(conf Config) (*Calc, error) {
	c := &Calc{
		Config:     conf,
		Modules:    make(map[string]*Env),
		Natives:    make(map[string]CalcFunc),
		ModuleDefs: make(map[string]ModuleDef),
		States:     make(map[string]any),
	}
	if c.MaxHistory == 0 {
		c.MaxHistory = DefaultMaxHistory
	}

	for _, def := range conf.ModuleDefs {
		c.ModuleDefs[def.Name] = def
	}

	c.Env = NewEnv(c, "zc")

	for _, name := range c.Preload {
		def, ok := c.ModuleDefs[name]
		if !ok {
			return nil, fmt.Errorf("no such preload module: %v", name)
		}
		if _, err := c.Load(def); err != nil {
			return nil, err
		}
	}

	for _, preName := range c.PreludeCLI {
		def, ok := c.ModuleDefs[preName]
		if !ok {
			return nil, fmt.Errorf("no such prelude module: %v", preName)
		}
		prefix := ""
		if !def.Include {
			prefix = def.Name
		}
		_, err := c.Env.Import(def, prefix)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Calc) Eval(name string, src []byte) error {
	c.Info = ""
	return Eval(c.Env, name, src)
}

func Eval(env *Env, name string, src []byte) error {
	root, err := parser.Parse(name, src)
	if err != nil {
		return err
	}
	env.Calc.Frames = append(env.Calc.Frames, Frame{
		Pos:  root.Pos(),
		Func: "",
		Env:  env,
	})
	err = env.evalFile(root)
	env.Calc.Frames = env.Calc.Frames[:len(env.Calc.Frames)-1]
	if err != nil {
		return err
	}
	return nil
}

func (c *Calc) EvalString(name string, src string) error {
	return c.Eval(name, []byte(src))
}

func (c *Calc) EvalLines(name string, lines []string) error {
	for _, line := range lines {
		if err := c.EvalString(name, line); err != nil {
			return err
		}
	}
	return nil
}

func (c *Calc) Load(def ModuleDef) (*Env, error) {
	if mod, ok := c.Modules[def.Name]; ok {
		return mod, nil
	}

	if os.Getenv("ZC_DEBUG_LOAD") != "" {
		log.Printf("load: %v", def.Name)
	}

	env := NewEnv(c, fmt.Sprintf("mod(%v)", def.Name))
	env.Module = def.Name

	for _, preName := range c.PreludeDev {
		mod, ok := c.Modules[preName]
		if !ok {
			continue
		}
		def := c.ModuleDefs[preName]
		for _, name := range mod.Exports {
			qName := name
			if !def.Include {
				qName = def.Name + "." + name
			}
			env.Funcs[qName] = mod.Funcs[name]
		}
	}

	if def.Natives != nil {
		for name, fn := range def.Natives {
			c.Natives[name] = fn
		}
	}

	if def.ScriptPath != "" {
		src, err := LoadFile(def.ScriptPath)
		if err != nil {
			return nil, err
		}
		ast, err := parser.Parse(def.ScriptPath, src)
		if err != nil {
			return nil, err
		}
		if err := env.evalFile(ast); err != nil {
			return nil, err
		}
	}

	if def.Init != nil {
		if err := def.Init(c.Env); err != nil {
			return nil, err
		}
	}

	c.Modules[def.Name] = env
	return env, nil
}

func (c *Calc) WordCompleter(line string, pos int) (string, []string, string) {
	endPos := pos
	for ; endPos < len(line); endPos++ {
		if line[endPos] == ' ' {
			break
		}
	}
	startPos := pos
	if startPos >= len(line) && len(line) > 0 {
		startPos = len(line) - 1
	}
	for ; startPos > 0; startPos-- {
		if line[startPos] == ' ' {
			startPos++
			break
		}
	}
	prefix := line[:startPos]
	word := line[startPos:endPos]
	suffix := line[endPos:]

	var candidates []string
	for name := range c.Env.Funcs {
		if strings.HasPrefix(name, word) {
			candidates = append(candidates, name)
		}
	}
	sort.Strings(candidates)
	//fmt.Printf("\n[%v] (%v)[%v] [%v]\n", prefix, word, candidates, suffix)
	return prefix, candidates, suffix
}

func (c *Calc) SetMode(name string) error {
	fileName := fmt.Sprintf("zc:modes/%v.zc", name)
	script, err := LoadFile(fileName)
	if err != nil {
		return fmt.Errorf("unable to load mode %v: %v", name, err)
	}
	if err := c.Eval(fileName, script); err != nil {
		return err
	}
	c.Mode = name
	return nil
}

type RoundingMode int

const (
	RoundingModeHalfUp RoundingMode = iota
	RoundingModeCeil
	RoundingModeDown
	RoundingModeFloor
	RoundingModeHalfEven
	RoundingModeUp
)

func (r RoundingMode) String() string {
	switch r {
	case RoundingModeHalfUp:
		return "half-up"
	case RoundingModeCeil:
		return "ceil"
	case RoundingModeDown:
		return "down"
	case RoundingModeFloor:
		return "floor"
	case RoundingModeHalfEven:
		return "half-even"
	case RoundingModeUp:
		return "up"
	}
	panic("unknown rounding mode")
}

func ParseRoundingMode(v string) (RoundingMode, bool) {
	switch strings.ToLower(v) {
	case "half-up":
		return RoundingModeHalfUp, true
	case "ceil":
		return RoundingModeCeil, true
	case "down":
		return RoundingModeDown, true
	case "floor":
		return RoundingModeFloor, true
	case "half-even":
		return RoundingModeHalfEven, true
	case "up":
		return RoundingModeUp, true
	}
	return 0, false
}

type RoundingFuncFix func(decimal.Decimal, int32) decimal.Decimal

var (
	RoundingFuncsFix = map[RoundingMode]RoundingFuncFix{
		RoundingModeCeil:     func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundCeil(places) },
		RoundingModeDown:     func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundDown(places) },
		RoundingModeFloor:    func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundFloor(places) },
		RoundingModeHalfUp:   func(d decimal.Decimal, places int32) decimal.Decimal { return d.Round(places) },
		RoundingModeHalfEven: func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundBank(places) },
		RoundingModeUp:       func(d decimal.Decimal, places int32) decimal.Decimal { return d.RoundUp(places) },
	}
)

func LoadFile(p string) ([]byte, error) {
	if strings.HasPrefix(p, "zc:") {
		p = p[3:]
		return internal.Files.ReadFile(p)
	}
	return os.ReadFile(p)
}

func ErrorWithStack(err error) string {
	if calcErr, ok := err.(CalcError); ok {
		var buf strings.Builder
		for _, frame := range calcErr.Frames {
			buf.WriteString(frame.String())
			buf.WriteRune('\n')
		}
		buf.WriteString(calcErr.Error())
		buf.WriteRune('\n')
		return buf.String()
	}
	return err.Error()
}

func Quote(v string) string {
	return lexer.Quote(v)
}
