package zc

import (
	"fmt"
	"log"
	"math/big"
	"math/bits"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/blackchip-org/zc/internal"
	"github.com/blackchip-org/zc/lang/parser"
	"github.com/blackchip-org/zc/lang/token"
	"github.com/shopspring/decimal"
)

type Config struct {
	ModuleDefs   []ModuleDef
	PreludeCLI   []string
	PreludeDev   []string
	Trace        bool
	Places       int32
	RoundingMode RoundingMode
	IntPat       string
	Point        rune
	FracPat      string
}

type ModuleDef struct {
	Name       string
	Include    bool
	ScriptPath string
	Natives    map[string]CalcFunc
}

type CalcFunc func(*Env) error

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

type Calc struct {
	Config
	Out        *strings.Builder
	Info       string
	Env        *Env
	ModuleDefs map[string]ModuleDef
	Modules    map[string]*Env
	Natives    map[string]CalcFunc
}

func NewCalc(conf Config) (*Calc, error) {
	c := &Calc{
		Config:     conf,
		Modules:    make(map[string]*Env),
		Natives:    make(map[string]CalcFunc),
		ModuleDefs: make(map[string]ModuleDef),
	}
	for _, def := range conf.ModuleDefs {
		c.ModuleDefs[def.Name] = def
	}

	c.Env = NewEnv(c)

	for _, preName := range c.PreludeDev {
		def, ok := c.ModuleDefs[preName]
		if !ok {
			return nil, fmt.Errorf("no such prelude module: %v", preName)
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
	root, err := parser.Parse(name, src)
	if err != nil {
		return err
	}
	return c.Env.evalNode(root)
}

func (c *Calc) EvalString(name string, src string) error {
	return c.Eval(name, []byte(src))
}

func (c *Calc) Load(def ModuleDef) (*Env, error) {
	if mod, ok := c.Modules[def.Name]; ok {
		return mod, nil
	}

	env := NewEnv(c)
	env.Module = def.Name

	for _, preName := range c.PreludeDev {
		mod, ok := c.Modules[preName]
		if !ok {
			continue
		}
		for _, name := range mod.Exports {
			env.Funcs[name] = mod.Funcs[name]
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
		if err := env.evalNode(ast); err != nil {
			return nil, err
		}
	}

	c.Modules[def.Name] = env
	return env, nil
}

func (c *Calc) parseDigits(sep rune, v string) ([]rune, []rune) {
	var intDigits, fracDigits []rune
	inInt := true
	for _, ch := range v {
		if ch == sep {
			if !inInt {
				fracDigits = append(fracDigits, ch)
			}
			inInt = false
		} else if inInt {
			intDigits = append(intDigits, ch)
		} else {
			fracDigits = append(fracDigits, ch)
		}
	}
	return intDigits, fracDigits
}

func (c *Calc) FormatNumberString(v string) string {
	var digits strings.Builder
	intDigits, fracDigits := c.parseDigits('.', v)

	if c.IntPat == "" {
		digits.WriteString(string(intDigits))
	} else {
		var intResult []rune
		intPat := []rune(c.IntPat)

		idxPat := len(c.IntPat) - 1
		idxDig := len(intDigits) - 1
		for idxDig >= 0 {
			if intDigits[idxDig] == '-' {
				intResult = append([]rune{intDigits[idxDig]}, intResult...)
				idxDig--
			} else if intPat[idxPat] == '0' {
				intResult = append([]rune{intDigits[idxDig]}, intResult...)
				idxDig--
				idxPat--
			} else {
				intResult = append([]rune{intPat[idxPat]}, intResult...)
				idxPat--
			}
			if idxPat < 0 {
				idxPat = len(intPat) - 1
			}
		}
		digits.WriteString(string(intResult))
	}

	if len(fracDigits) == 0 {
		return digits.String()
	}

	point := c.Point
	if point == 0 {
		point = '.'
	}
	digits.WriteRune(point)

	if c.FracPat == "" {
		digits.WriteString(string(fracDigits))
	} else {
		var fracResult []rune
		fracPat := []rune(c.FracPat)

		idxPat := 0
		idxDig := 0
		for idxDig < len(fracDigits) {
			if fracPat[idxPat] == '0' {
				fracResult = append(fracDigits, fracResult[idxDig])
				idxDig++
				idxPat++
			} else {
				fracResult = append(fracDigits, fracPat[idxPat])
				idxPat++
			}
			if idxPat >= len(fracDigits) {
				idxPat = 0
			}
		}
		digits.WriteString(string(fracResult))
	}

	return digits.String()

}

func (c *Calc) FormatBigInt(v *big.Int) string {
	return fmt.Sprintf("%d", v)
}

func (c *Calc) FormatBigIntWithRadix(v *big.Int, radix int) string {
	sign := ""
	if v.Sign() < 0 {
		sign = "-"
	}
	var absV big.Int
	absV.Abs(v)

	switch radix {
	case 16:
		return fmt.Sprintf("%v0x%x", sign, &absV)
	case 8:
		return fmt.Sprintf("%v0o%o", sign, &absV)
	case 2:
		return fmt.Sprintf("%v0b%b", sign, &absV)
	}
	return c.FormatBigInt(v)
}

func (c *Calc) FormatBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func (c *Calc) FormatFixed(v decimal.Decimal) string {
	fn, ok := RoundingFuncsFix[c.RoundingMode]
	if !ok {
		log.Panicf("invalid rounding mode: %v", c.RoundingMode)
	}

	return fn(v, c.Places).String()
}

func (c *Calc) FormatInt64(i int64) string {
	return fmt.Sprintf("%v", i)
}

func (c *Calc) FormatInt32(i int32) string {
	return c.FormatInt64(int64(i))
}

func (c *Calc) FormatInt(i int) string {
	return c.FormatInt64(int64(i))
}

func (c *Calc) FormatUint64(i uint64) string {
	return fmt.Sprintf("%v", i)
}

func (c *Calc) FormatUint32(i uint32) string {
	return c.FormatUint64(uint64(i))
}

func (c *Calc) FormatUint(i uint) string {
	return c.FormatUint64(uint64(i))
}

func (c *Calc) FormatValue(v string) string {
	r := ParseRadix(v)
	switch {
	case r != 10:
		return v
	case c.IsBigInt(v):
		v := c.FormatBigIntWithRadix(c.MustParseBigInt(v), r)
		return c.FormatNumberString(v)
	case c.IsFixed(v):
		v := c.FormatFixed(c.MustParseFixed(v))
		return c.FormatNumberString(v)
	}
	return v
}

func (c *Calc) cleanNumString(v string) string {
	var sb strings.Builder
	// FIXME
	// seps := c.Settings.NumberFormat.Separators()
	for _, ch := range v {
		// if _, ok := seps[ch]; ok {
		// 	continue
		// }
		if ch == ',' {
			continue
		}
		if unicode.Is(unicode.Sc, ch) {
			continue
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func (c *Calc) ParseBigInt(v string) (*big.Int, error) {
	i := new(big.Int)
	_, ok := i.SetString(c.cleanNumString(v), 0)
	if !ok {
		return i, fmt.Errorf("expecting BigInt but got %v", v)
	}
	return i, nil
}

func (c *Calc) IsBigInt(v string) bool {
	_, err := c.ParseBigInt(v)
	return err == nil
}

func (c *Calc) MustParseBigInt(v string) *big.Int {
	i, err := c.ParseBigInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

func (c *Calc) ParseBool(v string) (bool, error) {
	vl := strings.ToLower(v)
	switch vl {
	case "true":
		return true, nil
	case "false":
		return false, nil
	}
	return false, fmt.Errorf("expecting Bool but got %v", v)
}

func (c *Calc) IsBool(v string) bool {
	_, err := c.ParseBool(v)
	return err == nil
}

func (c *Calc) MustParseBool(v string) bool {
	b, err := c.ParseBool(v)
	if err != nil {
		panic(err)
	}
	return b
}

func (c *Calc) ParseFixed(v string) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(c.cleanNumString(v))
	if err != nil {
		return decimal.Zero, fmt.Errorf("expecting Fixed but got %v", v)
	}
	return d, nil
}

func (c *Calc) IsFixed(v string) bool {
	_, err := c.ParseFixed(v)
	return err == nil
}

func (c *Calc) MustParseFixed(v string) decimal.Decimal {
	d, err := c.ParseFixed(v)
	if err != nil {
		panic(err)
	}
	return d
}

func (c *Calc) ParseInt(v string) (int, error) {
	i, err := strconv.ParseInt(c.cleanNumString(v), 0, 64)
	if err != nil {
		return 0, fmt.Errorf("expecting Int but got %v", v)
	}
	return int(i), nil
}

func (c *Calc) IsInt(v string) bool {
	_, err := c.ParseInt(v)
	return err == nil
}

func (c *Calc) MustParseInt(v string) int {
	i, err := c.ParseInt(v)
	if err != nil {
		panic(err)
	}
	return i
}

func (c *Calc) ParseInt32(v string) (int32, error) {
	i, err := strconv.ParseInt(c.cleanNumString(v), 0, 32)
	if err != nil {
		return 0, fmt.Errorf("expecting Int32 but got %v", v)
	}
	return int32(i), nil
}

func (c *Calc) ParseUint(v string) (uint, error) {
	i, err := strconv.ParseUint(c.cleanNumString(v), 0, bits.UintSize)
	if err != nil {
		return 0, fmt.Errorf("expecting Uint but got %v", v)
	}
	return uint(i), nil
}

func ParseRadix(v string) int {
	if len(v) < 2 {
		return 10
	}
	prefix := strings.ToLower(v[:2])
	switch prefix {
	case "0b":
		return 2
	case "0o":
		return 8
	case "0x":
		return 16
	}
	return 10
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
