package ints

import (
	"github.com/blackchip-org/zc/v6"
)

func addIntArch(e *zc.OpEnv) {
	x := e.Args[0].(int)
	y := e.Args[1].(int)
	e.Returns = []any{x + y}
}

func addInt8(e *zc.OpEnv) {
	x := e.Args[0].(int8)
	y := e.Args[1].(int8)
	e.Returns = []any{x + y}
}

func addInt16(e *zc.OpEnv) {
	x := e.Args[0].(int16)
	y := e.Args[1].(int16)
	e.Returns = []any{x + y}
}

func addInt32(e *zc.OpEnv) {
	x := e.Args[0].(int32)
	y := e.Args[1].(int32)
	e.Returns = []any{x + y}
}

func addInt64(e *zc.OpEnv) {
	x := e.Args[0].(int64)
	y := e.Args[1].(int64)
	e.Returns = []any{x + y}
}

func addIntArchU(e *zc.OpEnv) {
	x := e.Args[0].(uint)
	y := e.Args[1].(uint)
	e.Returns = []any{x + y}
}

func addInt8U(e *zc.OpEnv) {
	x := e.Args[0].(uint8)
	y := e.Args[1].(uint8)
	e.Returns = []any{x + y}
}

func addInt16U(e *zc.OpEnv) {
	x := e.Args[0].(uint16)
	y := e.Args[1].(uint16)
	e.Returns = []any{x + y}
}

func addInt32U(e *zc.OpEnv) {
	x := e.Args[0].(uint32)
	y := e.Args[1].(uint32)
	e.Returns = []any{x + y}
}

func addInt64U(e *zc.OpEnv) {
	x := e.Args[0].(uint64)
	y := e.Args[1].(uint64)
	e.Returns = []any{x + y}
}

func isIntArch(e *zc.OpEnv) {
	x := e.Args[0].(string)
	_, ok := IntArchKind.To(x)
	e.Returns = []any{ok}
}

func int64ToInt8(e *zc.OpEnv) {
	x := e.Args[0].(int64)
	e.Returns = []any{int8(x)}
}

func uint64ToInt8(e *zc.OpEnv) {
	x := e.Args[0].(uint64)
	e.Returns = []any{int8(x)}
}

func isInt8(e *zc.OpEnv) {
	x := e.Args[0].(string)
	_, ok := Int8Kind.To(x)
	e.Returns = []any{ok}
}

// func int8_(e *zc.OpEnv) {
// 	x := e.Args[0].(int64)
// 	e.Returns = []any{int8(x)}
// }

// func isInt8(e *zc.OpEnv) {
// 	x := e.Args[0].(string)
// 	_, ok := Int8Kind.To(x)
// 	e.Returns = []any{ok}
// }

// func int8_(e *zc.OpEnv) {
// 	x := e.Args[0].(int64)
// 	e.Returns = []any{int8(x)}
// }

// func isInt8(e *zc.OpEnv) {
// 	x := e.Args[0].(string)
// 	_, ok := Int8Kind.To(x)
// 	e.Returns = []any{ok}
// }

// func int8_(e *zc.OpEnv) {
// 	x := e.Args[0].(int64)
// 	e.Returns = []any{int8(x)}
// }

// func isInt8(e *zc.OpEnv) {
// 	x := e.Args[0].(string)
// 	_, ok := Int8Kind.To(x)
// 	e.Returns = []any{ok}
// }
