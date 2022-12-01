package lang

import "strings"

type Errors []error

func (e Errors) Error() string {
	var b strings.Builder
	for i, err := range e {
		b.WriteString(err.Error())
		if i != len(e)-1 {
			b.WriteString("\n")
		}
	}
	return b.String()
}

// type ErrorHandler func(error)
