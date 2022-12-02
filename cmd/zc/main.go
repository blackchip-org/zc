package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/app"
	"github.com/blackchip-org/zc/internal/modules"
	"github.com/blackchip-org/zc/lang"
)

var (
	evalFile  string
	parseFile string
	scanFile  string
	trace     bool
)

func init() {
	flag.StringVar(&evalFile, "eval", "", "evaluate file")
	flag.StringVar(&parseFile, "parse", "", "parse file and print out the AST")
	flag.StringVar(&scanFile, "scan", "", "scan file and print out the tokens")
	flag.BoolVar(&trace, "trace", false, "trace execution")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	switch {
	case scanFile != "":
		scan()
	case parseFile != "":
		parse()
	case evalFile != "":
		eval()
	default:
		app.RunConsole()
	}
}

func eval() {
	src, err := ioutil.ReadFile(evalFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	config := zc.Config{
		ModuleDefs: modules.All,
	}
	calc := zc.NewCalc(config)
	calc.Trace = trace
	if err := calc.Eval(src); err != nil {
		log.Fatal(err)
	}
	fmt.Print(calc.Stack().Items())
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
	ast, err := lang.Parse(parseFile, src)
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
	scanner := lang.NewScanner(scanFile, src)
	fmt.Println("line col  token")
	for tok := scanner.Next(); tok.Type != lang.EndToken; tok = scanner.Next() {
		fmt.Printf("%4d %3d  %v\n", tok.At.Line, tok.At.Column, tok)
	}
}
