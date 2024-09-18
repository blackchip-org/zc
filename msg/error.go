package msg

import "fmt"

func ErrInvalidRoundingMode(m string) error {
	return fmt.Errorf("invalid rounding mode: %v", m)
}
