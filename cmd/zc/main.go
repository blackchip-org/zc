package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/blackchip-org/zc/pkg/ansi"
	"github.com/blackchip-org/zc/pkg/calc"
	"github.com/blackchip-org/zc/pkg/repl"
)

func main() {
	log.SetFlags(0)
	c := calc.New()

	if len(os.Args) == 1 {
		if os.Getenv("ZC_NO_ANSI") != "" {
			ansi.Enabled = false
		}
		repl.Run(c)
		return
	}

	line := strings.Join(os.Args[1:], " ")
	err := c.Eval(line)
	if err != nil {
		log.Fatalf("(!) %v", err)
	}
	for _, item := range c.Stack() {
		fmt.Println(item)
	}
}
