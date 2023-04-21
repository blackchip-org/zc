package main

import (
	"fmt"
	"log"
	"time"

	"github.com/blackchip-org/zc/pkg/ptime"
	"github.com/blackchip-org/zc/pkg/ptime/locale"
)

func main() {
	p := ptime.For(locale.EnUS)

	parsed, err := p.Parse("3:04:05pm MST")
	if err != nil {
		log.Panic(err)
	}
	t, err := p.Time(parsed, time.Now())
	if err != nil {
		log.Panic(err)
	}
	f := p.Format("[hour]:[minute]:[second] [offset]", t)
	fmt.Println(f)
}
