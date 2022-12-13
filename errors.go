package zc

import (
	"fmt"
)

func ErrUnsupported(name string) error {
	return fmt.Errorf("unsupported operation: %v", name)
}
