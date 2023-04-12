package types

import (
	"errors"
	"fmt"

	"github.com/blackchip-org/zc/scanner"
)

var (
	ErrNotSupported = errors.New("not supported")
	ErrParse        = errors.New("unable to parse")
)

type Type interface {
	String() string
	ParseValue(string) (Value, error)
}

type Value interface {
	String() string
	Type() Type
	Native() any
}

func isFormatting(ch rune) bool {
	if ch == ',' || ch == '_' || ch == ' ' {
		return true
	}
	if scanner.IsCurrency(ch) {
		return true
	}
	return false
}

func cleanNumber(str string) string {
	var s scanner.Scanner
	s.SetString(str)
	for s.Ok() {
		s.SkipIf(isFormatting)
	}
	return s.Token()
}

var QuoteFunc = func(s string) string { return "'" + s + "'" }

func parseErr(t Type, s string) error {
	return fmt.Errorf("expecting %v but got %v", t, QuoteFunc(s))
}
