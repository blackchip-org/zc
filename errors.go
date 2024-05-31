package zc

import (
	"errors"
	"fmt"
)

var ErrStackEmpty = errors.New("stack empty")

func ErrDivisionByZero(env *OpEnv) error {
	return fmt.Errorf("%v: division by zero", env.Op.Name)
}

func ErrInvalidArg(env *OpEnv, format string, args ...any) error {
	args2 := append([]any{env.Op.Name}, args...)
	return fmt.Errorf("%v: invalid argument, "+format, args2...)
}

func ErrOp(env *OpEnv, err error) error {
	return fmt.Errorf("%v: %v", env.Op.Name, err)
}
