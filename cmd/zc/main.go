package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/blackchip-org/zc/v5/pkg/ansi"
	"github.com/blackchip-org/zc/v5/pkg/repl"
	"github.com/blackchip-org/zc/v5/pkg/zc"
	"github.com/blackchip-org/zc/v5/pkg/zcalc"
)

func main() {
	log.SetFlags(0)

	c := zcalc.New()

	if len(os.Args) == 1 {
		if os.Getenv("ZC_NO_ANSI") != "" {
			ansi.Enabled = false
		}
		repl.Run(c)
		return
	}

	if os.Getenv("ZC_TRACE") != "" {
		c.Trace = true
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
		fmt.Println(zc.FormatStackItem(item))
	}
	if c.Info() != "" {
		fmt.Println(c.Info())
	}
}
