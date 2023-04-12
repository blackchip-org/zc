package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/blackchip-org/zc/ptime"
)

var (
	dateOnly   bool
	format     string
	localeName string
	timeOnly   bool
	verbose    bool
)

func main() {
	log.SetFlags(0)
	flag.BoolVar(&dateOnly, "d", false, "only parse date")
	flag.StringVar(&format, "f", "", "format the result with `layout`")
	flag.StringVar(&localeName, "l", "en-US", "set `locale`")
	flag.BoolVar(&timeOnly, "t", false, "only parse time")
	flag.BoolVar(&verbose, "v", false, "verbose")

	flag.Parse()

	text := strings.Join(flag.Args(), " ")
	p, err := ptime.ForLocale(localeName)
	if err != nil {
		log.Fatal(err)
	}

	if verbose {
		p.Parser.Trace = true
	}

	var parseFn func(string) (ptime.Parsed, error)
	switch {
	case dateOnly:
		parseFn = p.ParseDate
	case timeOnly:
		parseFn = p.ParseTime
	default:
		parseFn = p.Parse
	}
	res, err := parseFn(text)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if format != "" {
		t, err := p.Time(res, time.Now())
		if err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
		fmt.Println(p.Format(format, t))
	} else {
		b, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(string(b))
	}
}
