package gfoo

import (
	"io"
)

type Symbol struct {
	name string
}

func NewSymbol(name string) *Symbol {
	return &Symbol{name: name}
}

func (sym *Symbol) Quote() *Value {
	return NewValue(&TSymbol, sym) 
}

var TSymbol SymbolType

func init() {
	TSymbol.Init()
}

type SymbolType struct {
	TypeBase
}

func (typ *SymbolType) Init() {
	typ.TypeBase.Init("Symbol")
}

func (typ *SymbolType) Dump(val interface{}, out io.Writer) error {
	_, err := io.WriteString(out, val.(*Symbol).name)
	return err
}

