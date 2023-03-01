package zlib

import (
	"github.com/blackchip-org/zc"
)

func BigInt(env *zc.Env) error {
	_, err := env.Stack.PopBigInt()
	if err != nil {
		return err
	}
	return nil
}

func BoolAssert(env *zc.Env) error {
	_, err := env.Stack.PopBool()
	if err != nil {
		return err
	}
	return nil
}

func Fixed(env *zc.Env) error {
	_, err := env.Stack.PopFixed()
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
