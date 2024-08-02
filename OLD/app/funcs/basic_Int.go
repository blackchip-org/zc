package funcs

import "github.com/blackchip-org/zc/v6"

func AddUint(e *zc.OpEnv) {
	y := zc.Uint.Pop(e)
	x := zc.Uint.Pop(e)
	z := x + y
	zc.Uint.Push(e, z)
}

func AddUint8(e *zc.OpEnv) {
	y := zc.Uint8.Pop(e)
	x := zc.Uint8.Pop(e)
	z := x + y
	zc.Uint8.Push(e, z)
}

func AddUint16(e *zc.OpEnv) {
	y := zc.Uint16.Pop(e)
	x := zc.Uint16.Pop(e)
	z := x + y
	zc.Uint16.Push(e, z)
}

func AddUint32(e *zc.OpEnv) {
	y := zc.Uint32.Pop(e)
	x := zc.Uint32.Pop(e)
	z := x + y
	zc.Uint32.Push(e, z)
}

func AddUint64(e *zc.OpEnv) {
	y := zc.Uint64.Pop(e)
	x := zc.Uint64.Pop(e)
	z := x + y
	zc.Uint64.Push(e, z)
}
