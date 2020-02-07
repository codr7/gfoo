package gfoo

import (
)

var Symbol SymbolType

func init() {
	Symbol.Init("Symbol")
}

type SymbolType struct {
	StringType
}
