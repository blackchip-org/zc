package zc

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/blackchip-org/zc/internal"
	"github.com/blackchip-org/zc/lang/parser"
	"github.com/blackchip-org/zc/lang/scanner"
	"github.com/blackchip-org/zc/lang/token"
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
	IntLayout    string
	Point        rune
	FracFormat   string
	MinDigits    uint
	AutoCurrency bool
	AutoFormat   bool
	Locale       string
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
	Pos  token.Pos
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

type FormatAttrs struct {
	Radix       int
	Currency    rune
	Fix         Fix
	ApplyLayout bool
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

	if c.Locale != "" {
		c.SetLocale(c.Locale)
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

func (c *Calc) SetLocale(name string) error {
	fileName := fmt.Sprintf("zc:locales/%v.zc", name)
	script, err := LoadFile(fileName)
	if err != nil {
		return fmt.Errorf("unable to load locale %v: %v", name, err)
	}
	if err := c.Eval(fileName, script); err != nil {
		return err
	}
	c.Locale = name
	return nil
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

func (c *Calc) FormatNumberString(v string, applyLayout bool) string {
	var digits strings.Builder
	intDigits, fracDigits := c.parseDigits('.', v)
	if c.IntLayout == "" || !applyLayout {
		digits.WriteString(string(intDigits))
	} else {
		var intResult []rune
		intPat := []rune(c.IntLayout)

		idxPat := len(c.IntLayout) - 1
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

	diff := int(c.MinDigits) - len(fracDigits)
	if c.MinDigits > 0 && diff > 0 {
		for i := 0; i < diff; i++ {
			fracDigits = append(fracDigits, '0')
		}
	}

	if len(fracDigits) == 0 {
		return digits.String()
	}

	point := c.Point
	if point == 0 {
		point = '.'
	}
	digits.WriteRune(point)

	if c.FracFormat == "" || !applyLayout {
		digits.WriteString(string(fracDigits))
	} else {
		var fracResult []rune
		fracPat := []rune(c.FracFormat)

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

func (c *Calc) FormatBigInt(v *big.Int, applyLayout bool) string {
	return c.FormatNumberString(v.String(), applyLayout)
}

func (c *Calc) FormatBigIntWithAttrs(v *big.Int, attrs FormatAttrs) string {
	sign := ""
	if v.Sign() < 0 {
		sign = "-"
	}
	var absV big.Int
	absV.Abs(v)

	switch attrs.Radix {
	case 16:
		return fmt.Sprintf("%v0x%x", sign, &absV)
	case 8:
		return fmt.Sprintf("%v0o%o", sign, &absV)
	case 2:
		return fmt.Sprintf("%v0b%b", sign, &absV)
	}

	s := c.FormatBigInt(v, attrs.ApplyLayout)
	return c.addCurrencySymbol(attrs, s)
}

func (c *Calc) FormatBool(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

func (c *Calc) FormatDecimal(v decimal.Decimal, applyLayout bool) string {
	return c.FormatNumberString(v.String(), applyLayout)
}

func (c *Calc) FormatDecimalWithAttrs(v decimal.Decimal, attrs FormatAttrs) string {
	s := c.FormatDecimal(v, attrs.ApplyLayout)
	return c.addCurrencySymbol(attrs, s)
}

func (c *Calc) FormatFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', 16, 64)
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

func (c *Calc) formatValueWithAttrs(v string, attrs FormatAttrs) string {
	switch {
	case attrs.Radix != 10:
		return v
	case c.IsBigInt(v):
		return c.FormatBigIntWithAttrs(c.MustParseBigInt(v), attrs)
	case c.IsDecimal(v):
		return c.FormatDecimalWithAttrs(c.MustParseDecimal(v), attrs)
	}
	return v
}

func (c *Calc) FormatValue(v string) string {
	attrs := ParseFormatAttrs(v)
	attrs.ApplyLayout = c.AutoFormat
	return c.formatValueWithAttrs(v, attrs)
}

func (c *Calc) ApplyLayout(v string) string {
	attrs := ParseFormatAttrs(v)
	attrs.ApplyLayout = true
	return c.formatValueWithAttrs(v, attrs)
}

func (c *Calc) cleanNumString(v string) string {
	var sb strings.Builder
	var buf strings.Builder

	seenPoint := false

	for _, ch := range v {
		if ch == c.Point {
			seenPoint = true
			buf.WriteRune('.')
			continue
		}
		if ch == '0' && seenPoint {
			buf.WriteRune(ch)
			continue
		}
		if strings.ContainsRune(ValidSeparators, ch) {
			continue
		}
		if unicode.Is(unicode.Sc, ch) {
			continue
		}
		if buf.Len() > 0 {
			sb.WriteString(buf.String())
			buf.Reset()
		}
		sb.WriteRune(ch)
	}
	return sb.String()
}

func (c *Calc) ParseBigInt(v string) (*big.Int, error) {
	i := new(big.Int)
	_, ok := i.SetString(c.cleanNumString(v), 0)
	if !ok {
		return i, fmt.Errorf("expecting BigInt but got %v", scanner.Quote(v))
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
	return false, fmt.Errorf("expecting Bool but got %v", scanner.Quote(v))
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

func (c *Calc) ParseDecimal(v string) (decimal.Decimal, error) {
	d, err := decimal.NewFromString(c.cleanNumString(v))
	if err != nil {
		return decimal.Zero, fmt.Errorf("expecting Decimal but got %v", scanner.Quote(v))
	}
	return d, nil
}

func (c *Calc) IsDecimal(v string) bool {
	_, err := c.ParseDecimal(v)
	return err == nil
}

func (c *Calc) MustParseDecimal(v string) decimal.Decimal {
	d, err := c.ParseDecimal(v)
	if err != nil {
		panic(err)
	}
	return d
}

func (c *Calc) ParseFloat(v string) (float64, error) {
	f, err := strconv.ParseFloat(c.cleanNumString(v), 64)
	if err != nil {
		return 0, fmt.Errorf("expecting Float but got %v", scanner.Quote(v))
	}
	return f, nil
}

func (c *Calc) IsFloat(v string) bool {
	_, err := c.ParseFloat(v)
	return err == nil
}

func (c *Calc) MustParseFloat(v string) float64 {
	d, err := c.ParseFloat(v)
	if err != nil {
		panic(err)
	}
	return d
}

func (c *Calc) ParseInt(v string) (int, error) {
	i, err := strconv.ParseInt(c.cleanNumString(v), 0, 64)
	if err != nil {
		return 0, fmt.Errorf("expecting Int but got %v", scanner.Quote(v))
	}
	return int(i), nil
}

func (c *Calc) IsInt(v string) bool {
	_, err := c.ParseInt(v)
	return err == nil
}

func (c *Calc) IsInt64(v string) bool {
	_, err := c.ParseInt64(v)
	return err == nil
}

func (c *Calc) IsInt32(v string) bool {
	_, err := c.ParseInt32(v)
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
		return 0, fmt.Errorf("expecting Int32 but got %v", scanner.Quote(v))
	}
	return int32(i), nil
}

func (c *Calc) ParseInt64(v string) (int64, error) {
	i, err := strconv.ParseInt(c.cleanNumString(v), 0, 64)
	if err != nil {
		return 0, fmt.Errorf("expecting Int64 but got %v", v)
	}
	return i, nil
}

func (c *Calc) ParseUint(v string) (uint, error) {
	i, err := strconv.ParseUint(c.cleanNumString(v), 0, bits.UintSize)
	if err != nil {
		return 0, fmt.Errorf("expecting Uint but got %v", scanner.Quote(v))
	}
	return uint(i), nil
}

func (c *Calc) ParseUint8(v string) (uint8, error) {
	i, err := strconv.ParseUint(c.cleanNumString(v), 0, 8)
	if err != nil {
		return 0, fmt.Errorf("expecting Uint8 but got %v", scanner.Quote(v))
	}
	return uint8(i), nil
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

func (c *Calc) ParseRune(v string) (rune, error) {
	if utf8.RuneCountInString(v) != 1 {
		return 0, fmt.Errorf("expected Rune but got %v", Quote(v))
	}
	r, _ := utf8.DecodeRuneInString(v)
	return r, nil
}

func (c *Calc) IsRune(v string) bool {
	_, err := c.ParseRune(v)
	return err == nil
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

type Fix int

const (
	NoFix Fix = iota
	Prefix
	Suffix
)

func ParseCurrencySymbol(v string) (rune, Fix) {
	prefix, _ := utf8.DecodeRuneInString(v)
	if unicode.Is(unicode.Sc, prefix) {
		return prefix, Prefix
	}
	suffix, _ := utf8.DecodeLastRuneInString(v)
	if unicode.Is(unicode.Sc, suffix) {
		return suffix, Suffix
	}
	return rune(0), NoFix
}

func (c *Calc) addCurrencySymbol(attrs FormatAttrs, v string) string {
	if !c.AutoCurrency {
		return v
	}
	switch attrs.Fix {
	case Prefix:
		return string(attrs.Currency) + v
	case Suffix:
		return v + string(attrs.Currency)
	}
	return v
}

func ParseFormatAttrs(xs ...string) FormatAttrs {
	attrs := FormatAttrs{}
	badCurrency := false

	for _, x := range xs {
		radix := ParseRadix(x)
		if attrs.Radix == 0 || attrs.Radix == 10 || radix > attrs.Radix {
			attrs.Radix = radix
		}
		sym, fix := ParseCurrencySymbol(x)
		if fix != NoFix && !badCurrency {
			if attrs.Currency != rune(0) && attrs.Currency != sym {
				badCurrency = true
			} else {
				attrs.Currency, attrs.Fix = sym, fix
			}
		}
	}
	return attrs
}

func (c *Calc) LocalizeNumber(v string) string {
	return strings.ReplaceAll(v, ".", string(c.Point))
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
	return scanner.Quote(v)
}
