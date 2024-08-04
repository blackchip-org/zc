package coll

import (
	"errors"
	"fmt"
)

var ErrStackEmpty = errors.New("stack empty")

func ErrIndexOutOfBounds(i int) error {
	return fmt.Errorf("index out of bounds: %v", i)
}

func ErrNotEnoughArgs(have int, want int) error {
	return fmt.Errorf("not enough arguments, expected %v but got %v", want, have)
}
