package app

import (
	"errors"

	"github.com/blackchip-org/zc"
)

var commands = map[string]func(*zc.Calc, string) error{
	"":     pop,
	"quit": quit,
	"r":    redo,
	"redo": redo,
	"u":    undo,
	"undo": undo,
}
var errQuit = errors.New("quit")

func eval(c *zc.Calc, line string) error {
	return c.EvalString("<cli>", line)
}

func pop(c *zc.Calc, line string) error {
	var err error
	if c.Env.Stack.Len() > 0 {
		_, err = c.Env.Stack.Pop()
	}
	return err
}

func redo(c *zc.Calc, line string) error {
	return c.Redo()
}

func quit(c *zc.Calc, line string) error {
	return errQuit
}

func undo(c *zc.Calc, line string) error {
	return c.Undo()
}
