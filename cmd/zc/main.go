package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/blackchip-org/zc/calc"
)

func main() {
	log.SetFlags(0)
	c := calc.New()
	flag.Parse()

	line := strings.Join(flag.Args(), " ")
	err := c.Eval(line)
	if err != nil {
		log.Fatalf("(!) %v", err)
	}

	for _, item := range c.Stack() {
		fmt.Println(item)
	}
}
