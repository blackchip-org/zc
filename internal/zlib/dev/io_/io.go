package io_

import (
	"fmt"

	"github.com/blackchip-org/zc"
)

func Print(env *zc.Env) error {
	a, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	fmt.Print(a)
	return nil
}
