package zc

import "fmt"

func ErrDivisionByZero(env *OpEnv) error {
	return fmt.Errorf("%v: division by zero", env.Op.Name)
}

func ErrInvalidArg(env *OpEnv, format string, args ...any) error {
	args2 := append([]any{env.Op.Name}, args...)
	return fmt.Errorf("%v: invalid argument, "+format, args2...)
}
