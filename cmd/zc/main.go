package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/blackchip-org/zc/calc"
	"github.com/blackchip-org/zc/repl"
)

func main() {
	log.SetFlags(0)
	c := calc.New()
	flag.Parse()

	if flag.NArg() == 0 {
		repl.Run(c)
		return
	}

	line := strings.Join(flag.Args(), " ")
	err := c.Eval(line)
	if err != nil {
		log.Fatalf("(!) %v", err)
	}
	for _, item := range c.Stack() {
		fmt.Println(item)
	}
}
