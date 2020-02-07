package gfoo

import (
	"io"
)

var Symbol SymbolType

func init() {
	Symbol.Init()
}

type SymbolType struct {
	TypeBase
}

func (typ *SymbolType) Init() {
	typ.TypeBase.Init("Symbol")
}

func (typ *SymbolType) Dump(val interface{}, out io.Writer) error {
	_, err := io.WriteString(out, val.(string))
	return err
}
