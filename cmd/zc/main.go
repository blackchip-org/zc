package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/blackchip-org/zc/app"
	"github.com/blackchip-org/zc/lang"
)

var (
	parseFile string
)

func init() {
	flag.StringVar(&parseFile, "parse", "", "parse file and print out AST")
}

func main() {
	log.SetFlags(0)
	flag.Parse()
	if parseFile != "" {
		src, err := ioutil.ReadFile(parseFile)
		if err != nil {
			log.Fatalf("unable to read file: %v", err)
		}
		ast, err := lang.Parse("", src)
		if err != nil {
			log.Fatalf("parse error:\n%v", err)
		}
		fmt.Println(ast)
		return
	}
	app.RunConsole()
}
