package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/app"
	"github.com/blackchip-org/zc/internal/ansi"
	"github.com/blackchip-org/zc/lang/lexer"
	"github.com/blackchip-org/zc/lang/parser"
	"github.com/blackchip-org/zc/lang/token"
)

var (
	noAnsi     bool
	noFileName bool
	mode       string
	trace      bool
	use        string
	verbose    bool
)

var cmdHelp = `
Commands:
  eval      Evaluate from command line arguments
  file      Evaluate from file
  parse     Show parse tree
  scan      Show tokens from scanner
  test      Run test file
`

var cmds map[string]*flag.FlagSet

func init() {
	cmds = make(map[string]*flag.FlagSet)

	main := flag.NewFlagSet("", flag.ContinueOnError)
	commonFlags(main)
	main.BoolVar(&noAnsi, "no-ansi", false, "disable ANSI control codes")
	cmds[""] = main

	eval := flag.NewFlagSet("eval", flag.ExitOnError)
	commonFlags(eval)
	cmds["eval"] = eval

	file := flag.NewFlagSet("file", flag.ExitOnError)
	commonFlags(file)
	cmds["file"] = file

	parse := flag.NewFlagSet("parse", flag.ExitOnError)
	commonFlags(parse)
	parse.BoolVar(&noFileName, "no-filename", false, "do not use filename in output")
	cmds["parse"] = parse

	scan := flag.NewFlagSet("scan", flag.ExitOnError)
	commonFlags(scan)
	cmds["scan"] = scan

	test := flag.NewFlagSet("test", flag.ExitOnError)
	commonFlags(test)
	cmds["test"] = scan
}

func commonFlags(fs *flag.FlagSet) {
	fs.StringVar(&mode, "m", "", "start calculator with this mode")
	fs.BoolVar(&trace, "trace", false, "trace execution")
	fs.StringVar(&use, "u", "", "use this module")
	fs.BoolVar(&verbose, "v", false, "print additional information to the console")
}

func main() {
	log.SetFlags(0)

	flags := cmds[""]
	args := os.Args[1:]
	cmd := ""

	if len(os.Args) > 1 {
		arg1 := os.Args[1]
		fs, ok := cmds[arg1]
		if ok {
			cmd = arg1
			flags = fs
			args = os.Args[2:]
		}
	}
	res := flags.Parse(args)
	if res == flag.ErrHelp {
		fmt.Println(cmdHelp)
		os.Exit(1)
	}

	config := app.DefaultConfig()
	config.Trace = trace
	calc, err := zc.NewCalc(config)
	if err != nil {
		log.Fatal(zc.ErrorWithStack(err))
	}

	if mode != "" {
		if err := calc.SetMode(mode); err != nil {
			log.Fatal(zc.ErrorWithStack(err))
		}
	}
	if use != "" {
		if err := calc.EvalString("<cli>", "use "+use); err != nil {
			log.Fatal(zc.ErrorWithStack(err))
		}
	}

	switch cmd {
	case "scan":
		scan(flags)
	case "parse":
		parse(flags)
	case "file":
		evalFile(flags, calc)
	case "test":
		testFile(flags, calc)
	case "eval":
		evalLines(flags, calc)
	default:
		if flags.NArg() > 0 {
			log.Fatal("extra command line arguments, did you intend to use 'eval'?")
		}
		if noAnsi {
			ansi.Enabled = false
		}
		c := app.NewConsole(calc)
		c.Init()
		for {
			line, err := c.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			ansi.Write(ansi.ClearScreen)
			if ok := c.Eval(line); !ok {
				break
			}
		}
		c.Close()
		for _, item := range calc.Env.Stack.Items() {
			fmt.Println(item)
		}
		fmt.Println()
	}
}

func evalLines(flags *flag.FlagSet, calc *zc.CalcImpl) {
	var err error
	for i, line := range flags.Args() {
		name := fmt.Sprintf("<cli:%v>", i)
		if err = calc.EvalString(name, line); err != nil {
			break
		}
	}
	evalResults(calc, err)
}

func evalFile(flags *flag.FlagSet, calc *zc.CalcImpl) {
	var calcErr error
	for _, fileName := range flags.Args() {
		src, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		if calcErr = calc.Eval(fileName, src); calcErr != nil {
			break
		}
	}
	evalResults(calc, calcErr)
}

func evalResults(calc *zc.CalcImpl, err error) {
	if err != nil {
		log.Print(err)
		if cErr, ok := err.(zc.CalcError); ok {
			for _, frame := range cErr.Frames {
				log.Println(frame)
			}
		}
		os.Exit(1)
	}
	for _, item := range calc.Env.Stack.Items() {
		fmt.Println(item)
	}
}

func parse(flags *flag.FlagSet) {
	for _, fileName := range flags.Args() {
		src, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		// This is useful for generating test files for the parser. In this case,
		// omit the filename from the parser output.
		if noFileName {
			fileName = ""
		}
		ast, err := parser.Parse(fileName, src)
		if err != nil {
			log.Fatalf("parse error:\n%v", err)
		}
		fmt.Println(ast)
	}
}

func scan(flags *flag.FlagSet) {
	for _, fileName := range flags.Args() {
		src, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		l := lexer.New(fileName, src)
		fmt.Println("line col  token")
		for tok := l.Next(); tok.Type != token.End; tok = l.Next() {
			fmt.Printf("%4d %3d  %v\n", tok.Pos.Line, tok.Pos.Column, tok)
		}
	}
}

func testFile(flags *flag.FlagSet, calc *zc.CalcImpl) {
	if err := calc.SetMode("dev"); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	cmd := []string{
		"import test",
		fmt.Sprintf("%v test.verbose", verbose),
	}
	if err := calc.EvalLines("", cmd); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	for _, fileName := range flags.Args() {
		cmd := fmt.Sprintf("'%v' test.file", fileName)
		if err := calc.EvalString(fileName, cmd); err != nil {
			log.Fatalf("error running test %v: %v", fileName, err)
		}
	}

	if err := calc.EvalString("", "test.report test.ok"); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	ok, err := calc.Env.Stack.PopBool()
	if err != nil {
		log.Fatalf("unexpected error: %v", err)
	}
	if !ok {
		os.Exit(1)
	}
}