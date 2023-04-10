//go:build proj

package geo

import (
	"fmt"

	"github.com/blackchip-org/zc"
	"github.com/twpayne/go-proj/v10"
)

func Transform(env *zc.Env) error {
	target, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	source, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	y, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	x, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}

	fmt.Printf("s %v t %v x %v y %v\n", source, target, x, y)
	sourceCoord := proj.NewCoord(x, y, 0, 0)
	pj, err := proj.NewCRSToCRS(source, target, nil)
	if err != nil {
		return err
	}
	targetCoord, err := pj.Forward(sourceCoord)
	if err != nil {
		return err
	}

	env.Stack.PushFloat(targetCoord.X())
	env.Stack.PushFloat(targetCoord.Y())
	return nil
}
