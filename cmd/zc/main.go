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
	"github.com/blackchip-org/zc/lang/parser"
	"github.com/blackchip-org/zc/lang/scanner"
	"github.com/blackchip-org/zc/lang/token"
)

var (
	evalFile  string
	noAnsi    bool
	parseFile string
	scanFile  string
	trace     bool
)

func init() {
	flag.StringVar(&evalFile, "eval", "", "evaluate file")
	flag.BoolVar(&noAnsi, "no-ansi", false, "disable ANSI control codes")
	flag.StringVar(&parseFile, "parse", "", "parse file and print out the AST")
	flag.StringVar(&scanFile, "scan", "", "scan file and print out the tokens")
	flag.BoolVar(&trace, "trace", false, "trace execution")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	config := app.DefaultConfig()
	config.Trace = trace
	calc, err := zc.NewCalc(config)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case scanFile != "":
		scan()
	case parseFile != "":
		parse()
	case evalFile != "":
		eval(calc)
	default:
		if noAnsi {
			ansi.Enabled = false
		}
		app.RunConsole(calc)
	}
}

func eval(calc *zc.Calc) {
	src, err := ioutil.ReadFile(evalFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	if err := calc.Eval(evalFile, src); err != nil {
		log.Print(err)
		if cErr, ok := err.(zc.CalcError); ok {
			for _, frame := range cErr.Frames {
				log.Println(frame)
			}
		}
		os.Exit(1)
	}
	for _, item := range calc.Stack.Items() {
		fmt.Println(item)
	}
}

func parse() {
	src, err := ioutil.ReadFile(parseFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	// This is useful for generating test files for the parser. In this case,
	// omit the filename from the parser output.
	if os.Getenv("ZC_TEST") == "true" {
		parseFile = ""
	}
	ast, err := parser.Parse(parseFile, src)
	if err != nil {
		log.Fatalf("parse error:\n%v", err)
	}
	fmt.Println(ast)
}

func scan() {
	src, err := ioutil.ReadFile(scanFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	scanner := scanner.New(scanFile, src)
	fmt.Println("line col  token")
	for tok := scanner.Next(); tok.Type != token.End; tok = scanner.Next() {
		fmt.Printf("%4d %3d  %v\n", tok.Pos.Line, tok.Pos.Column, tok)
	}
}
