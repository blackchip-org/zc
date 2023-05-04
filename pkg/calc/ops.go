package calc

import (
	"fmt"

	"github.com/blackchip-org/zc/pkg/ops"
	"github.com/blackchip-org/zc/pkg/zc"
)

var opsTable map[string]zc.CalcFunc

func addMacros(table map[string]string) {
	for k, v := range table {
		opsTable[k] = evalOp(zc.Macro(k, v))
	}
}

func init() {
	opsTable = make(map[string]zc.CalcFunc)
	for _, v := range opsList {
		k := v.Name
		if _, exists := opsTable[k]; exists {
			panic(fmt.Sprintf("duplicate operation: %v", k))
		}
		opsTable[k] = evalOp(v)
	}
	addMacros(ops.TimeZones)
	addMacros(ops.Emoji)
	addMacros(ops.Entities)
}
