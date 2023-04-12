package zc

import (
	"github.com/blackchip-org/zc/lang/lexer"
	"github.com/blackchip-org/zc/types"
)

// FIXME: this is a hack
func init() {
	types.QuoteFunc = lexer.Quote
}
