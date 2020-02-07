package gfoo

import (
	"io"
	"strings"
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

func (typ *SymbolType) Compare(x, y interface{}) Order {
	return Order(strings.Compare(x.(string), y.(string)))
}
