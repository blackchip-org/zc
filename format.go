package zc

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/blackchip-org/scan"
	"github.com/blackchip-org/zc/v6/pkg/stack"
)

func FormatList(vals ...any) string {
	var strs []string
	for _, val := range vals {
		str := fmt.Sprintf("%v", val)
		strs = append(strs, str)
	}
	return strings.Join(strs, " | ")
}

func FormatStack(s stack.Stack[Item]) string {
	var strs []string
	for _, item := range s.Items() {
		str := fmt.Sprintf("%v", item.Value)
		strs = append(strs, str)
	}
	return strings.Join(strs, " | ")
}

func String(a any) string {
	return fmt.Sprintf("%v", a)
}

func Quote(v string) string {
	var s scan.Scanner
	s.InitFromString("", v)

	needsQuotes := false
	if !IsValuePrefix(s.This, s.Next) {
		needsQuotes = true
	} else {
		scan.Until(&s, scan.IsSpace, s.Discard)
		if s.HasMore() {
			needsQuotes = true
		}
	}

	if !needsQuotes {
		return v
	}

	s.InitFromString("", v)
	s.Val.WriteRune('\'')
	for s.HasMore() {
		if s.This == '\'' {
			s.Val.WriteString("\\'")
			s.Skip()
		} else {
			s.Keep()
		}
	}
	s.Val.WriteRune('\'')
	return s.Emit().Val
}

func TypeName(v any) string {
	var name strings.Builder
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
		name.WriteRune('*')
	}
	name.WriteString(t.Name())
	return name.String()
}
