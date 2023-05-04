package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

//go:generate go run main.go

const (
	EmojiGo = "../../pkg/ops/emoji.go"
	EmojiMd = "../../doc/ops/emoji.md"
)

type Emoji struct {
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Group           string `json:"group"`
	EmojiVersion    string `json:"emoji_version"`
	UnicodeVersion  string `json:"unicode_version"`
	SkinToneSupport bool   `json:"skin_tone_support"`
}

func main() {
	var names []string
	nameToCh := make(map[string]string)

	emojii := make(map[string]Emoji)
	data, err := os.ReadFile("data-by-emoji.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &emojii)

	fent, err := os.Create(EmojiGo)
	if err != nil {
		panic(err)
	}
	defer fent.Close()

	fmt.Fprintf(fent, "package ops\n\n")
	fmt.Fprintf(fent, "var Emoji = map[string]string {")
	for ch, emoji := range emojii {
		name := slugToName(emoji.Slug)
		fmt.Fprintf(fent, "\n\t\":%v:\": \"[%v]\",", name, ch)
		names = append(names, name)
		nameToCh[name] = ch

		if emoji.SkinToneSupport {
			for i := 0; i < 5; i++ {
				ch2 := ch + string(rune(0x1f3fb+i))
				name2 := fmt.Sprintf("%v-%v", name, i+1)
				fmt.Fprintf(fent, "\n\t\":%v:\": \"[%v]\",", name2, ch2)
				names = append(names, name2)
				nameToCh[name2] = ch2
			}
		}
	}
	fmt.Fprintf(fent, "\n}\n")

	fdoc, err := os.Create(EmojiMd)
	if err != nil {
		log.Panic(err)
	}
	defer fdoc.Close()

	fmt.Fprintf(fdoc, `
# emoji

Unicode emoji characters.

| Operation | Description
|-----------|------------
`)

	sort.Strings(names)
	for _, name := range names {
		ch := nameToCh[name]
		fmt.Printf("NAME %v CH %v\n", name, ch)
		fmt.Fprintf(fdoc, "| `:%v:` | %v\n", name, ch)
	}

}

func slugToName(slug string) string {
	return fmt.Sprintf("%v", strings.ReplaceAll(slug, "_", "-"))
}
