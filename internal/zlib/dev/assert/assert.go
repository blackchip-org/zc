package assert

import (
	"github.com/blackchip-org/zc"
	"github.com/blackchip-org/zc/types"
)

func BigInt(env *zc.Env) error {
	s, err := env.Stack.Pop()
	if err != nil {
		return err
	}
	g := types.Parse(s)
	_, err = types.To(g, types.BigInt)
	return err
}

func BoolAssert(env *zc.Env) error {
	_, err := env.Stack.PopBool()
	if err != nil {
		return err
	}
	return nil
}

func Decimal(env *zc.Env) error {
	_, err := env.Stack.PopDecimal()
	if err != nil {
		return err
	}
	return nil
}

func Float(env *zc.Env) error {
	_, err := env.Stack.PopFloat()
	if err != nil {
		return err
	}
	return nil
}

func Int(env *zc.Env) error {
	_, err := env.Stack.PopInt()
	if err != nil {
		return err
	}
	return nil
}

func Int64(env *zc.Env) error {
	_, err := env.Stack.PopInt64()
	if err != nil {
		return err
	}
	return nil
}

func Int32(env *zc.Env) error {
	_, err := env.Stack.PopInt32()
	if err != nil {
		return err
	}
	return nil
}
