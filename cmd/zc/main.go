package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/blackchip-org/zc/v6"
	"github.com/blackchip-org/zc/v6/app"
	"github.com/blackchip-org/zc/v6/app/repl"
	"github.com/blackchip-org/zc/v6/pkg/ansi"
)

func main() {
	log.SetFlags(0)

	c := app.NewCalc()

	if len(os.Args) == 1 {
		if os.Getenv("ZC_NO_ANSI") != "" {
			ansi.Enabled = false
		}
		repl.Run(c)
		return
	}

	if os.Getenv("ZC_TRACE") != "" {
		c.Listener = zc.ConsoleLogger
		ansi.Enabled = false
	}

	var source string
	if len(os.Args) == 2 && os.Args[1] == "-" {
		in, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		source = string(in)
	} else {
		source = strings.Join(os.Args[1:], " ")
	}

	err := c.Eval(source)
	if err != nil {
		log.Fatalf("(!) %v", err)
	}
	for _, item := range c.Stack() {
		fmt.Println(item)
	}
	if c.Notice != "" {
		fmt.Println(c.Notice)
	}
}
